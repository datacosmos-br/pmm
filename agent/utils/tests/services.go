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

package tests

import (
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	defaultServiceHost = "127.0.0.1"

	mysqlServicePort              = 3306
	mongoServicePort              = 27017
	mongoTLSServicePort           = 27018
	mongoReplicaFirstServicePort  = 27020
	mongoReplicaSecondServicePort = 27021
	mongoReplicaTLSFirstPort      = 27022
	mongoReplicaTLSSecondPort     = 27023
)

// ServiceHost returns the host where Docker-published integration services are reachable.
func ServiceHost(tb testing.TB) string {
	tb.Helper()

	if host := strings.TrimSpace(os.Getenv("PMM_TEST_SERVICE_HOST")); host != "" {
		return host
	}

	dockerHost := strings.TrimSpace(os.Getenv("DOCKER_HOST"))
	if dockerHost == "" || strings.HasPrefix(dockerHost, "unix://") || strings.HasPrefix(dockerHost, "npipe://") {
		return defaultServiceHost
	}

	u, err := url.Parse(dockerHost)
	require.NoErrorf(tb, err, "failed to parse DOCKER_HOST=%q; set PMM_TEST_SERVICE_HOST", dockerHost)
	if u.Scheme != "tcp" {
		require.FailNowf(
			tb,
			"unsupported DOCKER_HOST for integration service discovery",
			"DOCKER_HOST=%q must be empty, unix://, npipe://, or tcp://host:port; set PMM_TEST_SERVICE_HOST",
			dockerHost,
		)
	}

	host := u.Hostname()
	require.NotEmptyf(tb, host, "DOCKER_HOST=%q did not include a host; set PMM_TEST_SERVICE_HOST", dockerHost)

	return host
}

// ServiceAddr returns a host:port address for an integration service.
func ServiceAddr(tb testing.TB, port int) string {
	tb.Helper()

	require.Positive(tb, port, "integration service port must be positive")

	return net.JoinHostPort(ServiceHost(tb), strconv.Itoa(port))
}
