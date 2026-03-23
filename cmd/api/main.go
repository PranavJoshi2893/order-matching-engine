package main

import (
	"log"

	"github.com/PranavJoshi2893/order-matching-engine/internal/config"
	"github.com/PranavJoshi2893/order-matching-engine/internal/handler"
	"github.com/PranavJoshi2893/order-matching-engine/internal/repository"
	"github.com/PranavJoshi2893/order-matching-engine/internal/routes"
	"github.com/PranavJoshi2893/order-matching-engine/internal/server"
	"github.com/PranavJoshi2893/order-matching-engine/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	orderRepos := repository.NewOrderRepo()
	orderService := service.NewOrderService(orderRepos)
	orderHandler := handler.NewOrderHandler(orderService)

	r := routes.Routes(orderHandler)

	srv := server.NewServer(cfg, r)

	log.Println("Server is running on port", cfg.Port)
	if err := srv.Run(); err != nil {
		log.Fatalln(err)
	}

}
