package models

type Permission struct {
	Users    bool `bson:"users" json:"users"`
	Events   bool `bson:"events" json:"events"`
	Tickets  bool `bson:"tickets" json:"tickets"`
	Sponsors bool `bson:"sponsors" json:"sponsors"`
}
