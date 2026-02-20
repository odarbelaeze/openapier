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

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	http.HandleFunc("/testapi/get-string-by-int/", api.GetStringByInt)
	http.HandleFunc("/testapi/get-struct-array-by-string/", api.GetStructArrayByString)
	http.HandleFunc("/testapi/upload", api.Upload)
	http.ListenAndServe(":8080", nil)
}
