package testing

import (
	"errors"
	"reflect"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/drivers/database"
	testingDatabase "github.com/ionian-uni-ieee/ieeesb-app/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/reflections"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Repository contains the necessary modules to communicate with the database
// Its fields are public so they can be inserted in a test
type Repository struct {
	Database   database.Driver
	Collection *testingDatabase.Collection
}

// MakeRepository is a sponsor repository factory
func MakeRepository(database database.Driver) *Repository {
	collection := database.GetCollection("sponsors").(*testingDatabase.Collection)

	sponsorFieldNames, err := reflections.GetFieldNames(reflect.TypeOf(models.Sponsor{}))

	if err != nil {
		panic(err)
	}

	for _, fieldName := range sponsorFieldNames {
		collection.Columns[fieldName] = []interface{}{}
	}

	return &Repository{database, collection}
}

// FindByID finds a sponsor by his primary ID
func (r *Repository) FindByID(sponsorID string) (*models.Sponsor, error) {
	objectIDs := r.Collection.Columns["ID"]

	for rowIndex, objectID := range objectIDs {
		if objectID.(primitive.ObjectID).Hex() == sponsorID {
			sponsorFound := &models.Sponsor{}

			// Take cell values of each column
			for key, value := range r.Collection.Columns {
				cellValue := value[rowIndex]
				err := reflections.SetFieldValue(sponsorFound, key, cellValue)
				if err != nil {
					return nil, err
				}
			}

			return sponsorFound, nil
		}
	}
	return nil, nil
}

// UpdateByID finds a sponsor by his primary ID and updates him
func (r *Repository) UpdateByID(sponsorID string, update interface{}) error {
	objectIDs := r.Collection.Columns["ID"]

	for rowIndex, objectID := range objectIDs {
		if objectID.(primitive.ObjectID).Hex() == sponsorID {
			sponsorFound := &models.Sponsor{}

			// Take cell values of each column
			for key, value := range r.Collection.Columns {
				cellValue := value[rowIndex]
				err := reflections.SetFieldValue(sponsorFound, key, cellValue)
				if err != nil {
					return err
				}
			}

			// Update sponsor found
			if updateMap, ok := update.(map[string]interface{}); ok {
				for key, value := range updateMap {
					err := reflections.SetFieldValue(sponsorFound, key, value)
					if err != nil {
						return err
					}
				}
			} else if updateModel, ok := update.(models.Sponsor); ok {
				fieldNames, err := reflections.GetFieldNames(reflect.TypeOf(models.Sponsor{}))
				if err != nil {
					return err
				}
				for _, fieldName := range fieldNames {
					field, err := reflections.GetFieldValue(updateModel, fieldName)
					if err != nil {
						return err
					}
					err = reflections.SetFieldValue(sponsorFound, fieldName, field)
					if err != nil {
						return err
					}
				}
			} else {
				return errors.New("Update interface is not a string map of interfaces")
			}

			// Update database
			for key := range r.Collection.Columns {
				cellValue, err := reflections.GetFieldValue(*sponsorFound, key)
				if err != nil {
					return err
				}
				r.Collection.Columns[key][rowIndex] = cellValue
			}
			return nil
		}
	}
	return errors.New("No document found with this id " + sponsorID)
}

// DeleteByID finds an sponsor by his ID and deletes it
func (r *Repository) DeleteByID(sponsorID string) error {
	objectIDs := r.Collection.Columns["ID"]

	for rowIndex, objectID := range objectIDs {
		if objectID.(primitive.ObjectID).Hex() == sponsorID {
			for key, column := range r.Collection.Columns {
				r.Collection.Columns[key] = append(column[:rowIndex], column[rowIndex+1:]...)
			}

			return nil
		}
	}
	return errors.New("No document found with this id " + sponsorID)
}

// Find finds all the sponsors that match the filter
func (r *Repository) Find(filter interface{}, skip int64, limit int64) ([]models.Sponsor, error) {

	totalItems := int64(len(r.Collection.Columns["ID"]))
	if totalItems < skip {
		skip = totalItems
	}
	if totalItems < limit+skip {
		limit = 0
	}

	modelsFound := make([]models.Sponsor, limit)

	// Iterating on the filter
	if filterMap, ok := filter.(map[string]interface{}); ok {
		for key, value := range filterMap {
			// Iterating on all the filtered's column's cell
			for rowIndex, cell := range r.Collection.Columns[key][skip : skip+limit] {
				if cell == value {
					// Create an user from this row's data
					for fieldName, column := range r.Collection.Columns {
						value := column[rowIndex]
						err := reflections.SetFieldValue(
							&modelsFound[rowIndex],
							fieldName,
							value,
						)
						if err != nil {
							return nil, err
						}
					}
				}
			}
		}
	} else if filter == nil {
		for fieldName, column := range r.Collection.Columns {
			for rowIndex := range column[skip : skip+limit] {
				// Create an user from this row's data
				value := column[rowIndex]
				err := reflections.SetFieldValue(&modelsFound[rowIndex], fieldName, value)
				if err != nil {
					return nil, err
				}
			}
		}
	} else {
		return nil, errors.New("Filter interface is not a string map of interfaces")
	}

	return modelsFound, nil
}

// FindOne finds the first sponsor that matches the filter
func (r *Repository) FindOne(filter interface{}) (*models.Sponsor, error) {
	// Iterating on the filter
	if filterMap, ok := filter.(map[string]interface{}); ok {
		for key, value := range filterMap {
			// Iterating on all the filtered's column's cell
			for rowIndex, cell := range r.Collection.Columns[key] {
				// Model fits a filter requirement
				if cell == value {
					// Create an sponsor from this row's data
					newSponsor := models.Sponsor{}
					for fieldName, column := range r.Collection.Columns {
						value := column[rowIndex]
						err := reflections.SetFieldValue(&newSponsor, fieldName, value)
						if err != nil {
							return nil, err
						}
					}
					return &newSponsor, nil
				}
			}
		}
	} else {
		return nil, errors.New("Filter interface is not a string map of interfaces")
	}
	return nil, errors.New("No sponsor found")
}

// UpdateMany finds the sponsors that match the filter and updates them
func (r *Repository) UpdateMany(filter interface{}, update interface{}) ([]string, error) {
	updatedIDs := []string{}
	// Iterating on the filter
	if filterMap, ok := filter.(map[string]interface{}); ok {
		for key, value := range filterMap {
			// Iterating on all the filtered's column's cell
			for cellIndex, cell := range r.Collection.Columns[key] {
				// Record fits a filter requirement
				if cell == value {
					if sponsorModel, ok := cell.(models.Sponsor); ok {
						// Update model
						if updateMap, ok := update.(map[string]interface{}); ok {
							// Iterate on the update map
							for key, value := range updateMap {
								if key == "ID" {
									return nil, errors.New("Can't update id in testing mode")
								}

								// Update found sponsor's field
								err := reflections.SetFieldValue(&sponsorModel, key, value)
								if err != nil {
									return nil, err
								}

								// Update cell
								r.Collection.Columns[key][cellIndex] = sponsorModel

								// Store ID of the updated object
								recordID := r.Collection.Columns["ID"][cellIndex].(primitive.ObjectID)
								// Insert ID if ID doesn't exist already in the array
								for updatedIDIndex, updatedID := range updatedIDs {
									notSameID := updatedID == recordID.Hex()
									lastItem := updatedIDIndex == len(updatedIDs)-1
									if !notSameID {
										break
									}
									if notSameID && lastItem {
										updatedIDs = append(updatedIDs, updatedID)
									}
								}
							}
						} else {
							return nil, errors.New("Update interface is not a string map of interfaces")
						}
					} else {
						return nil, errors.New("Cell could not be parsed as Sponsor model")
					}
				}
			}
		}
		return nil, errors.New("Filter interface is not a string map of interfaces")
	}

	return updatedIDs, nil
}

// DeleteMany finds the sponsors that match the filter and deletes them
func (r *Repository) DeleteMany(filter interface{}) (int64, error) {
	deletedCount := int64(0)

	// Iterating on the filter
	if filterMap, ok := filter.(map[string]interface{}); ok {
		for key, value := range filterMap {
			// Iterating on all the filtered's column's cell
			for cellIndex, cell := range r.Collection.Columns[key] {
				// Model fits a filter requirement
				if cell == value {
					for key := range r.Collection.Columns {
						// Remove column's value of the record
						r.Collection.Columns[key] = append(
							r.Collection.Columns[key][:cellIndex],
							r.Collection.Columns[key][cellIndex+1:])
					}
					deletedCount++
				}
			}
		}
		return 0, errors.New("Filter interface is not a string map of interfaces")
	}
	return deletedCount, nil
}

// InsertOne adds a new sponsor to the repository
func (r *Repository) InsertOne(document models.Sponsor) (string, error) {
	if document.ID.IsZero() {
		document.ID = primitive.NewObjectID()
	}

	for key := range r.Collection.Columns {
		docField, err := reflections.GetFieldValue(document, key)

		if err != nil {
			return "", err
		}

		r.Collection.Columns[key] = append(r.Collection.Columns[key], docField)
	}

	return document.ID.Hex(), nil
}

// InsertMany iterates and adds all the sponsors to the repository
func (r *Repository) InsertMany(documents []models.Sponsor) ([]string, error) {
	insertedIDs := []string{}
	for _, doc := range documents {
		id, err := r.InsertOne(doc)
		if err != nil {
			return nil, err
		}
		insertedIDs = append(insertedIDs, id)
	}

	return insertedIDs, nil
}

// IsDuplicate checks if the primary key values are unique in the repository
func (r *Repository) IsDuplicate(name string) bool {
	namesColumn := r.Collection.Columns["Name"]

	for _, cellSponsorname := range namesColumn {
		if cellSponsorname == name {
			return true
		}
	}

	return false
}
