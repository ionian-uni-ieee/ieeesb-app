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

// MakeRepository is a event repository factory
func MakeRepository(database database.Driver) *Repository {
	collection := database.GetCollection("events").(*testingDatabase.Collection)

	eventFieldNames, err := reflections.GetFieldNames(reflect.TypeOf(models.Event{}))

	if err != nil {
		panic(err)
	}

	for _, fieldName := range eventFieldNames {
		collection.Columns[fieldName] = []interface{}{}
	}

	return &Repository{database, collection}
}

// FindByID finds a event by his primary ID
func (r *Repository) FindByID(eventID string) (*models.Event, error) {
	objectIDs := r.Collection.Columns["ID"]

	for rowIndex, objectID := range objectIDs {
		if objectID.(primitive.ObjectID).Hex() == eventID {
			eventFound := &models.Event{}

			// Take cell values of each column
			for key, value := range r.Collection.Columns {
				cellValue := value[rowIndex]
				err := reflections.SetFieldValue(eventFound, key, cellValue)
				if err != nil {
					return nil, err
				}
			}

			return eventFound, nil
		}
	}
	return nil, nil
}

// UpdateByID finds a event by his primary ID and updates him
func (r *Repository) UpdateByID(eventID string, update interface{}) error {
	objectIDs := r.Collection.Columns["ID"]

	for rowIndex, objectID := range objectIDs {
		if objectID.(primitive.ObjectID).Hex() == eventID {
			eventFound := &models.Event{}

			// Take cell values of each column
			for key, value := range r.Collection.Columns {
				cellValue := value[rowIndex]
				err := reflections.SetFieldValue(eventFound, key, cellValue)
				if err != nil {
					return err
				}
			}

			// Update event found
			if updateMap, ok := update.(map[string]interface{}); ok {
				for key, value := range updateMap {
					err := reflections.SetFieldValue(eventFound, key, value)
					if err != nil {
						return err
					}
				}
			} else if updateModel, ok := update.(models.Event); ok {
				fieldNames, err := reflections.GetFieldNames(reflect.TypeOf(models.Event{}))
				if err != nil {
					return err
				}
				for _, fieldName := range fieldNames {
					field, err := reflections.GetFieldValue(updateModel, fieldName)
					if err != nil {
						return err
					}
					err = reflections.SetFieldValue(eventFound, fieldName, field)
					if err != nil {
						return err
					}
				}
			} else {
				return errors.New("Update interface is not a string map of interfaces")
			}

			// Update database
			for key := range r.Collection.Columns {
				cellValue, err := reflections.GetFieldValue(*eventFound, key)
				if err != nil {
					return err
				}
				r.Collection.Columns[key][rowIndex] = cellValue
			}
			return nil
		}
	}
	return errors.New("No document found with this id " + eventID)
}

// DeleteByID finds an event by his ID and deletes it
func (r *Repository) DeleteByID(eventID string) error {
	objectIDs := r.Collection.Columns["ID"]

	for rowIndex, objectID := range objectIDs {
		if objectID.(primitive.ObjectID).Hex() == eventID {
			for key, column := range r.Collection.Columns {
				r.Collection.Columns[key] = append(column[:rowIndex], column[rowIndex+1:]...)
			}

			return nil
		}
	}
	return errors.New("No document found with this id " + eventID)
}

// Find finds all the events that match the filter
func (r *Repository) Find(filter interface{}) ([]models.Event, error) {
	modelsFound := []models.Event{}

	// Iterating on the filter
	if filterMap, ok := filter.(map[string]interface{}); ok {
		for key, value := range filterMap {
			// Iterating on all the filtered's column's cell
			for rowIndex, cell := range r.Collection.Columns[key] {
				if cell == value {
					// Create an event from this row's data
					newEvent := models.Event{}
					for fieldName, column := range r.Collection.Columns {
						value := column[rowIndex]
						err := reflections.SetFieldValue(&newEvent, fieldName, value)
						if err != nil {
							return nil, err
						}
					}
					modelsFound = append(modelsFound, newEvent)
				}
			}
		}
	} else {
		return nil, errors.New("Filter interface is not a string map of interfaces")
	}
	return modelsFound, nil
}

// FindOne finds the first event that matches the filter
func (r *Repository) FindOne(filter interface{}) (*models.Event, error) {
	// Iterating on the filter
	if filterMap, ok := filter.(map[string]interface{}); ok {
		for key, value := range filterMap {
			// Iterating on all the filtered's column's cell
			for rowIndex, cell := range r.Collection.Columns[key] {
				// Model fits a filter requirement
				if cell == value {
					// Create an event from this row's data
					newEvent := models.Event{}
					for fieldName, column := range r.Collection.Columns {
						value := column[rowIndex]
						err := reflections.SetFieldValue(&newEvent, fieldName, value)
						if err != nil {
							return nil, err
						}
					}
					return &newEvent, nil
				}
			}
		}
	} else {
		return nil, errors.New("Filter interface is not a string map of interfaces")
	}
	return nil, errors.New("No event found")
}

// UpdateMany finds the events that match the filter and updates them
func (r *Repository) UpdateMany(filter interface{}, update interface{}) ([]string, error) {
	updatedIDs := []string{}
	// Iterating on the filter
	if filterMap, ok := filter.(map[string]interface{}); ok {
		for key, value := range filterMap {
			// Iterating on all the filtered's column's cell
			for cellIndex, cell := range r.Collection.Columns[key] {
				// Record fits a filter requirement
				if cell == value {
					if eventModel, ok := cell.(models.Event); ok {
						// Update model
						if updateMap, ok := update.(map[string]interface{}); ok {
							// Iterate on the update map
							for key, value := range updateMap {
								if key == "ID" {
									return nil, errors.New("Can't update id in testing mode")
								}

								// Update found event's field
								err := reflections.SetFieldValue(&eventModel, key, value)
								if err != nil {
									return nil, err
								}

								// Update cell
								r.Collection.Columns[key][cellIndex] = eventModel

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
						return nil, errors.New("Cell could not be parsed as Event model")
					}
				}
			}
		}
		return nil, errors.New("Filter interface is not a string map of interfaces")
	}

	return updatedIDs, nil
}

// DeleteMany finds the events that match the filter and deletes them
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

// InsertOne adds a new event to the repository
func (r *Repository) InsertOne(document models.Event) (string, error) {
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

// InsertMany iterates and adds all the events to the repository
func (r *Repository) InsertMany(documents []models.Event) ([]string, error) {
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

	for _, cellEventname := range namesColumn {
		if cellEventname == name {
			return true
		}
	}

	return false
}
