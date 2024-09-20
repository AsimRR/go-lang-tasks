package main

import (
	"fmt"
	"time"
)

func main() {
	var projectName string = "Task Management System"

	var startDate = time.Date(2024, 9, 18, 19, 0, 0, 0, time.UTC)

	fmt.Printf("Welcome to the %s\n", projectName)

	fmt.Println("Project start date is:", startDate)

	fmt.Printf("Project: %s", projectName)

}
