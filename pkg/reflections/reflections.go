package reflections

import (
	"errors"
	"reflect"
)

// SetFieldValue sets a map's field value to a new value
func SetFieldValue(d interface{}, key string, value interface{}) error {
	r := reflect.ValueOf(d)

	if r.Kind() != reflect.Ptr {
		return errors.New("Interface is not a pointer")
	}

	s := r.Elem()

	if s.Kind() != reflect.Struct {
		return errors.New("Pointed interface is not a struct")
	}

	field := s.FieldByName(key)

	if !field.IsValid() {
		return errors.New("Field is not valid")
	}

	if !field.CanSet() {
		return errors.New("Field cannot be set")
	}

	valueReflection := reflect.ValueOf(value)

	if field.Kind() != valueReflection.Kind() {
		return errors.New("Value type is not the same type with struct's type")
	}

	field.Set(valueReflection)

	return nil
}

// GetFieldValue returns a value of a field of a map
func GetFieldValue(d interface{}, key string) (interface{}, error) {
	r := reflect.ValueOf(d)

	if r.Kind() != reflect.Struct {
		return nil, errors.New("Interface is not a struct")
	}

	field := reflect.Indirect(r).FieldByName(key)

	if !field.IsValid() {
		return nil, errors.New("Field is not valid")
	}

	return field.Interface(), nil
}

// GetField returns a reflect type of a struct's field
func GetField(rType reflect.Type, key string) (reflect.StructField, error) {
	if rType.Kind() != reflect.Struct {
		return reflect.StructField{}, errors.New("Expected type to be a struct, instead got " + rType.Kind().String())
	}

	field, ok := rType.FieldByName(key)

	if !ok {
		return reflect.StructField{}, errors.New("Field was not found")
	}

	return field, nil
}

// GetTagByFieldName returns a field's tag's name
func GetTagByFieldName(rType reflect.Type, fieldName string, tagKey string) (string, error) {
	field, err := GetField(rType, fieldName)

	if err != nil {
		return "", err
	}

	tag := field.Tag.Get(tagKey)

	return tag, nil
}

// GetFieldNames returns all the field keys of a struct
func GetFieldNames(rType reflect.Type) ([]string, error) {
	if rType.Kind() != reflect.Struct {
		return nil, errors.New("Expected type to be a struct, instead got " + rType.Kind().String())
	}

	fieldNames := []string{}

	for fieldIndex := 0; fieldIndex < rType.NumField(); fieldIndex++ {
		fieldName := rType.Field(fieldIndex).Name
		fieldNames = append(fieldNames, fieldName)
	}

	return fieldNames, nil
}

// GetTagNames returns the struct's fields' tags
func GetTagNames(rType reflect.Type, tagKey string) ([]string, error) {
	if rType.Kind() != reflect.Struct {
		return nil, errors.New("Expected type to be a struct, instead got " + rType.Kind().String())
	}

	tagNames := []string{}

	for fieldIndex := 0; fieldIndex < rType.NumField(); fieldIndex++ {
		tag := rType.Field(fieldIndex).Tag.Get(tagKey)
		tagNames = append(tagNames, tag)
	}

	return tagNames, nil
}

// GetBSONTagNames returns all the bson tags of each struct's field
// It's the same with GetTagNames but calls it with key "bson"
func GetBSONTagNames(rType reflect.Type) ([]string, error) {
	return GetTagNames(rType, "bson")
}

// GetJSONTagNames returns all the bson tags of each struct's field
// It's the same with GetTagNames but calls it with key "json"
func GetJSONTagNames(rType reflect.Type) ([]string, error) {
	return GetTagNames(rType, "json")
}
