package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import "errors"

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

// NewEvent is an event factory that generates an event
func NewEvent (
	name string,
	description string,
	tags []string,
	category string,
	sponsors  []Sponsor,
	logo MediaMeta,
	media []MediaMeta,
) (*Event, error) {
		if name == "" {
			return nil,  errors.New("Name can't be empty")
		}
		if description == "" {
			return nil, errors.New("Description can't be empty")
		}
		return &Event{
			ID:				primitive.NewObjectID(),
			Name:			name,
			Description:	description,
			Tags:			tags,
			Type:			category,
			Sponsors:		sponsors,
			Logo:			logo,
			Media:			media,
		}, nil
}

// GetID returns the hex of the event's id
func (e *Event) GetID() string {
	return e.ID.Hex()
}

//GetObjectID returns the event's object id
func (e *Event) GetObjectID() primitive.ObjectID {
	return e.ID
}

// GetName returns the event's name
func (e *Event) GetName() string {
	return e.Name
}

// SetName changes the event's name
func (e *Event) SetName(newName string) error {
	if newName == "" {
		return errors.New("Event's name can't be empty")
	}
	e.Name = newName
	return nil
}

// GetDescription returns the event's description
func (e *Event) GetDescription() string {
	return e.Description
}

// SetDescriotion changes the event's description
func (e *Event) SetDescription(newDescription string) error {
	if newDescription == "" {
		return errors.New("Event's description can't be empty")
	}
	e.Description = newDescription
	return nil
}

// GetTags returns the event's tags
func (e *Event) GetTags() []string {
	return e.Tags
}

// SetTags changes the event's tags
func (e *Event) SetTags(newTags []string) error {
	for _, tag := range e.Tags {
		if tag == "" {
			return errors.New("Event's tag can't be empty")
		}
	}
	e.Tags = newTags
	return nil
}

// AddTag adds a tag at the end of the event's tags
func (e *Event) AddTag(newTag string) error {
	if newTag == "" {
		return errors.New("Event's tag can't be empty")
	}
	e.Tags = append(e.Tags, newTag)
	return nil
}

// RemoveTag removes a tag (by the name)
func (e *Event) RemoveTag(oddTag string) error {
	if oddTag == "" {
		return errors.New("Event's tag can't be empty")
	}
	for i, existingTag := range e.Tags {
		if existingTag == oddTag {
			e.Sponsors = append(e.Sponsors[:i], e.Sponsors[i+1:]...)
		}
	}
	return nil
}

// GetType returns the event's type
func (e *Event) GetType()string {
	return e.Type
}

// SetType changes the event's type
func (e *Event) SetType(newType string) error {
	if newType == "" {
		return errors.New("Event's type can't be empty")
	}
	e.Type = newType
	return nil
}

// GetSponsors returns the event's sponsors
func (e *Event) GetSponsors() []Sponsor {
	return e.Sponsors
}

// SetSponsors changes the event's sponsors
func (e *Event) SetSponsors(NewSponsors []Sponsor) error {
	for _, sponsor := range NewSponsors {
		if sponsor.isEmpty() == true {
			return errors.New("Event's sponsor can't be empty")
		}
	}
	e.Sponsors = NewSponsors
	return nil
}

// AddSponsor adds a sponsor at the bottom of the event's sponsors 
func (e *Event) AddSponsor(newSponsor Sponsor) error {
	if newSponsor.isEmpty() == true {
		return errors.New("Event's sponsor can't be empty")
	}
	e.Sponsors = append(e.Sponsors, newSponsor)
	return nil
}

// RemoveSponsor removes a sponsor (by the name)
func (e *Event) RemoveSponsor(oddSponsor Sponsor) error {
	if oddSponsor.isEmpty() == true {
		return errors.New("Event's sponsor can't be empty")
	}
	for i, existingSponsor := range e.Sponsors {
		if existingSponsor.areEqual(oddSponsor) == true {
			e.Sponsors = append(e.Sponsors[:i], e.Sponsors[i+1:]...)
			return nil
		}
	}
	return errors.New( "Event's sponsor was not found. Sponsor couldn't be deleted")
}

// GetLogo returns event's logo
func (e *Event) GetLogo() MediaMeta {
	return e.Logo
}

// SetLogo changes the event's logo
func (e *Event) SetLogo(newLogo MediaMeta) error {
	if newLogo.isEmpty() == true {
		return errors.New("Event's logo can't be empty")
	}
	e.Logo = newLogo
	return nil
}

// GetMedia returns event's Media
func (e *Event) GetMedia() []MediaMeta {
	return e.Media
}

// SetMedia changes the event's media
func (e *Event) SetMedia(newMedia []MediaMeta) error {
	for _, medium := range newMedia {
		if medium.isEmpty() == true {
			return errors.New("Event's medium can't be empty")
		}
	}
	e.Media = newMedia
	return nil
}

// AddMedium adds a medium at the bottom of the event's media
func (e *Event) AddMedium(newMedium MediaMeta) error {
	if newMedium.isEmpty() == true {
		return errors.New("Event's medium cant't be empty")
	}
	e.Media = append(e.Media, newMedium)
	return nil
}

// RemoveMedium removes a medium (by the name)
func (e *Event) RemoveMedium(oddMeidum MediaMeta) error {
	if oddMeidum.isEmpty() == true {
		return errors.New("Event's medium can't be empty")
	}
	for i, existingMedium := range e.Media {
		if existingMedium.areEqual(oddMeidum) == true {
			e.Media = append(e.Media[:i], e.Media[i+1:]...)
			return nil
		}
	}
	return errors.New("Event's medium was not found. Medium couldn't be deleted ")
}
