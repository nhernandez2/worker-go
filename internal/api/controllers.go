package api

import (
	"worker-go/internal/controllers"
)

type Controllers struct {
	ClientController  *controllers.ClientController
	ProductController *controllers.ProductController
	OrderController   *controllers.OrderController
}
