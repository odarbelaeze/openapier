package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// @info.title API that exposes types
// @info.version 1.0
// @info.description This is a sample API that exposes types.
// @info.termsOfService http://swagger.io/terms/
func main() {
	http.ListenAndServe(":8080", nil)
}

type RestError struct {
	// Success is a boolean indicating if the request was successful
	Success bool `json:"success" validate:"eq=false"`

	// Code is the HTTP status code
	Code int `json:"code,string" example:"500"`

	// Message is the error message
	Message string `json:"message"`
}

func errorResponse(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(RestError{Code: code, Message: http.StatusText(code)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type BaseTodo struct {
	// Title is the title of the todo
	Title string `json:"title" example:"Buy milk" validate:"required,min=3,max=100"`

	// Description is a string with an exact length of 50 characters
	Description string `json:"description" example:"Buy 2 liters of milk from the store." validate:"required,len=50"`

	// Completed is the completion status of the todo
	Completed string `json:"completed"`
}

// Todo is a model for to do items
type Todo struct {
	BaseTodo

	// ID is the unique identifier for the todo
	ID uuid.UUID `json:"id"`

	// Attributes is a map of additional attributes for the todo
	Attributes map[string]string `json:"attributes" validate:"len=12"`

	// Coordinates is the geographical coordinates of the todo
	Coordinates [3]float64 `json:"coordinates"`

	// Mailto is the email address of the todo owner
	Mailto *string `json:"mailto" validate:"email"`

	// Created is the time the todo was created
	Created time.Time `json:"created"`

	// Updated is the time the todo was last updated
	Updated time.Time `json:"updated"`
}

// Pagination is a model for cursors and pages
type Pagination struct {
	// Cursor is the cursor for the next page
	Cursor string `json:"cursor"`

	// HasMore indicates if there are more pages
	HasMore bool `json:"hasMore"`

	// Next is the URL for the next page
	Next *string `json:"next" validate:"omitempty,url"`
}

// PaginatedTodos is a model for paginated todos
type PaginatedTodos struct {
	// Todos is the list of todos
	Todos []Todo `json:"todos" validate:"max=100"`

	// Meta is the pagination metadata
	Meta Pagination `json:"meta"`
}

// @summary List todos
// @response 200 application/json PaginatedTodos Paginated list of todos
// @response 500 application/json RestError Internal server error
// @router /todos [get]
func TodoList(w http.ResponseWriter, r *http.Request) {
	var paginatedTodos PaginatedTodos
	bytes, err := json.Marshal(paginatedTodos)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
