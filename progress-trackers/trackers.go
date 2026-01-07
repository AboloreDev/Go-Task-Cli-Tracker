package progresstrackers

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker/storage"
	"time"
)

// Status change
// In- progres
// done
// list all trackers

func UpdateStatusChange(id int, status string) {
	tasklist, err := storage.LoadTasks()
	if err != nil {
		fmt.Printf("Unable to load tasks: %v", err)
		os.Exit(1)
	}

	// Loop over the task using index
	found := false
	for i := range tasklist.Tasks {
		if tasklist.Tasks[i].ID == id {
			tasklist.Tasks[i].Status = status
			tasklist.Tasks[i].UpdatedAt = time.Now()
			break
		}
	}

	if !found {
		fmt.Printf("Task with the id %v, not found", err)
	}

	// save task
	err = storage.SaveTasks(tasklist)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task %d marked as %s\n", id, status)
}

// handle mark in progress
func MarkInProgress() {
	// Validation
	if len(os.Args) < 3{
		fmt.Println("Error: Please provide task ID")
		fmt.Println("Usage: task-tracker mark-in-progress <id>")
		os.Exit(1)
	}

	// coNvert the id to int
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Task id must be a number")
	}

	// Update the task
	UpdateStatusChange(id, "in-progress")
}

// Handle done
func MarkDone() {
	if len(os.Args) < 3 {
	fmt.Println("Error: Please provide task ID")
		fmt.Println("Usage: task-tracker done <id>")
		os.Exit(1)	
	}

	// convert the id to string
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Task must be a number")
	}

	UpdateStatusChange(id, "done")
}

// list all tasks
func List() {
	// Load the task
	tasklist, err := storage.LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks %v\n", err)
		os.Exit(1)
	}

	filter := "all"
	if len(os.Args) >= 3 {
		filter = os.Args[2]
	}

	// Create a slice of task for filtered task
	var filteredTasks []storage.Task
	// Range over each task
	for _, task := range tasklist.Tasks {
		// Using switch, fiter each list based on the case
		switch filter {
		case "all":
			filteredTasks = append(filteredTasks, task)
		case "done":
			if task.Status == "done" {
				filteredTasks = append(filteredTasks, task)
			}
		case "todo":
			if task.Status == "todo" {
				filteredTasks = append(filteredTasks, task)
			}
		case "in-progress":
			if task.Status == "in-progress" {
				filteredTasks = append(filteredTasks, task)
			}
		default:
			fmt.Printf("Unknown filter: %s\n", filter)
			fmt.Println("Available filters: all, done, todo, in-progress")
			os.Exit(1)
		}
	}

	if len(filteredTasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Printf("\n%-5s %-20s %-15s %-20s\n", "ID", "Description", "Status", "Updated")
	fmt.Println("----------------------------------------------------------------------")

	for _, task := range filteredTasks {
		desc := task.Description
		if len(desc) > 20 {
			desc = desc[:17] + "..."
		}
		fmt.Printf("%-5d %-20s %-15s %-20s\n",
			task.ID,
			desc,
			task.Status,
			task.UpdatedAt.Format("2006-01-02 15:04"))
	}
	fmt.Println()
}