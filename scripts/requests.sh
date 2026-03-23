#!/usr/bin/bash

curl -s -X POST http://localhost:3000/api/v1/orders/ -H "Content-Type: application/json" -d '{"side":"BUY","price":10.11,"qty":12}' | jq

ID=$(curl -s -X POST http://localhost:3000/api/v1/orders/ -H "Content-Type: application/json" -d '{"side":"SELL","price":10.11,"qty":12}' | jq .data.id)

curl -s -X DELETE http://localhost:3000/api/v1/orders/$ID | jq