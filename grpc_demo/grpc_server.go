package main

import (
	service2 "gomicroL/grpc_demo/grpc_client/service"
	"gomicroL/grpc_demo/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func main() {
	tls, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatalln(err)
	}

	rpcServer := grpc.NewServer(grpc.Creds(tls))

	service2.RegisterProdServiceServer(rpcServer, new(service.ProdService))

	listener, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatalln(err)
	}

	_ = rpcServer.Serve(listener)

}
