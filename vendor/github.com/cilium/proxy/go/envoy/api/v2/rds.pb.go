// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/rds.proto

package v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	route "github.com/cilium/proxy/go/envoy/api/v2/route"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RouteConfiguration struct {
	// The name of the route configuration. For example, it might match
	// :ref:`route_config_name
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.Rds.route_config_name>` in
	// :ref:`envoy_api_msg_config.filter.network.http_connection_manager.v2.Rds`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// An array of virtual hosts that make up the route table.
	VirtualHosts []*route.VirtualHost `protobuf:"bytes,2,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
	// Optionally specifies a list of HTTP headers that the connection manager
	// will consider to be internal only. If they are found on external requests they will be cleaned
	// prior to filter invocation. See :ref:`config_http_conn_man_headers_x-envoy-internal` for more
	// information.
	InternalOnlyHeaders []string `protobuf:"bytes,3,rep,name=internal_only_headers,json=internalOnlyHeaders,proto3" json:"internal_only_headers,omitempty"`
	// Specifies a list of HTTP headers that should be added to each response that
	// the connection manager encodes. Headers specified at this level are applied
	// after headers from any enclosed :ref:`envoy_api_msg_route.VirtualHost` or
	// :ref:`envoy_api_msg_route.RouteAction`. For more information, including details on
	// header value syntax, see the documentation on :ref:`custom request headers
	// <config_http_conn_man_headers_custom_request_headers>`.
	ResponseHeadersToAdd []*core.HeaderValueOption `protobuf:"bytes,4,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each response
	// that the connection manager encodes.
	ResponseHeadersToRemove []string `protobuf:"bytes,5,rep,name=response_headers_to_remove,json=responseHeadersToRemove,proto3" json:"response_headers_to_remove,omitempty"`
	// Specifies a list of HTTP headers that should be added to each request
	// routed by the HTTP connection manager. Headers specified at this level are
	// applied after headers from any enclosed :ref:`envoy_api_msg_route.VirtualHost` or
	// :ref:`envoy_api_msg_route.RouteAction`. For more information, including details on
	// header value syntax, see the documentation on :ref:`custom request headers
	// <config_http_conn_man_headers_custom_request_headers>`.
	RequestHeadersToAdd []*core.HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	// An optional boolean that specifies whether the clusters that the route
	// table refers to will be validated by the cluster manager. If set to true
	// and a route refers to a non-existent cluster, the route table will not
	// load. If set to false and a route refers to a non-existent cluster, the
	// route table will load and the router filter will return a 404 if the route
	// is selected at runtime. This setting defaults to true if the route table
	// is statically defined via the :ref:`route_config
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.HttpConnectionManager.route_config>`
	// option. This setting default to false if the route table is loaded dynamically via the
	// :ref:`rds
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.HttpConnectionManager.rds>`
	// option. Users may which to override the default behavior in certain cases (for example when
	// using CDS with a static route table).
	ValidateClusters     *wrappers.BoolValue `protobuf:"bytes,7,opt,name=validate_clusters,json=validateClusters,proto3" json:"validate_clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_78812f46dcff924a, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetVirtualHosts() []*route.VirtualHost {
	if m != nil {
		return m.VirtualHosts
	}
	return nil
}

