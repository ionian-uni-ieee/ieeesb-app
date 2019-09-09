package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Ticket is an instance of a "contact us" message
type Ticket struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Email    string             `bson:"email" json:"email"`
	Fullname string             `bson:"fullname" json:"fullname"`
	Message  string             `bson:"message" json:"message"`
	State    string             `bson:"state" json:"state"`
}
