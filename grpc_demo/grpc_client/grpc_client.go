package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials"

	"gomicroL/grpc_demo/grpc_client/service"
	"google.golang.org/grpc"
	"log"
)

func main() {
	tls, err := credentials.NewClientTLSFromFile("server.crt", "localhost")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := grpc.Dial(":8087", grpc.WithTransportCredentials(tls))
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	productServiceClient := service.NewProdServiceClient(conn)

	resp, err := productServiceClient.GetProductStock(context.Background(), &service.ProductRequest{ProdId: 2333})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("调用grpc方法成功, pordStock=", resp.ProdStock)
}
