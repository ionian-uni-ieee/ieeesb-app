package models

// Permissions contains the moderator's permissions
// of a moderator user in the control panel
type Permissions struct {
	Users    bool `bson:"users" json:"users"`
	Events   bool `bson:"events" json:"events"`
	Tickets  bool `bson:"tickets" json:"tickets"`
	Sponsors bool `bson:"sponsors" json:"sponsors"`
}
