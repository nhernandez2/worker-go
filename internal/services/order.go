package services

import (
	"worker-go/internal/crud"
	"worker-go/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(order models.Order) error {
	return crud.SaveOrder(order)
}

func (s *OrderService) GetOrder(id string) (*models.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return crud.GetOrder(objectID)
}

func (s *OrderService) UpdateOrder(id string, order models.Order) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return crud.UpdateOrder(objectID, order)
}

func (s *OrderService) DeleteOrder(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return crud.DeleteOrder(objectID)
}
