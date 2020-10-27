package service

import (
	"context"
	"gomicroL/grpc_demo/grpc_client/service"
)

type ProdService struct {
}

func (ps *ProdService) GetProductStock(ctx context.Context, request *service.ProductRequest) (*service.ProductResponse, error) {
	return &service.ProductResponse{ProdStock: request.ProdId}, nil
}
