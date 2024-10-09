package main

import (
	"errors"
	"fmt"
)

const AvailableTasks = 100

var (
	errExceededTasks = errors.New("error: completedTasks exceeds total available tasks")
	taskList         = []string{"TASK 1", "TASK 2", "TASK 3", "TASK 4", "TASK 5"}
)

func main() {
	defer recoverFromPanic()

	completedTasks := 105

	validateCompletedTasks(completedTasks)

	isCompleted := checkCompletion(completedTasks)
	remainingTasks, isNearCompletion := getRemainingTasks(completedTasks)
	projectPhase := getProjectPhase(completedTasks)

	if isNearCompletion {
		fmt.Println("More than 90 tasks completed")
	}

	fmt.Printf("Project is in the: %s\n", projectPhase)

	showTaskList(taskList...)

	performCountdown(remainingTasks)
	fmt.Printf("Is the project completed? %v\n", isCompleted)
}

func validateCompletedTasks(completedTasks int) {
	if completedTasks > AvailableTasks {
		panic(errExceededTasks)
	}
}

func recoverFromPanic() {
	if recoveryMessage := recover(); recoveryMessage != nil {
		fmt.Println(recoveryMessage)
	}
}

func getRemainingTasks(completedTasks int) (int, bool) {
	remainingTasks := AvailableTasks - completedTasks
	isNearCompletion := completedTasks >= 90
	return remainingTasks, isNearCompletion
}

func showTaskList(tasks ...string) {
	for i, task := range tasks {
		fmt.Printf("%d: %s\n", i+1, task)
	}
}

func getProjectPhase(completedTasks int) string {
	switch {
	case completedTasks < 30:
		return "Starting phase"
	case completedTasks >= 30 && completedTasks < 60:
		return "Midway"
	case completedTasks >= 60:
		return "Final phase"
	default:
		return "Unknown phase"
	}
}

func checkCompletion(completedTasks int) bool {
	fmt.Println("Status check completed")
	return AvailableTasks == completedTasks
}

func performCountdown(tasksRemaining int) {
	if tasksRemaining == 0 {
		fmt.Println("All tasks completed!")
		return
	}
	fmt.Printf("Tasks remaining: %d\n", tasksRemaining)
	performCountdown(tasksRemaining - 1)
}
