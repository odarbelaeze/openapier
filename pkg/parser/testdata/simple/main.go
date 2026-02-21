package main

import (
	"net/http"

	"github.com/swaggo/swag/v2/testdata/simple/api"
)

// @info.title Swagger Example API
// @info.version 1.0
// @info.description This is a sample server Petstore server.
// @info.termsOfService http://swagger.io/terms/

// @info.contact.name API Support
// @info.contact.url http://www.swagger.io/support
// @info.contact.email support@swagger.io

// @info.license.name Apache 2.0
// @info.license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @externalDocs.url http://www.swagger.io/support
// @externalDocs.description Find more info here

// @servers.url https://petstore.swagger.io/v2

// @tag.name pets
// @tag.description These are some sample pets
// @tag.externalDocs.url http://swagger.io/petstore
// @tag.externalDocs.description Find more info here
func main() {
	http.HandleFunc("/testapi/get-string-by-int/", api.GetStringByInt)
	http.HandleFunc("/testapi/get-struct-array-by-string/", api.GetStructArrayByString)
	http.HandleFunc("/testapi/upload", api.Upload)
	http.ListenAndServe(":8080", nil)
}
