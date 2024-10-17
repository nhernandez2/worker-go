package main

import (
	"log"
	"net/http"
	"worker-go/internal/api"
	"worker-go/internal/controllers"
	"worker-go/internal/crud"
	"worker-go/internal/services"
	"worker-go/pkg/config"
	database "worker-go/pkg/mongo"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()

	mongoClient := database.ConnectMongoDB()

	crud.InitClientCollection(mongoClient)
	crud.InitProductCollection(mongoClient)
	crud.InitOrderCollection(mongoClient)

	clientService := services.NewClientService()
	orderService := services.NewOrderService()
	productService := services.NewProductService()

	ctrls := api.Controllers{
		ClientController:  controllers.NewClientController(clientService),
		ProductController: controllers.NewProductController(productService),
		OrderController:   controllers.NewOrderController(orderService),
	}

	r := mux.NewRouter()
	api.SetupRoutes(r, ctrls)

	port := config.Port
	log.Println("Iniciando servidor en " + port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
