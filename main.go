package main

import (
	"fmt"
)

type ToDo struct {
	ID     int
	Title  string
	Status string
}


type ToDoList struct {
	todos []ToDo
}

func (tl *ToDoList) addTodo(todo ToDo) {
	tl.todos = append(tl.todos, todo)
}

func (tl *ToDoList) updateTodo(id int, newStatus string) {
	for i := range tl.todos {
		if tl.todos[i].ID == id {
			tl.todos[i].Status = newStatus
			break
		}
	}
}

func (tl *ToDoList) deleteTodo(id int) {
	for i := range tl.todos {
		if tl.todos[i].ID == id {
			tl.todos = append(tl.todos[:i], tl.todos[i+1:]...)
			break
		}
	}
}

func (tl *ToDoList) getAllTodos() []ToDo {
	return tl.todos
}

func main() {
	todoList := ToDoList{}

	todo1 := ToDo{ID: 1, Title: "Learn Golang", Status: "Not completed"}
	todo2 := ToDo{ID: 2, Title: "Make a ToDo List", Status: "Doing"}
	todo3 := ToDo{ID: 3, Title: "Solve quizzes", Status: "Started"}

	todoList.addTodo(todo1)
	todoList.addTodo(todo2)
	todoList.addTodo(todo3)

	allTodos := todoList.getAllTodos()
	fmt.Println("All Todos:")
	fmt.Println(allTodos)

	todoList.updateTodo(2, "Complete")

	updatedTodos := todoList.getAllTodos()
	fmt.Println("Updated Todos:")
	fmt.Println(updatedTodos)

	todoList.deleteTodo(1)

	remainingTodos := todoList.getAllTodos()
	fmt.Println("Remaining Todos:")
	fmt.Println(remainingTodos)
}
