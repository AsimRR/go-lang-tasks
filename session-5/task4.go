package main

import "fmt"

func main() {
	a, b, c, d := 1, 2, 3, 4

	fmt.Println("Before doubling:", a, b, c, d)

	doubleVariadicElements(&a, &b, &c, &d)

	fmt.Println("After doubling:", a, b, c, d)
}

func doubleVariadicElements(numbers ...*int) {
	for _, num := range numbers {
		*num = *num * 2
	}
}
