package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-cli/model"
)

const dataFile = "task.json"

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
	return nil
}

func generateID(tasks []model.Task) int {
	return 0
}

func addTask(description string) {

}

func updateTask(id int, newDescription string) {

}

func deleteTask(id int) {

}

func markTaskStatus(id int, status string) {

}

func listTasks(status string) {

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
