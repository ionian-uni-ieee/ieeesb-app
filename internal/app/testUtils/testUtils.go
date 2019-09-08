package testUtils

import (
	"reflect"

	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/pkg/reflections"
)

// IsInterfaceEqualToCollectionRow checks if the collection at a specific row
// provides exactly the same data with the provided interface
func IsInterfaceEqualToCollectionRow(db *testingDatabase.DatabaseSession, collectionName string, d interface{}, row int) bool {

	fieldNames, err := reflections.GetFieldNames(reflect.TypeOf(d))

	if err != nil {
		panic(err)
	}

	users := db.GetCollection(collectionName).(*testingDatabase.Collection)

	for _, fieldName := range fieldNames {
		field, err := reflections.GetFieldValue(d, fieldName)

		if err != nil {
			panic(err)
		}

		if !reflect.DeepEqual(users.Columns[fieldName][row], field) {
			return false
		}
	}

	return true
}

// GetInterfaceAtCollectionRow returns the interface that is located in a collection's row
func GetInterfaceAtCollectionRow(db *testingDatabase.DatabaseSession, collectionName string, structType reflect.Type, row int) interface{} {
	users := db.GetCollection(collectionName).(*testingDatabase.Collection)

	if IsCollectionEmpty(db, collectionName) {
		return reflections.CreateEmptyInstance(structType)
	}

	fieldNames, err := reflections.GetFieldNames(structType)

	if err != nil {
		panic(err)
	}

	val := reflect.New(structType).Elem()

	for _, fieldName := range fieldNames {
		fieldValue := users.Columns[fieldName][row]
		field := val.FieldByName(fieldName)
		field.Set(reflect.ValueOf(fieldValue))
	}

	return val.Interface()
}

// IsCollectionEmpty checks every collection's column's length
// if at least one has more than 0 length, it returns false
// otherwise if no columns are filled
// or the collection or columns point to an nil pointer, it returns true
func IsCollectionEmpty(db *testingDatabase.DatabaseSession, collectionName string) bool {
	users := db.GetCollection(collectionName).(*testingDatabase.Collection)

	if users == nil ||
		users.Columns == nil ||
		len(users.Columns) == 0 {
		return true
	}

	for _, column := range users.Columns {
		if len(column) != 0 {
			return false
		}
	}

	return true
}

// GetCollectionLength returns a collection's first column's length
// Not safe if columns might have different lengths in an irregular (probably buggy) circumstance
func GetCollectionLength(db *testingDatabase.DatabaseSession, collectionName string) (length int) {

	users := db.GetCollection(collectionName).(*testingDatabase.Collection)

	if users == nil {
		panic("Collection points to nil")
	}
	if users.Columns == nil {
		panic("Columns point to nil")
	}
	if len(users.Columns) == 0 {
		panic("Columns have zero length when they should be filled with field names from the testing driver")
	}

	for _, column := range users.Columns {
		length = len(column)
		break
	}

	return length
}

// ResetCollection Clears the collection's data
func ResetCollection(db *testingDatabase.DatabaseSession, collectionName string) {
	collection := db.GetCollection(collectionName).(*testingDatabase.Collection)

	if collection == nil {
		panic("Can't reset collection, collection points to nil")
	}

	for key := range collection.Columns {
		collection.Columns[key] = []interface{}{}
	}
}

// SetupData resets the collection and inserts an array of data in it
// if no data is provided, it's the same as calling resetCollection()
func SetupData(db *testingDatabase.DatabaseSession, collectionName string, data ...interface{}) {
	ResetCollection(db, collectionName)

	if len(data) == 0 {
		return
	}

	// Getting first element's type
	interfaceType := reflect.TypeOf(data[0])
	// Getting all the struct type's field names
	ticketFieldNames, err := reflections.GetFieldNames(interfaceType)
	if err != nil {
		panic(err)
	}

	// Getting pointer to database's requested collection
	collection := db.GetCollection(collectionName).(*testingDatabase.Collection)

	for _, item := range data {
		for _, fieldName := range ticketFieldNames {
			fieldValue, err := reflections.GetFieldValue(item, fieldName)
			if err != nil {
				panic(err)
			}

			collection.Columns[fieldName] = append(
				collection.Columns[fieldName],
				fieldValue)
		}
	}
}
