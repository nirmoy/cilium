// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/eds.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_endpoint "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/endpoint"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Each route from RDS will map to a single cluster or traffic split across
// clusters using weights expressed in the RDS WeightedCluster.
//
// With EDS, each cluster is treated independently from a LB perspective, with
// LB taking place between the Localities within a cluster and at a finer
// granularity between the hosts within a locality. For a given cluster, the
// effective weight of a host is its load_balancing_weight multiplied by the
// load_balancing_weight of its Locality.
type ClusterLoadAssignment struct {
	// Name of the cluster. This will be the :ref:`service_name
	// <envoy_api_field_Cluster.EdsClusterConfig.service_name>` value if specified
	// in the cluster :ref:`EdsClusterConfig
	// <envoy_api_msg_Cluster.EdsClusterConfig>`.
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	// List of endpoints to load balance to.
	Endpoints []*envoy_api_v2_endpoint.LocalityLbEndpoints `protobuf:"bytes,2,rep,name=endpoints" json:"endpoints,omitempty"`
	// Load balancing policy settings.
	Policy *ClusterLoadAssignment_Policy `protobuf:"bytes,4,opt,name=policy" json:"policy,omitempty"`
}

func (m *ClusterLoadAssignment) Reset()                    { *m = ClusterLoadAssignment{} }
func (m *ClusterLoadAssignment) String() string            { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment) ProtoMessage()               {}
func (*ClusterLoadAssignment) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ClusterLoadAssignment) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterLoadAssignment) GetEndpoints() []*envoy_api_v2_endpoint.LocalityLbEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetPolicy() *ClusterLoadAssignment_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

// Load balancing policy settings.
type ClusterLoadAssignment_Policy struct {
	// Percentage of traffic (0-100) that should be dropped. This
	// action allows protection of upstream hosts should they unable to
	// recover from an outage or should they be unable to autoscale and hence
	// overall incoming traffic volume need to be trimmed to protect them.
	// [#v2-api-diff: This is known as maintenance mode in v1.]
	DropOverload float64 `protobuf:"fixed64,1,opt,name=drop_overload,json=dropOverload" json:"drop_overload,omitempty"`
}

func (m *ClusterLoadAssignment_Policy) Reset()                    { *m = ClusterLoadAssignment_Policy{} }
func (m *ClusterLoadAssignment_Policy) String() string            { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy) ProtoMessage()               {}
func (*ClusterLoadAssignment_Policy) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

func (m *ClusterLoadAssignment_Policy) GetDropOverload() float64 {
	if m != nil {
		return m.DropOverload
	}
	return 0
}

