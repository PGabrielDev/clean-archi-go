package usecases

import (
	"github.com/PGabrielDev/clean-archi-go/internal/entity"
	events2 "github.com/PGabrielDev/clean-archi-go/internal/events"
	"github.com/PGabrielDev/clean-archi-go/pkg/events"
)

type OrderInputDTO struct {
	ID    string  `json:"ID"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutputDTO struct {
	ID         string  `json:"ID"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"finalPrice"`
}
type CalculateFinalPriceUseCase struct {
	Repository     entity.OrderRepositoryInterface
	OrderEvent     events2.OrderCreatedEvent
	EventDispacher events.EventDispatcherInterface
}

func NewCalculateFinalPriceUseCase(repository entity.OrderRepositoryInterface, event events2.OrderCreatedEvent, eventD events.EventDispatcherInterface) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{
		Repository:     repository,
		OrderEvent:     event,
		EventDispacher: eventD,
	}
}

func (c *CalculateFinalPriceUseCase) Execute(dto OrderInputDTO) (OrderOutputDTO, error) {
	panic("")
}
