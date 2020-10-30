package main

import (
	"context"
	pb "gomicroL/hello/proto/hello_http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const Address = "127.0.0.1:50052"

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	c := pb.NewHelloHTTPClient(conn)

	req := &pb.HelloHTTPRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}
	grpclog.Println(res.Message)

}
