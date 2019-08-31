package models

// MediaMeta contains information for tracking
// a media object (image, video & audio) in the storage
type MediaMeta struct {
	ID   string `bson:"_id" json:"id"`
	MIME string `bson:"MIME" json:"MIME"`
}
