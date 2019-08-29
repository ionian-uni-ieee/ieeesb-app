package reflections

import (
	"errors"
	"reflect"
)

// SetField sets a map's field value to a new value
func SetField(d interface{}, key string, value interface{}) error {
	r := reflect.ValueOf(d)

	if r.Kind() == reflect.Struct {
		return errors.New("Struct is not a pointer")
	}

	s := r.Elem()

	if s.Kind() == reflect.Struct {
		field := s.FieldByName(key)
		if field.IsValid() {
			if field.CanSet() {
				valueReflection := reflect.ValueOf(value)
				if field.Kind() == valueReflection.Kind() {
					field.Set(valueReflection)
					return nil
				}
				return errors.New("Value type is not the same type with struct's type")
			}
			return errors.New("Field cannot be set")
		}
		return errors.New("Field is not valid")
	}
	return errors.New("Interface is not a struct")
}

// GetField returns a value of a field of a map
func GetField(d interface{}, key string) (interface{}, error) {
	r := reflect.ValueOf(d)

	if r.Kind() == reflect.Struct {
		return nil, errors.New("Struct is not a pointer")
	}

	s := r.Elem()

	if s.Kind() == reflect.Struct {
		field := reflect.Indirect(r).FieldByName(key)
		if field.IsValid() {
			return field.Interface(), nil
		}
		return nil, errors.New("Field is not valid")
	}
	return nil, errors.New("Interface is not a struct")

}

// GetFieldNames returns all the field keys of a struct
func GetFieldNames(d interface{}) ([]string, error) {
	r := reflect.ValueOf(d)

	if r.Kind() == reflect.Struct {
		return nil, errors.New("Struct is not a pointer")
	}

	s := r.Elem()

	fieldNames := []string{}

	if s.Kind() == reflect.Struct {
		for fieldIndex := 0; fieldIndex < s.NumField(); fieldIndex++ {
			fieldName := s.Type().Field(fieldIndex).Name
			fieldNames = append(fieldNames, fieldName)
		}
	} else {
		return nil, errors.New("Interface is not a struct")
	}

	return fieldNames, nil
}
