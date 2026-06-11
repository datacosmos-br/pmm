// Copyright (C) 2023 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clickhouseconn

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseProtocol(t *testing.T) {
	tests := []struct {
		name        string
		raw         string
		want        string
		wantErr     bool
		errContains string
	}{
		{
			name: "empty defaults to native",
			raw:  "",
			want: "native",
		},
		{
			name: "native",
			raw:  "native",
			want: "native",
		},
		{
			name: "http",
			raw:  "http",
			want: "http",
		},
		{
			name: "https",
			raw:  "https",
			want: "https",
		},
		{
			name: "mixed case native",
			raw:  "Native",
			want: "native",
		},
		{
			name: "mixed case HTTP",
			raw:  "HTTP",
			want: "http",
		},
		{
			name: "mixed case Https",
			raw:  "Https",
			want: "https",
		},
		{
			name: "whitespace around native",
			raw:  "  native  ",
			want: "native",
		},
		{
			name: "whitespace around http",
			raw:  "\thttp\n",
			want: "http",
		},
		{
			name:        "invalid protocol",
			raw:         "tcp",
			wantErr:     true,
			errContains: "unknown clickhouse protocol",
		},
		{
			name:    "whitespace only defaults to native",
			raw:     "   ",
			wantErr: false,
			want:    "native",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseProtocol(tt.raw)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConfig_Scheme(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		want string
	}{
		{name: "native", cfg: Config{Protocol: "native"}, want: "clickhouse"},
		{name: "empty", cfg: Config{Protocol: ""}, want: "clickhouse"},
		{name: "http", cfg: Config{Protocol: "http"}, want: "http"},
		{name: "https", cfg: Config{Protocol: "https"}, want: "https"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.cfg.Scheme())
		})
	}
}

func TestConfig_ExporterScheme(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		want string
	}{
		{name: "https protocol", cfg: Config{Protocol: "https"}, want: "https"},
		{name: "native with TLS", cfg: Config{Protocol: "", TLS: true}, want: "https"},
		{name: "native without TLS", cfg: Config{Protocol: "", TLS: false}, want: "http"},
		{name: "http", cfg: Config{Protocol: "http"}, want: "http"},
		{name: "native explicit", cfg: Config{Protocol: "native", TLS: false}, want: "http"},
		{name: "native explicit with TLS", cfg: Config{Protocol: "native", TLS: true}, want: "http"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.cfg.ExporterScheme())
		})
	}
}

func TestConfig_IsSecure(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		want bool
	}{
		{name: "https protocol", cfg: Config{Protocol: "https"}, want: true},
		{name: "TLS enabled", cfg: Config{Protocol: "native", TLS: true}, want: true},
		{name: "http with TLS", cfg: Config{Protocol: "http", TLS: true}, want: true},
		{name: "native without TLS", cfg: Config{Protocol: "native", TLS: false}, want: false},
		{name: "http without TLS", cfg: Config{Protocol: "http", TLS: false}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.cfg.IsSecure())
		})
	}
}

func TestConfig_effectivePort(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		want uint16
	}{
		{name: "explicit port", cfg: Config{Protocol: "native", Port: 1234}, want: 1234},
		{name: "empty protocol explicit port", cfg: Config{Protocol: "", Port: 1234}, want: 1234},
		{name: "native default", cfg: Config{Protocol: "native"}, want: 9000},
		{name: "empty protocol default", cfg: Config{Protocol: ""}, want: 9000},
		{name: "http default", cfg: Config{Protocol: "http"}, want: 8123},
		{name: "https default", cfg: Config{Protocol: "https"}, want: 8443},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.cfg.effectivePort())
		})
	}
}

func TestConfig_DSN(t *testing.T) {
	tests := []struct {
		name        string
		cfg         Config
		want        string
		wantErr     bool
		errContains string
	}{
		{
			name: "native default",
			cfg: Config{
				Protocol: "native",
				Host:     "127.0.0.1",
				Port:     9000,
				Database: "pmm",
				User:     "default",
				Password: "clickhouse",
			},
			want: "clickhouse://default:clickhouse@127.0.0.1:9000/pmm",
		},
		{
			name: "native with TLS",
			cfg: Config{
				Protocol:      "native",
				Host:          "127.0.0.1",
				Port:          9440,
				Database:      "pmm",
				User:          "default",
				Password:      "clickhouse",
				TLS:           true,
				TLSSkipVerify: true,
			},
			want: "clickhouse://default:clickhouse@127.0.0.1:9440/pmm?secure=true&skip_verify=true",
		},
		{
			name: "http no tls",
			cfg: Config{
				Protocol: "http",
				Host:     "127.0.0.1",
				Port:     8123,
				Database: "pmm",
				User:     "default",
			},
			want: "http://default@127.0.0.1:8123/pmm",
		},
		{ //nolint:gosec // deterministic DSN fixture, not real credentials.
			name: "https with ca",
			cfg: Config{
				Protocol: "https",
				Host:     "ch.example.com",
				Port:     8443,
				Database: "metrics",
				User:     "admin",
				Password: "secret",
				TLSCa:    "/etc/ssl/ca.crt",
			},
			want: "https://admin:secret@ch.example.com:8443/metrics?secure=true&sslrootcert=%2Fetc%2Fssl%2Fca.crt",
		},
		{
			name: "empty protocol defaults to native",
			cfg: Config{
				Protocol: "",
				Host:     "127.0.0.1",
				Port:     0,
				Database: "pmm",
			},
			want: "clickhouse://127.0.0.1:9000/pmm",
		},
		{
			name: "zero port uses default for https",
			cfg: Config{
				Protocol: "https",
				Host:     "127.0.0.1",
				Port:     0,
				Database: "pmm",
			},
			want: "https://127.0.0.1:8443/pmm?secure=true",
		},
		{
			name: "empty host errors",
			cfg: Config{
				Protocol: "",
				Host:     "",
				Port:     0,
				Database: "",
			},
			wantErr:     true,
			errContains: "clickhouse host is required",
		},
		{ //nolint:gosec // deterministic DSN fixture, not real credentials.
			name: "native tls with ca",
			cfg: Config{
				Protocol: "native",
				Host:     "127.0.0.1",
				Port:     9000,
				Database: "pmm",
				User:     "u",
				Password: "p",
				TLS:      true,
				TLSCa:    "/ca",
			},
			want: "clickhouse://u:p@127.0.0.1:9000/pmm?secure=true&sslrootcert=%2Fca",
		},
		{
			name: "empty user with password",
			cfg: Config{
				Protocol: "native",
				Host:     "127.0.0.1",
				Port:     9000,
				Database: "pmm",
				User:     "",
				Password: "secret",
			},
			want: "clickhouse://:secret@127.0.0.1:9000/pmm",
		},
		{ //nolint:gosec // deterministic DSN fixture, not real credentials.
			name: "native tls with cert and key",
			cfg: Config{
				Protocol: "native",
				Host:     "127.0.0.1",
				Port:     9000,
				Database: "pmm",
				User:     "u",
				Password: "p",
				TLS:      true,
				TLSCa:    "/ca",
				TLSCert:  "/cert",
				TLSKey:   "/key",
			},
			want: "clickhouse://u:p@127.0.0.1:9000/pmm?secure=true&sslcert=%2Fcert&sslkey=%2Fkey&sslrootcert=%2Fca",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cfg.DSN()
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
