// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: managementpb/service/service.proto

package servicev1beta1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	inventorypb "github.com/percona/pmm/api/inventorypb"
	agent "github.com/percona/pmm/api/managementpb/agent"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Service status.
type UniversalService_Status int32

const (
	// In case we don't support the db vendor yet.
	UniversalService_STATUS_INVALID UniversalService_Status = 0
	// The service is up.
	UniversalService_UP UniversalService_Status = 1
	// The service is down.
	UniversalService_DOWN UniversalService_Status = 2
	// The service's status cannot be known (e.g. there are no metrics yet).
	UniversalService_UNKNOWN UniversalService_Status = 3
)

// Enum value maps for UniversalService_Status.
var (
	UniversalService_Status_name = map[int32]string{
		0: "STATUS_INVALID",
		1: "UP",
		2: "DOWN",
		3: "UNKNOWN",
	}
	UniversalService_Status_value = map[string]int32{
		"STATUS_INVALID": 0,
		"UP":             1,
		"DOWN":           2,
		"UNKNOWN":        3,
	}
)

func (x UniversalService_Status) Enum() *UniversalService_Status {
	p := new(UniversalService_Status)
	*p = x
	return p
}

func (x UniversalService_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UniversalService_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_service_service_proto_enumTypes[0].Descriptor()
}

func (UniversalService_Status) Type() protoreflect.EnumType {
	return &file_managementpb_service_service_proto_enumTypes[0]
}

func (x UniversalService_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UniversalService_Status.Descriptor instead.
func (UniversalService_Status) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_service_service_proto_rawDescGZIP(), []int{0, 0}
}

type UniversalService struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique service identifier.
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	// Service type.
	ServiceType string `protobuf:"bytes,2,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	// User-defined name unique across all Services.
	ServiceName string `protobuf:"bytes,3,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Database name.
	DatabaseName string `protobuf:"bytes,4,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	// Node identifier where this instance runs.
	NodeId string `protobuf:"bytes,5,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Node name where this instance runs.
	NodeName string `protobuf:"bytes,6,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Environment name.
	Environment string `protobuf:"bytes,7,opt,name=environment,proto3" json:"environment,omitempty"`
	// Cluster name.
	Cluster string `protobuf:"bytes,8,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Replication set name.
	ReplicationSet string `protobuf:"bytes,9,opt,name=replication_set,json=replicationSet,proto3" json:"replication_set,omitempty"`
	// Custom user-assigned labels for Service.
	CustomLabels map[string]string `protobuf:"bytes,10,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// External group name.
	ExternalGroup string `protobuf:"bytes,11,opt,name=external_group,json=externalGroup,proto3" json:"external_group,omitempty"`
	// Access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `protobuf:"bytes,12,opt,name=address,proto3" json:"address,omitempty"`
	// Access port.
	// Port is required when the address present.
	Port uint32 `protobuf:"varint,13,opt,name=port,proto3" json:"port,omitempty"`
	// Access unix socket.
	// Address (and port) or socket is required.
	Socket string `protobuf:"bytes,14,opt,name=socket,proto3" json:"socket,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// List of agents related to this service.
	Agents []*agent.UniversalAgent `protobuf:"bytes,17,rep,name=agents,proto3" json:"agents,omitempty"`
	// The health status of the service.
	Status UniversalService_Status `protobuf:"varint,18,opt,name=status,proto3,enum=service.v1beta1.UniversalService_Status" json:"status,omitempty"`
	// The service/database version.
	Version       string `protobuf:"bytes,19,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UniversalService) Reset() {
	*x = UniversalService{}
	mi := &file_managementpb_service_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UniversalService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniversalService) ProtoMessage() {}

func (x *UniversalService) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_service_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniversalService.ProtoReflect.Descriptor instead.
func (*UniversalService) Descriptor() ([]byte, []int) {
	return file_managementpb_service_service_proto_rawDescGZIP(), []int{0}
}

func (x *UniversalService) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *UniversalService) GetServiceType() string {
	if x != nil {
		return x.ServiceType
	}
	return ""
}

func (x *UniversalService) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *UniversalService) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *UniversalService) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *UniversalService) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *UniversalService) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *UniversalService) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *UniversalService) GetReplicationSet() string {
	if x != nil {
		return x.ReplicationSet
	}
	return ""
}

func (x *UniversalService) GetCustomLabels() map[string]string {
	if x != nil {
		return x.CustomLabels
	}
	return nil
}

func (x *UniversalService) GetExternalGroup() string {
	if x != nil {
		return x.ExternalGroup
	}
	return ""
}

func (x *UniversalService) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UniversalService) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *UniversalService) GetSocket() string {
	if x != nil {
		return x.Socket
	}
	return ""
}

func (x *UniversalService) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UniversalService) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UniversalService) GetAgents() []*agent.UniversalAgent {
	if x != nil {
		return x.Agents
	}
	return nil
}

func (x *UniversalService) GetStatus() UniversalService_Status {
	if x != nil {
		return x.Status
	}
	return UniversalService_STATUS_INVALID
}

func (x *UniversalService) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type ListServiceRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Return only Services running on that Node.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Return only services filtered by service type.
	ServiceType inventorypb.ServiceType `protobuf:"varint,2,opt,name=service_type,json=serviceType,proto3,enum=inventory.ServiceType" json:"service_type,omitempty"`
	// Return only services in this external group.
	ExternalGroup string `protobuf:"bytes,3,opt,name=external_group,json=externalGroup,proto3" json:"external_group,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListServiceRequest) Reset() {
	*x = ListServiceRequest{}
	mi := &file_managementpb_service_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceRequest) ProtoMessage() {}

