package service

import (
	"context"
	"github.com/PGabrielDev/clean-archi-go/internal/infra/grpc/internal/infra/grpc/pb"
	usecase "github.com/PGabrielDev/clean-archi-go/internal/usecases"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
}

func (o *OrderService) CreateOrder(ctx context.Context, createOrderRequest *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	orderDTO := usecase.OrderInputDTO{ID: createOrderRequest.Id, Price: float64(createOrderRequest.Price), Tax: float64(createOrderRequest.Tax)}
	order, err := o.CreateOrderUseCase.Execute(orderDTO)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         order.ID,
		Price:      float32(order.Price),
		Tax:        float32(order.Tax),
		FinalPrice: float32(order.FinalPrice),
	}, nil
}
