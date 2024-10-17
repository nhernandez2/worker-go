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

const clients = "clients"

var clientCollection *mongo.Collection

func InitClientCollection(mongoClient *mongo.Client) {
	clientCollection = mongoClient.Database(config.MongoDB).Collection(clients)
}

func SaveClient(clientData models.Client) error {
	_, err := clientCollection.InsertOne(context.TODO(), clientData)
	if err != nil {
		log.Printf("Error al insertar cliente en MongoDB: %v", err)
		return err
	}
	return nil
}

func GetClient(id primitive.ObjectID) (models.Client, error) {
	var clientData models.Client
	err := clientCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&clientData)
	if err != nil {
		log.Printf("Error al obtener el cliente: %v", err)
		return clientData, err
	}
	return clientData, nil
}

func UpdateClient(id primitive.ObjectID, updatedData models.Client) error {
	_, err := clientCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": updatedData},
	)
	if err != nil {
		log.Printf("Error al actualizar el cliente: %v", err)
		return err
	}
	return nil
}

func DeleteClient(id primitive.ObjectID) error {
	_, err := clientCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Printf("Error al eliminar el cliente: %v", err)
		return err
	}
	return nil
}

func ValidateClient(clientID string) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(clientID)
	if err != nil {
		log.Printf("Error al convertir el ID: %v", err)
		return false, err
	}
	client, err := GetClient(objectID)
	if err != nil {
		return false, err
	}
	return client.Active, nil
}
