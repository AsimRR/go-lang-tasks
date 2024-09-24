package main

import (
	"fmt"
	"time"
)

const AvailableTask int = 100

// Task priority levels
const (
	Low int = iota
	Medium
	High
)

const (
	Active     string = "Active"
	InProgress string = "InProgress"
	Completed  string = "Completed"
)

func main() {

	var projectName string = "Task Management System"
	var startDate = time.Date(2024, 9, 18, 19, 0, 0, 0, time.UTC)

	fmt.Printf("Welcome to the %s\n", projectName)
	fmt.Println("Project start date is:", startDate)
	fmt.Printf("Project: %s \n", projectName)

	/*I know it must be string, I change it when learning function,
	I write getStatus function and it return string data for status */

	var projectStatus string = InProgress
	var tasksCompleted int = 25

	fmt.Println("Current project status: ", projectStatus)

	fmt.Printf("Tasks completed: %d out of %d\n", tasksCompleted, AvailableTask)

	fmt.Printf("Task priorities: %d-Low, %d-Medium, %d-High\n", Low, Medium, High)

	isCompleted := projectStatus == "Completed"

	fmt.Printf("Is the project completed? %v", isCompleted)
}
