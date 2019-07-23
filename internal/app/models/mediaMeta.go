package models

type MediaMeta struct {
	ID   string `bson:"_id" json:"id"`
	MIME string `bson:"MIME" json:"MIME"`
}
