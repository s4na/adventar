// Code generated by protoc-gen-go. DO NOT EDIT.
// source: adventar/v1/adventar.proto

package adventar_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("adventar/v1/adventar.proto", fileDescriptor_c3b1639048342dcd) }

var fileDescriptor_c3b1639048342dcd = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6f, 0xda, 0x30,
	0x18, 0x86, 0xd1, 0x0e, 0x0c, 0x39, 0x82, 0x09, 0x4b, 0xec, 0x10, 0x26, 0xb6, 0x21, 0x6d, 0xc7,
	0x44, 0x6c, 0xa7, 0x69, 0x27, 0x04, 0x08, 0x2a, 0xb5, 0x3d, 0x80, 0x2a, 0xf5, 0x86, 0x4c, 0xf2,
	0x35, 0x44, 0x0d, 0x76, 0x1a, 0x1b, 0x24, 0xae, 0xfd, 0x0b, 0xfd, 0x69, 0xbd, 0xf5, 0xdc, 0x1f,
	0x52, 0x25, 0x8e, 0x23, 0x3b, 0x4a, 0xc8, 0x8d, 0xe4, 0x79, 0xfd, 0xe0, 0x7c, 0xef, 0x87, 0x6c,
	0xe2, 0x9f, 0x80, 0x0a, 0x92, 0xb8, 0xa7, 0x89, 0xab, 0x7e, 0x3b, 0x71, 0xc2, 0x04, 0xc3, 0x56,
	0xf1, 0x7c, 0x9a, 0xd8, 0xc3, 0x80, 0xb1, 0x20, 0x02, 0x37, 0x43, 0xbb, 0xe3, 0x83, 0x0b, 0x87,
	0x58, 0x9c, 0x65, 0xd2, 0xfe, 0x96, 0x43, 0x12, 0x87, 0x2e, 0xa1, 0x94, 0x09, 0x22, 0x42, 0x46,
	0x79, 0x4e, 0x87, 0xfa, 0x7f, 0x24, 0xc0, 0xd9, 0x31, 0xf1, 0x40, 0xc1, 0x91, 0x01, 0x63, 0x6f,
	0x7b, 0x00, 0xce, 0x49, 0xa0, 0xf8, 0x9f, 0xb7, 0x0e, 0xea, 0x4c, 0xf3, 0x08, 0xbe, 0x47, 0xdd,
	0xeb, 0x90, 0x8b, 0x19, 0x89, 0x80, 0xfa, 0x24, 0xe1, 0xf8, 0xa7, 0xa3, 0xdd, 0xd1, 0x31, 0xd8,
	0x1a, 0x9e, 0x8e, 0xc0, 0x85, 0x3d, 0xbe, 0x14, 0xe1, 0x31, 0xa3, 0x1c, 0xc6, 0x2d, 0xbc, 0x46,
	0xd6, 0x12, 0x0a, 0x82, 0xbf, 0x1b, 0x87, 0x34, 0xa2, 0xac, 0x3f, 0xea, 0x03, 0x85, 0xf3, 0x06,
	0xf5, 0x66, 0x09, 0x10, 0x01, 0x85, 0xd6, 0xbc, 0x8b, 0x09, 0x95, 0x79, 0x60, 0x66, 0x72, 0x2a,
	0x75, 0x77, 0xb1, 0x5f, 0xaf, 0x33, 0x61, 0xa3, 0xee, 0x16, 0xf5, 0xe6, 0x10, 0x41, 0xad, 0xce,
	0x84, 0x4a, 0xf7, 0xd5, 0x91, 0x55, 0x3b, 0x6a, 0x0f, 0x9c, 0x45, 0xba, 0x07, 0xe3, 0x16, 0xf6,
	0x51, 0x5f, 0x1f, 0xee, 0x46, 0x10, 0xc1, 0xf1, 0xaf, 0xda, 0xe1, 0x67, 0x5c, 0x59, 0x7f, 0x37,
	0xc5, 0xf4, 0x9e, 0x52, 0xbc, 0xa0, 0x22, 0x09, 0x81, 0x97, 0x7a, 0xd2, 0x48, 0x75, 0x4f, 0x46,
	0xa0, 0x70, 0xce, 0x91, 0x25, 0xab, 0x48, 0xd1, 0xb9, 0xe4, 0xd4, 0x88, 0x72, 0x62, 0x23, 0x90,
	0x21, 0x69, 0x91, 0x0d, 0x54, 0x59, 0x34, 0x72, 0xd9, 0xb2, 0x42, 0x96, 0x1c, 0x7c, 0x95, 0x45,
	0x23, 0xcd, 0x7d, 0xfc, 0x43, 0xed, 0x4d, 0x18, 0xd0, 0x2b, 0x8a, 0x6d, 0x43, 0x22, 0x5f, 0xaa,
	0xf3, 0x7d, 0xf3, 0x9a, 0x1c, 0xd2, 0xd5, 0xf8, 0x8f, 0x3e, 0x2f, 0x41, 0xa4, 0x0f, 0x78, 0x58,
	0xde, 0xf3, 0xf4, 0xed, 0xc5, 0xc3, 0x53, 0x84, 0xe4, 0xd7, 0x66, 0xe7, 0x47, 0x15, 0x63, 0x68,
	0x54, 0x04, 0xc8, 0x5a, 0x01, 0x89, 0xc4, 0x7e, 0xb6, 0x07, 0xef, 0xb1, 0x34, 0x04, 0x8d, 0x54,
	0x97, 0x6c, 0x04, 0xf2, 0x92, 0x07, 0xcf, 0xaf, 0xef, 0x2f, 0x9f, 0xbe, 0xe0, 0xae, 0xbb, 0xcf,
	0xe8, 0xd6, 0x4b, 0xf1, 0xae, 0x9d, 0x4d, 0xed, 0xef, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46,
	0xb5, 0x40, 0xce, 0x06, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdventarClient is the client API for Adventar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdventarClient interface {
	ListCalendars(ctx context.Context, in *ListCalendarsRequest, opts ...grpc.CallOption) (*ListCalendarsResponse, error)
	GetCalendar(ctx context.Context, in *GetCalendarRequest, opts ...grpc.CallOption) (*GetCalendarResponse, error)
	CreateCalendar(ctx context.Context, in *CreateCalendarRequest, opts ...grpc.CallOption) (*Calendar, error)
	UpdateCalendar(ctx context.Context, in *UpdateCalendarRequest, opts ...grpc.CallOption) (*Calendar, error)
	DeleteCalendar(ctx context.Context, in *DeleteCalendarRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListCalendarStats(ctx context.Context, in *ListCalendarStatsRequest, opts ...grpc.CallOption) (*ListCalendarStatsResponse, error)
	ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error)
	CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*Entry, error)
	UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*Entry, error)
	DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*User, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type adventarClient struct {
	cc *grpc.ClientConn
}

