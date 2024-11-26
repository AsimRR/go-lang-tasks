package reflection

import (
	"fmt"
	"reflect"
)

func InspectStruct(value interface{}) {

	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		fmt.Printf("Field Name: %s, Type: %s, Value: %v\n", field.Name, field.Type, fieldValue)
	}
}
