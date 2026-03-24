package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PranavJoshi2893/order-matching-engine/internal/model"
	"github.com/PranavJoshi2893/order-matching-engine/internal/service"
	"github.com/google/uuid"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req model.CreateOrderRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		log.Println("decode error:", err)

		WriteJSON(w, http.StatusBadRequest, APIResponse{
			Message: "invalid data",
		})
		return
	}

	if decoder.More() {
		WriteJSON(w, http.StatusBadRequest, APIResponse{
			Message: "unexpected extra data",
		})
		return
	}

	if errs := req.Validate(); len(errs) > 0 {
		WriteJSON(w, http.StatusBadRequest, APIResponse{
			Message: "invalid data",
			Errors:  errs,
		})
		return
	}

	resp, err := h.orderService.Create(&req)
	if err != nil {
		log.Println("Internal server error:", err)

		WriteJSON(w, http.StatusInternalServerError, APIResponse{
			Message: "internal server error",
		})
	}

	WriteJSON(w, http.StatusCreated, APIResponse{
		Message: "Order Created",
		Data:    resp,
	})
}

func (h *OrderHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIResponse{
			Message: "invalid id",
		})
		return
	}

	if err := h.orderService.Cancel(id); err != nil {
		switch err {
		case model.ErrorNotFound:
			WriteJSON(w, http.StatusNotFound, APIResponse{
				Message: err.Error(),
			})
		default:
			log.Println("Internal server error:", err)
			WriteJSON(w, http.StatusInternalServerError, APIResponse{
				Message: "internal server error",
			})
		}
	}

	WriteJSON(w, http.StatusOK, APIResponse{
		Message: "Order Cancelled",
	})
}
