package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sponsor struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	Name   string             `bson:"name" json:"name"`
	Emails []string           `bson:"emails" json:"emails"`
	Phones []string           `bson:"phones" json:"phones"`
	Logo   MediaMeta          `bson:"logo" json:"logo"`
}
