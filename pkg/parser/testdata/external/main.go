package main

import (
	"net/http"

	"github.com/odarbelaeze/rest"
)

// @info.title API that exposes types
// @info.version 1.0
// @info.description Loads external packages.
// @info.termsOfService http://swagger.io/terms/
func main() {
	http.ListenAndServe(":8080", nil)
}

type HelloHandler struct{}

// @summary Hello World handler
// @description Returns a greeting message.
// @response 200 text/plain rest.Error Hello, World!
// @tags hello
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = &rest.Error{}
	_, _ = w.Write([]byte("Hello, World!"))
}
