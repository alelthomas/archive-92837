package main

import (
	"bufio"
	"fmt"
	"github.com/alelthomas/go-todolist/todolist"
	"os"
	"strings"
)

func main() {
	todoList := todolist.TodoList{}

	scanner := bufio.NewScanner(os.Stdin)

	// Load the to-do list from a file (if it exists)
	filename := "todolist.txt"
	if err := todoList.LoadFromFile(filename); err != nil {
		fmt.Println("Error loading to-do list:", err)
	}

	fmt.Println("Welcome to the Todo List App!")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		switch {
		case strings.HasPrefix(input, "add"):
			description := strings.TrimPrefix(input, "add")
			todoList.AddTask(strings.TrimSpace(description))
		case strings.HasPrefix(input, "list"):
			todoList.ListTasks()
		case strings.HasPrefix(input, "mark"), strings.HasPrefix(input, "delete"):
			todoList.MarkOrDeleteTask(input)
		case strings.HasPrefix(input, "deleteall"):
			todoList.DeleteAllTasks()
			fmt.Println("All tasks deleted successfully.")
		case strings.HasPrefix(input, "save"):
			if err := todoList.SaveToFile(filename); err != nil {
				fmt.Println("Error saving to-do list:", err)
			} else {
				fmt.Println("To-do list saved successfully.")
			}
		case strings.HasPrefix(input, "exit"):
			// Save the to-do list before exiting
			if err := todoList.SaveToFile(filename); err != nil {
				fmt.Println("Error saving to-do list:", err)
			}
			fmt.Println("Exiting the Todo List App. Have a great day!")
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Please enter 'add', 'list', 'mark', 'delete', 'deleteall', 'save', or 'exit'.")
		}
	}
}
