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
}

func TestGetField(t *testing.T) {
	type testStruct struct {
		Name string
	}

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
}
