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

package connectionchecker

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/percona/pmm/agent/config"
	"github.com/percona/pmm/agent/utils/tests"
	agentv1 "github.com/percona/pmm/api/agent/v1"
	inventoryv1 "github.com/percona/pmm/api/inventory/v1"
)

func mysqlConnectionDSN(tb testing.TB, user, password, timeout string) string {
	tb.Helper()

	return fmt.Sprintf("%s:%s@tcp(%s)/?clientFoundRows=true&parseTime=true&timeout=%s", user, password, tests.ServiceAddr(tb, 3306), timeout)
}

func mongoConnectionDSN(tb testing.TB, credentials string, port int, database string, timeoutMS int) string {
	tb.Helper()

	auth := ""
	if credentials != "" {
		auth = credentials + "@"
	}
	if database == "" {
		return fmt.Sprintf("mongodb://%s%s?connectTimeoutMS=%d", auth, tests.ServiceAddr(tb, port), timeoutMS)
	}

	return fmt.Sprintf("mongodb://%s%s/%s?connectTimeoutMS=%d", auth, tests.ServiceAddr(tb, port), database, timeoutMS)
}

func postgresqlConnectionDSN(tb testing.TB, password, timeout string) string {
	tb.Helper()

	return fmt.Sprintf("postgres://pmm-agent:%s@%s/postgres?connect_timeout=%s&sslmode=disable", password, tests.ServiceAddr(tb, 5432), timeout)
}

func valkeyConnectionDSN(tb testing.TB, password string) string {
	tb.Helper()

	return fmt.Sprintf("redis://default:%s@%s", password, tests.ServiceAddr(tb, 6379))
}

