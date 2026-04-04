package main

import (
	"encoding/json"
	"net/http"
)

// @info.title API that exposes types
// @info.version 1.0
// @info.description This is a sample API that exposes types.
// @info.termsOfService http://swagger.io/terms/
func main() {
	http.ListenAndServe(":8080", nil)
}

// Todo is a model for to do items
type Todo struct {
	// ID is the unique identifier for the todo
	ID string `json:"id"`

	// Title is the title of the todo
	Title string `json:"title"`

	// Completed is the completion status of the todo
	Completed string `json:"completed"`
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
// @router /todos [get]
func TodoList(w http.ResponseWriter, r *http.Request) {
	var paginatedTodos PaginatedTodos
	bytes, err := json.Marshal(paginatedTodos)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

// TodoCreatePayload is the payload to create a todo
type TodoCreatePayload struct {
	// Title is the title of the todo
	Title string `json:"title"`
}

// @summary Create a todo
// @requestBody application/json TodoCreatePayload The payload to create the todo
// @response 201 application/json Todo The recently created todo
// @router /todos [post]
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var createdTodo Todo
	bytes, err := json.Marshal(createdTodo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

// @summary Get a todo
// @param id string path Todo ID
// @response 200 application/json Todo The requested todo
// @router /todos/{id} [get]
func TodoGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	todo := Todo{ID: id}
	bytes, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

// TodoUpdatePayload is the payload for updating a todo
type TodoUpdatePayload struct {
	// Title is the title of the todo
	Title string `json:"title"`

	// Completed is the completion status of the todo
	Completed bool `json:"completed"`
}

// @summary Update a todo
// @param id string path Todo ID
// @requestBody application/json TodoUpdatePayload The payload to update the todo
// @response 200 application/json Todo The updated todo
// @router /todos/{id} [put]
func TodoUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	todo := Todo{ID: id}
	bytes, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

// TodoPatchPayload is the payload for patching a todo
type TodoPatchPayload struct {
	// Optional field for patching a todo's title
	Title *string `json:"title"`

	// Optional field for patching a todo's completion status
	Completed *bool `json:"completed"`
}

// @summary Patch a todo
// @param id string path Todo ID
// @requestBody application/json TodoPatchPayload The payload to update the todo
// @response 200 application/json Todo The updated todo
// @router /todos/{id} [patch]
func TodoPatch(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	todo := Todo{ID: id}
	bytes, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
