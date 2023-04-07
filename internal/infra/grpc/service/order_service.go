package service

import (
	"context"
	pb2 "github.com/PGabrielDev/clean-archi-go/internal/infra/grpc/pb"
	usecase "github.com/PGabrielDev/clean-archi-go/internal/usecases"
)

type OrderService struct {
	pb2.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
}

func (o *OrderService) CreateOrder(ctx context.Context, createOrderRequest *pb2.CreateOrderRequest) (*pb2.CreateOrderResponse, error) {
	orderDTO := usecase.OrderInputDTO{ID: createOrderRequest.Id, Price: float64(createOrderRequest.Price), Tax: float64(createOrderRequest.Tax)}
	order, err := o.CreateOrderUseCase.Execute(orderDTO)
	if err != nil {
		return nil, err
	}
	return &pb2.CreateOrderResponse{
		Id:         order.ID,
		Price:      float32(order.Price),
		Tax:        float32(order.Tax),
		FinalPrice: float32(order.FinalPrice),
	}, nil
}
