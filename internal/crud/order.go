package crud

import (
	"context"
	"errors"
	"log"
	"worker-go/internal/models"
	"worker-go/pkg/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const orders = "orders"

var orderCollection *mongo.Collection

func InitOrderCollection(client *mongo.Client) {
	orderCollection = client.Database(config.MongoDB).Collection(orders)
}

func SaveOrder(order models.Order) error {
	exists, err := OrderExists(order.OrderId)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("el order_id ya existe")
	}

	_, err = orderCollection.InsertOne(context.TODO(), order)
	if err != nil {
		log.Printf("Error al insertar orden en MongoDB: %v", err)
		return err
	}
	return nil
}

func GetOrder(id primitive.ObjectID) (*models.Order, error) {
	var order models.Order
	err := orderCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&order)
	if err != nil {
		log.Printf("Error al obtener la orden: %v", err)
		return nil, err
	}
	return &order, nil
}

func UpdateOrder(id primitive.ObjectID, updatedData models.Order) error {
	_, err := orderCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": updatedData},
	)
	if err != nil {
		log.Printf("Error al actualizar la orden: %v", err)
		return err
	}
	return nil
}

func DeleteOrder(id primitive.ObjectID) error {
	_, err := orderCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Printf("Error al eliminar la orden: %v", err)
		return err
	}
	return nil
}

func OrderExists(orderID string) (bool, error) {
	filter := bson.M{"order_id": orderID}

	var existingOrder models.Order
	err := orderCollection.FindOne(context.TODO(), filter).Decode(&existingOrder)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
