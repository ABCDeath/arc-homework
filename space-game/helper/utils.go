package helper

import "reflect"

func GetStructTypeName(instance interface{}) string {
	return GetStructType(instance).Name()
}

func GetStructType(instance interface{}) reflect.Type {
	actualValue := reflect.ValueOf(instance).Elem()

	return actualValue.Type()
}
