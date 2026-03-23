# Order Matching Engine

A high-performance order matching engine microservice built in Go, implementing price-time priority matching with bid/ask order book management.

---

## Prerequisites

- Go 1.22+
- Make

---

## Setup

1. Clone the repository

```bash
git clone https://github.com/PranavJoshi2893/order-matching-engine.git
cd order-matching-engine
```

2. Create a `.env` file in the root directory

```bash
cp .env.example .env
```

3. Fill in the environment variables

```env
PORT=<PORT NUMBER>
```

---

## Run

```bash
make run
```

---

## Build

```bash
make build
```

---

## API

### Place Order

```
POST /orders
```

Request body:
```json
{
  "side": "BUY",
  "price": 100.00,
  "qty": 10
}
```

Response:
```json
{
  "message": "Order Created",
  "data": null
}
```

### Cancel Order

```
DELETE /orders/{id}
```

Response:
```json
{
  "message": "Order Cancelled",
  "data": null
}
```

### Validation Errors

```json
{
  "errors": [
    "invalid side",
    "invalid price",
    "invalid qty"
  ]
}
```

---

## Project Structure

```
order-matching-engine/
├── internal/
│   ├── handler/       # HTTP handlers
│   └── model/         # Order, OrderBook types
├── .env.example
├── Makefile
└── main.go
```