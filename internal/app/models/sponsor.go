package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import "errors"

// Sponsor describes an IEEE Sponsor
type Sponsor struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	Name   string             `bson:"name" json:"name"`
	Emails []string           `bson:"emails" json:"emails"`
	Phones []string           `bson:"phones" json:"phones"`
	Logo   MediaMeta          `bson:"logo" json:"logo"`
}

// NewSponsor is a sponsor factory that generates a sponsor
func NewSponsor(
	name string,
	emails []string,
	phones []string,
	logo MediaMeta,
) (*Sponsor) {
	return &Sponsor{
		ID:			primitive.NewObjectID(),
		Name:		name,
		Emails:		emails,
		Phones:		phones,
		Logo:		logo,
	}
}

// GetObjectID returns the object id of the sponsor
func (s *Sponsor) GetObjectID() primitive.ObjectID {
	return s.ID
}

// GetName returns the sponsor's name
func (s *Sponsor) GetName() string {
	return s.Name
}

// SetName changes the sponsor's name
func (s *Sponsor) SetName(newName string) error {
	if newName == "" {
		return errors.New("Sponsor's name cant't be empty")
	}
	s.Name = newName
	return nil
}

// GetEmails returns the sponsor's emails
func (s *Sponsor) GetEmail() []string {
	return s.Emails
}

// SetEmails changes the sponsor's emails
func (s *Sponsor) SetEmails(newEmails []string) error {
	for _, email := range newEmails {
		if email == "" {
			return errors.New("Sponsor's email can't be empty")
		}
	}
	s.Emails = newEmails
	return nil
}

// AddEmail adds email at the bottom of the sponsor's emails
func (s *Sponsor) AddEmail(newEmail string) error {
	if newEmail == "" {
		return errors.New("Sponsor's email can't be empty")
	}
	s.Emails = append(s.Emails,newEmail)
	return nil
}

// RemoveEmail removes an email (by the name)
func (s *Sponsor) RemoveEmail(oddEmail string) error {
	for i, existingEmail := range s.Emails {
		if existingEmail == oddEmail {
			s.Emails = append(s.Emails[:i], s.Emails[i+1:]...)
			return nil
		}
	}
	return errors.New("Sponsor's email was not found. Email couldn't be deleted")
}

// GetPhones returns the sponsor's phones
func (s *Sponsor) GetPhones() []string {
	return s.Phones
}

// SetPhones changes the sponsor's phones
func (s *Sponsor) SetPhones(newPhones []string) error {
	for _, phone := range newPhones {
		if phone == "" {
			return errors.New("Sponsor's phone can't be empty")
		}
	}
	s.Phones = newPhones
	return nil
}

// AddPhone adds phone at the bottom of the sponsor's phones
func (s *Sponsor) AddPhone(newPhone string) error {
	if newPhone == "" {
		return errors.New("Sponsor's phone can't be empty")
	}
	s.Phones = append(s.Phones,newPhone)
	return nil
}

// RemovePhone removes an phone (by the name)
func (s *Sponsor) RemovePhone(oddPhone string) error {
	for i, existingPhone := range s.Phones {
		if existingPhone == oddPhone {
			s.Phones = append(s.Phones[:i], s.Phones[i+1:]...)
			return nil
		}
	}
	return errors.New("Sponsor's phone was not found. Phone couldn't be deleted")
}

// GetLogo returns the sponsor's logo
func (s *Sponsor) GetLogo() MediaMeta {
	return s.Logo
}

// SetLogo changes the sponsor's logo
func (s *Sponsor) SetLogo(newLogo MediaMeta) error {
	if newLogo.isEmpty == true {
		return errors.New("Sponsor's logo can't be empty")
	}
	s.Logo = newLogo
	return nil
}
