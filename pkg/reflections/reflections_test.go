package reflections_test

import (
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/pkg/reflections"
)

func TestSetField(t *testing.T) {
	type testStruct struct {
		Name string
	}

	s := testStruct{
		Name: "jack",
	}

	err := reflections.SetField(&s, "Name", "john")

	if err != nil {
		t.Error(err)
	}

	if s.Name != "john" {
		t.Error("Expected 'Name' to be equal to 'john', but is " + s.Name)
	}

	// Check for non struct type
	invalidInterface := "string"

	err = reflections.SetField(&invalidInterface, "Name", "")

	if err == nil {
		t.Error("Expected an error for invalid type, but no error was returned")
	}

	if err.Error() != "Interface is not a struct" {
		t.Error("Expected an error for invalid interface type, but got ", err.Error())
	}

	// Get a non existant field
	s = testStruct{}

	err = reflections.SetField(&s, "NotName", "")

	if err.Error() != "Field is not valid" {
		t.Error("Expected invalid field error but got", err)
	}
}

func TestGetField(t *testing.T) {
	type testStruct struct {
		Name string
	}

	// Check for 'john'
	s := testStruct{
		Name: "john",
	}

	name, err := reflections.GetField(&s, "Name")

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

	name, err = reflections.GetField(&s, "Name")

	if err != nil {
		t.Error(err)
	}

	if name.(string) != "" {
		t.Error("Expected 'Name' to be equal to '', but it's ", name)
	}

	// Check for non struct type
	invalidInterface := "string"

	name, err = reflections.GetField(&invalidInterface, "Name")

	if err == nil {
		t.Error("Expected an error for invalid type, but no error was returned")
	}

	if err.Error() != "Interface is not a struct" {
		t.Error("Expected an error for invalid interface type, but got ", err.Error())
	}

	// Get a non existant field
	s = testStruct{}

	name, err = reflections.GetField(&s, "NotName")

	if err.Error() != "Field is not valid" {
		t.Error("Expected invalid field error but got", err)
	}
}

func TestGetFieldNames(t *testing.T) {
	// Struct with one field
	type testingStruct struct {
		Name string
	}

	s := testingStruct{}

	fieldNames, err := reflections.GetFieldNames(&s)

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

	fieldNames, err = reflections.GetFieldNames(&s2)

	if err != nil {
		t.Error(err)
	}

	if fieldNames[0] != "Name" || fieldNames[1] != "Address" || fieldNames[2] != "Number" {
		t.Error("Expected first field names to be 'Name', 'Address' & 'Number' but they are ",
			fieldNames[0], fieldNames[1], fieldNames[2])
	}

	// Non struct
	s3 := ""

	fieldNames, err = reflections.GetFieldNames(&s3)

	if err.Error() != "Interface is not a struct" {
		t.Error("Expected 'interface is not struct' error, instead got", err.Error())
	}
}
