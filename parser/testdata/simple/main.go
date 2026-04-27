package main

import (
	"net/http"

	"github.com/odarbelaeze/test/simple/api"
)

// Unused type to test that the parser ignores it
type Unused int

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

// @server.url https://petstore.swagger.io/v2

// @server.url {scheme}://petstore.swagger.io/v2
// @server.description Production Server
// @server.variable.default scheme http
// @server.variable.enum scheme http https
// @server.variable.description scheme description

// @tag.name pets
// @tag.description These are some sample pets
// @tag.externalDocs.url http://swagger.io/petstore
// @tag.externalDocs.description Find more info here

// @securityScheme auth apiKey An api key auth
// @securityScheme.in auth header
// @securityScheme.name auth X-API-Key

// @securityScheme queryApiKey apiKey An api key auth
// @securityScheme.in queryApiKey query
// @securityScheme.name queryApiKey access_token

// @securityScheme bearer http A bearer token auth
// @securityScheme.scheme bearer Bearer
// @securityScheme.format bearer JWT

// @security auth list get create update
// @security something foo bar

// @tag.name users
// @tag.description These are some sample users
func main() {
	http.HandleFunc("/testapi/get-string-by-int/", api.GetStringByInt)
	http.HandleFunc("/testapi/get-struct-array-by-string/", api.GetStructArrayByString)
	http.HandleFunc("/testapi/upload", api.Upload)
	http.ListenAndServe(":8080", nil)
}
