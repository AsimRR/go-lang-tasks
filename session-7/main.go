package main

import (
	"fmt"
	"session-7/model/embedding"
	"session-7/model/methods"
	"session-7/model/structs"
	"session-7/utils"
)

func main() {
	// TASK1
	book := structs.Book{
		Title:  "Go Programming Language",
		Author: "Alan A. A. Donovan",
		Pages:  100,
	}

	structs.ShowBooks(book)

	fmt.Println("*****************")

	// TASK2
	rectangle := structs.Rectangle{
		Width:  10.5,
		Height: 5,
	}

	structs.CalculateRectangle(rectangle)

	fmt.Println("*****************")

	// TASK3

	circle := methods.Circle{
		Radius: 7,
	}

	area := circle.Area()

	fmt.Printf("Circle Radius: %f\nArea: %f\n", circle.Radius, area)

	fmt.Println("*****************")

	// TASK4

	std1 := methods.Student{
		Name:   "Ali",
		Grades: []int{10, 20, 30, 40, 50},
	}

	std2 := methods.Student{
		Name:   "Vali",
		Grades: []int{15, 25, 35, 45, 55},
	}
	utils.CompareStudents(std1, std2)

	// TASK5
	employee := embedding.Employee{
		EmployeeID: 12345,
		Position:   "Software Engineer",
		Person: embedding.Person{
			FirstName: "Ali",
			LastName:  "Aliyev",
		},
	}

	employee.GetEmployeeDetails()

	fmt.Println("*****************")

	// TASK6
	myCar := embedding.Car{
		Vehicle: embedding.Vehicle{
			Make:  "Kia",
			Model: "K5",
			Year:  2022,
		},
		FuelType: "Gasoline",
	}

	myCar.GetCarDetails()

}