func (x *ListServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_service_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceRequest.ProtoReflect.Descriptor instead.
func (*ListServiceRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_service_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListServiceRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *ListServiceRequest) GetServiceType() inventorypb.ServiceType {
	if x != nil {
		return x.ServiceType
	}
	return inventorypb.ServiceType(0)
}

func (x *ListServiceRequest) GetExternalGroup() string {
	if x != nil {
		return x.ExternalGroup
	}
	return ""
}

type ListServiceResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of Services.
	Services      []*UniversalService `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListServiceResponse) Reset() {
	*x = ListServiceResponse{}
	mi := &file_managementpb_service_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceResponse) ProtoMessage() {}

func (x *ListServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_service_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceResponse.ProtoReflect.Descriptor instead.
func (*ListServiceResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_service_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListServiceResponse) GetServices() []*UniversalService {
	if x != nil {
		return x.Services
	}
	return nil
}

var File_managementpb_service_service_proto protoreflect.FileDescriptor

var file_managementpb_service_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x70,
	0x62, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x07, 0x0a, 0x10, 0x55, 0x6e,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x74, 0x12, 0x58, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x35, 0x0a, 0x06, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x06, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x1a, 0x3f, 0x0a, 0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x3b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x55, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f,
	0x57, 0x4e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x03, 0x22, 0x8f, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x39, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x22, 0x54, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x32, 0xc9, 0x01, 0x0a, 0x0b, 0x4d, 0x67,
	0x6d, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb9, 0x01, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5e, 0x92, 0x41, 0x35, 0x12, 0x0d, 0x4c, 0x69, 0x73, 0x74,
	0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x24, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x73, 0x20, 0x61, 0x20, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x20, 0x6c, 0x69,
	0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0xc0, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e,
	0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x53, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_service_service_proto_rawDescOnce sync.Once
	file_managementpb_service_service_proto_rawDescData = file_managementpb_service_service_proto_rawDesc
)

func file_managementpb_service_service_proto_rawDescGZIP() []byte {
	file_managementpb_service_service_proto_rawDescOnce.Do(func() {
		file_managementpb_service_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_service_service_proto_rawDescData)
	})
	return file_managementpb_service_service_proto_rawDescData
}

var (
	file_managementpb_service_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_managementpb_service_service_proto_msgTypes  = make([]protoimpl.MessageInfo, 4)
	file_managementpb_service_service_proto_goTypes   = []any{
		(UniversalService_Status)(0),  // 0: service.v1beta1.UniversalService.Status
		(*UniversalService)(nil),      // 1: service.v1beta1.UniversalService
		(*ListServiceRequest)(nil),    // 2: service.v1beta1.ListServiceRequest
		(*ListServiceResponse)(nil),   // 3: service.v1beta1.ListServiceResponse
		nil,                           // 4: service.v1beta1.UniversalService.CustomLabelsEntry
		(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
		(*agent.UniversalAgent)(nil),  // 6: agent.v1beta1.UniversalAgent
		(inventorypb.ServiceType)(0),  // 7: inventory.ServiceType
	}
)

var file_managementpb_service_service_proto_depIdxs = []int32{
	4, // 0: service.v1beta1.UniversalService.custom_labels:type_name -> service.v1beta1.UniversalService.CustomLabelsEntry
	5, // 1: service.v1beta1.UniversalService.created_at:type_name -> google.protobuf.Timestamp
	5, // 2: service.v1beta1.UniversalService.updated_at:type_name -> google.protobuf.Timestamp
	6, // 3: service.v1beta1.UniversalService.agents:type_name -> agent.v1beta1.UniversalAgent
	0, // 4: service.v1beta1.UniversalService.status:type_name -> service.v1beta1.UniversalService.Status
	7, // 5: service.v1beta1.ListServiceRequest.service_type:type_name -> inventory.ServiceType
	1, // 6: service.v1beta1.ListServiceResponse.services:type_name -> service.v1beta1.UniversalService
	2, // 7: service.v1beta1.MgmtService.ListServices:input_type -> service.v1beta1.ListServiceRequest
	3, // 8: service.v1beta1.MgmtService.ListServices:output_type -> service.v1beta1.ListServiceResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_managementpb_service_service_proto_init() }
func file_managementpb_service_service_proto_init() {
	if File_managementpb_service_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managementpb_service_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_service_service_proto_goTypes,
		DependencyIndexes: file_managementpb_service_service_proto_depIdxs,
		EnumInfos:         file_managementpb_service_service_proto_enumTypes,
		MessageInfos:      file_managementpb_service_service_proto_msgTypes,
	}.Build()
	File_managementpb_service_service_proto = out.File
	file_managementpb_service_service_proto_rawDesc = nil
	file_managementpb_service_service_proto_goTypes = nil
	file_managementpb_service_service_proto_depIdxs = nil
}
