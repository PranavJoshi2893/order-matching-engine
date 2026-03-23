package routes

import (
	"net/http"

	"github.com/PranavJoshi2893/order-matching-engine/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(orderHandler *handler.OrderHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/orders", func(r chi.Router) {
			r.Post("/", orderHandler.Create)
			r.Delete("/{id}", orderHandler.Cancel)
		})
	})

	return r
}
