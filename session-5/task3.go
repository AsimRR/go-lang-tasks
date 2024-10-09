package main

import "fmt"

func main() {
	var x = 10
	var y = 20

	fmt.Printf("Before swapping: x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swapping: x=%d, y=%d\n", x, y)

}

func swap(x, y *int) {
	*x, *y = *y, *x
}
