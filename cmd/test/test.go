package main

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	ID        int
	Text      string `json:"-"`
	something string
}

func main() {
	todo := Todo{ID: 1, Text: "hello", something: "world"}
	bytes, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}
