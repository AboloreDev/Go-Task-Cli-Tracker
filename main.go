package main

import (
	"fmt"
	"os"
	"task-tracker/helpers"
	progresstrackers "task-tracker/progress-trackers"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		helpers.AddNewTask()
	case "update":
		helpers.UpdateTask()
	case "delete":
		helpers.DeleteTask()
	case "mark-in-progress":
		progresstrackers.MarkInProgress()
	case "mark-done":
		progresstrackers.MarkDone()
	case "list":
		progresstrackers.List()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Task Tracker CLI")
	fmt.Println("\nUsage:")
	fmt.Println("  task-tracker add <description>           - Add a new task")
	fmt.Println("  task-tracker update <id> <description>   - Update a task")
	fmt.Println("  task-tracker delete <id>                 - Delete a task")
	fmt.Println("  task-tracker mark-in-progress <id>       - Mark task as in progress")
	fmt.Println("  task-tracker mark-done <id>              - Mark task as done")
	fmt.Println("  task-tracker list                        - List all tasks")
	fmt.Println("  task-tracker list done                   - List completed tasks")
	fmt.Println("  task-tracker list todo                   - List pending tasks")
	fmt.Println("  task-tracker list in-progress            - List in-progress tasks")

}
