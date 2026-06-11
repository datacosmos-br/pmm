// Copyright (C) 2023 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

// Package testutils contains integration-test helpers for qan-api2.
package testutils

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	defaultServiceHost   = "127.0.0.1"
	clickHouseNativePort = 19000
)

// ServiceHost returns the host where Docker-published integration services are reachable.
func ServiceHost(tb testing.TB) string {
	tb.Helper()

	if host := strings.TrimSpace(os.Getenv("PMM_TEST_SERVICE_HOST")); host != "" {
		return host
	}

	dockerHost := strings.TrimSpace(os.Getenv("DOCKER_HOST"))
	if dockerHost == "" || strings.HasPrefix(dockerHost, "unix://") || strings.HasPrefix(dockerHost, "npipe://") {
		return defaultServiceHost
	}

	u, err := url.Parse(dockerHost)
	require.NoErrorf(tb, err, "failed to parse DOCKER_HOST=%q; set PMM_TEST_SERVICE_HOST", dockerHost)
	if u.Scheme != "tcp" {
		require.FailNowf(
			tb,
			"unsupported DOCKER_HOST for integration service discovery",
			"DOCKER_HOST=%q must be empty, unix://, npipe://, or tcp://host:port; set PMM_TEST_SERVICE_HOST",
			dockerHost,
		)
	}

	host := u.Hostname()
	require.NotEmptyf(tb, host, "DOCKER_HOST=%q did not include a host; set PMM_TEST_SERVICE_HOST", dockerHost)

	return host
}

// ServiceAddr returns a host:port address for an integration service.
func ServiceAddr(tb testing.TB, port int) string {
	tb.Helper()

	require.Positive(tb, port, "integration service port must be positive")

	return net.JoinHostPort(ServiceHost(tb), strconv.Itoa(port))
}

// ClickHouseDSN returns the ClickHouse DSN used by qan-api2 integration tests.
func ClickHouseDSN(tb testing.TB, database string) string {
	tb.Helper()

	require.NotEmpty(tb, database, "ClickHouse database must be set")

	if dsn := strings.TrimSpace(os.Getenv("QANAPI_DSN_TEST")); dsn != "" {
		u, err := url.Parse(dsn)
		require.NoErrorf(tb, err, "failed to parse QANAPI_DSN_TEST=%q", dsn)
		u.Path = "/" + database

		return u.String()
	}

	return fmt.Sprintf("clickhouse://default:clickhouse@%s/%s", ServiceAddr(tb, clickHouseNativePort), database)
}
