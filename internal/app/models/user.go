package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User is a moderator user that has access to the control panel
type User struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Username    string             `bson:"username" json:"username"`
	Password    string             `bson:"password" json:"password"`
	Email       string             `bson:"email" json:"email"`
	Fullname    string             `bson:"fullname" json:"fullname"`
	Permissions Permissions        `bson:"permissions" json:"permissions"`
}
