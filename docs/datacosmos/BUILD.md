# PMM — datacosmos build & fork model

This is a datacosmos fork of [`percona/pmm`](https://github.com/percona/pmm),
tracking the upstream **`v3`** branch.

## Branch model

| Branch | Purpose | How to keep current |
|---|---|---|
| `main` | Integration branch — the datacosmos default branch; merges `feat/build` and `feat/clickhouse-collector`. | `git merge upstream/v3` |
| `feat/build` | datacosmos OCI/RPM packaging pipeline (`Makefile.datacosmos`, `datacosmos-release.yml`) + agent coordination protocol. The branch the team builds and releases from. | `git merge upstream/v3` |
| `feat/clickhouse-collector` | Isolated **draft** of a ClickHouse metrics collector. Compilable, not yet upstream-ready. | `git merge upstream/v3` |

Upstream is tracked via the `upstream/v3` remote branch — there is no local
mirror branch. Custom commits live only on the fork branches, so periodic
`git merge upstream/v3` keeps conflicts confined to the fork-specific files.
This mirrors the standard long-lived-fork practice
([Atlassian](https://www.atlassian.com/git/tutorials/git-forks-and-upstreams),
[GitHub Docs](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/syncing-a-fork)).

## Remotes

```
origin    https://github.com/datacosmos-br/pmm.git   # the fork
upstream  https://github.com/percona/pmm.git          # Percona
```

## Syncing with upstream (periodic)

```bash
make dc-upstream                         # merge upstream/v3 head
make dc-upstream UPSTREAM_TARGET=stable  # merge the latest stable vX.Y.Z tag
```

## Releases

datacosmos releases are tagged with the **semver-`dcN`** scheme:
`v<upstream version>-dc<N>` — e.g. `v3.8.0-dc2`: the upstream PMM release the
build tracks (`3.8.0`), and `dcN` the datacosmos build counter on top of it
(`dc1`, `dc2`, …). The image/RPM version drops the leading `v` (`3.8.0-dc2`);
the git tag keeps it (`v3.8.0-dc2`).

For each upstream version, the datacosmos counter starts at `dc1` and increments
from the highest existing `v<upstream version>-dc<N>` tag. The release helper
fetches tags first, so it accounts for releases already pushed to `origin`.
When the upstream commit used by the branch is exactly a stable `vX.Y.Z` tag,
the release helper uses `vX.Y.Z-dc<N>`. Otherwise it uses
`v3-<upstream commit ISO date>-dc<N>`, with `N` incrementing for that date.

```bash
make dc-next                  # next tag, e.g. v3.8.0-dc1 or v3-2026-06-11-dc1
make dc-release               # creates and pushes that tag
```

The earlier date scheme `v3-<ISO date>-<upstream commit count>` still works if
such a tag is pushed manually — the workflow triggers on both `v*-dc*` and
`v3-*` tags — but `dcN` is the current scheme.

Pushing the tag runs `.github/workflows/datacosmos-release.yml`, which builds
linux/amd64 images and a GitHub Release. The build's internal `PMM_VERSION`
(written to `VERSION`, used for the upstream S3 RPM cache) is derived
separately from the nearest upstream semver tag — it stays a clean `X.Y.Z`.

## Building (datacosmos pipeline)

`Makefile.datacosmos` is included by the root `Makefile`, but only exposes
datacosmos-specific `dc-*` targets. Common targets such as `release`, `check`,
`clean`, and `gen` stay owned by upstream Makefiles.

```bash
make dc-build   # prepare + upstream client/server build + publish amd64 images + artifacts
make dc-clean   # remove only datacosmos external build/artifact dirs
```

### Upstream-aligned default

`Makefile.datacosmos` deliberately follows upstream's build strategy where the
fork can safely reuse public Percona infrastructure:

- `RPMBUILD_DOCKER_IMAGE` defaults to upstream's public
  `public.ecr.aws/e7j3v3n0/rpmbuild:3` image from `build/scripts/vars`.
- `SKIP_S3_CACHE` is unset by default, so amd64 server RPMs can reuse
  `s3://pmm-build-cache` anonymously and avoid rebuilding Grafana and other
  heavy components.
- The current datacosmos release workflow publishes linux/amd64 only.
- Images are built locally first and published to `ghcr.io/datacosmos-br` as
  part of `dc-build`.

### Local-only mode

For a fully local rebuild, set the overrides explicitly:

```bash
docker build --pull --tag pmm-rpmbuild:local --file build/docker/rpmbuild/Dockerfile.el9 build/docker/rpmbuild
RPMBUILD_DOCKER_IMAGE=pmm-rpmbuild:local SKIP_S3_CACHE=1 make dc-build
```

`make dc-build` materialises the build tree under `$(ROOT_DIR)` (default
`../pmm-build-root`, **outside** the repo).

### ⚠️ Build status — verified vs. pending

The default release path is the GitHub Actions workflow. Local full builds still
need Docker, submodules, and enough disk for PMM's upstream RPM/image pipeline.

## ClickHouse collector (draft — `feat/clickhouse-collector`)

`agent/agents/clickhouse/{collector.go,config.go}` is a `prometheus.Collector`
draft. It **compiles** (after `go get github.com/ClickHouse/clickhouse-go/v2`)
but is **not wired into pmm-agent** and is **not upstream-ready**. Before any
PR to `percona/pmm` it needs: integration into pmm-agent's exporter framework,
English comments, configuration via `pmm-agent.yaml`, and tests.
It is intentionally kept off `datacosmos/build` to avoid carrying the
`clickhouse-go/v2` dependency into the build before the feature is real.

## Agent coordination

This repo adopts the mandatory multi-agent coordination protocol —
see `.agents/skills/agent-coordination/SKILL.md` and the live ledger
`.agents/coordination/LEDGER.md`. Details in `AGENTS.md` §W-9.
