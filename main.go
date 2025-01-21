package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func AddTask(Description string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	newTask := Task{
		ID:          len(tasks) + 1,
		Description: Description,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	return SaveTasks(tasks)
}

func LoadTasks() ([]Task, error) {
	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	return tasks, err
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.jason", data, 0644)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add <description>")
			return
		}
		description := os.Args[2]
		err := AddTask(description)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Task added!!!")
		}

	case "list":
		tasks, err := LoadTasks()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for _, task := range tasks {
			fmt.Printf("%d: %s [%s]\n", task.ID, task.Description, task.Status)
		}

	default:
		fmt.Println("Unknown command:", command)
	}
}
