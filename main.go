package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type TODO struct {
	Id   string
	Todo string
}

var TODOLIST []TODO
var menuNumber int

func main() {
	welcome()
	menu()
}

func newLine(numberOfLines int) {
	i := 0
	for i < numberOfLines {
		fmt.Println("\n")
		i++
	}
}

func welcome() {
	fmt.Println("*******{ Welcome to the TODO CLI app 🍾 }*******")
}

func listTodos() {
	fmt.Println("List of items in your TODO list are:")
	newLine(1)
	for index, item := range TODOLIST {
		fmt.Printf("TODO #%v: %v %v", index+1, item.Id, item.Todo)
		newLine(1)
	}
	menu()
}

func createTodo() {
	fmt.Println("Please enter your Todo Item")
	var todoInput string
	_, err := fmt.Scan(&todoInput)
	if err != nil {
		fmt.Println("Error: Please enter a valid todo item")
	}

	var todo TODO
	todo.Id = strconv.Itoa(rand.Intn(100000))
	todo.Todo = todoInput

	TODOLIST = append(TODOLIST, todo)
	fmt.Println("Todo item added.")
	menu()
}

func updateTodo() {
	var updateItem string
	var updateInput string
	fmt.Println("Enter the id of the item you want to update: ")
	_, err := fmt.Scan(&updateItem)
	if err != nil {
		fmt.Println("Error:, please select the correct menu item")
	}
	for index, item := range TODOLIST {
		if item.Id == updateItem {
			TODOLIST = append(TODOLIST[:index], TODOLIST[index+1:]...)
			var newTodo TODO
			//Todo: display current todo item to user

			fmt.Println("Enter your todo: ")
			_, err := fmt.Scan(&updateInput)
			if err != nil {
				fmt.Println("Error:, please select the correct menu item")
			}

			newTodo.Id = updateItem
			newTodo.Todo = updateInput

			TODOLIST = append(TODOLIST, newTodo)

			fmt.Println("Todo item updated")
			menu()
		}
	}

	fmt.Println("Todo item not found!")
	menu()
}

func deleteTodo() {
	var deleteItem string
	fmt.Println("Enter the id of the todo item to delete: ")
	_, err := fmt.Scan(&deleteItem)
	if err != nil {
		fmt.Println("Error:, please select the correct menu item")
	}

	for index, item := range TODOLIST {
		if item.Id == deleteItem {
			TODOLIST = append(TODOLIST[:index], TODOLIST[index+1:]...)
			fmt.Println("Todo Item deleted")
			menu()
		}
	}

	fmt.Println("Todo Item not found")
	menu()
}

func exitProgram() {
	fmt.Println("Goodbye 👋")
	os.Exit(0)
}

func menu() {
	newLine(1)
	fmt.Println("Select Operation:")
	fmt.Println("1. Create Todo \t\t 2. List Todo Items")
	fmt.Println("3. Edit Todo Item \t 4. Delete Todo Item")
	fmt.Println("0. Exit program \t")

	_, err := fmt.Scan(&menuNumber)
	if err != nil {
		fmt.Println("Error:, please select the correct menu item")
		//menu()
	}

	switch menuNumber {
	case 1:
		createTodo()
	case 2:
		listTodos()
	case 3:
		updateTodo()
	case 4:
		deleteTodo()
	case 0:
		exitProgram()
	default:
		fmt.Println("Error: Invalid input")
	}
}
