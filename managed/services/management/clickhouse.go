// Copyright (C) 2023 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package management

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/AlekSi/pointer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"

	inventoryv1 "github.com/percona/pmm/api/inventory/v1"
	managementv1 "github.com/percona/pmm/api/management/v1"
	"github.com/percona/pmm/managed/models"
	"github.com/percona/pmm/managed/services"
)

// defaultClickHouseNativeMetricsPort is the default port of the ClickHouse
// native Prometheus endpoint (the <prometheus> server config section).
const defaultClickHouseNativeMetricsPort = 9363

// clickHouseNativeProbeTimeout bounds the auto-probe of the native endpoint.
const clickHouseNativeProbeTimeout = 3 * time.Second

const (
	clickHouseProtocolHTTP  = "http"
	clickHouseProtocolHTTPS = "https"
	clickHouseMetricsPath   = "/metrics"
)

// probeClickHouseNativeEndpoint reports whether the ClickHouse native
// Prometheus endpoint answers an HTTP GET on {scheme}://{address}:{port}/metrics.
func probeClickHouseNativeEndpoint(ctx context.Context, scheme, address string, port uint16, tlsSkipVerify bool) bool {
	if address == "" {
		return false
	}
	if scheme == "" {
		scheme = clickHouseProtocolHTTP
	}

	probeCtx, cancel := context.WithTimeout(ctx, clickHouseNativeProbeTimeout)
	defer cancel()

	urlStr := fmt.Sprintf("%s://%s%s", scheme, net.JoinHostPort(address, strconv.Itoa(int(port))), clickHouseMetricsPath)
	client := http.DefaultClient
	if scheme == clickHouseProtocolHTTPS && tlsSkipVerify {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec
			},
		}
	}
	req, err := http.NewRequestWithContext(probeCtx, http.MethodGet, urlStr, nil)
	if err != nil {
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close() //nolint:errcheck

	return resp.StatusCode == http.StatusOK
}

// addClickHouse adds a new ClickHouse service. Its metrics reach PMM either
// through the ClickHouse native Prometheus endpoint (modelled as an external
// exporter) or through a PMM-managed clickhouse_exporter, selected by
// req.MetricsSource (auto-probe when unspecified).
func (s *ManagementService) addClickHouse(ctx context.Context, req *managementv1.AddClickHouseServiceParams) (*managementv1.AddServiceResponse, error) {
	clickhouse := &managementv1.ClickHouseServiceResult{}
	var pmmAgentID *string

	nativePort := uint16(req.NativeMetricsPort) //nolint:gosec
	if nativePort == 0 {
		nativePort = defaultClickHouseNativeMetricsPort
	}

	errTx := s.db.InTransactionContext(ctx, nil, func(tx *reform.TX) error {
		var err error
		pmmAgentID, err = s.addClickHouseInTransaction(ctx, tx, req, clickhouse, nativePort)
		if err != nil {
			return err
		}
		return nil
	})
	if errTx != nil {
		return nil, errTx
	}

	if pmmAgentID != nil {
		s.state.RequestStateUpdate(ctx, *pmmAgentID)
	} else {
		s.vmdb.RequestConfigurationUpdate()
	}
	res := &managementv1.AddServiceResponse{
		Service: &managementv1.AddServiceResponse_Clickhouse{
			Clickhouse: clickhouse,
		},
	}
	return res, nil
}

func (s *ManagementService) addClickHouseInTransaction(
	ctx context.Context,
	tx *reform.TX,
	req *managementv1.AddClickHouseServiceParams,
	clickhouse *managementv1.ClickHouseServiceResult,
	nativePort uint16,
) (*string, error) {
	nodeIDValue, service, err := createClickHouseService(tx, req)
	if err != nil {
		return nil, err
	}

	inventoryService, err := services.ToAPIService(service)
	if err != nil {
		return nil, err
	}
	clickhouse.Service = inventoryService.(*inventoryv1.ClickHouseService) //nolint:forcetypeassert

	source, err := clickHouseMetricsSource(ctx, req, nativePort)
	if err != nil {
		return nil, err
	}
	if source == managementv1.MetricsSource_METRICS_SOURCE_NATIVE {
		return addClickHouseNativeMetrics(tx, req, clickhouse, nodeIDValue, service.ServiceID, nativePort)
	}
	return s.addClickHouseManagedExporter(ctx, tx, req, clickhouse, service)
}

func createClickHouseService(
	tx *reform.TX,
	req *managementv1.AddClickHouseServiceParams,
) (string, *models.Service, error) {
	nodeIDValue, err := nodeID(tx, req.NodeId, req.NodeName, req.AddNode, req.Address)
	if err != nil {
		return "", nil, err
	}

	service, err := models.AddNewService(tx.Querier, models.ClickHouseServiceType, &models.AddDBMSServiceParams{
		ServiceName:    req.ServiceName,
		NodeID:         nodeIDValue,
		Environment:    req.Environment,
		Cluster:        req.Cluster,
		ReplicationSet: req.ReplicationSet,
		Address:        pointer.ToStringOrNil(req.Address),
		Port:           pointer.ToUint16OrNil(uint16(req.Port)), //nolint:gosec
		Socket:         pointer.ToStringOrNil(req.Socket),
		CustomLabels:   req.CustomLabels,
	})
	if err != nil {
		return "", nil, err
	}
	return nodeIDValue, service, nil
}

