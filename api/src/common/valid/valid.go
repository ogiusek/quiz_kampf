package valid

import (
	"reflect"
)

func Valid(object any) error {
	if customValid, ok := object.(Validable); ok {
		return customValid.Valid()
	}

	val := reflect.ValueOf(object)
	var fields = make([]Validable, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if vo, ok := field.Interface().(Validable); ok {
			fields = append(fields, vo)
		}
	}

	res := ValidMany(fields)
	return res
}
