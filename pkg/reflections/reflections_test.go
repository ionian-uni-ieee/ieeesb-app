package reflections_test

import (
	"reflect"
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/pkg/reflections"
)

func TestSetFieldValue(t *testing.T) {
	type testStruct struct {
		Name string
	}

	s := testStruct{
		Name: "jack",
	}

	err := reflections.SetFieldValue(&s, "Name", "john")

	if err != nil {
		t.Error(err)
	}

	if s.Name != "john" {
		t.Error("Expected 'Name' to be equal to 'john', but is " + s.Name)
	}

	// Check for non struct type
	invalidInterface := "string"

	err = reflections.SetFieldValue(&invalidInterface, "Name", "")

	if err == nil {
		t.Error("Expected an error for invalid type, but no error was returned")
	}

	expectedError := "Pointed interface is not a struct"
	if err.Error() != expectedError {
		t.Error("Expected '"+expectedError+"' error, but got ", err.Error())
	}

	// Get a non existant field
	s = testStruct{}

	err = reflections.SetFieldValue(&s, "NotName", "")

	if err.Error() != "Field is not valid" {
		t.Error("Expected invalid field error but got", err)
	}

	// No struct pointer provided
	s = testStruct{}

	err = reflections.SetFieldValue(s, "Name", "")

	expectedError = "Interface is not a pointer"
	if err.Error() != expectedError {
		t.Error("Expected '"+expectedError+"' error, but got ", err.Error())
	}
}

func TestGetFieldValue(t *testing.T) {
	type testStruct struct {
		Name string
	}

	// Check for 'john'
	s := testStruct{
		Name: "john",
	}

	name, err := reflections.GetFieldValue(s, "Name")

	if err != nil {
		t.Error(err)
	}

	if name.(string) != "john" {
		t.Error("Expected 'Name' to be equal to 'john', but it's ", name)
	}

	// Check for empty field
	s = testStruct{
		Name: "",
	}

	name, err = reflections.GetFieldValue(s, "Name")

	if err != nil {
		t.Error(err)
	}

	if name.(string) != "" {
		t.Error("Expected 'Name' to be equal to '', but it's ", name)
	}

	// Check for non struct type
	invalidInterface := "string"

	name, err = reflections.GetFieldValue(invalidInterface, "Name")

	if err == nil {
		t.Error("Expected an error for invalid type, but no error was returned")
	}

	if err.Error() != "Interface is not a struct" {
		t.Error("Expected an error for invalid interface type, but got ", err.Error())
	}

	// Get a non existant field
	s = testStruct{}

	name, err = reflections.GetFieldValue(s, "NotName")

	if err.Error() != "Field is not valid" {
		t.Error("Expected invalid field error but got", err)
	}
}

func TestGetTagByFieldName(t *testing.T) {
	// Normal
	type foo struct {
		Username string `json:"uname"`
	}

	tag, err := reflections.GetTagByFieldName(reflect.TypeOf(foo{}), "Username", "json")

	if err != nil {
		t.Error(err)
	}

	if tag != "uname" {
		t.Error("Expected tag name 'uname' but instead got " + tag)
	}
}

func TestGetFieldNames(t *testing.T) {
	// Struct with one field
	type testingStruct struct {
		Name string
	}

	s := testingStruct{}

	fieldNames, err := reflections.GetFieldNames(reflect.TypeOf(s))

	if err != nil {
		t.Error(err)
	}

	if fieldNames[0] != "Name" {
		t.Error("Expected first field name to be 'Name' but it's ", fieldNames[0])
	}

	// Struct with many fields
	type testingStruct2 struct {
		Name    string
		Address string
		Number  int
	}

	s2 := testingStruct2{}

	fieldNames, err = reflections.GetFieldNames(reflect.TypeOf(s2))

	if err != nil {
		t.Error(err)
	}

	if fieldNames[0] != "Name" ||
		fieldNames[1] != "Address" ||
		fieldNames[2] != "Number" {
		t.Error("Expected first field names to be 'Name', 'Address' & 'Number' but they are ",
			fieldNames[0], fieldNames[1], fieldNames[2])
	}

	// Non struct
	s3 := ""

	fieldNames, err = reflections.GetFieldNames(reflect.TypeOf(s3))

	expectedError := "Expected type to be a struct, instead got " + reflect.TypeOf(s3).Kind().String()
	if err.Error() != expectedError {
		t.Error("Expected error '" + expectedError + "', instead got " + err.Error())
	}
}

func TestGetTagNames(t *testing.T) {
	// Single tag name
	type foo struct {
		Name string `json:"name"`
	}

	tagNames, err := reflections.GetTagNames(reflect.TypeOf(foo{}), "json")

	if err != nil {
		t.Error(err)
	}

	if tagNames[0] != "name" {
		t.Error("Expected first string element to be equal to 'name', but got ", tagNames[0])
	}

	// Multiple tag keys in a struct's field
	type foo2 struct {
		Name string `json:"name" bson:"_name"`
	}

	tagNames, err = reflections.GetTagNames(reflect.TypeOf(foo2{}), "bson")

	if err != nil {
		t.Error(err)
	}

	if tagNames[0] != "_name" {
		t.Error("Expected first string element to be equal to '_name', but got ", tagNames[0])
	}

	// Multiple fields with tag names
	type foo3 struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
	}

	tagNames, err = reflections.GetTagNames(reflect.TypeOf(foo3{}), "json")

	if err != nil {
		t.Error(err)
	}

	if tagNames[0] != "name" && tagNames[1] != "surname" {
		t.Error("Expected string elements to be equal to 'name' & 'surname', but got ", tagNames[0], tagNames[1])
	}

}

func TestConvertFieldNamesToTagNames(t *testing.T) {
	// Normal
	type foo struct {
		Username string `json:"uname"`
		Password string `json:"pass"`
	}

	m := map[string]interface{}{
		"Username": "joe123",
		"Password": "easy",
	}

	tagsM, err := reflections.ConvertFieldNamesToTagNames(
		m,
		reflect.TypeOf(foo{}),
		"json",
	)

	if err != nil {
		t.Error(err)
	}

	didChangeToTagNames := tagsM["uname"] != m["Username"] && tagsM["pass"] != m["Password"]
	if didChangeToTagNames {
		t.Error("Expected converted map's keys to have same values with the map, instead it's value is", tagsM)
	}
}
