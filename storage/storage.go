package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Create a type struct for tasks
type Task struct {
	ID int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// tasklist struct that holds all the task
type TaskList struct {
	Tasks []Task `json:"tasks"`
}

// Filename
const filename = "tasks.json"

// Load all tasks 
func LoadTasks() (*TaskList, error) {
	// Create an empty tasklist in memeory that returns a pointer
	tasklist := &TaskList{Tasks: []Task{}}

	// Check if the file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return tasklist, nil
	}

	// Read the file
	data, err := os.ReadFile(filename) 
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	
	if len(data) == 0 {
		return  tasklist, nil
	}

	// Unmarshal from json into struct type
	err = json.Unmarshal(data, tasklist)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return tasklist, nil
}

// Save ttasks
func SaveTasks(tasklist *TaskList) error {
	// Marshal the struct created into json file
	data, err := json.MarshalIndent(tasklist, "", " ")
	if err != nil {
		return fmt.Errorf("error creating JSON: %v", err)
	}

	// Write the savedTask to the file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}

// ID generator
func GenNextID (tasklist *TaskList) int {
	maxID := 0

	// Loop over each task, compare the ID and increment the id
	for _, task := range tasklist.Tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return  maxID + 1
}
