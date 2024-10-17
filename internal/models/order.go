package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CustomerId primitive.ObjectID `bson:"customer_id,omitempty" json:"customer_id,omitempty"`
	OrderId    string             `bson:"order_id,omitempty" json:"order_id,omitempty"`
	Products   []Product          `bson:"products,omitempty" json:"products,omitempty"`
}
