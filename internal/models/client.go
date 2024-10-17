package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name   string             `bson:"name" json:"name"`
	Active bool               `bson:"active" json:"active"`
}
