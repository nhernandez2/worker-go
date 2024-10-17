package api

import (
	"worker-go/internal/services"
)

type Services struct {
	OrderService   *services.OrderService
	ProductService *services.ProductService
	ClientService  *services.ClientService
}
