#!/usr/bin/env bash
# Runs the ClickHouse collector integration tests across the full support
# matrix: every supported version × {single-node, cluster} topology.
#
# Each combination is started, validated and torn down in turn, so the matrix
# fits within modest memory. Override the matrix with env vars:
#
#   CLICKHOUSE_VERSIONS="25.3 24.8"  CLICKHOUSE_TOPOLOGIES="single" bash run-matrix.sh
#   CLICKHOUSE_SINGLE_TCP_PORT=39000 CLICKHOUSE_CLUSTER_NODE1_TCP_PORT=39010 \
#     CLICKHOUSE_CLUSTER_NODE2_TCP_PORT=39011 bash run-matrix.sh
#
# Usage:  bash run-matrix.sh
set -uo pipefail

cd "$(dirname "$0")"
repo_root=$(cd ../../../.. && pwd)
compose=(docker compose -f docker-compose.matrix.yml)

cleanup_profile() {
    local profile="$1"
    if "${compose[@]}" --profile "$profile" down -v >/dev/null 2>&1; then
        return 0
    fi

    echo "!!! failed to clean up ClickHouse ${profile} profile" >&2
    return 1
}

cleanup() {
    local cleanup_rc=0
    if ! cleanup_profile single; then
        cleanup_rc=1
    fi
    if ! cleanup_profile cluster; then
        cleanup_rc=1
    fi
    return "$cleanup_rc"
}

on_exit() {
    local rc=$?
    if ! cleanup && [ "$rc" -eq 0 ]; then
        rc=1
    fi
    exit "$rc"
}

on_signal() {
    cleanup
    exit 130
}

trap on_exit EXIT
trap on_signal INT TERM

# Build the clickhouse_exporter binary once, up front, so the exporter matrix
# test (TestClickHouseExporterMatrix) can launch the packaged-equivalent binary
# per endpoint. Its path is handed to the tests via CLICKHOUSE_EXPORTER_BIN.
exporter_bin="${repo_root}/bin/clickhouse_exporter"
echo ">>> building clickhouse_exporter"
if ! CGO_ENABLED=0 go -C "$repo_root" build -o bin/clickhouse_exporter ./agent/cmd/clickhouse_exporter; then
    echo "!!! failed to build clickhouse_exporter"
    exit 1
fi
export CLICKHOUSE_EXPORTER_BIN="$exporter_bin"

require_test() {
    local pkg="$1"
    local test_name="$2"
    local listed
    if ! listed=$(go -C "$repo_root" test -tags clickhouse_integration -run '^$' -list "^${test_name}$" "$pkg"); then
        echo "!!! failed to list ${test_name} in ${pkg}" >&2
        return 1
    fi
    if [[ "$listed" != *"$test_name"* ]]; then
        echo "!!! required integration test ${test_name} is missing from ${pkg}" >&2
        return 1
    fi
}

if ! require_test ./agent/agents/clickhouse TestClickHouseMatrix; then
    exit 1
fi
if ! require_test ./agent/agents/clickhouse TestClickHouseExporterMatrix; then
    exit 1
fi
if ! require_test ./agent/agents/clickhouse/querylog TestClickHouseQANMatrix; then
    exit 1
fi

# Supported ClickHouse versions and topologies (override via env).
read -r -a versions <<<"${CLICKHOUSE_VERSIONS:-26.3 25.8 25.3 24.8 24.3}"
read -r -a topologies <<<"${CLICKHOUSE_TOPOLOGIES:-single cluster}"
export CLICKHOUSE_SINGLE_TCP_PORT="${CLICKHOUSE_SINGLE_TCP_PORT:-39000}"
export CLICKHOUSE_CLUSTER_NODE1_TCP_PORT="${CLICKHOUSE_CLUSTER_NODE1_TCP_PORT:-39010}"
export CLICKHOUSE_CLUSTER_NODE2_TCP_PORT="${CLICKHOUSE_CLUSTER_NODE2_TCP_PORT:-39011}"

if [[ "${DOCKER_HOST:-}" == tcp://* ]]; then
    export PMM_TEST_BIND_HOST="${PMM_TEST_BIND_HOST:-0.0.0.0}"
fi

if [ -z "${PMM_TEST_SERVICE_HOST:-}" ]; then
    if [[ "${DOCKER_HOST:-}" == tcp://* ]]; then
        docker_host="${DOCKER_HOST#tcp://}"
        docker_host="${docker_host%%/*}"
        docker_host="${docker_host%%:*}"
        if [ -z "$docker_host" ]; then
            echo "!!! DOCKER_HOST did not contain a host: ${DOCKER_HOST}" >&2
            exit 1
        fi
        export PMM_TEST_SERVICE_HOST="$docker_host"
    else
        export PMM_TEST_SERVICE_HOST="127.0.0.1"
    fi
fi

rc=0
for v in "${versions[@]}"; do
    for topo in "${topologies[@]}"; do
        echo ">>> ClickHouse ${v} / ${topo}"
        export CLICKHOUSE_IMAGE="clickhouse/clickhouse-server:${v}"

        if ! "${compose[@]}" --profile "$topo" up -d --wait --wait-timeout 240; then
            echo "!!! ${v}/${topo}: containers did not become healthy"
            if ! "${compose[@]}" --profile "$topo" logs --tail 30; then
                echo "!!! ${v}/${topo}: failed to collect compose logs" >&2
            fi
            if ! cleanup_profile "$topo"; then
                echo "!!! ${v}/${topo}: failed to clean up unhealthy containers" >&2
            fi
            rc=1
            continue
        fi

        if [ "$topo" = "single" ]; then
            endpoints="single-${v}=clickhouse://default:clickhouse@${PMM_TEST_SERVICE_HOST}:${CLICKHOUSE_SINGLE_TCP_PORT}/default"
        else
            endpoints="cluster-${v}-node1=clickhouse://default:clickhouse@${PMM_TEST_SERVICE_HOST}:${CLICKHOUSE_CLUSTER_NODE1_TCP_PORT}/default"
            endpoints+=",cluster-${v}-node2=clickhouse://default:clickhouse@${PMM_TEST_SERVICE_HOST}:${CLICKHOUSE_CLUSTER_NODE2_TCP_PORT}/default"
        fi

        if ! CLICKHOUSE_TEST_ENDPOINTS="$endpoints" \
            go -C "$repo_root" test -tags clickhouse_integration -count=1 -v \
            -run '^(TestClickHouseMatrix|TestClickHouseExporterMatrix)$' ./agent/agents/clickhouse; then
            echo "!!! ${v}/${topo}: collector/exporter integration tests failed"
            rc=1
        fi

        if ! CLICKHOUSE_TEST_ENDPOINTS="$endpoints" \
            go -C "$repo_root" test -tags clickhouse_integration -count=1 -v \
            -run '^TestClickHouseQANMatrix$' ./agent/agents/clickhouse/querylog; then
            echo "!!! ${v}/${topo}: QAN integration tests failed"
            rc=1
        fi

        if ! cleanup_profile "$topo"; then
            rc=1
        fi
    done
done

if [ "$rc" -eq 0 ]; then
    echo "=== ClickHouse matrix: ALL PASSED ==="
else
    echo "=== ClickHouse matrix: FAILURES (see above) ==="
fi
exit "$rc"
