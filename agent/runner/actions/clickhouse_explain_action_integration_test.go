// Copyright (C) 2023 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package actions

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"
	"strings"
	"testing"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2" // register database/sql driver "clickhouse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/percona/pmm/agent/utils/tests"
	agentv1 "github.com/percona/pmm/api/agent/v1"
)

// clickHouseExplainDSN returns the DSN to exercise the EXPLAIN action against,
// from CLICKHOUSE_TEST_ENDPOINTS ("name=dsn" pairs) — the same env the collector
// matrix uses — falling back to the shared integration service host.
func clickHouseExplainDSN(t *testing.T) string {
	t.Helper()

	raw := strings.TrimSpace(os.Getenv("CLICKHOUSE_TEST_ENDPOINTS"))
	if raw == "" {
		return "clickhouse://default:clickhouse@" + tests.ServiceAddr(t, 9000) + "/default"
	}
	first := strings.TrimSpace(strings.Split(raw, ",")[0])
	if _, dsn, ok := strings.Cut(first, "="); ok {
		require.NotEmpty(t, strings.TrimSpace(dsn), "CLICKHOUSE_TEST_ENDPOINTS contained an empty DSN in %q", raw)
		return strings.TrimSpace(dsn)
	}
	require.NotEmpty(t, first, "CLICKHOUSE_TEST_ENDPOINTS contained an empty endpoint")
	return strings.TrimSpace(first)
}

// TestClickHouseExplainActionIntegration exercises the action's live Run against
// a real ClickHouse: it covers the generic multi-column scan that the unit tests
// cannot reach — single-column output (PLAN/AST) AND the 5-column ESTIMATE output
// (database, table, parts, rows, marks) that the QAN "Explain" panel consumes.
func TestClickHouseExplainActionIntegration(t *testing.T) {
	dsn := clickHouseExplainDSN(t)

	db, err := sql.Open("clickhouse", dsn)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, db.Close())
	})

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	require.NoError(t, err, "ClickHouse endpoint %s must be reachable", dsn)

	// A MergeTree table is required for EXPLAIN ESTIMATE (it reports per-part
	// row/mark estimates). Created+dropped per run so the test is idempotent.
	const tbl = "pmm_explain_action_it"
	_, err = db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS "+tbl+" (id UInt64) ENGINE = MergeTree ORDER BY id")
	require.NoError(t, err)
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cleanupCancel()
		_, cleanupErr := db.ExecContext(cleanupCtx, "DROP TABLE IF EXISTS "+tbl)
		require.NoError(t, cleanupErr)
	})
	_, err = db.ExecContext(ctx, "INSERT INTO "+tbl+" SELECT number FROM numbers(10)")
	require.NoError(t, err)

	run := func(t *testing.T, query, explainType string) map[string]any {
		t.Helper()
		action, err := NewClickHouseExplainAction("it-"+explainType, 30*time.Second,
			&agentv1.StartActionRequest_ClickHouseExplainParams{Dsn: dsn, Query: query, ExplainType: explainType}, "")
		require.NoError(t, err)

		out, err := action.Run(ctx)
		require.NoError(t, err, "EXPLAIN %s must execute against live ClickHouse", explainType)

		var result map[string]any
		require.NoError(t, json.Unmarshal(out, &result))
		require.Equal(t, query, result["explained_query"])
		require.Equal(t, explainType, result["explain_type"])
		lines, ok := result["explain_result"].([]any)
		require.True(t, ok, "explain_result must be a list of lines")
		require.NotEmpty(t, lines, "EXPLAIN %s must return at least one line", explainType)
		return result
	}

	t.Run("PLAN single-column", func(t *testing.T) {
		run(t, "SELECT count() FROM "+tbl, "PLAN")
	})

	t.Run("default type (PLAN)", func(t *testing.T) {
		run(t, "SELECT count() FROM "+tbl, "")
	})

	t.Run("AST", func(t *testing.T) {
		run(t, "SELECT 1", "AST")
	})

	t.Run("ESTIMATE multi-column", func(t *testing.T) {
		result := run(t, "SELECT count() FROM "+tbl, "ESTIMATE")
		lines := result["explain_result"].([]any)
		// ESTIMATE is multi-column -> Run prepends a tab-joined header row.
		header, _ := lines[0].(string)
		assert.Contains(t, header, "\t", "ESTIMATE output must carry a multi-column (tab-joined) header")
		assert.Contains(t, strings.ToLower(header), "rows", "ESTIMATE header must include the rows estimate column")
	})
}
