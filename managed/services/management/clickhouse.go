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

const (
	// DefaultClickHouseNativeMetricsPort is the default port of the ClickHouse
	// native Prometheus endpoint (the <prometheus> server config section).
	defaultClickHouseNativeMetricsPort = 9363
	maxClickHouseNativeMetricsPort     = 1<<16 - 1

	// ClickHouseNativeProbeTimeout bounds the auto-probe of the native endpoint.
	clickHouseNativeProbeTimeout = 3 * time.Second

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

	nativePort, err := clickHouseNativeMetricsPort(req.NativeMetricsPort)
	if err != nil {
		return nil, err
	}

	pmmAgentID, err := s.addClickHouseInTransaction(ctx, req, clickhouse, nativePort)
	if err != nil {
		return nil, err
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

func clickHouseNativeMetricsPort(port uint32) (uint16, error) {
	if port == 0 {
		return defaultClickHouseNativeMetricsPort, nil
	}
	if port > maxClickHouseNativeMetricsPort {
		return 0, status.Error(codes.InvalidArgument, "ClickHouse native metrics port should be between 1 and 65535.")
	}
	return uint16(port), nil
}

func (s *ManagementService) addClickHouseInTransaction(
	ctx context.Context,
	req *managementv1.AddClickHouseServiceParams,
	clickhouse *managementv1.ClickHouseServiceResult,
	nativePort uint16,
) (*string, error) {
	var pmmAgentID *string
	errTx := s.db.InTransactionContext(ctx, nil, func(tx *reform.TX) error {
		service, resolvedNodeID, err := addClickHouseService(tx, req, clickhouse)
		if err != nil {
			return err
		}

		source, err := resolveClickHouseMetricsSource(ctx, req, nativePort)
		if err != nil {
			return err
		}

		if source == managementv1.MetricsSource_METRICS_SOURCE_NATIVE {
			pmmAgentID, err = addClickHouseNativeExporter(tx, req, resolvedNodeID, service.ServiceID, nativePort, clickhouse)
			return err
		}

		pmmAgentID, err = s.addClickHouseManagedExporter(ctx, tx, req, service, clickhouse)
		return err
	})
	return pmmAgentID, errTx
}

func addClickHouseService(
	tx *reform.TX,
	req *managementv1.AddClickHouseServiceParams,
	clickhouse *managementv1.ClickHouseServiceResult,
) (*models.Service, string, error) {
	resolvedNodeID, err := nodeID(tx, req.NodeId, req.NodeName, req.AddNode, req.Address)
	if err != nil {
		return nil, "", err
	}

	service, err := models.AddNewService(tx.Querier, models.ClickHouseServiceType, &models.AddDBMSServiceParams{
		ServiceName:    req.ServiceName,
		NodeID:         resolvedNodeID,
		Environment:    req.Environment,
		Cluster:        req.Cluster,
		ReplicationSet: req.ReplicationSet,
		Address:        pointer.ToStringOrNil(req.Address),
		Port:           pointer.ToUint16OrNil(uint16(req.Port)), //nolint:gosec
		Socket:         pointer.ToStringOrNil(req.Socket),
		CustomLabels:   req.CustomLabels,
	})
	if err != nil {
		return nil, "", err
	}

	inventoryService, err := services.ToAPIService(service)
	if err != nil {
		return nil, "", err
	}
	clickhouse.Service = inventoryService.(*inventoryv1.ClickHouseService) //nolint:forcetypeassert

	return service, resolvedNodeID, nil
}

func resolveClickHouseMetricsSource(
	ctx context.Context,
	req *managementv1.AddClickHouseServiceParams,
	nativePort uint16,
) (managementv1.MetricsSource, error) {
	probeScheme := clickHouseProtocolHTTP
	if req.Protocol == clickHouseProtocolHTTPS {
		probeScheme = clickHouseProtocolHTTPS
	}

	switch req.MetricsSource {
	case managementv1.MetricsSource_METRICS_SOURCE_UNSPECIFIED:
		if probeClickHouseNativeEndpoint(ctx, probeScheme, req.Address, nativePort, req.TlsSkipVerify) {
			return managementv1.MetricsSource_METRICS_SOURCE_NATIVE, nil
		}
		return managementv1.MetricsSource_METRICS_SOURCE_EXPORTER, nil
	case managementv1.MetricsSource_METRICS_SOURCE_NATIVE:
		if !probeClickHouseNativeEndpoint(ctx, probeScheme, req.Address, nativePort, req.TlsSkipVerify) {
			return managementv1.MetricsSource_METRICS_SOURCE_UNSPECIFIED, status.Errorf(codes.FailedPrecondition,
				"ClickHouse native Prometheus endpoint is not reachable at %s:%d; "+
					"enable the <prometheus> server config section or use --metrics-source=exporter",
				req.Address, nativePort)
		}
		return managementv1.MetricsSource_METRICS_SOURCE_NATIVE, nil
	case managementv1.MetricsSource_METRICS_SOURCE_EXPORTER:
		return managementv1.MetricsSource_METRICS_SOURCE_EXPORTER, nil
	default:
		return req.MetricsSource, nil
	}
}

func addClickHouseNativeExporter(
	tx *reform.TX,
	req *managementv1.AddClickHouseServiceParams,
	nodeID string,
	serviceID string,
	nativePort uint16,
	clickhouse *managementv1.ClickHouseServiceResult,
) (*string, error) {
	row, err := models.CreateExternalExporter(tx.Querier, &models.CreateExternalExporterParams{
		RunsOnNodeID:  nodeID,
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
	service *models.Service,
	clickhouse *managementv1.ClickHouseServiceResult,
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
