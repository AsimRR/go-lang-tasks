package main

import "fmt"

func main() {

	var x int = 10
	var ptr *int = &x

	fmt.Printf("Value of x: %d \n", x)
	fmt.Printf("Adress of x: %d \n", &x)
	fmt.Printf("Value at pointer: %d \n", *ptr)
}
