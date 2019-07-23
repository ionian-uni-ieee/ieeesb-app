package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	ID      primitive.ObjectID `bson:"_id" json:"id"`
	From    string             `bson:"from" json:"from"`
	Message string             `bson:"message" json:"message"`
}
