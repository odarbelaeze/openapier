package main

import (
	"net/http"

	"github.com/google/uuid"
)

// @info.title API that exposes types
// @info.version 1.0
// @info.description This demonstrates embedding.
// @info.termsOfService http://swagger.io/terms/
func main() {
	http.Handle("GET /user/{id}", &UserHandler{})
	http.ListenAndServe(":8080", nil)
}

type Base struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type User struct {
	Base
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserHandler struct{}

// @summary Get a user by ID
// @param id string path User ID
// @response 200 application/json User The user with the given ID
// @router /user/{id} [get]
func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = &User{}
}
