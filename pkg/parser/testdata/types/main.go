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
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed string `json:"completed"`
}

// Pagination is a model for cursors and pages
type Pagination struct {
	Cursor  string `json:"cursor"`
	HasMore bool   `json:"hasMore"`
}

// PaginatedTodos is a model for paginated todos
type PaginatedTodos struct {
	Todos []Todo     `json:"todos"`
	Meta  Pagination `json:"meta"`
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

type TodoCreatePayload struct {
	Title string `json:"title"`
}

// @summary List todos
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

type TodoUpdatePayload struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// @summary Update a todo
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

type TodoPatchPayload struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

// @summary Patch a todo
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
