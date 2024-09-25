package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"task-cli/model"
	"time"
)

const dataFile = "task.json"

// loadTasks from dataFile
func loadTasks() ([]model.Task, error) {
	data, err := os.ReadFile(dataFile)
	if os.IsNotExist(err) {
		return []model.Task{}, nil
	} else if err != nil {
		return nil, err
	}

	var tasks []model.Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func saveTasks(tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

func generateID(tasks []model.Task) int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

func addTask(description string) {
	tasks, err := loadTasks()
	if err != nil {
		log.Fatalf("Failed to load tasks: %v", err)
	}
	newTask := model.Task{
		ID:          generateID(tasks),
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)
	err = saveTasks(tasks)
	if err != nil {
		log.Fatalf("Failed to save tasks: %v", err)
	}

	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
}

func updateTask(id int, newDescription string) {
	tasks, err := loadTasks()
	if err != nil {
		log.Fatalf("Failed to load tasks: %v", err)
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = newDescription
			tasks[i].Status = "todo"
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Task with ID %d not found\n", id)
		return
	}

	err = saveTasks(tasks)
	if err != nil {
		log.Fatalf("Failed to save task: %v", id)
	}

	fmt.Printf("Task updated successfully (ID: %d)\n", id)
}

func deleteTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		log.Fatalf("Failed to load tasks: %v", err)
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Task with ID %d not found\n", id)
		return
	}

	err = saveTasks(tasks)
	if err != nil {
		log.Fatalf("Failed to save task: %v", id)
	}

	fmt.Printf("Task deleted successfully (ID: %d)\n", id)
}

func markTaskStatus(id int, status string) {
	tasks, err := loadTasks()
	if err != nil {
		log.Fatalf("Failed to load tasks: %v", err)
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Task with Id %d not found\n", id)
		return
	}

	err = saveTasks(tasks)
	if err != nil {
		log.Fatalf("Failed to save tasks: %v", err)
	}

	fmt.Printf("Task marked as %s (ID: %d)\n", status, id)
}

func listTasks(status string) {
	tasks, err := loadTasks()
	if err != nil {
		log.Fatalf("Failed to load tasks: %v", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	filteredTasks := tasks
	if status != "" {
		filteredTasks = []model.Task{}
		for _, task := range tasks {
			if task.Status == status {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}

	for _, task := range filteredTasks {
		fmt.Printf("ID: %d, Description: %s, Status: %s, Created At: %s, Updated At: %s\n",
			task.ID, task.Description, task.Status, task.CreatedAt.Format(time.RFC3339), task.UpdatedAt.Format(time.RFC3339))
	}
}

func Execute() int {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: task-cli [command] [args...]")
		return 1
	}

	command := args[0]
	switch command {
	case "add":
		if len(args) < 2 {
			fmt.Println("Usage: task-cli add [description]")
			return 1
		}
		addTask(strings.Join(args[1:], " "))
	case "update":
		if len(args) < 3 {
			fmt.Println("Usage: task-cli update [id] [new description]")
			return 1
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Invaild task ID: %v\n", err)
			return 1
		}
		updateTask(id, strings.Join(args[2:], " "))
	case "delete":
		if len(args) < 2 {
			fmt.Println("Usage: task-cli delete [id]")
			return 1
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Invalid task ID: %v\n", err)
			return 1
		}
		deleteTask(id)
	case "mark-in-progress":
		if len(args) < 2 {
			fmt.Println("Usage: task-cli mark-in-progress [id]")
			return 1
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Invalid task ID: %v\n", err)
			return 1
		}
		markTaskStatus(id, "in-progress")
	case "mark-done":
		if len(args) < 2 {
			fmt.Println("Usage: task-cli mark-done [id]")
			return 1
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Invalid task ID: %v\n", err)
			return 1
		}
		markTaskStatus(id, "done")
	case "list":
		if len(args) > 1 {
			listTasks(args[1])
		} else {
			listTasks("")
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
		return 1
	}
	return 0
}
