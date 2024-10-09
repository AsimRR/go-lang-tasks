package main

import "fmt"

func main() {

	var x int = 5
	fmt.Printf("Value of x before incrementing by value: %d \n", x)
	incrementByValue(x)
	fmt.Printf("Value of x after incrementing by value: %d \n", x)

	fmt.Printf("Value of x before incrementing by pointer: %d \n", x)
	incrementByPointer(&x)
	fmt.Printf("Value of x after incrementing by pointer: %d \n", x)
}

func incrementByValue(val int) {
	val++
}

func incrementByPointer(val *int) {
	*val++
}
