package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"worker-go/internal/models"
	"worker-go/pkg/config"

	"github.com/cenkalti/backoff/v4"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const urlApi = "http://localhost:8080"

type RetryFunction func() error

func getClientRequest(clientId string) (models.Client, error) {
	client := resty.New()
	clientModel := &models.Client{}
	resp, err := client.R().
		SetResult(clientModel).
		Get(fmt.Sprintf("%s/%s/%s", urlApi, "clients", clientId))

	if err != nil {
		return *clientModel, err
	}

	if !resp.IsSuccess() {
		return *clientModel, fmt.Errorf("status code: %d", resp.StatusCode())
	}

	return *clientModel, nil
}

func getProductRequest(productIDs []string) (models.ValidateProductsResponse, error) {
	client := resty.New()
	validateProducts := &models.ValidateProductsResponse{}
	validateRequest := models.ValidateProductsRequest{
		ProductIDs: productIDs,
	}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(validateRequest).
		SetResult(&validateProducts).
		Post(fmt.Sprintf("%s/%s/%s", urlApi, "products", "validate-products"))

	if err != nil {
		return *validateProducts, err
	}

	if !resp.IsSuccess() {
		return *validateProducts, fmt.Errorf("status code: %d", resp.StatusCode())
	}

	return *validateProducts, nil
}

func getOrderRequest(order models.Order) (models.Order, error) {
	client := resty.New()
	newOrder := &models.Order{}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(order).
		SetResult(&newOrder).
		Post(fmt.Sprintf("%s/%s", urlApi, "orders"))

	if err != nil {
		return *newOrder, err
	}

	if !resp.IsSuccess() {
		return *newOrder, fmt.Errorf("status code: %d", resp.StatusCode())
	}

	return *newOrder, nil
}

func retryWithExponentialBackUp(fn RetryFunction) error {
	retryCount := 0
	maxRetries := 3

	operation := func() error {
		retryCount++
		err := fn()

		if retryCount >= maxRetries {
			return backoff.Permanent(err)
		}

		return err
	}

	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.Multiplier = 2.0
	expBackoff.InitialInterval = 5 * time.Second
	expBackoff.MaxElapsedTime = 30 * time.Second

	err := backoff.Retry(operation, expBackoff)
	if err != nil {
		return err
	}

	return nil
}

func ConsumeKafka() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.KafkaBroker,
		"group.id":          "my-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("Error al crear el consumidor de Kafka: %v", err)
	}

	consumer.Subscribe(config.KafkaTopic, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {

			var orderMessage models.OrderMessage
			err := json.Unmarshal(msg.Value, &orderMessage)
			if err != nil {
				log.Printf("Error al deserializar el mensaje: %v", err)
				continue
			}

			var clientModel models.Client
			var validateProducts models.ValidateProductsResponse
			var orderModel models.Order

			err = retryWithExponentialBackUp(func() error {
				clientModel, err = getClientRequest(orderMessage.ClientID)
				return err
			})

			if err != nil {
				fmt.Println("Error en llamada Cliente")
				fmt.Println(err)
				continue
			}

			productIDs := make([]string, len(orderMessage.Products))
			for i, product := range orderMessage.Products {
				productIDs[i] = product.ProductID
			}

			err = retryWithExponentialBackUp(func() error {
				validateProducts, err = getProductRequest(productIDs)
				return err
			})

			if err != nil {
				fmt.Println("Error en llamada Productos")
				fmt.Println(err)
				continue
			}

			if !validateProducts.IsValid {
				log.Printf("Productos no encontrados: %v", validateProducts.NotFoundItems)
				continue
			}

			var products []models.Product
			for _, prodDetail := range orderMessage.Products {
				productID, _ := primitive.ObjectIDFromHex(prodDetail.ProductID)
				product := models.Product{
					ID:    productID,
					Name:  prodDetail.Name,
					Price: prodDetail.Price,
				}
				products = append(products, product)
			}

			order := models.Order{
				OrderId:    orderMessage.OrderID,
				CustomerId: clientModel.ID,
				Products:   products,
			}

			err = retryWithExponentialBackUp(func() error {
				orderModel, err = getOrderRequest(order)
				return err
			})

			if err != nil {
				fmt.Println("Error en llamada Orden")
				fmt.Println(err)
				continue
			}

			log.Printf("Orden guardada exitosamente: %+v", orderModel)
		} else {
			log.Printf("Error al recibir mensaje: %v\n", err)
		}
	}
}
