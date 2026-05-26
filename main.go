package main

import (
	"bufio"
	"fmt"
	"go-task-tracker-cli/internal/handler"
	"go-task-tracker-cli/internal/model"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=== Go Task Tracker CLI ===")
	fmt.Println("Type 'help' to see command")
	fmt.Println("Type 'exit' to close")
	fmt.Println("=================================")

	for {
		fmt.Print("\ntask-cli ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading command: ", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := parts[0]
		args := parts[1:]

		switch command {
		case "list":
			err := handler.ListTasks(args)
			if err != nil {
				fmt.Println("Error reading tasks: ", err)
			}
		case "add":
			err := handler.AddTask(args)
			if err != nil {
				fmt.Println("Error adding task: ", err)
			}
		case "delete":
			err := handler.DeleteTask(args)
			if err != nil {
				fmt.Println("Error deleting task: ", err)
			}
		case "update":
			err := handler.UpdateTask(args)
			if err != nil {
				fmt.Println("Error updating task: ", err)
			}
		case "mark-done":
			err := handler.MarkTask(args, model.TaskStatusDone)
			if err != nil {
				fmt.Println("Error updating task mark: ", err)
			}
		case "mark-in-progress":
			err := handler.MarkTask(args, model.TaskStatusInProgress)
			if err != nil {
				fmt.Println("Error updating task mark: ", err)
			}
		case "mark-todo":
			err := handler.MarkTask(args, model.TaskStatusTodo)
			if err != nil {
				fmt.Println("Error updating task mark: ", err)
			}
		case "exit":
			fmt.Println("Exit")
			return
		case "help":
			fmt.Println("Commands:")
			fmt.Println("add <description>")
			fmt.Println("list")
			fmt.Println("list <status>")
			fmt.Println("update <id> <description>")
			fmt.Println("delete <id>")
			fmt.Println("mark-done <id>")
			fmt.Println("mark-in-progress <id>")
			fmt.Println("mark-todo <id>")
		default:
			fmt.Println("Command not valid")
		}
	}

}
