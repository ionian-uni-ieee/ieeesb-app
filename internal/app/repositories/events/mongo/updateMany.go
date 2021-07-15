package mongo

import "context"

func (r *Repository) UpdateMany(filter interface{}, update interface{}) ([]string, error) {
	result, err := r.collection.UpdateMany(context.Background(), filter, update)

	return result.UpsertedID.([]string), err
}
