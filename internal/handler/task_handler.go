package handler

import (
	"fmt"
	"go-task-tracker-cli/internal/helper"
	"go-task-tracker-cli/internal/model"
	"strings"
	"time"
)

func ListTasks(args []string) error {
	var tasks []model.Task
	var err error
	if len(args) == 0 {
		tasks, err = helper.ReadTasks()
		if err != nil {
			return err
		}
	} else {
		tasks, err = helper.ReadTasksByStatus(model.TaskStatus(args[0]))
		if err != nil {
			return err
		}
	}

	if len(tasks) == 0 {
		fmt.Println("You don't have any task")
		return nil
	}
	fmt.Printf("%-4s %-12s %s\n", "ID", "Status", "Description")
	for _, task := range tasks {
		fmt.Printf("%-4s %-12s %s\n",
			fmt.Sprintf("[%d]", task.ID),
			task.Status,
			task.Description,
		)
	}

	return nil
}

func AddTask(args []string) error {
	err := helper.CheckLengthArgs(args, 2)
	if err != nil {
		return err
	}

	tasks, err := helper.ReadTasks()
	if err != nil {
		return err
	}

	now := time.Now()
	taskID := helper.ReadLastID(tasks) + 1
	task := model.Task{
		ID:          taskID,
		Description: strings.Join(args, " "),
		Status:      model.TaskStatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, task)
	err = helper.SaveTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)\n", taskID)
	return nil
}

func DeleteTask(args []string) error {
	err := helper.CheckLengthArgs(args, 1)
	if err != nil {
		return err
	}

	tasks, err := helper.ReadTasks()
	if err != nil {
		return err
	}

	tasks, err = helper.RemoveTaskByID(tasks, args[0])
	if err != nil {
		return err
	}

	err = helper.SaveTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("Task deleted successfully (ID: %s)\n", args[0])
	return nil
}

func UpdateTask(args []string) error {
	err := helper.CheckLengthArgs(args, 2)
	if err != nil {
		return err
	}

	tasks, err := helper.ReadTasks()
	if err != nil {
		return err
	}

	taskIndex, err := helper.FindTaskIndexByID(tasks, args[0])
	if err != nil {
		return err
	}

	taskDescriptionUpd := strings.Join(args[1:], " ")
	tasks[taskIndex].Description = taskDescriptionUpd
	tasks[taskIndex].UpdatedAt = time.Now()

	err = helper.SaveTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("Task updated successfully (ID: %s)\n", args[0])
	return nil
}

func MarkTask(args []string, status model.TaskStatus) error {
	err := helper.CheckLengthArgs(args, 1)
	if err != nil {
		return err
	}

	tasks, err := helper.ReadTasks()
	if err != nil {
		return err
	}

	taskIndex, err := helper.FindTaskIndexByID(tasks, args[0])
	if err != nil {
		return err
	}

	tasks[taskIndex].Status = status
	tasks[taskIndex].UpdatedAt = time.Now()

	err = helper.SaveTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("Task updated successfully (ID: %s)\n", args[0])
	return nil
}
