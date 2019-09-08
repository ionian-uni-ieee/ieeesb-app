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

	t.Run("Should set field value", func(t *testing.T) {

		s := testStruct{
			Name: "jack",
		}

		gotErr := reflections.SetFieldValue(&s, "Name", "john")

		if gotErr != nil {
			t.Error(gotErr)
		}

		if s.Name != "john" {
			t.Error("Expected 'Name' to be equal to 'john', but is " + s.Name)
		}
	})

	// Check for non struct type
	t.Run("Should return no interface pointer error", func(t *testing.T) {
		invalidInterface := "string"

		gotErr := reflections.SetFieldValue(&invalidInterface, "Name", "")

		if gotErr == nil {
			t.Error("Expected an error for invalid type, but no error was returned")
		}

		expectedError := "Pointed interface is not a struct"
		if gotErr.Error() != expectedError {
			t.Error("Expected '"+expectedError+"' error, but got ", gotErr.Error())
		}
	})

	// Get a non existant field
	t.Run("Should return field is not valid error", func(t *testing.T) {
		s := testStruct{}

		gotErr := reflections.SetFieldValue(&s, "NotName", "")

		if gotErr.Error() != "Field is not valid" {
			t.Error("Expected invalid field error but got", gotErr)
		}
	})

	t.Run("Should return interface is not a pointer error", func(t *testing.T) {
		s := testStruct{}

		gotErr := reflections.SetFieldValue(s, "Name", "")

		expectedError := "Interface is not a pointer"
		if gotErr.Error() != expectedError {
			t.Error("Expected '"+expectedError+"' error, but got ", gotErr.Error())
		}
	})
}

func TestGetFieldValue(t *testing.T) {
	type testStruct struct {
		Name string
	}

	t.Run("Should return 'john'", func(t *testing.T) {
		s := testStruct{
			Name: "john",
		}

		gotName, gotErr := reflections.GetFieldValue(s, "Name")

		if gotErr != nil {
			t.Error(gotErr)
		}

		if gotName.(string) != "john" {
			t.Error("Expected 'Name' to be equal to 'john', but it's ", gotName)
		}
	})

	t.Run("Should return empty string", func(t *testing.T) {
		s := testStruct{
			Name: "",
		}

		gotName, gotErr := reflections.GetFieldValue(s, "Name")

		if gotErr != nil {
			t.Error(gotErr)
		}

		if gotName.(string) != "" {
			t.Error("Expected 'Name' to be equal to '', but it's ", gotName)
		}
	})

	t.Run("Should return interface is not a struct error", func(t *testing.T) {
		invalidInterface := "string"

		_, gotErr := reflections.GetFieldValue(invalidInterface, "Name")

		if gotErr.Error() != "Interface is not a struct" {
			t.Error("Expected an error for invalid interface type, but got ", gotErr.Error())
		}
	})

	t.Run("Should return field is not valid error", func(t *testing.T) {
		s := testStruct{}

		_, gotErr := reflections.GetFieldValue(s, "NotName")

		if gotErr.Error() != "Field is not valid" {
			t.Error("Expected invalid field error but got", gotErr)
		}
	})
}

func TestGetTagByFieldName(t *testing.T) {
	t.Run("Should return 'uname' tag name", func(t *testing.T) {
		type foo struct {
			Username string `json:"uname"`
		}

		gotTag, gotErr := reflections.GetTagByFieldName(reflect.TypeOf(foo{}), "Username", "json")

		if gotErr != nil {
			t.Error(gotErr)
		}

		if gotTag != "uname" {
			t.Error("Expected tag name 'uname' but instead got " + gotTag)
		}
	})
}

func TestGetFieldNames(t *testing.T) {
	t.Run("Should return array with 1 field name", func(t *testing.T) {
		type testingStruct struct {
			Name string
		}

		s := testingStruct{}

		gotFieldNames, gotErr := reflections.GetFieldNames(reflect.TypeOf(s))

		if gotErr != nil {
			t.Error(gotErr)
		}

		if len(gotFieldNames) != 1 {
			t.Error("Expected array length of 1 but got", len(gotFieldNames))
		}

		if gotFieldNames[0] != "Name" {
			t.Error("Expected first field name to be 'Name' but it's ", gotFieldNames[0])
		}
	})

	t.Run("Should return all field names", func(t *testing.T) {
		type testingStruct2 struct {
			Name    string
			Address string
			Number  int
		}

		s2 := testingStruct2{}

		gotFieldNames, gotErr := reflections.GetFieldNames(reflect.TypeOf(s2))

		if gotErr != nil {
			t.Error(gotErr)
		}

		if gotFieldNames[0] != "Name" ||
			gotFieldNames[1] != "Address" ||
			gotFieldNames[2] != "Number" {
			t.Error("Expected first field names to be 'Name', 'Address' & 'Number' but they are ",
				gotFieldNames[0], gotFieldNames[1], gotFieldNames[2])
		}
	})

	t.Run("Should return type is not a struct error", func(t *testing.T) {
		s3 := ""

		_, gotErr := reflections.GetFieldNames(reflect.TypeOf(s3))

		expectedError := "Expected type to be a struct, instead got " + reflect.TypeOf(s3).Kind().String()
		if gotErr.Error() != expectedError {
			t.Error("Expected error '" + expectedError + "', instead got " + gotErr.Error())
		}
	})
}

