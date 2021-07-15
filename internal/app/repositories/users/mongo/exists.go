package mongo

import (
	"context"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) Exists(email string, username string, fullname string) bool {
	sameKeysFilter := &bson.M{
		"$or": bson.A{
			bson.M{"email": email},
			bson.M{"username": username},
			bson.M{"fullname": fullname},
		}}

	userFound := models.User{}
	r.collection.FindOne(context.Background(), sameKeysFilter).Decode(&userFound)

	return !userFound.ID.IsZero()
}