func (m *RouteConfiguration) GetInternalOnlyHeaders() []string {
	if m != nil {
		return m.InternalOnlyHeaders
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToAdd() []*core.HeaderValueOption {
	if m != nil {
		return m.ResponseHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToRemove() []string {
	if m != nil {
		return m.ResponseHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToAdd() []*core.HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetValidateClusters() *wrappers.BoolValue {
	if m != nil {
		return m.ValidateClusters
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.api.v2.RouteConfiguration")
}

func init() { proto.RegisterFile("envoy/api/v2/rds.proto", fileDescriptor_78812f46dcff924a) }

var fileDescriptor_78812f46dcff924a = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcf, 0x6e, 0xd3, 0x4c,
	0x10, 0xaf, 0x93, 0x7c, 0xad, 0xba, 0xc9, 0x27, 0xd1, 0x6d, 0xda, 0x86, 0x08, 0x25, 0x51, 0xc4,
	0x21, 0xf4, 0x60, 0x23, 0x73, 0x2b, 0x27, 0x52, 0x04, 0x85, 0x4b, 0x25, 0x17, 0x2a, 0x21, 0x0e,
	0xd6, 0xc6, 0x9e, 0x3a, 0x96, 0x36, 0x3b, 0x66, 0x77, 0x6d, 0x94, 0x2b, 0x27, 0xee, 0xbc, 0x04,
	0x8f, 0xc2, 0x15, 0x71, 0xe7, 0x80, 0x10, 0xcf, 0x81, 0xbc, 0x6b, 0x43, 0xdd, 0x82, 0x04, 0x12,
	0x97, 0x68, 0x33, 0xbf, 0x3f, 0xf3, 0xdb, 0x99, 0x35, 0xd9, 0x07, 0x51, 0xe0, 0xda, 0x63, 0x59,
	0xea, 0x15, 0xbe, 0x27, 0x63, 0xe5, 0x66, 0x12, 0x35, 0xd2, 0x9e, 0xa9, 0xbb, 0x2c, 0x4b, 0xdd,
	0xc2, 0x1f, 0xde, 0x6a, 0xb0, 0x22, 0x94, 0xe0, 0x2d, 0x98, 0x02, 0xcb, 0xbd, 0x82, 0xc6, 0xa9,
	0x8a, 0xb0, 0x00, 0xb9, 0xae, 0xd0, 0x51, 0xb3, 0x03, 0xe6, 0x1a, 0xec, 0x6f, 0xad, 0x4e, 0x10,
	0x13, 0x0e, 0x86, 0xc0, 0x84, 0x40, 0xcd, 0x74, 0x8a, 0x42, 0xd5, 0xea, 0x0a, 0x35, 0xff, 0x16,
	0xf9, 0x85, 0xf7, 0x5a, 0xb2, 0x2c, 0x03, 0x59, 0xe3, 0xfd, 0x04, 0x13, 0x34, 0x47, 0xaf, 0x3c,
	0xd9, 0xea, 0xf4, 0x5b, 0x9b, 0xd0, 0xa0, 0xec, 0x71, 0x8c, 0xe2, 0x22, 0x4d, 0x72, 0x69, 0x3c,
	0x29, 0x25, 0x1d, 0xc1, 0x56, 0x30, 0x70, 0x26, 0xce, 0x6c, 0x3b, 0x30, 0x67, 0xfa, 0x94, 0xfc,
	0x5f, 0xa4, 0x52, 0xe7, 0x8c, 0x87, 0x4b, 0x54, 0x5a, 0x0d, 0x5a, 0x93, 0xf6, 0xac, 0xeb, 0x8f,
	0xdd, 0xcb, 0x03, 0x70, 0x6d, 0xe0, 0x73, 0x4b, 0x3c, 0x41, 0xa5, 0xe7, 0x9d, 0x0f, 0x9f, 0xc7,
	0x1b, 0x41, 0xaf, 0xf8, 0x59, 0x52, 0xd4, 0x27, 0x7b, 0xa9, 0xd0, 0x20, 0x05, 0xe3, 0x21, 0x0a,
	0xbe, 0x0e, 0x97, 0xc0, 0x62, 0x90, 0x6a, 0xd0, 0x9e, 0xb4, 0x67, 0xdb, 0xc1, 0x6e, 0x0d, 0x9e,
	0x0a, 0xbe, 0x3e, 0xb1, 0x10, 0x7d, 0x49, 0x0e, 0x24, 0xa8, 0x0c, 0x85, 0x82, 0x9a, 0x1e, 0x6a,
	0x0c, 0x59, 0x1c, 0x0f, 0x3a, 0x26, 0xc9, 0xed, 0x66, 0x92, 0x72, 0xf8, 0xae, 0x15, 0x9f, 0x33,
	0x9e, 0xc3, 0x69, 0x56, 0x5e, 0x2d, 0xe8, 0xd7, 0x26, 0x95, 0xef, 0x33, 0x7c, 0x10, 0xc7, 0xf4,
	0x3e, 0x19, 0xfe, 0xca, 0x5c, 0xc2, 0x0a, 0x0b, 0x18, 0xfc, 0x67, 0x52, 0x1d, 0x5c, 0x53, 0x06,
	0x06, 0xa6, 0x2f, 0xc8, 0xbe, 0x84, 0x57, 0x39, 0x28, 0x7d, 0x35, 0xd8, 0xe6, 0x5f, 0x04, 0xdb,
	0xad, 0x3c, 0x1a, 0xb9, 0x1e, 0x93, 0x9d, 0x82, 0xf1, 0x34, 0x66, 0x1a, 0xc2, 0x88, 0xe7, 0x4a,
	0x97, 0x43, 0xda, 0x9a, 0x38, 0xb3, 0xae, 0x3f, 0x74, 0xed, 0xc6, 0xdd, 0x7a, 0xe3, 0xee, 0x1c,
	0x91, 0x1b, 0xc7, 0xe0, 0x46, 0x2d, 0x3a, 0xae, 0x34, 0xfe, 0xc7, 0x16, 0xd9, 0x33, 0x8b, 0x7e,
	0x58, 0xbf, 0xba, 0x33, 0x90, 0x45, 0x1a, 0x01, 0x7d, 0x4e, 0x7a, 0x67, 0x5a, 0x02, 0x5b, 0x19,
	0x58, 0xd1, 0x51, 0x33, 0xed, 0x0f, 0x7e, 0x60, 0xe3, 0x0d, 0xc7, 0xbf, 0xc5, 0xed, 0x74, 0xa6,
	0x1b, 0x33, 0xe7, 0xae, 0x43, 0x33, 0xb2, 0xf3, 0x44, 0x44, 0x12, 0x56, 0x20, 0x34, 0xe3, 0x95,
	0xf7, 0x9d, 0xa6, 0xf6, 0x12, 0xe1, 0x5a, 0x9b, 0xc3, 0x3f, 0xa1, 0x36, 0x3a, 0x22, 0xe9, 0x3e,
	0x02, 0x1d, 0x2d, 0xff, 0xd5, 0x3d, 0xc6, 0x6f, 0x3e, 0x7d, 0x7d, 0xd7, 0xba, 0x39, 0xed, 0x37,
	0x3e, 0xd6, 0x23, 0xf3, 0xc2, 0xd5, 0x91, 0x73, 0x38, 0xdf, 0x7a, 0xff, 0x65, 0xe4, 0xbc, 0x75,
	0x9c, 0xc5, 0xa6, 0x59, 0xc1, 0xbd, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x06, 0x6f, 0x84, 0x4f,
	0x24, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouteDiscoveryServiceClient is the client API for RouteDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteDiscoveryServiceClient interface {
	StreamRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_StreamRoutesClient, error)
	IncrementalRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_IncrementalRoutesClient, error)
	FetchRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type routeDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRouteDiscoveryServiceClient(cc *grpc.ClientConn) RouteDiscoveryServiceClient {
	return &routeDiscoveryServiceClient{cc}
}

func (c *routeDiscoveryServiceClient) StreamRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_StreamRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.RouteDiscoveryService/StreamRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeDiscoveryServiceStreamRoutesClient{stream}
	return x, nil
}

type RouteDiscoveryService_StreamRoutesClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type routeDiscoveryServiceStreamRoutesClient struct {
	grpc.ClientStream
}

func (x *routeDiscoveryServiceStreamRoutesClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeDiscoveryServiceStreamRoutesClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeDiscoveryServiceClient) IncrementalRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_IncrementalRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v2.RouteDiscoveryService/IncrementalRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeDiscoveryServiceIncrementalRoutesClient{stream}
	return x, nil
}

type RouteDiscoveryService_IncrementalRoutesClient interface {
	Send(*IncrementalDiscoveryRequest) error
	Recv() (*IncrementalDiscoveryResponse, error)
	grpc.ClientStream
}

type routeDiscoveryServiceIncrementalRoutesClient struct {
	grpc.ClientStream
}

func (x *routeDiscoveryServiceIncrementalRoutesClient) Send(m *IncrementalDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeDiscoveryServiceIncrementalRoutesClient) Recv() (*IncrementalDiscoveryResponse, error) {
	m := new(IncrementalDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeDiscoveryServiceClient) FetchRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.RouteDiscoveryService/FetchRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteDiscoveryServiceServer is the server API for RouteDiscoveryService service.
type RouteDiscoveryServiceServer interface {
	StreamRoutes(RouteDiscoveryService_StreamRoutesServer) error
	IncrementalRoutes(RouteDiscoveryService_IncrementalRoutesServer) error
	FetchRoutes(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterRouteDiscoveryServiceServer(s *grpc.Server, srv RouteDiscoveryServiceServer) {
	s.RegisterService(&_RouteDiscoveryService_serviceDesc, srv)
}

func _RouteDiscoveryService_StreamRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteDiscoveryServiceServer).StreamRoutes(&routeDiscoveryServiceStreamRoutesServer{stream})
}

type RouteDiscoveryService_StreamRoutesServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type routeDiscoveryServiceStreamRoutesServer struct {
	grpc.ServerStream
}

func (x *routeDiscoveryServiceStreamRoutesServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeDiscoveryServiceStreamRoutesServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteDiscoveryService_IncrementalRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteDiscoveryServiceServer).IncrementalRoutes(&routeDiscoveryServiceIncrementalRoutesServer{stream})
}

type RouteDiscoveryService_IncrementalRoutesServer interface {
	Send(*IncrementalDiscoveryResponse) error
	Recv() (*IncrementalDiscoveryRequest, error)
	grpc.ServerStream
}

type routeDiscoveryServiceIncrementalRoutesServer struct {
	grpc.ServerStream
}

func (x *routeDiscoveryServiceIncrementalRoutesServer) Send(m *IncrementalDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeDiscoveryServiceIncrementalRoutesServer) Recv() (*IncrementalDiscoveryRequest, error) {
	m := new(IncrementalDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteDiscoveryService_FetchRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteDiscoveryServiceServer).FetchRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.RouteDiscoveryService/FetchRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteDiscoveryServiceServer).FetchRoutes(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.RouteDiscoveryService",
	HandlerType: (*RouteDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRoutes",
			Handler:    _RouteDiscoveryService_FetchRoutes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRoutes",
			Handler:       _RouteDiscoveryService_StreamRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "IncrementalRoutes",
			Handler:       _RouteDiscoveryService_IncrementalRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/rds.proto",
}
