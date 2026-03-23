run:
	go run cmd/api/main.go || true

build:
	go build -o order-matching-engine cmd/api/main.go