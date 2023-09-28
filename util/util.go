package util

import (
	"reflect"
)

func ObjectAssign(target interface{}, object interface{}) error {
	t := reflect.ValueOf(target).Elem()
	o := reflect.ValueOf(object).Elem()

	for i := 0; i < o.NumField(); i++ {
		fieldName := o.Type().Field(i).Name
		targetField := t.FieldByName(fieldName)
		if targetField.IsValid() && targetField.CanSet() {
			targetField.Set(o.Field(i))
		}
	}

	return nil
}
