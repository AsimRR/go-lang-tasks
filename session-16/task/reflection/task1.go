package reflection

import (
	"fmt"
	"reflect"
)

func IdentifyTypeAndKind(value interface{}) {

	typeOfValue := reflect.TypeOf(value)
	kindOfValue := typeOfValue.Kind()

	fmt.Printf("Value: %v, Type: %s, Kind: %s\n", value, typeOfValue, kindOfValue)
}
