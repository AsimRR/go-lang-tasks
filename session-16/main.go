package main

import (
	"fmt"
	"session-16/task/go_generics"
	"session-16/task/reflection"
	"session-16/task/using_reflection"
)

type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) Greet() string {
	return "Hello, I am " + p.Name
}

func main() {
	p := Person{Name: "Ali", Age: 30, City: "Baku"}

	fmt.Println("*******TASK1*******")
	reflection.IdentifyTypeAndKind(42)
	reflection.IdentifyTypeAndKind("Hello")
	reflection.IdentifyTypeAndKind([]int{1, 2, 3})
	reflection.IdentifyTypeAndKind(true)

	fmt.Println("*******TASK2*******")

	reflection.InspectStruct(p)

	fmt.Println("*******TASK3*******")
	using_reflection.SetFieldValue(&p, "City", "Khazak")
	fmt.Println(p)

	fmt.Println("*******TASK4*******")
	using_reflection.InvokeMethod(p, "Greet")

	fmt.Println("*******TASK5*******")
	fmt.Println(go_generics.Sum([]int{1, 2, 3, 4}))
	fmt.Println(go_generics.Sum([]float64{1.5, 2.5, 3.5}))

	fmt.Println("*******TASK6*******")
	min, max := go_generics.MinMax([]int{10, 20, 5, 8, 15})
	fmt.Println(min, max)

}
