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

type UserSettings struct {
	Alpha       string   `json:"alpha" validate:"required,alpha"`
	Alphanum    string   `json:"alphanumeric" validate:"required,alphanum"`
	Base64      string   `json:"base64" validate:"required,base64"`
	CIDR        string   `json:"cidr" validate:"required,cidr"`
	DateTime    string   `json:"datetime" validate:"required,datetime"`
	EndsWith    string   `json:"ends_with" validate:"required,endswith=foo"`
	Hexadecimal string   `json:"hexadecimal" validate:"required,hexadecimal"`
	HexColor    string   `json:"hex_color" validate:"required,hexcolor"`
	Hostname    string   `json:"hostname" validate:"required,hostname"`
	IP          string   `json:"ip" validate:"required,ip"`
	IPV4        string   `json:"ipv4" validate:"required,ipv4"`
	IPV6        string   `json:"ipv6" validate:"required,ipv6"`
	JSON        string   `json:"json" validate:"required,json"`
	Latitude    float64  `json:"latitude" validate:"required,latitude"`
	Longitude   float64  `json:"longitude" validate:"required,longitude"`
	Numeric     string   `json:"numeric" validate:"required,numeric"`
	StartsWith  string   `json:"starts_with" validate:"required,startswith=foo"`
	Unique      []string `json:"unique" validate:"required,unique"`
	URI         string   `json:"uri" validate:"required,uri"`
}

type User struct {
	Base
	Name     string       `json:"name"`
	Email    string       `json:"email"`
	Settings UserSettings `json:"settings"`
}

type UserHandler struct{}

// @summary Get a user by ID
// @param id string path User ID
// @response 200 application/json User The user with the given ID
// @router /user/{id} [get]
func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = &User{}
}
