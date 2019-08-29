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
		t.Error("Expected 'Name' to be equal to 'john', but is ", name)
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
		t.Error("Expected 'Name' to be equal to '', but is ", name)
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
