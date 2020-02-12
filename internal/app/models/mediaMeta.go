package models

import (
	"errors"
)

// MediaMeta contains information for tracking
// a media object (image, video & audio) in the storage
type MediaMeta struct {
	ID   string `bson:"_id" json:"id"`
	MIME string `bson:"MIME" json:"MIME"`
}

// NewMediaMeta is a mediameta factory that generates a mediameta
func NewMediaMeta(
	id string,
	mine string,
) *MediaMeta {
	return &MediaMeta{
		ID:   id,
		MIME: mine,
	}
}

// GetID returns the hex of the media's id
func (m *MediaMeta) GetID() string {
	return m.ID
}

// SetID changes the Media's id
func (m *MediaMeta) SetID(newID string) error {
	if newID == "" {
		return errors.New("Media's id can't be empty")
	}
	m.ID = newID
	return nil
}

// GetMime returns the mediameta's mime
func (m *MediaMeta) GetMime() string {
	return m.MIME
}

// SetMime changes the mediameta's mime
func (m *MediaMeta) SetMime(newMime string) error {
	if newMime == "" {
		return errors.New("MediaMeta's mime can't be empty")
	}
	m.MIME = newMime
	return nil
}

// isEmpty returns true if the mediameta is empty
func (m *MediaMeta) isEmpty() bool {
	if m.ID == "" && m.MIME == "" {
		return true
	}
	return false
}

// isEqual returns true if the mediameta is equal to another mediameta
func (firstMediameta *MediaMeta) isEqual(secondMediameta MediaMeta) bool {
	if firstMediameta.ID == secondMediameta.GetID() && firstMediameta.MIME == secondMediameta.GetMime() {
		return true
	}
	return false
}