func NewAdventarClient(cc *grpc.ClientConn) AdventarClient {
	return &adventarClient{cc}
}

func (c *adventarClient) ListCalendars(ctx context.Context, in *ListCalendarsRequest, opts ...grpc.CallOption) (*ListCalendarsResponse, error) {
	out := new(ListCalendarsResponse)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/ListCalendars", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) GetCalendar(ctx context.Context, in *GetCalendarRequest, opts ...grpc.CallOption) (*GetCalendarResponse, error) {
	out := new(GetCalendarResponse)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/GetCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) CreateCalendar(ctx context.Context, in *CreateCalendarRequest, opts ...grpc.CallOption) (*Calendar, error) {
	out := new(Calendar)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/CreateCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) UpdateCalendar(ctx context.Context, in *UpdateCalendarRequest, opts ...grpc.CallOption) (*Calendar, error) {
	out := new(Calendar)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/UpdateCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) DeleteCalendar(ctx context.Context, in *DeleteCalendarRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/DeleteCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) ListCalendarStats(ctx context.Context, in *ListCalendarStatsRequest, opts ...grpc.CallOption) (*ListCalendarStatsResponse, error) {
	out := new(ListCalendarStatsResponse)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/ListCalendarStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error) {
	out := new(ListEntriesResponse)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/ListEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/CreateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/UpdateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/DeleteEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventarClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/adventar.v1.Adventar/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdventarServer is the server API for Adventar service.
type AdventarServer interface {
	ListCalendars(context.Context, *ListCalendarsRequest) (*ListCalendarsResponse, error)
	GetCalendar(context.Context, *GetCalendarRequest) (*GetCalendarResponse, error)
	CreateCalendar(context.Context, *CreateCalendarRequest) (*Calendar, error)
	UpdateCalendar(context.Context, *UpdateCalendarRequest) (*Calendar, error)
	DeleteCalendar(context.Context, *DeleteCalendarRequest) (*empty.Empty, error)
	ListCalendarStats(context.Context, *ListCalendarStatsRequest) (*ListCalendarStatsResponse, error)
	ListEntries(context.Context, *ListEntriesRequest) (*ListEntriesResponse, error)
	CreateEntry(context.Context, *CreateEntryRequest) (*Entry, error)
	UpdateEntry(context.Context, *UpdateEntryRequest) (*Entry, error)
	DeleteEntry(context.Context, *DeleteEntryRequest) (*empty.Empty, error)
	SignIn(context.Context, *SignInRequest) (*User, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterAdventarServer(s *grpc.Server, srv AdventarServer) {
	s.RegisterService(&_Adventar_serviceDesc, srv)
}

func _Adventar_ListCalendars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCalendarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).ListCalendars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/ListCalendars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).ListCalendars(ctx, req.(*ListCalendarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_GetCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).GetCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/GetCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).GetCalendar(ctx, req.(*GetCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_CreateCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).CreateCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/CreateCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).CreateCalendar(ctx, req.(*CreateCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_UpdateCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).UpdateCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/UpdateCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).UpdateCalendar(ctx, req.(*UpdateCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_DeleteCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).DeleteCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/DeleteCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).DeleteCalendar(ctx, req.(*DeleteCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_ListCalendarStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCalendarStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).ListCalendarStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/ListCalendarStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).ListCalendarStats(ctx, req.(*ListCalendarStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_ListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).ListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/ListEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).ListEntries(ctx, req.(*ListEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_CreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).CreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/CreateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).CreateEntry(ctx, req.(*CreateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).UpdateEntry(ctx, req.(*UpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_DeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).DeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/DeleteEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).DeleteEntry(ctx, req.(*DeleteEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adventar_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventarServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventar.v1.Adventar/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventarServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Adventar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "adventar.v1.Adventar",
	HandlerType: (*AdventarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCalendars",
			Handler:    _Adventar_ListCalendars_Handler,
		},
		{
			MethodName: "GetCalendar",
			Handler:    _Adventar_GetCalendar_Handler,
		},
		{
			MethodName: "CreateCalendar",
			Handler:    _Adventar_CreateCalendar_Handler,
		},
		{
			MethodName: "UpdateCalendar",
			Handler:    _Adventar_UpdateCalendar_Handler,
		},
		{
			MethodName: "DeleteCalendar",
			Handler:    _Adventar_DeleteCalendar_Handler,
		},
		{
			MethodName: "ListCalendarStats",
			Handler:    _Adventar_ListCalendarStats_Handler,
		},
		{
			MethodName: "ListEntries",
			Handler:    _Adventar_ListEntries_Handler,
		},
		{
			MethodName: "CreateEntry",
			Handler:    _Adventar_CreateEntry_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _Adventar_UpdateEntry_Handler,
		},
		{
			MethodName: "DeleteEntry",
			Handler:    _Adventar_DeleteEntry_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Adventar_SignIn_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Adventar_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Adventar_UpdateUser_Handler,
		},
		{
			MethodName: "HealthCheck",
			Handler:    _Adventar_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "adventar/v1/adventar.proto",
}
