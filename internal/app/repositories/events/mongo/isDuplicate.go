package mongo

import (
	"context"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) IsDuplicate(name string) bool {
	sameKeysFilter := &bson.M{
		"name": name,
	}

	eventFound := models.Sponsor{}
	r.collection.FindOne(context.Background(), sameKeysFilter).Decode(&eventFound)

	return !eventFound.ID.IsZero()
}
