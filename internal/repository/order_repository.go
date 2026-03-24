package repository

import (
	"sort"
	"sync"
	"time"

	"github.com/PranavJoshi2893/order-matching-engine/internal/model"
	"github.com/google/uuid"
)

type OrderRepository struct {
	mu         sync.RWMutex
	orders     map[uuid.UUID]*model.Order
	sortedBuy  []*model.Order
	sortedSell []*model.Order
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
	r.rebuildSorted()
	return nil
}

func (r *OrderRepository) CancelOrder(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.orders[id]; !ok {
		return model.ErrorNotFound
	}

	delete(r.orders, id)
	r.rebuildSorted()
	return nil
}

func (r *OrderRepository) rebuildSorted() {
	buy := make([]*model.Order, 0)
	sell := make([]*model.Order, 0)

	for _, o := range r.orders {
		if o.Side == model.Buy {
			buy = append(buy, o)
		} else {
			sell = append(sell, o)
		}
	}

	sort.Slice(buy, func(i, j int) bool {
		if buy[i].Price != buy[j].Price {
			return buy[i].Price > buy[j].Price
		}
		return buy[i].CreatedAt.Before(buy[j].CreatedAt)
	})

	sort.Slice(sell, func(i, j int) bool {
		if sell[i].Price != sell[j].Price {
			return sell[i].Price < sell[j].Price
		}
		return sell[i].CreatedAt.Before(sell[j].CreatedAt)
	})

	r.sortedBuy = buy
	r.sortedSell = sell
}
