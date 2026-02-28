package api

import (
	"net/http"

	. "github.com/swaggo/swag/v2/testdata/simple/cross"
	_ "github.com/swaggo/swag/v2/testdata/simple/web"
)

// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-string-by-int
// @Param   some_id int path
// @Param.description some_id the ID of the string to return
// @Router /testapi/get-string-by-int/{some_id} [get]
func GetStringByInt(w http.ResponseWriter, r *http.Request) {
	_ = Cross{}
	//write your code
}

// @Description get struct array by ID
// @ID get-struct-array-by-string
// @Param some_id string path
// @Param category int query
// @Param offset int query
// @Param limit int query
// @Param q string query
// @Security ApiKeyAuth scope1 scope2
// @Security BasicAuth scope1 scope2
// @Security OAuth2Application write scope3
// @Security OAuth2Implicit read admin scope4
// @Security OAuth2AccessCode read scope5
// @Security OAuth2Password admin scope6
// @Security OAuth2Implicit read write scope7
// @Security Firebase scope8
// @Router /testapi/get-struct-array-by-string/{some_id} [get]
func GetStructArrayByString(w http.ResponseWriter, r *http.Request) {
	//write your code
}

// @Summary Upload file
// @Description Upload file
// @ID file.upload

// @Router /file/upload [post]
func Upload(w http.ResponseWriter, r *http.Request) {
	//write your code
}

// @Summary use Anonymous field
// @Router /AnonymousField [get]
func AnonymousField() {

}

// @Summary use pet2
// @Router /Pet2 [get]
func Pet2() {

}

// @Summary Use IndirectRecursiveTest
// @Router /IndirectRecursiveTest [get]
func IndirectRecursiveTest() {
}

// @Summary Use Tags
// @Router /Tags [get]
func Tags() {
}

// @Summary Use CrossAlias
// @Router /CrossAlias [get]
func CrossAlias() {
}

// @Summary Use AnonymousStructArray
// @Router /AnonymousStructArray [get]
func AnonymousStructArray() {
}

type Pet3 struct {
	ID int `json:"id"`
}

// @Router /GetPet5a [options]
func GetPet5a() {

}

// @Router /GetPet5b [head]
func GetPet5b() {

}

// @Router /GetPet5c [patch]
func GetPet5c() {

}

type SwagReturn []map[string]string

// @Router /GetPet6MapString [get]
func GetPet6MapString() {

}

// @Router /GetPet6FunctionScopedResponse [get]
func GetPet6FunctionScopedResponse() {
	type response struct {
		Name string
	}
}

// @Router /GetPet6FunctionScopedComplexResponse [get]
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
