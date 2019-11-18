package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Sponsor describes an IEEE Sponsor
type Sponsor struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	Name   string             `bson:"name" json:"name"`
	Emails []string           `bson:"emails" json:"emails"`
	Phones []string           `bson:"phones" json:"phones"`
	Logo   MediaMeta          `bson:"logo" json:"logo"`
}

// GetID returns the hex of the sponsor's id
func (s *Sponsor) GetID() string {
	return s.ID.Hex()
}

// isEmpty returns true if the name is empty
func (s *Sponsor) isEmpty() bool {
	if s.Name == "" {
		return true
	}
	return false
}

// areEqual returns true if two sponsors are equal
// two sponsors are equals if their IDs are equal
func (firstSponsor *Sponsor) areEqual(secondSponsor Sponsor) bool {
	if firstSponsor.ID.Hex() == secondSponsor.GetID() {
		return true
	}
	return false
}
