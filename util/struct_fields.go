package util

import (
	"reflect"
)

func GetStructFields(o interface{}) []string {
	v := reflect.Indirect(reflect.ValueOf(&o)).Elem().Type()
	r := make([]string, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		r[i] = v.Field(i).Name
	}

	return r
}
