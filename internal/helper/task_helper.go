package helper

import (
	"encoding/json"
	"errors"
	"go-task-tracker-cli/internal/model"
	"os"
	"strconv"
)

func ReadLastID(tasks []model.Task) int {
	if len(tasks) == 0 {
		return 0
	}

	max := tasks[0].ID

	for _, task := range tasks {
		if task.ID > max {
			max = task.ID
		}
	}

	return max
}

func FindTaskIndexByID(tasks []model.Task, strTaskID string) (int, error) {
	taskID, err := strconv.Atoi(strTaskID)
	if err != nil {
		return -1, err
	}

	for i, task := range tasks {
		if task.ID == taskID {
			return i, nil
		}
	}

	return -1, errors.New("task id not found")
}

func RemoveTaskByID(tasks []model.Task, strTaskID string) ([]model.Task, error) {
	taskID, err := strconv.Atoi(strTaskID)
	if err != nil {
		return nil, err
	}

	for i, task := range tasks {
		if task.ID == taskID {
			return append(tasks[:i], tasks[i+1:]...), nil
		}
	}

	return nil, errors.New("failed to delete due by id not found")
}

func ReadTasks() ([]model.Task, error) {
	_, err := os.Stat("./internal/tasks.json")

	if os.IsNotExist(err) {
		err = os.WriteFile("./internal/tasks.json", []byte("[]"), 0644)
		if err != nil {
			return nil, err
		}
	}

	data, err := os.ReadFile("./internal/tasks.json")
	if err != nil {
		return nil, err
	}

	var tasks []model.Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func ReadTasksByStatus(status model.TaskStatus) ([]model.Task, error) {
	_, err := os.Stat("./internal/tasks.json")

	if os.IsNotExist(err) {
		err = os.WriteFile("./internal/tasks.json", []byte("[]"), 0644)
		if err != nil {
			return nil, err
		}
	}

	data, err := os.ReadFile("./internal/tasks.json")
	if err != nil {
		return nil, err
	}

	var tasks []model.Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	var tasksByStatus []model.Task

	for _, task := range tasks {
		if task.Status == status {
			tasksByStatus = append(tasksByStatus, task)
		}
	}

	return tasksByStatus, nil
}

func WriteTasks(tasks []byte) error {
	err := os.WriteFile("./internal/tasks.json", tasks, 0644)
	if err != nil {
		return err
	}

	return nil
}

func SaveTasks(tasks []model.Task) error {
	bytesTasks, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = WriteTasks(bytesTasks)
	if err != nil {
		return err
	}

	return nil
}

func CheckLengthArgs(args []string, total int) error {
	if len(args) < total {
		return errors.New("invalid arguments")
	}

	return nil
}
