// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: mgmt/v1alpha1/transformer.proto

package mgmtv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// TransformersServiceName is the fully-qualified name of the TransformersService service.
	TransformersServiceName = "mgmt.v1alpha1.TransformersService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TransformersServiceGetTransformersProcedure is the fully-qualified name of the
	// TransformersService's GetTransformers RPC.
	TransformersServiceGetTransformersProcedure = "/mgmt.v1alpha1.TransformersService/GetTransformers"
)

// TransformersServiceClient is a client for the mgmt.v1alpha1.TransformersService service.
type TransformersServiceClient interface {
	GetTransformers(context.Context, *connect.Request[v1alpha1.GetTransformersRequest]) (*connect.Response[v1alpha1.GetTransformersResponse], error)
}

// NewTransformersServiceClient constructs a client for the mgmt.v1alpha1.TransformersService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTransformersServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TransformersServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &transformersServiceClient{
		getTransformers: connect.NewClient[v1alpha1.GetTransformersRequest, v1alpha1.GetTransformersResponse](
			httpClient,
			baseURL+TransformersServiceGetTransformersProcedure,
			opts...,
		),
	}
}

// transformersServiceClient implements TransformersServiceClient.
type transformersServiceClient struct {
	getTransformers *connect.Client[v1alpha1.GetTransformersRequest, v1alpha1.GetTransformersResponse]
}

// GetTransformers calls mgmt.v1alpha1.TransformersService.GetTransformers.
func (c *transformersServiceClient) GetTransformers(ctx context.Context, req *connect.Request[v1alpha1.GetTransformersRequest]) (*connect.Response[v1alpha1.GetTransformersResponse], error) {
	return c.getTransformers.CallUnary(ctx, req)
}

// TransformersServiceHandler is an implementation of the mgmt.v1alpha1.TransformersService service.
type TransformersServiceHandler interface {
	GetTransformers(context.Context, *connect.Request[v1alpha1.GetTransformersRequest]) (*connect.Response[v1alpha1.GetTransformersResponse], error)
}

// NewTransformersServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTransformersServiceHandler(svc TransformersServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	transformersServiceGetTransformersHandler := connect.NewUnaryHandler(
		TransformersServiceGetTransformersProcedure,
		svc.GetTransformers,
		opts...,
	)
	return "/mgmt.v1alpha1.TransformersService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TransformersServiceGetTransformersProcedure:
			transformersServiceGetTransformersHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTransformersServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTransformersServiceHandler struct{}

func (UnimplementedTransformersServiceHandler) GetTransformers(context.Context, *connect.Request[v1alpha1.GetTransformersRequest]) (*connect.Response[v1alpha1.GetTransformersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.TransformersService.GetTransformers is not implemented"))
}
