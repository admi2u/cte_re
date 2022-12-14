// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: datasets.proto

package datasets

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

// DatasetsClient is the client API for Datasets service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatasetsClient interface {
	DatasetCreate(ctx context.Context, in *DatasetCreateReq, opts ...grpc.CallOption) (*DatasetCreateResp, error)
	DatasetList(ctx context.Context, in *DatasetListReq, opts ...grpc.CallOption) (*DatasetListResp, error)
}

type datasetsClient struct {
	cc grpc.ClientConnInterface
}

func NewDatasetsClient(cc grpc.ClientConnInterface) DatasetsClient {
	return &datasetsClient{cc}
}

func (c *datasetsClient) DatasetCreate(ctx context.Context, in *DatasetCreateReq, opts ...grpc.CallOption) (*DatasetCreateResp, error) {
	out := new(DatasetCreateResp)
	err := c.cc.Invoke(ctx, "/datasets.datasets/datasetCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetsClient) DatasetList(ctx context.Context, in *DatasetListReq, opts ...grpc.CallOption) (*DatasetListResp, error) {
	out := new(DatasetListResp)
	err := c.cc.Invoke(ctx, "/datasets.datasets/datasetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasetsServer is the server API for Datasets service.
// All implementations must embed UnimplementedDatasetsServer
// for forward compatibility
type DatasetsServer interface {
	DatasetCreate(context.Context, *DatasetCreateReq) (*DatasetCreateResp, error)
	DatasetList(context.Context, *DatasetListReq) (*DatasetListResp, error)
	mustEmbedUnimplementedDatasetsServer()
}

// UnimplementedDatasetsServer must be embedded to have forward compatible implementations.
type UnimplementedDatasetsServer struct {
}

func (UnimplementedDatasetsServer) DatasetCreate(context.Context, *DatasetCreateReq) (*DatasetCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatasetCreate not implemented")
}
func (UnimplementedDatasetsServer) DatasetList(context.Context, *DatasetListReq) (*DatasetListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatasetList not implemented")
}
func (UnimplementedDatasetsServer) mustEmbedUnimplementedDatasetsServer() {}

// UnsafeDatasetsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatasetsServer will
// result in compilation errors.
type UnsafeDatasetsServer interface {
	mustEmbedUnimplementedDatasetsServer()
}

func RegisterDatasetsServer(s grpc.ServiceRegistrar, srv DatasetsServer) {
	s.RegisterService(&Datasets_ServiceDesc, srv)
}

func _Datasets_DatasetCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatasetCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetsServer).DatasetCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datasets.datasets/datasetCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetsServer).DatasetCreate(ctx, req.(*DatasetCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datasets_DatasetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatasetListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetsServer).DatasetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datasets.datasets/datasetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetsServer).DatasetList(ctx, req.(*DatasetListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Datasets_ServiceDesc is the grpc.ServiceDesc for Datasets service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Datasets_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "datasets.datasets",
	HandlerType: (*DatasetsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "datasetCreate",
			Handler:    _Datasets_DatasetCreate_Handler,
		},
		{
			MethodName: "datasetList",
			Handler:    _Datasets_DatasetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datasets.proto",
}
