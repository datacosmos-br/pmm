package migrations

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGetEngine pins the cluster-vs-single MergeTree engine selection that drives
// every QAN table's DDL. A wrong engine on a replicated ClickHouse cluster silently
// breaks replication (data lands on one shard only) — the exact failure this client
// must prevent, so it is asserted explicitly.
func TestGetEngine(t *testing.T) {
	t.Parallel()
	require.Equal(t, "ReplicatedMergeTree", GetEngine(true), "cluster topology must use the replicated engine")
	require.Equal(t, "MergeTree", GetEngine(false), "single-node topology must use the plain engine")
}

// TestAddSchemaMigrationsParams verifies the golang-migrate schema_migrations table
// is forced onto the replicated engine for clusters (so the migration bookkeeping
// table itself replicates), while pre-existing DSN query parameters are preserved.
func TestAddSchemaMigrationsParams(t *testing.T) {
	t.Parallel()

	t.Run("forces replicated engine and preserves existing params", func(t *testing.T) {
		t.Parallel()
		out, err := addSchemaMigrationsParams("clickhouse://default:pw@ch-host:9000/pmm?dial_timeout=10s")
		require.NoError(t, err)

		u, err := url.Parse(out)
		require.NoError(t, err)
		q := u.Query()
		require.Equal(t, "ReplicatedMergeTree ORDER BY version", q.Get("x-migrations-table-engine"))
		require.Equal(t, "10s", q.Get("dial_timeout"), "existing query params must survive")
		require.Equal(t, "/pmm", u.Path, "database path must be preserved")
	})

	t.Run("rejects an unparseable dsn", func(t *testing.T) {
		t.Parallel()
		_, err := addSchemaMigrationsParams("://not-a-valid-dsn")
		require.Error(t, err)
	})
}
