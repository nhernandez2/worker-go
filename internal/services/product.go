package services

import (
	"worker-go/internal/crud"
	"worker-go/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) CreateProduct(product models.Product) error {
	return crud.SaveProduct(product)
}

func (s *ProductService) GetProduct(id string) (*models.Product, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return crud.GetProduct(objectID)
}

func (s *ProductService) UpdateProduct(id string, product models.Product) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return crud.UpdateProduct(objectID, product)
}

func (s *ProductService) DeleteProduct(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return crud.DeleteProduct(objectID)
}

func (s *ProductService) ValidateProducts(productIDs []string) (bool, []string) {
	var notFoundProducts []string

	for _, productID := range productIDs {
		objectID, err := primitive.ObjectIDFromHex(productID)
		if err != nil {
			notFoundProducts = append(notFoundProducts, productID)
			continue
		}

		_, err = crud.GetProduct(objectID)
		if err != nil {
			notFoundProducts = append(notFoundProducts, productID)
		}
	}

	if len(notFoundProducts) > 0 {
		return false, notFoundProducts
	}

	return true, nil
}
