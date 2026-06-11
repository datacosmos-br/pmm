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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	agentv1 "github.com/percona/pmm/api/agent/v1"
)

func TestClickHouseExplainAction(t *testing.T) {
	t.Parallel()

	t.Run("EmptyQueryRejected", func(t *testing.T) {
		t.Parallel()
		_, err := NewClickHouseExplainAction("id", time.Second, &agentv1.StartActionRequest_ClickHouseExplainParams{
			Query: "",
		}, "")
		require.Error(t, err)
	})

	t.Run("TrimmedQueryRejected", func(t *testing.T) {
		t.Parallel()
		_, err := NewClickHouseExplainAction("id", time.Second, &agentv1.StartActionRequest_ClickHouseExplainParams{
			Query: "SELECT * FROM system.numbers WHERE...",
		}, "")
		require.Error(t, err)
	})

	t.Run("ExplainTypeWhitelist", func(t *testing.T) {
		t.Parallel()
		for _, tt := range []struct {
			explainType string
			valid       bool
		}{
			{"", true},
			{"PLAN", true},
			{"plan", true},
			{" Pipeline ", true},
			{"AST", true},
			{"ESTIMATE", true},
			{"SYNTAX", true},
			{"PLAN; DROP TABLE x; --", false},
			{"QUERY TREE", false},
			{"bogus", false},
		} {
			_, err := NewClickHouseExplainAction("id", time.Second, &agentv1.StartActionRequest_ClickHouseExplainParams{
				Query:       "SELECT 1",
				ExplainType: tt.explainType,
			}, "")
			if tt.valid {
				assert.NoError(t, err, "explainType %q should be accepted", tt.explainType)
			} else {
				assert.Error(t, err, "explainType %q must be rejected", tt.explainType)
			}
		}
	})

	t.Run("BuildExplainQuery", func(t *testing.T) {
		t.Parallel()
		a, err := NewClickHouseExplainAction("id", time.Second, &agentv1.StartActionRequest_ClickHouseExplainParams{
			Query:       "SELECT 1",
			ExplainType: "pipeline",
		}, "")
		require.NoError(t, err)
		assert.Equal(t, "EXPLAIN PIPELINE SELECT 1", a.(*clickHouseExplainAction).buildExplainQuery())

		a, err = NewClickHouseExplainAction("id", time.Second, &agentv1.StartActionRequest_ClickHouseExplainParams{
			Query: "  SELECT 1  ",
		}, "")
		require.NoError(t, err)
		assert.Equal(t, "EXPLAIN SELECT 1", a.(*clickHouseExplainAction).buildExplainQuery())
	})

	t.Run("MetadataAndDSN", func(t *testing.T) {
		t.Parallel()
		a, err := NewClickHouseExplainAction("test-id", 42*time.Second, &agentv1.StartActionRequest_ClickHouseExplainParams{
			Dsn:   "clickhouse://user:pass@host:9000/db",
			Query: "SELECT 1",
		}, "")
		require.NoError(t, err)
		assert.Equal(t, "test-id", a.ID())
		assert.Equal(t, 42*time.Second, a.Timeout())
		assert.Equal(t, "clickhouse-explain", a.Type())
		assert.Equal(t, "clickhouse://user:pass@host:9000/db", a.DSN())
	})
}