func clickHouseMetricsSource(
	ctx context.Context,
	req *managementv1.AddClickHouseServiceParams,
	nativePort uint16,
) (managementv1.MetricsSource, error) {
	probeScheme := clickHouseProtocolHTTP
	if req.Protocol == clickHouseProtocolHTTPS {
		probeScheme = clickHouseProtocolHTTPS
	}

	source := req.MetricsSource
	switch source {
	case managementv1.MetricsSource_METRICS_SOURCE_UNSPECIFIED:
		if probeClickHouseNativeEndpoint(ctx, probeScheme, req.Address, nativePort, req.TlsSkipVerify) {
			return managementv1.MetricsSource_METRICS_SOURCE_NATIVE, nil
		}
		return managementv1.MetricsSource_METRICS_SOURCE_EXPORTER, nil
	case managementv1.MetricsSource_METRICS_SOURCE_NATIVE:
		if !probeClickHouseNativeEndpoint(ctx, probeScheme, req.Address, nativePort, req.TlsSkipVerify) {
			return source, status.Errorf(codes.FailedPrecondition,
				"ClickHouse native Prometheus endpoint is not reachable at %s:%d; "+
					"enable the <prometheus> server config section or use --metrics-source=exporter",
				req.Address, nativePort)
		}
	case managementv1.MetricsSource_METRICS_SOURCE_EXPORTER:
		return source, nil
	}
	return source, nil
}

func addClickHouseNativeMetrics(
	tx *reform.TX,
	req *managementv1.AddClickHouseServiceParams,
	clickhouse *managementv1.ClickHouseServiceResult,
	nodeIDValue string,
	serviceID string,
	nativePort uint16,
) (*string, error) {
	row, err := models.CreateExternalExporter(tx.Querier, &models.CreateExternalExporterParams{
		RunsOnNodeID:  nodeIDValue,
		ServiceID:     serviceID,
		Username:      req.Username,
		Password:      req.Password,
		Scheme:        clickhouseExporterScheme(req.Protocol, req.Tls),
		MetricsPath:   clickHouseMetricsPath,
		ListenPort:    uint32(nativePort),
		CustomLabels:  req.CustomLabels,
		TLSSkipVerify: req.TlsSkipVerify,
	})
	if err != nil {
		return nil, err
	}

	row.ClickHouseOptions = models.ClickHouseOptions{
		NativeEndpoint:    true,
		NativeMetricsPort: nativePort,
	}
	err = tx.Update(row)
	if err != nil {
		return nil, err
	}

	agent, err := services.ToAPIAgent(tx.Querier, row)
	if err != nil {
		return nil, err
	}
	clickhouse.ExternalExporter = agent.(*inventoryv1.ExternalExporter) //nolint:forcetypeassert
	return row.PMMAgentID, nil
}

func (s *ManagementService) addClickHouseManagedExporter(
	ctx context.Context,
	tx *reform.TX,
	req *managementv1.AddClickHouseServiceParams,
	clickhouse *managementv1.ClickHouseServiceResult,
	service *models.Service,
) (*string, error) {
	var err error
	req.MetricsMode, err = supportedMetricsMode(req.MetricsMode, req.PmmAgentId)
	if err != nil {
		return nil, err
	}

	row, err := models.CreateAgent(tx.Querier, models.ClickHouseExporterType, &models.CreateAgentParams{
		PMMAgentID:    req.PmmAgentId,
		ServiceID:     service.ServiceID,
		Username:      req.Username,
		Password:      req.Password,
		TLS:           req.Tls,
		TLSSkipVerify: req.TlsSkipVerify,
		ExporterOptions: models.ExporterOptions{
			ExposeExporter: req.ExposeExporter,
			PushMetrics:    isPushMode(req.MetricsMode),
		},
		ClickHouseOptions: models.ClickHouseOptionsFromRequest(req),
	})
	if err != nil {
		return nil, err
	}

	if !req.SkipConnectionCheck {
		err = s.cc.CheckConnectionToService(ctx, tx.Querier, service, row)
		if err != nil {
			return nil, err
		}

		err = s.sib.GetInfoFromService(ctx, tx.Querier, service, row)
		if err != nil {
			return nil, err
		}
	}

	agent, err := services.ToAPIAgent(tx.Querier, row)
	if err != nil {
		return nil, err
	}
	clickhouse.ClickhouseExporter = agent.(*inventoryv1.ClickHouseExporter) //nolint:forcetypeassert
	return row.PMMAgentID, nil
}

// clickhouseExporterScheme returns the scheme for the ClickHouse native metrics endpoint.
// It mirrors the logic in clickhouseconn.Config.ExporterScheme.
func clickhouseExporterScheme(protocol string, tls bool) string {
	if protocol == clickHouseProtocolHTTPS || (protocol == "" && tls) {
		return clickHouseProtocolHTTPS
	}
	return clickHouseProtocolHTTP
}
