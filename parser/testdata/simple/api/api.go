package api

import (
	"net/http"

	"github.com/odarbelaeze/test/simple/cross"
)

// @summary Add a new pet to the store
// @description get string by ID
// @id get-string-by-int
// @param   some_id int path the ID of the string to return
// @security.none
// @router /testapi/get-string-by-int/{some_id} [get]
func GetStringByInt(w http.ResponseWriter, r *http.Request) {
	//write your code
}

// @description get struct array by ID
// @id get-struct-array-by-string
// @param some_id string path
// @param.allowReserved some_id
// @param category int query
// @param.allowEmptyValue category
// @param offset int query
// @param.deprecated offset
// @param limit int query
// @param.required limit
// @param q string query
// @security auth scope1 scope2
// @security BasicAuth scope1 scope2
// @security OAuth2Application write scope3
// @security OAuth2Implicit read admin scope4
// @security OAuth2AccessCode read scope5
// @security OAuth2Password admin scope6
// @security OAuth2Implicit read write scope7
// @security Firebase scope8
// @router /testapi/get-struct-array-by-string/{some_id} [get]
func GetStructArrayByString(w http.ResponseWriter, r *http.Request) {
	//write your code
}

// @summary Upload file
// @description Upload file
// @id file.upload
// @requestBody application/json string The request body
// @response 200 application/json string The response body
// @router /file/upload [post]
func Upload(w http.ResponseWriter, r *http.Request) {
	//write your code
}

// @summary use Anonymous field
// @router /AnonymousField [get]
func AnonymousField() {

}

// @summary use pet2
// @router /Pet2 [get]
func Pet2() {

}

// @summary Use IndirectRecursiveTest
// @router /IndirectRecursiveTest [get]
func IndirectRecursiveTest() {
}

// @summary Use Tags
// @router /Tags [get]
func Tags() {
}

// @summary Use CrossAlias
// @router /CrossAlias [get]
func CrossAlias() {
}

// @summary Use AnonymousStructArray
// @router /AnonymousStructArray [get]
func AnonymousStructArray() {
}

type Pet3 struct {
	ID int `json:"id"`
}

// @requestBody application/json Pet3 The request body
// @router /GetPet5a [options]
func GetPet5a() {

}

// @requestBody application/json []Pet3 The request body
// @router /GetPet5b [head]
func GetPet5b() {

}

// @requestBody application/json []cross.Cross The request body
// @router /GetPet5c [patch]
func GetPet5c() {
	_ = cross.Cross{}
}

type SwagReturn []map[string]string

// @router /GetPet6MapString [get]
func GetPet6MapString() {

}

// @router /GetPet6FunctionScopedResponse [get]
func GetPet6FunctionScopedResponse() {
	type response struct {
		Name string
	}
}

// @router /GetPet6FunctionScopedComplexResponse [get]
func GetPet6FunctionScopedComplexResponse() {
	type pet struct {
		Name string
	}

	type pointerPet struct {
		Name string
	}

	type response struct {
		Pets       []pet
		PointerPet *pointerPet
	}
}
