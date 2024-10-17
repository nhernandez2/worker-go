package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, ctrl Controllers) {
	r.HandleFunc("/health", HealthCheck).Methods("GET")

	r.HandleFunc("/clients", ctrl.ClientController.CreateClient).Methods("POST")
	r.HandleFunc("/clients/{id}", ctrl.ClientController.GetClient).Methods("GET")
	r.HandleFunc("/clients/{id}", ctrl.ClientController.UpdateClient).Methods("PUT")
	r.HandleFunc("/clients/{id}", ctrl.ClientController.DeleteClient).Methods("DELETE")

	r.HandleFunc("/products", ctrl.ProductController.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", ctrl.ProductController.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id}", ctrl.ProductController.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", ctrl.ProductController.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/products/validate-products", ctrl.ProductController.ValidateProducts).Methods("POST")

	r.HandleFunc("/orders", ctrl.OrderController.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", ctrl.OrderController.GetOrder).Methods("GET")
	r.HandleFunc("/orders/{id}", ctrl.OrderController.UpdateOrder).Methods("PUT")
	r.HandleFunc("/orders/{id}", ctrl.OrderController.DeleteOrder).Methods("DELETE")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status OK"))
}
