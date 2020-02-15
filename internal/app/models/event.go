package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import "errors"

// Event describes all the necessary information
// about what an IEEE event
type Event struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Date        uint               `bson:"date" json:"date"`
	Tags        []string           `bson:"tags" json:"tags"`
	Type        string             `bson:"type" json:"type"`
	Sponsors    []Sponsor          `bson:"sponsors" json:"sponsors"`
	Logo        string             `bson:"logo" json:"logo"`
	Media       []string           `bson:"media" json:"media"`
}

// NewEvent is an event factory that generates an event
func NewEvent(
	name string,
	description string,
	date uint,
	tags []string,
	category string,
	sponsors []Sponsor,
	logo string,
	media []string,
) (*Event, error) {
	if name == "" {
		return nil, errors.New("Name can't be empty")
	}
	return &Event{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Description: description,
		Date:        date,
		Tags:        tags,
		Type:        category,
		Sponsors:    sponsors,
		Logo:        logo,
		Media:       media,
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
func (e *Event) SetDescription(newDescription string) {
	e.Description = newDescription
}

// GetDate returns the event's date
func (e *Event) GetDate() uint {
	return e.Date
}

// SetDate change the event's date
func (e *Event) SetDate(newDate uint) {
	e.Date = newDate
}

// GetTags returns the event's tags
func (e *Event) GetTags() []string {
	return e.Tags
}

// SetTags changes the event's tags
func (e *Event) SetTags(newTags []string) error {
	for tagIndex, tag := range e.Tags {
		if tag == "" {
			return errors.New("Event's tag can't be empty")
		}
		for tmpIndex, tmpTag := range e.Tags {
			if tag == tmpTag && tagIndex != tmpIndex {
				return errors.New("Duplicate tags")
			}
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
	for _, tag := range e.Tags {
		if tag == newTag {
			return errors.New("Event's tag already exist")
		}
	}
	e.Tags = append(e.Tags, newTag)
	return nil
}

// RemoveTag removes a tag (by the name)
func (e *Event) RemoveTag(tag string) error {
	for i, existingTag := range e.Tags {
		if existingTag == tag {
			e.Sponsors = append(e.Sponsors[:i], e.Sponsors[i+1:]...)
			return nil
		}
	}
	return errors.New("Event's tag was not found")
}

// GetType returns the event's type
func (e *Event) GetType() string {
	return e.Type
}

// SetType changes the event's type
func (e *Event) SetType(newType string) {
	e.Type = newType
}

// GetSponsors returns the event's sponsors
func (e *Event) GetSponsors() []Sponsor {
	return e.Sponsors
}

// SetSponsors changes the event's sponsors
func (e *Event) SetSponsors(NewSponsors []Sponsor) error {
	for sponsorIndex, sponsor := range NewSponsors {
		if sponsor.isEmpty() {
			return errors.New("Event's sponsor can't be empty")
		}
		for tmpIndex, tmpSponsor := range NewSponsors {
			if sponsor.isEqual(tmpSponsor) && sponsorIndex != tmpIndex {
				return errors.New("Duplicate Sponsor")
			}
		}
	}
	e.Sponsors = NewSponsors
	return nil
}

// AddSponsor adds a sponsor at the bottom of the event's sponsors
func (e *Event) AddSponsor(newSponsor Sponsor) error {
	if newSponsor.isEmpty() {
		return errors.New("Event's sponsor can't be empty")
	}
	for _, sponsor := range e.Sponsors {
		if sponsor.isEqual(newSponsor) {
			return errors.New("Event's sponsor already exist")
		}
	}
	e.Sponsors = append(e.Sponsors, newSponsor)
	return nil
}

// RemoveSponsor removes a sponsor (by the name)
func (e *Event) RemoveSponsor(sponsor Sponsor) error {
	for i, existingSponsor := range e.Sponsors {
		if existingSponsor.isEqual(sponsor) {
			e.Sponsors = append(e.Sponsors[:i], e.Sponsors[i+1:]...)
			return nil
		}
	}
	return errors.New("Event's sponsor was not found.")
}

// GetLogo returns event's logo
func (e *Event) GetLogo() string {
	return e.Logo
}

// SetLogo changes the event's logo
func (e *Event) SetLogo(newLogo string) {
	e.Logo = newLogo
}

// GetMedia returns event's Media
func (e *Event) GetMedia() []string {
	return e.Media
}

// SetMedia changes the event's media
func (e *Event) SetMedia(newMedia []string) error {
	for mediaIndex, media := range newMedia {
		if media == "" {
			return errors.New("Event's media can't be empty")
		}
		for tmpIndex, tmpMedia := range newMedia {
			if media == tmpMedia && mediaIndex != tmpIndex {
				return errors.New("Duplicate media")
			}
		}
	}
	e.Media = newMedia
	return nil
}

// AddMedia adds a media at the bottom of the event's media
func (e *Event) AddMedia(media string) error {
	if media == "" {
		return errors.New("Event's media cant't be empty")
	}
	for _, existingMedia := range e.Media {
		if existingMedia == media {
			return errors.New("Event's media already exist")
		}
	}
	e.Media = append(e.Media, media)
	return nil
}

// RemoveMedia removes a media (by the name)
func (e *Event) RemoveMedia(Media string) error {
	for i, existingMedia := range e.Media {
		if existingMedia == Media {
			e.Media = append(e.Media[:i], e.Media[i+1:]...)
			return nil
		}
	}
	return errors.New("Event's media was not found")
}
