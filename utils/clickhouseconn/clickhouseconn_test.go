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
			name:        "whitespace only defaults to native",
			raw:         "   ",
			wantErr:     false,
			want:        "native",
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