func TestGetTagNames(t *testing.T) {
	t.Run("Should return tag name", func(t *testing.T) {
		type foo struct {
			Name string `json:"name"`
		}

		gotTagNames, err := reflections.GetTagNames(reflect.TypeOf(foo{}), "json")

		if err != nil {
			t.Error(err)
		}

		if gotTagNames[0] != "name" {
			t.Error("Expected first string element to be equal to 'name', but got ", gotTagNames[0])
		}
	})

	t.Run("Should return the correct tag name", func(t *testing.T) {
		type foo2 struct {
			Name string `json:"name" bson:"_name"`
		}

		gotTagNames, err := reflections.GetTagNames(reflect.TypeOf(foo2{}), "bson")

		if err != nil {
			t.Error(err)
		}

		if gotTagNames[0] != "_name" {
			t.Error("Expected first string element to be equal to '_name', but got ", gotTagNames[0])
		}
	})

	t.Run("Should return all tag names", func(t *testing.T) {
		type foo3 struct {
			Name    string `json:"name"`
			Surname string `json:"surname"`
		}

		gotTagNames, err := reflections.GetTagNames(reflect.TypeOf(foo3{}), "json")

		if err != nil {
			t.Error(err)
		}

		if gotTagNames[0] != "name" && gotTagNames[1] != "surname" {
			t.Error("Expected string elements to be equal to 'name' & 'surname', but got ", gotTagNames[0], gotTagNames[1])
		}

	})
}

func TestConvertFieldNamesToTagNames(t *testing.T) {
	t.Run("Should return map with converted fields to tag names", func(t *testing.T) {
		type foo struct {
			Username string `json:"uname"`
			Password string `json:"pass"`
		}

		m := map[string]interface{}{
			"Username": "joe123",
			"Password": "easy",
		}

		gotTagMap, err := reflections.ConvertFieldNamesToTagNames(
			m,
			reflect.TypeOf(foo{}),
			"json",
		)

		if err != nil {
			t.Error(err)
		}

		didChangeToTagNames := gotTagMap["uname"] != m["Username"] &&
			gotTagMap["pass"] != m["Password"]
		if didChangeToTagNames {
			t.Error("Expected converted map's keys to have same values with the map, instead it's value is", gotTagMap)
		}
	})
}

func TestIsFieldValid(t *testing.T) {
	t.Run("Should return that existing field is valid", func(t *testing.T) {
		type foo struct {
			Name string
		}

		gotFieldExists := reflections.IsFieldValid(reflect.TypeOf(foo{}), "Name")

		if !gotFieldExists {
			t.Error("Expected fieldExists to return true")
		}
	})

	t.Run("Should return that non-existant field is invalid", func(t *testing.T) {
		type foo struct {
			Name string
		}

		gotFieldExists := reflections.IsFieldValid(reflect.TypeOf(foo{}), "Password")

		if gotFieldExists {
			t.Error("Expected fieldExists to return false")
		}
	})
}

func TestAreMapFieldsValid(t *testing.T) {
	t.Run("Should return that map is valid", func(t *testing.T) {
		type foo struct {
			Name string
		}

		m := map[string]interface{}{
			"Name": "joe",
		}

		gotIsMapValid := reflections.AreMapFieldsValid(
			reflect.TypeOf(foo{}),
			m,
		)

		if !gotIsMapValid {
			t.Error("Expected isMapValid to return true")
		}
	})

	t.Run("Should return that map with mixed non-existant fields is invalid", func(t *testing.T) {
		type foo struct {
			Name string
		}

		m := map[string]interface{}{
			"Name":     "joe",
			"Password": "passjoe",
		}

		gotIsMapValid := reflections.AreMapFieldsValid(
			reflect.TypeOf(foo{}),
			m,
		)

		if gotIsMapValid {
			t.Error("Expected isMapValid to return false")
		}
	})

	t.Run("Should return that map with non-existant fields is invalid", func(t *testing.T) {
		type foo struct {
			Name string
		}

		m := map[string]interface{}{
			"Password": "passjoe",
		}

		gotIsMapValid := reflections.AreMapFieldsValid(
			reflect.TypeOf(foo{}),
			m,
		)

		if gotIsMapValid {
			t.Error("Expected isMapValid to return false")
		}
	})
}

func TestCreateEmptyStructInstance(t *testing.T) {
	t.Run("Should return an empty struct", func(t *testing.T) {
		type foo struct {
			Name string
		}

		gotInstance := reflections.CreateEmptyInstance(
			reflect.TypeOf(foo{}),
		).(foo)

		if gotInstance.Name != "" {
			t.Error("Expected empty instance name, but got " + gotInstance.Name)
		}
	})
}
