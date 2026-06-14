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
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2" // register database/sql driver "clickhouse"

	agentv1 "github.com/percona/pmm/api/agent/v1"
)

const clickHouseExplainActionType = "clickhouse-explain"

// allowedClickHouseExplainTypes is the EXPLAIN type whitelist (matches the
// api/agent/v1 proto contract). The type is interpolated into the statement,
// so anything outside this set is rejected to keep injection impossible.
var allowedClickHouseExplainTypes = map[string]struct{}{
	"":         {}, // ClickHouse default (PLAN)
	"PLAN":     {},
	"PIPELINE": {},
	"AST":      {},
	"ESTIMATE": {},
	"SYNTAX":   {},
}

type clickHouseExplainAction struct {
	id      string
	timeout time.Duration
	params  *agentv1.StartActionRequest_ClickHouseExplainParams
}

// NewClickHouseExplainAction creates a ClickHouse EXPLAIN Action.
func NewClickHouseExplainAction(id string, timeout time.Duration, params *agentv1.StartActionRequest_ClickHouseExplainParams, _ string) (Action, error) {
	if params.Query == "" {
		return nil, errors.New("query to EXPLAIN is empty")
	}

	if strings.HasSuffix(params.Query, "...") {
		return nil, errors.New(
			"explain failed because the query exceeded max length and got trimmed; set max-query-length to a larger value",
		)
	}

	explainType := strings.ToUpper(strings.TrimSpace(params.ExplainType))
	if _, ok := allowedClickHouseExplainTypes[explainType]; !ok {
		return nil, fmt.Errorf("invalid EXPLAIN type %q; allowed: PLAN, PIPELINE, AST, ESTIMATE, SYNTAX", params.ExplainType)
	}

	return &clickHouseExplainAction{
		id:      id,
		timeout: timeout,
		params:  params,
	}, nil
}

// ID returns an Action ID.
func (a *clickHouseExplainAction) ID() string {
	return a.id
}

// Timeout returns Action timeout.
func (a *clickHouseExplainAction) Timeout() time.Duration {
	return a.timeout
}

// Type returns an Action type.
func (a *clickHouseExplainAction) Type() string {
	return clickHouseExplainActionType
}

// DSN returns a DSN for the Action.
func (a *clickHouseExplainAction) DSN() string {
	return a.params.Dsn
}

// Run runs an Action and returns output and error.
func (a *clickHouseExplainAction) Run(ctx context.Context) ([]byte, error) {
	db, err := sql.Open("clickhouse", a.params.Dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open ClickHouse connection: %w", err)
	}
	defer db.Close() //nolint:errcheck

	rows, err := db.QueryContext(ctx, a.buildExplainQuery())
	if err != nil {
		return nil, fmt.Errorf("failed to execute EXPLAIN: %w", err)
	}
	defer rows.Close() //nolint:errcheck

	// EXPLAIN output is single-column for PLAN/PIPELINE/AST/SYNTAX but
	// multi-column for ESTIMATE (database, table, parts, rows, marks) —
	// scan generically and join columns with a tab.
	cols, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to read EXPLAIN columns: %w", err)
	}

	var lines []string
	if len(cols) > 1 {
		lines = append(lines, strings.Join(cols, "\t"))
	}
	values := make([]sql.RawBytes, len(cols))
	scanArgs := make([]any, len(cols))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			return nil, fmt.Errorf("failed to scan EXPLAIN result: %w", err)
		}
		fields := make([]string, len(values))
		for i, v := range values {
			fields[i] = string(v)
		}
		lines = append(lines, strings.Join(fields, "\t"))
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("error reading EXPLAIN result: %w", err)
	}

	result := map[string]any{
		"explain_result":  lines,
		"explained_query": a.params.Query,
		"explain_type":    a.params.ExplainType,
	}

	return json.Marshal(result)
}

func (a *clickHouseExplainAction) buildExplainQuery() string {
	query := strings.TrimSpace(a.params.Query)
	explainType := strings.ToUpper(strings.TrimSpace(a.params.ExplainType))

	if explainType != "" {
		return fmt.Sprintf("EXPLAIN %s %s", explainType, query)
	}
	return "EXPLAIN " + query
}

func (a *clickHouseExplainAction) sealed() {}
