// Copyright (C) 2023 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package serviceinfobroker

import (
	"context"
	"database/sql"
	"math"

	_ "github.com/ClickHouse/clickhouse-go/v2" // register clickhouse driver
	agentv1 "github.com/percona/pmm/api/agent/v1"
)

func (sib *ServiceInfoBroker) getClickHouseInfo(ctx context.Context, dsn string) *agentv1.ServiceInfoResponse {
	var res agentv1.ServiceInfoResponse

	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		sib.l.Debugf("getClickHouseInfo: failed to open connection: %s", err)
		res.Error = err.Error()
		return &res
	}
	defer db.Close() //nolint:errcheck

	var tableCount uint64
	if err = db.QueryRowContext(ctx, "SELECT count() FROM system.tables").Scan(&tableCount); err != nil {
		res.Error = err.Error()
		return &res
	}
	res.TableCount = int32(tableCount)
	if tableCount > math.MaxInt32 {
		res.TableCount = math.MaxInt32
	}

	var version string
	if err = db.QueryRowContext(ctx, "SELECT version()").Scan(&version); err != nil {
		res.Error = err.Error()
		return &res
	}
	res.Version = version

	return &res
}
