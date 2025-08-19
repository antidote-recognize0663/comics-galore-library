package utils

import (
	"fmt"
	"reflect"
)

func Filter[T any](slice []T, condition func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

// StructToMap transforms any struct into a map based on its fields and JSON tags
func StructToMap(input interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(input)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be a struct or a pointer to a struct")
	}
	result := make(map[string]interface{})
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i).Interface()
		jsonKey := field.Tag.Get("json")
		if jsonKey == "" {
			jsonKey = field.Name
		}
		result[jsonKey] = fieldValue
	}
	return result, nil
}