func TestConnectionChecker(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		req         *agentv1.CheckConnectionRequest
		expectedErr string
		panic       bool
	}{
		{
			name: "MySQL",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     mysqlConnectionDSN(t, "root", "root-password", "1s"),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_MYSQL_SERVICE,
				Timeout: durationpb.New(3 * time.Second),
			},
		},
		{
			name: "MySQL wrong params",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     mysqlConnectionDSN(t, "pmm-agent", "pmm-agent-wrong-password", "1s"),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_MYSQL_SERVICE,
				Timeout: durationpb.New(3 * time.Second),
			},
			expectedErr: `Error 1045 \(28000\): Access denied for user 'pmm-agent'@'.+' \(using password: YES\)`,
		},
		{
			name: "MySQL timeout",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     mysqlConnectionDSN(t, "root", "root-password", "10s"),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_MYSQL_SERVICE,
				Timeout: durationpb.New(time.Nanosecond),
			},
			expectedErr: `context deadline exceeded`,
		},

		{
			name: "MongoDB with no auth",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     mongoConnectionDSN(t, "", 27019, "admin", 1000),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_MONGODB_SERVICE,
				Timeout: durationpb.New(3 * time.Second),
			},
		},
		{
			name: "MongoDB with no auth with params",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     mongoConnectionDSN(t, "root:root-password", 27019, "admin", 1000),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_MONGODB_SERVICE,
				Timeout: durationpb.New(3 * time.Second),
			},
			expectedErr: `.*auth error: (sasl conversation error: )?unable to authenticate using mechanism "[\w-]+": ` +
				`\(AuthenticationFailed\) Authentication failed.`,
		},
		{
			name: "MongoDB",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     mongoConnectionDSN(t, "root:root-password", 27017, "admin", 1000),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_MONGODB_SERVICE,
				Timeout: durationpb.New(3 * time.Second),
			},
		},
		{
			name: "MongoDB no params",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     mongoConnectionDSN(t, "", 27017, "admin", 1000),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_MONGODB_SERVICE,
				Timeout: durationpb.New(3 * time.Second),
			},
			expectedErr: `\(Unauthorized\) (?:command getDiagnosticData requires authentication|` +
				`Command buildInfo requires authentication|` +
				`there are no users authenticated|` +
				`not authorized on admin to execute command \{ getDiagnosticData\: 1 \})`,
		},
		{
			name: "MongoDB wrong params",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     mongoConnectionDSN(t, "root:root-password-wrong", 27017, "admin", 1000),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_MONGODB_SERVICE,
				Timeout: durationpb.New(3 * time.Second),
			},
			expectedErr: `.*auth error: (sasl conversation error: )?unable to authenticate using mechanism "[\w-]+": ` +
				`\(AuthenticationFailed\) Authentication failed.`,
		},
		{
			name: "MongoDB timeout",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     mongoConnectionDSN(t, "root:root-password", 27017, "admin", 10000),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_MONGODB_SERVICE,
				Timeout: durationpb.New(time.Nanosecond),
			},
			expectedErr: `.*context deadline exceeded.*`,
		},
		{
			name: "MongoDB no database",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     mongoConnectionDSN(t, "root:root-password", 27017, "", 1000),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_MONGODB_SERVICE,
				Timeout: durationpb.New(3 * time.Second),
			},
			expectedErr: `error parsing uri: must have a / before the query \?`,
		},

		{
			name: "PostgreSQL",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     postgresqlConnectionDSN(t, "pmm-agent-password", "1"),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_POSTGRESQL_SERVICE,
				Timeout: durationpb.New(3 * time.Second),
			},
		},
		{
			name: "PostgreSQL wrong params",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     postgresqlConnectionDSN(t, "pmm-agent-wrong-password", "1"),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_POSTGRESQL_SERVICE,
				Timeout: durationpb.New(3 * time.Second),
			},
			expectedErr: `pq: password authentication failed for user "pmm-agent"( \(\w+\))?`,
		},
		{
			name: "PostgreSQL timeout",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     postgresqlConnectionDSN(t, "pmm-agent-password", "10"),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_POSTGRESQL_SERVICE,
				Timeout: durationpb.New(time.Nanosecond),
			},
			expectedErr: `context deadline exceeded`,
		},
		{
			name: "Valkey",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     valkeyConnectionDSN(t, "pmm-agent_password"),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_VALKEY_SERVICE,
				Timeout: durationpb.New(3 * time.Second),
			},
		},
		{
			name: "Valkey wrong params",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     valkeyConnectionDSN(t, "pmm-agent_wrong_password"),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_VALKEY_SERVICE,
				Timeout: durationpb.New(3 * time.Second),
			},
			expectedErr: `WRONGPASS invalid username-password pair or user is disabled.`,
		},
		{
			name: "Valkey timeout",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     valkeyConnectionDSN(t, "pmm-agent_password"),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_VALKEY_SERVICE,
				Timeout: durationpb.New(time.Nanosecond),
			},
			expectedErr: `(?:dial tcp .+:6379: i/o timeout|dial tcp: lookup .+: i/o timeout|context deadline exceeded)`,
		},

		// Use MySQL for ProxySQL tests for now.
		// TODO https://jira.percona.com/browse/PMM-4930
		{
			name: "ProxySQL/MySQL",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     mysqlConnectionDSN(t, "root", "root-password", "1s"),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_PROXYSQL_SERVICE,
				Timeout: durationpb.New(3 * time.Second),
			},
		},
		{
			name: "ProxySQL/MySQL wrong params",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     mysqlConnectionDSN(t, "pmm-agent", "pmm-agent-wrong-password", "1s"),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_PROXYSQL_SERVICE,
				Timeout: durationpb.New(3 * time.Second),
			},
			expectedErr: `Error 1045 \(28000\): Access denied for user 'pmm-agent'@'.+' \(using password: YES\)`,
		},
		{
			name: "ProxySQL/MySQL timeout",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     mysqlConnectionDSN(t, "root", "root-password", "10s"),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_PROXYSQL_SERVICE,
				Timeout: durationpb.New(time.Nanosecond),
			},
			expectedErr: `context deadline exceeded`,
		},
		{
			name: "Invalid service type",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     mysqlConnectionDSN(t, "root", "root-password", "10s"),
				Type:    inventoryv1.ServiceType_SERVICE_TYPE_UNSPECIFIED,
				Timeout: durationpb.New(time.Nanosecond),
			},
			expectedErr: `unknown service type: SERVICE_TYPE_UNSPECIFIED`,
			panic:       true,
		},
		{
			name: "Unknown service type",
			req: &agentv1.CheckConnectionRequest{
				Dsn:     mysqlConnectionDSN(t, "root", "root-password", "10s"),
				Type:    inventoryv1.ServiceType(12345),
				Timeout: durationpb.New(time.Nanosecond),
			},
			expectedErr: `unknown service type: 12345`,
			panic:       true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cfgStorage := config.NewStorage(&config.Config{
				Paths: config.Paths{TempDir: t.TempDir()},
			})
			c := New(cfgStorage)

			if tt.panic {
				require.PanicsWithValue(t, tt.expectedErr, func() {
					c.Check(context.Background(), tt.req, 0)
				})
				return
			}

			resp := c.Check(context.Background(), tt.req, 0)
			require.NotNil(t, resp)
			if tt.expectedErr == "" {
				assert.Empty(t, resp.Error)
			} else {
				require.NotEmpty(t, resp.Error)
				assert.Regexp(t, `(?i)^`+tt.expectedErr+`$`, resp.Error)
			}
		})
	}

	t.Run("Stats should be empty", func(t *testing.T) {
		cfgStorage := config.NewStorage(&config.Config{
			Paths: config.Paths{TempDir: t.TempDir()},
		})
		c := New(cfgStorage)
		resp := c.Check(context.Background(), &agentv1.CheckConnectionRequest{
			Dsn:  mysqlConnectionDSN(t, "root", "root-password", "1s"),
			Type: inventoryv1.ServiceType_SERVICE_TYPE_MYSQL_SERVICE,
		}, 0)
		require.NotNil(t, resp)
	})

	t.Run("MongoDBWithSSL", func(t *testing.T) {
		mongoDBDSNWithSSL, mongoDBTextFiles := tests.GetTestMongoDBWithSSLDSN(t, "../")

		cfgStorage := config.NewStorage(&config.Config{
			Paths: config.Paths{TempDir: t.TempDir()},
		})

		c := New(cfgStorage)
		resp := c.Check(context.Background(), &agentv1.CheckConnectionRequest{
			Dsn:       mongoDBDSNWithSSL,
			Type:      inventoryv1.ServiceType_SERVICE_TYPE_MONGODB_SERVICE,
			Timeout:   durationpb.New(30 * time.Second),
			TextFiles: mongoDBTextFiles,
		}, rand.Uint32()) //nolint:gosec
		require.NotNil(t, resp)
		assert.Empty(t, resp.Error)
	})
}
