package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Session contains information for a login session
// of a user in the control panel
type Session struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Username string             `bson:"username" json:"username"`
	Expires  int64              `bson:"expires" json:"expires"`
}
