package services

import (
	"log"
	"worker-go/internal/crud"
	"worker-go/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientService struct{}

func NewClientService() *ClientService {
	return &ClientService{}
}

func (s *ClientService) CreateClient(client models.Client) error {
	return crud.SaveClient(client)
}

func (s *ClientService) GetClient(id string) (models.Client, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Client{}, err
	}
	return crud.GetClient(objectID)
}

func (s *ClientService) UpdateClient(id string, updatedData models.Client) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return crud.UpdateClient(objectID, updatedData)
}

func (s *ClientService) DeleteClient(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return crud.DeleteClient(objectID)
}

func (s *ClientService) ValidateClient(id string) (bool, error) {
	clientID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error al convertir el ID: %v", err)
		return false, err
	}
	client, err := crud.GetClient(clientID)
	if err != nil {
		return false, err
	}
	return client.Active, nil
}
