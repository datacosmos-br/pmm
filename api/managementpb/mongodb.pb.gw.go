// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: managementpb/mongodb.proto

/*
Package managementpb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package managementpb

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var (
	_ codes.Code
	_ io.Reader
	_ status.Status
	_ = runtime.String
	_ = utilities.NewDoubleArray
	_ = metadata.Join
)

func request_MongoDB_AddMongoDB_0(ctx context.Context, marshaler runtime.Marshaler, client MongoDBClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddMongoDBRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.AddMongoDB(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_MongoDB_AddMongoDB_0(ctx context.Context, marshaler runtime.Marshaler, server MongoDBServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddMongoDBRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.AddMongoDB(ctx, &protoReq)
	return msg, metadata, err
}

// RegisterMongoDBHandlerServer registers the http handlers for service MongoDB to "mux".
// UnaryRPC     :call MongoDBServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterMongoDBHandlerFromEndpoint instead.
func RegisterMongoDBHandlerServer(ctx context.Context, mux *runtime.ServeMux, server MongoDBServer) error {
	mux.Handle("POST", pattern_MongoDB_AddMongoDB_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/management.MongoDB/AddMongoDB", runtime.WithHTTPPathPattern("/v1/management/MongoDB/Add"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_MongoDB_AddMongoDB_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MongoDB_AddMongoDB_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	return nil
}

// RegisterMongoDBHandlerFromEndpoint is same as RegisterMongoDBHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterMongoDBHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterMongoDBHandler(ctx, mux, conn)
}

// RegisterMongoDBHandler registers the http handlers for service MongoDB to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterMongoDBHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterMongoDBHandlerClient(ctx, mux, NewMongoDBClient(conn))
}

// RegisterMongoDBHandlerClient registers the http handlers for service MongoDB
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "MongoDBClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "MongoDBClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "MongoDBClient" to call the correct interceptors.
func RegisterMongoDBHandlerClient(ctx context.Context, mux *runtime.ServeMux, client MongoDBClient) error {
	mux.Handle("POST", pattern_MongoDB_AddMongoDB_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/management.MongoDB/AddMongoDB", runtime.WithHTTPPathPattern("/v1/management/MongoDB/Add"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_MongoDB_AddMongoDB_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MongoDB_AddMongoDB_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	return nil
}

var pattern_MongoDB_AddMongoDB_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v1", "management", "MongoDB", "Add"}, ""))

var forward_MongoDB_AddMongoDB_0 = runtime.ForwardResponseMessage
