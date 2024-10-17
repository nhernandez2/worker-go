package controllers

import (
	"encoding/json"
	"net/http"
	"worker-go/internal/models"
	"worker-go/internal/services"

	"github.com/gorilla/mux"
)

type OrderController struct {
	orderService *services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := c.orderService.CreateOrder(order); err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *OrderController) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	order, err := c.orderService.GetOrder(id)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func (c *OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := c.orderService.UpdateOrder(id, order); err != nil {
		http.Error(w, "Failed to update order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *OrderController) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := c.orderService.DeleteOrder(id); err != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
