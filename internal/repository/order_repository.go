package repository

import (
	"sync"
	"time"

	"github.com/PranavJoshi2893/order-matching-engine/internal/model"
	"github.com/google/uuid"
)

type OrderRepository struct {
	mu     sync.RWMutex
	orders map[uuid.UUID]*model.Order
}

func NewOrderRepo() *OrderRepository {
	return &OrderRepository{
		orders: make(map[uuid.UUID]*model.Order),
	}
}

func (r *OrderRepository) AddOrder(id uuid.UUID, data *model.CreateOrderRequest) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	order := &model.Order{
		ID:           id,
		Side:         data.Side,
		Price:        data.Price,
		Qty:          data.Qty,
		RemainingQty: data.Qty,
		Status:       model.Open,
		CreatedAt:    time.Now(),
	}

	r.orders[id] = order

	return nil
}

func (r *OrderRepository) CancelOrder(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.orders[id]; !ok {
		return model.ErrorNotFound
	}

	delete(r.orders, id)

	return nil
}
