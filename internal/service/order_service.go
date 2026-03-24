package service

import (
	"fmt"

	"github.com/PranavJoshi2893/order-matching-engine/internal/model"
	"github.com/PranavJoshi2893/order-matching-engine/internal/repository"
	"github.com/google/uuid"
)

type OrderService struct {
	orderRepos *repository.OrderRepository
}

func NewOrderService(orderRepos *repository.OrderRepository) *OrderService {
	return &OrderService{
		orderRepos: orderRepos,
	}
}

func (s *OrderService) Create(data *model.CreateOrderRequest) (*model.CreateOrderResponse, error) {
	id, err := uuid.NewV7()

	if err != nil {
		return nil, fmt.Errorf("failed to generate id: %w", err)
	}

	if err := s.orderRepos.AddOrder(id, data); err != nil {
		return nil, fmt.Errorf("failed to store data: %v", err)
	}

	resp := &model.CreateOrderResponse{
		ID: id,
	}

	return resp, nil

}

func (s *OrderService) Cancel(id uuid.UUID) error {

	if err := s.orderRepos.CancelOrder(id); err != nil {
		return err
	}

	return nil
}
