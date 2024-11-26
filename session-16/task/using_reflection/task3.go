package using_reflection

import (
	"reflect"
)

func SetFieldValue(value interface{}, fieldName string, newValue interface{}) {
	v := reflect.ValueOf(value)

	elements := v.Elem()
	field := elements.FieldByName(fieldName)

	newValueReflect := reflect.ValueOf(newValue)

	field.Set(newValueReflect)
}
