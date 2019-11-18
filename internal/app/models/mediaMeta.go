package models

// MediaMeta contains information for tracking
// a media object (image, video & audio) in the storage
type MediaMeta struct {
	ID   string `bson:"_id" json:"id"`
	MIME string `bson:"MIME" json:"MIME"`
}

// GetID returns the hex of the mediameta's id
func (m *MediaMeta) GetID() string {
	return m.ID
}

// isEmpty returns true if the MIME is empty
func (m *MediaMeta) isEmpty() bool {
	if m.MIME == "" {
		return true
	}
	return false
}

// areEqual returns true if two mediameta are equal
// two mediameta are equal if their IDs are equal
func (firstMediaMeta *MediaMeta) areEqual(secondMediaMeta MediaMeta) bool {
	if firstMediaMeta.ID == secondMediaMeta.GetID() {
		return true
	}
	return false
}
