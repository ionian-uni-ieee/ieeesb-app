package reflections

import (
	"errors"
	"fmt"
	"reflect"
)

// SetField sets a map's field value to a new value
func SetField(d interface{}, key string, value interface{}) error {
	r := reflect.ValueOf(d)
	s := r.Elem()

	if s.Kind() == reflect.Struct {
		field := s.FieldByName(key)
		if field.IsValid() {
			if field.CanSet() {
				valueReflection := reflect.ValueOf(value)
				if field.Kind() == valueReflection.Kind() {
					fmt.Println(r, field, valueReflection)
					field.Set(valueReflection)
					fmt.Println(r, field, valueReflection)
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
