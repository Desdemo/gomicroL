package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net/http"

	gw "gomicroL/hello/proto/hello_http"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	endpoint := "127.0.0.1:8088"
	mux := runtime.NewServeMux()
	ops := []grpc.DialOption{grpc.WithInsecure()}

	err := gw.RegisterHelloHTTPHandlerFromEndpoint(ctx, mux, endpoint, ops)
	if err != nil {
		grpclog.Fatalf("Register handler err : %s", err)
	}

	grpclog.Println("Http Listen on 8084")
	http.ListenAndServe(":8084", mux)

}
