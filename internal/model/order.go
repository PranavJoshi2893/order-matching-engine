package model

import (
	"time"

	"github.com/google/uuid"
)

type Side string

const (
	Buy  Side = "BUY"
	Sell Side = "SELL"
)

type Status string

const (
	Open      Status = "OPEN"
	Partial   Status = "PARTIAL"
	Filled    Status = "FILLED"
	Cancelled Status = "CANCELLED"
)

type Order struct {
	ID           uuid.UUID
	Side         Side
	Price        float64
	Qty          int
	RemainingQty int
	Status       Status
	CreatedAt    time.Time
}

type OrderBook struct {
	Bids []*Order
	Asks []*Order
}

type CreateOrderRequest struct {
	Side  Side    `json:"side"`
	Price float64 `json:"price"`
	Qty   int     `json:"qty"`
}

type CreateOrderResponse struct {
	ID uuid.UUID `json:"id"`
}

func (m *CreateOrderRequest) Validate() []string {
	var errs []string

	if m.Side != "BUY" && m.Side != "SELL" {
		errs = append(errs, "invalid order type")
	}

	if m.Price <= 0 {
		errs = append(errs, "invalid order price")
	}

	if m.Qty <= 0 {
		errs = append(errs, "invalid order quantity")
	}

	return errs
}
