package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Pagination struct {
	Cursor  string
	HasMore bool
}

type Todo struct {
	ID        int    `json:",string"`
	Text      string `json:"-"`
	something string
}

type PaginatedTodo struct {
	Todos []Todo
	Meta  Pagination
}

func main() {
	paginatedTodos := PaginatedTodo{
		Todos: []Todo{
			{
				ID:        0,
				Text:      "Remember the milk",
				something: "something",
			},
		},
		Meta: Pagination{
			Cursor:  "af32d",
			HasMore: false,
		},
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	err := enc.Encode(paginatedTodos)
	if err != nil {
		fmt.Println(err)
		return
	}
}