func init() {
	proto.RegisterType((*ClusterLoadAssignment)(nil), "envoy.api.v2.ClusterLoadAssignment")
	proto.RegisterType((*ClusterLoadAssignment_Policy)(nil), "envoy.api.v2.ClusterLoadAssignment.Policy")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EndpointDiscoveryService service

type EndpointDiscoveryServiceClient interface {
	// The resource_names field in DiscoveryRequest specifies a list of clusters
	// to subscribe to updates for.
	StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error)
	FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type endpointDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewEndpointDiscoveryServiceClient(cc *grpc.ClientConn) EndpointDiscoveryServiceClient {
	return &endpointDiscoveryServiceClient{cc}
}

func (c *endpointDiscoveryServiceClient) StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[0], c.cc, "/envoy.api.v2.EndpointDiscoveryService/StreamEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceStreamEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_StreamEndpointsClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceStreamEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EndpointDiscoveryService service

type EndpointDiscoveryServiceServer interface {
	// The resource_names field in DiscoveryRequest specifies a list of clusters
	// to subscribe to updates for.
	StreamEndpoints(EndpointDiscoveryService_StreamEndpointsServer) error
	FetchEndpoints(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterEndpointDiscoveryServiceServer(s *grpc.Server, srv EndpointDiscoveryServiceServer) {
	s.RegisterService(&_EndpointDiscoveryService_serviceDesc, srv)
}

func _EndpointDiscoveryService_StreamEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).StreamEndpoints(&endpointDiscoveryServiceStreamEndpointsServer{stream})
}

type EndpointDiscoveryService_StreamEndpointsServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceStreamEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_FetchEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.EndpointDiscoveryService",
	HandlerType: (*EndpointDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchEndpoints",
			Handler:    _EndpointDiscoveryService_FetchEndpoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEndpoints",
			Handler:       _EndpointDiscoveryService_StreamEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/eds.proto",
}

func init() { proto.RegisterFile("envoy/api/v2/eds.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x3d, 0x8b, 0xd4, 0x40,
	0x18, 0xde, 0xc9, 0x85, 0x85, 0x9d, 0x5b, 0x15, 0x86, 0xd3, 0xcb, 0x85, 0xe5, 0x2e, 0x04, 0x8b,
	0xf5, 0x90, 0x44, 0x62, 0x77, 0x85, 0x68, 0xfc, 0xc0, 0x62, 0x39, 0x25, 0xd7, 0x68, 0x75, 0xcc,
	0x26, 0x2f, 0x71, 0x20, 0x99, 0x37, 0x66, 0x66, 0x03, 0x69, 0xfd, 0x0b, 0xfe, 0x04, 0x1b, 0x7f,
	0x83, 0x95, 0xa5, 0xbd, 0xbd, 0x85, 0xd8, 0x88, 0x7f, 0x42, 0xf2, 0xa9, 0x41, 0xed, 0x9c, 0xea,
	0xe5, 0x7d, 0x3e, 0xf2, 0x3c, 0x93, 0xa1, 0x37, 0x40, 0x56, 0x58, 0xfb, 0xbc, 0x10, 0x7e, 0x15,
	0xf8, 0x90, 0x28, 0xaf, 0x28, 0x51, 0x23, 0x5b, 0xb6, 0x7b, 0x8f, 0x17, 0xc2, 0xab, 0x02, 0x7b,
	0x35, 0x61, 0x25, 0x42, 0xc5, 0x58, 0x41, 0x59, 0x77, 0x5c, 0xfb, 0xe6, 0xd4, 0x43, 0x26, 0x05,
	0x0a, 0xa9, 0xc7, 0xa1, 0x67, 0xad, 0x52, 0xc4, 0x34, 0x83, 0x96, 0xc6, 0xa5, 0x44, 0xcd, 0xb5,
	0x40, 0xd9, 0x7f, 0xcf, 0x3e, 0xac, 0x78, 0x26, 0x12, 0xae, 0xc1, 0x1f, 0x86, 0x1e, 0x38, 0x48,
	0x31, 0xc5, 0x76, 0xf4, 0x9b, 0xa9, 0xdb, 0xba, 0xef, 0x0c, 0x7a, 0xfd, 0x61, 0xb6, 0x53, 0x1a,
	0xca, 0x0d, 0xf2, 0xe4, 0x81, 0x52, 0x22, 0x95, 0x39, 0x48, 0xcd, 0x6e, 0xd3, 0x65, 0xdc, 0x01,
	0x97, 0x92, 0xe7, 0x60, 0x11, 0x87, 0xac, 0x17, 0xe1, 0xe2, 0xc3, 0xf7, 0x8f, 0x7b, 0x66, 0x69,
	0x38, 0x24, 0xda, 0xef, 0xe1, 0x73, 0x9e, 0x03, 0x3b, 0xa7, 0x8b, 0x21, 0xa6, 0xb2, 0x0c, 0x67,
	0x6f, 0xbd, 0x1f, 0x9c, 0x7a, 0xbf, 0x57, 0xf7, 0xc6, 0x16, 0x1b, 0x8c, 0x79, 0x26, 0x74, 0xbd,
	0xd9, 0x3e, 0x1e, 0x14, 0xa1, 0xf9, 0xe9, 0xcb, 0xc9, 0x2c, 0xfa, 0x65, 0xc1, 0x42, 0x3a, 0x2f,
	0x30, 0x13, 0x71, 0x6d, 0x99, 0x0e, 0xf9, 0xd3, 0xec, 0xaf, 0x91, 0xbd, 0xe7, 0xad, 0x22, 0xea,
	0x95, 0xf6, 0x53, 0x3a, 0xef, 0x36, 0xec, 0x1e, 0xbd, 0x92, 0x94, 0x58, 0x5c, 0x36, 0x97, 0x9d,
	0x21, 0x4f, 0xda, 0x32, 0x24, 0x3c, 0x6a, 0xca, 0x1c, 0x30, 0x76, 0x34, 0x6b, 0xcf, 0xcb, 0xfb,
	0xb7, 0x66, 0xfd, 0x89, 0x96, 0x0d, 0xff, 0x59, 0x4f, 0x0f, 0x7e, 0x10, 0x6a, 0x0d, 0x61, 0x1f,
	0x0d, 0x3f, 0xed, 0x02, 0xca, 0x4a, 0xc4, 0xc0, 0x5e, 0xd0, 0x6b, 0x17, 0xba, 0x04, 0x9e, 0x8f,
	0x75, 0xd8, 0xf1, 0x34, 0xed, 0x28, 0x89, 0xe0, 0xf5, 0x0e, 0x94, 0xb6, 0x4f, 0xfe, 0x89, 0xab,
	0x02, 0xa5, 0x02, 0x77, 0xb6, 0x26, 0x77, 0x08, 0xdb, 0xd1, 0xab, 0x4f, 0x40, 0xc7, 0xaf, 0xfe,
	0xa3, 0xb1, 0xfb, 0xe6, 0xf3, 0xb7, 0xb7, 0xc6, 0xca, 0x3d, 0x9c, 0xbc, 0xbf, 0xb3, 0xf1, 0xe2,
	0xcf, 0xc8, 0x69, 0x68, 0xbe, 0xff, 0x7a, 0x4c, 0xb6, 0xf3, 0xf6, 0x81, 0xdc, 0xfd, 0x19, 0x00,
	0x00, 0xff, 0xff, 0x0a, 0x92, 0x5a, 0xa4, 0xd9, 0x02, 0x00, 0x00,
}
