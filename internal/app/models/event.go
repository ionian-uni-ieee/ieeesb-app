package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Event describes all the necessary information
// about what an IEEE event
type Event struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Tags        []string           `bson:"tags" json:"tags"`
	Type        string             `bson:"type" json:"type"`
	Sponsors    []Sponsor          `bson:"sponsors" json:"sponsors"`
	Logo        MediaMeta          `bson:"logo" json:"logo"`
	Media       []MediaMeta        `bson:"media" json:"media"`
}
