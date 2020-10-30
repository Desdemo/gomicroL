package main

import (
	"context"
	"fmt"

	pb "gomicroL/hello/proto/hello_http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

const Address = "127.0.0.1:8088"

type helloService struct {
}

var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloHTTPRequest) (*pb.HelloHTTPResponse, error) {
	resp := new(pb.HelloHTTPResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterHelloHTTPServer(s, HelloService)
	grpclog.Println("listen on " + Address)

	_ = s.Serve(listen)

}
