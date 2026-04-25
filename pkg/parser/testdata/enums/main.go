package main

import (
	"encoding/json"
	"net/http"
)

type (
	Status    string
	Priority  int
	Something int
)

const (
	StatusOpen   Status = "open"
	StatusClosed Status = "closed"
)

const (
	PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
)

const (
	SomethingSomething Something = 1
	SomethingElse      Something = 2
)

// @info.title API that exposes types
// @info.version 1.0
// @info.description This is a sample API that exposes types.
// @info.termsOfService http://swagger.io/terms/
func main() {
	http.HandleFunc("GET /items", ItemList)
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

// Item is a model for items
type Item struct {
	// Status is the status of the item
	Status Status `json:"status"`

	// Priority is the priority of the item
	Priority Priority `json:"priority"`

	// Something is a thing
	Something Something `json:"something"`
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

// Paginated is a generic type for paginated responses
type Paginated[T any] struct {
	// Items is the list of items
	Items []T `json:"items" validate:"max=100"`

	// Meta is the pagination metadata
	Meta Pagination `json:"meta"`
}

// @summary List items
// @param status Status query Item status
// @param priority Priority query Item priority
// @response 200 application/json Paginated[Item] Paginated list of items
// @response 500 application/json RestError Internal server error
// @router /items [get]
func ItemList(w http.ResponseWriter, r *http.Request) {
	var paginatedTodos Paginated[Item]
	bytes, err := json.Marshal(paginatedTodos)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
