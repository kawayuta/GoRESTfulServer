package main

import "fmt"

var (
	todos Todos
	currentID int
)

func init() {
	RepoCreateTodo(Todo{Name: "犬の散歩"})
	RepoCreateTodo(Todo{Name: "御飯食べる"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentID +=1
	t.ID = currentID
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:1], todos[i+1])
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}