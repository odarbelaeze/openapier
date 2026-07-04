package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type EnumThing int

const (
	_ EnumThing = iota
	// EnumThingOne is the first enum value
	EnumThingOne

	// EnumThingTwo is the second enum value
	EnumThingTwo

	// EnumThingThree is the third enum value
	EnumThingThree
)

type ExponentialEnumThing int

const (
	_ ExponentialEnumThing = 1 << iota
	// ExponentialEnumThingOne is the first enum value
	ExponentialEnumThingOne

	// ExponentialEnumThingTwo is the second enum value
	ExponentialEnumThingTwo

	// ExponentialEnumThingThree is the third enum value
	ExponentialEnumThingThree
)

type OnePlusExponentialEnumThing int

const (
	_ OnePlusExponentialEnumThing = 1 + (1 << iota)
	// OnePlusExponentialEnumThingOne is the first enum value
	OnePlusExponentialEnumThingOne

	// OnePlusExponentialEnumThingTwo is the second enum value
	OnePlusExponentialEnumThingTwo

	// OnePlusExponentialEnumThingThree is the third enum value
	OnePlusExponentialEnumThingThree
)

type MixedEnumThing int

const (
	_ MixedEnumThing = iota
	// MixedEnumThingOne is the first enum value
	MixedEnumThingOne

	// MixedEnumThingTwo is the second enum value
	MixedEnumThingTwo = 1 << iota

	// MixedEnumThingThree is the third enum value
	MixedEnumThingThree = 1 + (1 << iota)
)

// @info.title API that exposes iota expressions
// @info.version 1.0
// @info.description This is a sample API that exposes iota expressions.
// @info.termsOfService http://swagger.io/terms/
func main() {
	http.ListenAndServe(":8080", nil)
}

type RestError struct {
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

// Todo is a model for to do items
type Todo struct {
	// ID is the unique identifier for the todo
	ID uuid.UUID `json:"id"`

	// Title is the title of the todo
	Title string `json:"title" example:"Buy milk"`

	// EnumThing is an enum field for the todo
	EnumThing EnumThing `json:"enumThing" example:"1"`

	// ExponentialEnumThing is an exponential enum field for the todo
	ExponentialEnumThing ExponentialEnumThing `json:"exponentialEnumThing" example:"2"`

	// OnePlusExponentialEnumThing is a one plus exponential enum field for the todo
	OnePlusExponentialEnumThing OnePlusExponentialEnumThing `json:"onePlusExponentialEnumThing" example:"3"`

	// MixedEnumThing is a mixed enum field for the todo
	MixedEnumThing MixedEnumThing `json:"mixedEnumThing" example:"4"`
}

// Pagination is a model for cursors and pages
type Pagination struct {
	// Cursor is the cursor for the next page
	Cursor string `json:"cursor"`

	// HasMore indicates if there are more pages
	HasMore bool `json:"hasMore"`
}

// PaginatedTodos is a model for paginated todos
type PaginatedTodos struct {
	// Todos is the list of todos
	Todos []Todo `json:"todos"`

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
