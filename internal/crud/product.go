package crud

import (
	"context"
	"log"
	"worker-go/internal/models"
	"worker-go/pkg/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const products = "products"

var productCollection *mongo.Collection

func InitProductCollection(client *mongo.Client) {
	productCollection = client.Database(config.MongoDB).Collection(products)
}

func SaveProduct(product models.Product) error {
	_, err := productCollection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Printf("Error al insertar producto en MongoDB: %v", err)
		return err
	}
	return nil
}

func GetProduct(id primitive.ObjectID) (*models.Product, error) {
	var product models.Product
	err := productCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&product)
	if err != nil {
		log.Printf("Error al obtener el producto: %v", err)
		return nil, err
	}
	return &product, nil
}

func UpdateProduct(id primitive.ObjectID, updatedData models.Product) error {
	_, err := productCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": updatedData},
	)
	if err != nil {
		log.Printf("Error al actualizar el producto: %v", err)
		return err
	}
	return nil
}

func DeleteProduct(id primitive.ObjectID) error {
	_, err := productCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Printf("Error al eliminar el producto: %v", err)
		return err
	}
	return nil
}
