package main

import (
	"errors"
	"fmt"
	"time"
)

const AvailableTask int = 100

const (
	_ = iota
	Low
	Medium
	High
)

var errNoTasks = errors.New("No tasks completed")

func main() {

	projectName := "Task Management System"
	startDate := time.Date(2024, 9, 18, 19, 0, 0, 0, time.UTC)

	fmt.Printf("Welcome to the %s\n", projectName)
	fmt.Printf("Project start date is: %v\n", startDate)
	fmt.Printf("Project: %s\n", projectName)

	fmt.Println("*****************************")

	completedTask := 75
	remainingTask := AvailableTask - completedTask

	if completedTask > AvailableTask {
		fmt.Printf("Error: Completed tasks (%d) exceed the total available tasks (%d) \n", completedTask, AvailableTask)
		completedTask = AvailableTask
		remainingTask = 0
	}

	if completedTask == 0 {
		err := errNoTasks
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

	fmt.Printf("Tasks remaining: %d out of %d\n", remainingTask, AvailableTask)

	isNearCompleted := completedTask >= 90
	isCompleted := completedTask == AvailableTask

	status := "In Progress"

	if isNearCompleted {
		fmt.Println("More than 90 tasks completed")
	}

	if isCompleted {
		status = "Completed"
	}

	fmt.Println("Current project status:", status)

	var projectPhase string
	switch {
	case completedTask < 30:
		projectPhase = "Starting phase"
	case completedTask >= 30 && completedTask < 60:
		projectPhase = "Midway"
	case completedTask >= 60:
		projectPhase = "Final phase"
	default:
		projectPhase = "Unknown phase"
	}

	fmt.Println("Project is in the:", projectPhase)

	fmt.Printf("Task priorities: %d-Low, %d-Medium, %d-High\n", Low, Medium, High)

	for i := 1; i <= remainingTask; i++ {
		fmt.Printf("Task %d\n", i)
	}

	fmt.Printf("Is the project completed? %v\n", isCompleted)
}
