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

// Package clickhouseconn provides ClickHouse connection configuration and DSN building utilities.
package clickhouseconn

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
)

// Config holds ClickHouse connection configuration.
type Config struct {
	Protocol      string
	Host          string
	Port          uint16
	Database      string
	User          string
	Password      string
	TLS           bool
	TLSSkipVerify bool
	// TLSCa is the path to the TLS CA certificate file.
	TLSCa string
	// TLSCert is the path to the TLS client certificate file.
	TLSCert string
	// TLSKey is the path to the TLS client key file.
	TLSKey string
}

// ParseProtocol normalizes a raw protocol string.
// Supported values: "native", "http", "https", "" (empty defaults to native).
// Returns an error for unknown values.
func ParseProtocol(raw string) (string, error) {
	switch strings.TrimSpace(strings.ToLower(raw)) {
	case "", "native":
		return "native", nil
	case "http":
		return "http", nil
	case "https":
		return "https", nil
	default:
		return "", fmt.Errorf("unknown clickhouse protocol: %q", raw)
	}
}

// Scheme returns the URI scheme for the configured protocol.
func (c *Config) Scheme() string {
	switch c.Protocol {
	case "http":
		return "http"
	case "https":
		return "https"
	default:
		return "clickhouse"
	}
}

// ExporterScheme returns the scheme used by exporters.
// It returns "https" if Protocol is "https" or if Protocol is empty and TLS is enabled, else "http".
func (c *Config) ExporterScheme() string {
	if c.Protocol == "https" || (c.Protocol == "" && c.TLS) {
		return "https"
	}
	return "http"
}

// IsSecure returns true when the connection should use TLS.
func (c *Config) IsSecure() bool {
	return c.Protocol == "https" || c.TLS
}

// effectivePort returns the configured Port if non-zero,
// otherwise the default port for the configured protocol.
func (c *Config) effectivePort() uint16 {
	if c.Port != 0 {
		return c.Port
	}
	switch c.Protocol {
	case "http":
		return 8123
	case "https":
		return 8443
	default:
		return 9000
	}
}

// DSN builds a ClickHouse DSN string from the configuration.
func (c *Config) DSN() (string, error) {
	if c.Host == "" {
		return "", fmt.Errorf("clickhouse host is required")
	}

	u := &url.URL{
		Scheme: c.Scheme(),
		Host:   net.JoinHostPort(c.Host, strconv.Itoa(int(c.effectivePort()))),
	}
	if c.Database != "" {
		u.Path = "/" + c.Database
	}

	if c.User != "" || c.Password != "" {
		if c.Password != "" {
			u.User = url.UserPassword(c.User, c.Password)
		} else {
			u.User = url.User(c.User)
		}
	}

	q := url.Values{}
	if c.IsSecure() {
		q.Set("secure", "true")
	}
	if c.TLSSkipVerify {
		q.Set("skip_verify", "true")
	}
	if c.TLSCa != "" {
		q.Set("sslrootcert", c.TLSCa)
	}
	if c.TLSCert != "" {
		q.Set("sslcert", c.TLSCert)
	}
	if c.TLSKey != "" {
		q.Set("sslkey", c.TLSKey)
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	return u.String(), nil
}
