package using_reflection

import (
	"fmt"
	"reflect"
)

func InvokeMethod(value interface{}, methodName string) {

	v := reflect.ValueOf(value)

	method := v.MethodByName(methodName)

	result := method.Call(nil)

	fmt.Printf("Invoked method: %s, Result: %s\n", methodName, result[0].String())
}
