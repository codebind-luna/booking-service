// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: booking/v1/booking_service.proto

package bookingv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TicketServiceClient is the client API for TicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketServiceClient interface {
	PurchaseTicket(ctx context.Context, in *PurchaseTicketRequest, opts ...grpc.CallOption) (*PurchaseTicketResponse, error)
	GetReceipt(ctx context.Context, in *GetReceiptRequest, opts ...grpc.CallOption) (*GetReceiptResponse, error)
	ViewSeatMap(ctx context.Context, in *ViewSeatMapRequest, opts ...grpc.CallOption) (*ViewSeatMapResponse, error)
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error)
	ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*ModifySeatResponse, error)
}

type ticketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketServiceClient(cc grpc.ClientConnInterface) TicketServiceClient {
	return &ticketServiceClient{cc}
}

func (c *ticketServiceClient) PurchaseTicket(ctx context.Context, in *PurchaseTicketRequest, opts ...grpc.CallOption) (*PurchaseTicketResponse, error) {
	out := new(PurchaseTicketResponse)
	err := c.cc.Invoke(ctx, "/booking.v1.TicketService/PurchaseTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) GetReceipt(ctx context.Context, in *GetReceiptRequest, opts ...grpc.CallOption) (*GetReceiptResponse, error) {
	out := new(GetReceiptResponse)
	err := c.cc.Invoke(ctx, "/booking.v1.TicketService/GetReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ViewSeatMap(ctx context.Context, in *ViewSeatMapRequest, opts ...grpc.CallOption) (*ViewSeatMapResponse, error) {
	out := new(ViewSeatMapResponse)
	err := c.cc.Invoke(ctx, "/booking.v1.TicketService/ViewSeatMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error) {
	out := new(RemoveUserResponse)
	err := c.cc.Invoke(ctx, "/booking.v1.TicketService/RemoveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*ModifySeatResponse, error) {
	out := new(ModifySeatResponse)
	err := c.cc.Invoke(ctx, "/booking.v1.TicketService/ModifySeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketServiceServer is the server API for TicketService service.
// All implementations must embed UnimplementedTicketServiceServer
// for forward compatibility
type TicketServiceServer interface {
	PurchaseTicket(context.Context, *PurchaseTicketRequest) (*PurchaseTicketResponse, error)
	GetReceipt(context.Context, *GetReceiptRequest) (*GetReceiptResponse, error)
	ViewSeatMap(context.Context, *ViewSeatMapRequest) (*ViewSeatMapResponse, error)
	RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error)
	ModifySeat(context.Context, *ModifySeatRequest) (*ModifySeatResponse, error)
	mustEmbedUnimplementedTicketServiceServer()
}

// UnimplementedTicketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTicketServiceServer struct {
}

func (UnimplementedTicketServiceServer) PurchaseTicket(context.Context, *PurchaseTicketRequest) (*PurchaseTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseTicket not implemented")
}
func (UnimplementedTicketServiceServer) GetReceipt(context.Context, *GetReceiptRequest) (*GetReceiptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedTicketServiceServer) ViewSeatMap(context.Context, *ViewSeatMapRequest) (*ViewSeatMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewSeatMap not implemented")
}
func (UnimplementedTicketServiceServer) RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTicketServiceServer) ModifySeat(context.Context, *ModifySeatRequest) (*ModifySeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeat not implemented")
}
func (UnimplementedTicketServiceServer) mustEmbedUnimplementedTicketServiceServer() {}

// UnsafeTicketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketServiceServer will
// result in compilation errors.
type UnsafeTicketServiceServer interface {
	mustEmbedUnimplementedTicketServiceServer()
}

func RegisterTicketServiceServer(s grpc.ServiceRegistrar, srv TicketServiceServer) {
	s.RegisterService(&TicketService_ServiceDesc, srv)
}

func _TicketService_PurchaseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).PurchaseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.v1.TicketService/PurchaseTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).PurchaseTicket(ctx, req.(*PurchaseTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.v1.TicketService/GetReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).GetReceipt(ctx, req.(*GetReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ViewSeatMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewSeatMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ViewSeatMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.v1.TicketService/ViewSeatMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ViewSeatMap(ctx, req.(*ViewSeatMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.v1.TicketService/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).RemoveUser(ctx, req.(*RemoveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ModifySeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ModifySeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.v1.TicketService/ModifySeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ModifySeat(ctx, req.(*ModifySeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketService_ServiceDesc is the grpc.ServiceDesc for TicketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.v1.TicketService",
	HandlerType: (*TicketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurchaseTicket",
			Handler:    _TicketService_PurchaseTicket_Handler,
		},
		{
			MethodName: "GetReceipt",
			Handler:    _TicketService_GetReceipt_Handler,
		},
		{
			MethodName: "ViewSeatMap",
			Handler:    _TicketService_ViewSeatMap_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TicketService_RemoveUser_Handler,
		},
		{
			MethodName: "ModifySeat",
			Handler:    _TicketService_ModifySeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking/v1/booking_service.proto",
}
