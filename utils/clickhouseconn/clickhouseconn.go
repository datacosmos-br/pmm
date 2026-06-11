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
	"errors"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
)

const (
	protocolNative  = "native"
	protocolHTTP    = "http"
	protocolHTTPS   = "https"
	schemeNative    = "clickhouse"
	defaultHTTPPort = 8123
	defaultTLSPort  = 8443
	defaultTCPPort  = 9000
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
	case "", protocolNative:
		return protocolNative, nil
	case protocolHTTP:
		return protocolHTTP, nil
	case protocolHTTPS:
		return protocolHTTPS, nil
	default:
		return "", fmt.Errorf("unknown clickhouse protocol: %q", raw)
	}
}

// Scheme returns the URI scheme for the configured protocol.
func (c *Config) Scheme() string {
	switch c.Protocol {
	case protocolHTTP:
		return protocolHTTP
	case protocolHTTPS:
		return protocolHTTPS
	default:
		return schemeNative
	}
}

// ExporterScheme returns the scheme used by exporters.
// It returns "https" if Protocol is "https" or if Protocol is empty and TLS is enabled, else "http".
func (c *Config) ExporterScheme() string {
	if c.Protocol == protocolHTTPS || (c.Protocol == "" && c.TLS) {
		return protocolHTTPS
	}
	return protocolHTTP
}

// IsSecure returns true when the connection should use TLS.
func (c *Config) IsSecure() bool {
	return c.Protocol == protocolHTTPS || c.TLS
}

// DSN builds a ClickHouse DSN string from the configuration.
func (c *Config) DSN() (string, error) {
	if c.Host == "" {
		return "", errors.New("clickhouse host is required")
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

// effectivePort returns the configured Port if non-zero,
// otherwise the default port for the configured protocol.
func (c *Config) effectivePort() uint16 {
	if c.Port != 0 {
		return c.Port
	}
	switch c.Protocol {
	case protocolHTTP:
		return defaultHTTPPort
	case protocolHTTPS:
		return defaultTLSPort
	default:
		return defaultTCPPort
	}
}
