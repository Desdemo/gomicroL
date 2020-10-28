package main

import (
	"fmt"
	se "gomicroL/grpc_demo/grpc_client/service"
	"gomicroL/grpc_demo/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net/http"
)

func main() {
	// 引用证书
	tls, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatalln(err)
	}
	// 加入证书
	rpcServer := grpc.NewServer(grpc.Creds(tls))

	// 服务注册
	se.RegisterProdServiceServer(rpcServer, new(service.ProdService))
	// 路由
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request)
		rpcServer.ServeHTTP(writer, request)
	})
	// server

	httpServer := http.Server{Addr: ":8087", Handler: mux}

	_ = httpServer.ListenAndServeTLS("server.crt", "server.key")
}
