package helpers

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker/storage"
	"time"
)

// AddNewTask
// UpdateTask
// DeleteTask

// os.Args[1]
// os.Args[2]
// os.Args[3]
// os.Args[4]

func AddNewTask() {
	// Check for the command length using os.Ags
	if len(os.Args) < 3 {
		fmt.Println("Error: Please provide a task description")
		fmt.Println("Usage: task-tracker add <description>")
		os.Exit(1)
	}

	// Load the task
	tasklist, err := storage.LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	// Create a new task
	newTask := storage.Task{
		ID:          storage.GenNextID(tasklist),
		Description: os.Args[2],
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Append the new task
	tasklist.Tasks = append(tasklist.Tasks, newTask)

	// Sve the task
	err = storage.SaveTasks(tasklist)
	if err != nil {
		fmt.Printf("Failed to save tasks: %v\n", err)
		os.Exit(1)
	}

	// Print success message
	fmt.Printf("Task successfully added: %v", newTask.ID)
}

func UpdateTask() {
	// this is taking 4 args
	if len(os.Args) < 4 {
		fmt.Println("Error: Please provide a task id")
		fmt.Println("Usage: task-tracker update <id> <description>")
		os.Exit(1)
	}

	// convert the id as a string into int
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: TaskID must be a string")
	}

	// Load the task
	tasklist, err := storage.LoadTasks()
	if err != nil {
		fmt.Printf("Failed to load task: %v", err)
		os.Exit(1)
	}

	// create a statement that loops over the task list and uses the id to update the whole data
	// to update we create a copy
	found := false
	for i := range tasklist.Tasks {
		if tasklist.Tasks[i].ID == id {
			tasklist.Tasks[i].Description = os.Args[3]
			tasklist.Tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Task not found with id: %v", id)
	}

	// Save the task
	err = storage.SaveTasks(tasklist)
	if err != nil {
		fmt.Printf("Failed to save task: %v", err)
	}

	// Return a success statement
	// Print if successful
	fmt.Printf("Task %d updated successfully\n", id)
}

func DeleteTask() {
	if len(os.Args) < 3 {
		fmt.Println("Error: Please provide a task id")
		fmt.Println("Usage: task-tracker delete <id>")
		os.Exit(1)
	}

	// Conver to string
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Task ID must be a number")
	}

	// Load the task
	tasklist, err := storage.LoadTasks()
	if err != nil {
		fmt.Printf("Unable to load task: %v", err)
		os.Exit(1)
	}

	// To delete a task we cant range over the task list and delete, we create a copy and delete
	found := false
	newTask := []storage.Task{}
	for _, task := range tasklist.Tasks {
		if task.ID == id {
			found = true
			continue
		}
		// Append task to the new task
		// Never append to the same slice youre ranging over
		newTask = append(newTask, task)
	}

	if !found {
		fmt.Printf("Task with the ID %v not found", id)
	}

	tasklist.Tasks = newTask

	// save the task
	err = storage.SaveTasks(tasklist)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task %d deleted successfully\n", id)
}
