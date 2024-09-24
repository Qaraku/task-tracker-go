package cmd

import (
	"encoding/json"
	"os"
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
	return 0
}
