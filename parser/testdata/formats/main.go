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
	Alpha          string   `json:"alpha" validate:"required,alpha"`
	Alphanum       string   `json:"alphanumeric" validate:"required,alphanum"`
	AlphaSpace     string   `json:"alpha_space" validate:"required,alphaspace"`
	AlphanumSpace  string   `json:"alphanumeric_space" validate:"required,alphanumspace"`
	ASCII          string   `json:"ascii" validate:"required,ascii"`
	Base64         string   `json:"base64" validate:"required,base64"`
	Base64URL      string   `json:"base64_url" validate:"required,base64url"`
	Base64RawURL   string   `json:"base64_raw_url" validate:"required,base64rawurl"`
	CIDR           string   `json:"cidr" validate:"required,cidr"`
	CIDRv4         string   `json:"cidrv4" validate:"required,cidrv4"`
	CIDRv6         string   `json:"cidrv6" validate:"required,cidrv6"`
	CMYK           string   `json:"cmyk" validate:"required,cmyk"`
	Contains       string   `json:"contains" validate:"required,contains=foo"`
	Cron           string   `json:"cron" validate:"required,cron"`
	CVE            string   `json:"cve" validate:"required,cve"`
	DataURI        string   `json:"data_uri" validate:"required,datauri"`
	DateTime       string   `json:"datetime" validate:"required,datetime"`
	E164           string   `json:"e164" validate:"required,e164"`
	EndsWith       string   `json:"ends_with" validate:"required,endswith=foo"`
	FQDN           string   `json:"fqdn" validate:"required,fqdn"`
	Hexadecimal    string   `json:"hexadecimal" validate:"required,hexadecimal"`
	HexColor       string   `json:"hex_color" validate:"required,hexcolor"`
	HSL            string   `json:"hsl" validate:"required,hsl"`
	HSLA           string   `json:"hsla" validate:"required,hsla"`
	Hostname       string   `json:"hostname" validate:"required,hostname"`
	HostnamePort   string   `json:"hostname_port" validate:"required,hostname_port"`
	HTTPURL        string   `json:"http_url" validate:"required,http_url"`
	HTTPSURL       string   `json:"https_url" validate:"required,https_url"`
	IP             string   `json:"ip" validate:"required,ip"`
	IPAddr         string   `json:"ip_addr" validate:"required,ip_addr"`
	IP4Addr        string   `json:"ip4_addr" validate:"required,ip4_addr"`
	IP6Addr        string   `json:"ip6_addr" validate:"required,ip6_addr"`
	IPV4           string   `json:"ipv4" validate:"required,ipv4"`
	IPV6           string   `json:"ipv6" validate:"required,ipv6"`
	ISBN           string   `json:"isbn" validate:"required,isbn"`
	ISO31661Alpha2 string   `json:"iso3166_1_alpha2" validate:"required,iso3166_1_alpha2"`
	ISO31661Alpha3 string   `json:"iso3166_1_alpha3" validate:"required,iso3166_1_alpha3"`
	ISO4217        string   `json:"iso4217" validate:"required,iso4217"`
	JSON           string   `json:"json" validate:"required,json"`
	JWT            string   `json:"jwt" validate:"required,jwt"`
	Latitude       float64  `json:"latitude" validate:"required,latitude"`
	Longitude      float64  `json:"longitude" validate:"required,longitude"`
	Lowercase      string   `json:"lowercase" validate:"required,lowercase"`
	MAC            string   `json:"mac" validate:"required,mac"`
	NE             string   `json:"ne" validate:"required,ne=foo"`
	NoneOf         string   `json:"none_of" validate:"required,noneof=foo bar"`
	Number         string   `json:"number" validate:"required,number"`
	Numeric        string   `json:"numeric" validate:"required,numeric"`
	Port           int      `json:"port" validate:"required,port"`
	PrintASCII     string   `json:"print_ascii" validate:"required,printascii"`
	RGB            string   `json:"rgb" validate:"required,rgb"`
	RGBA           string   `json:"rgba" validate:"required,rgba"`
	Semver         string   `json:"semver" validate:"required,semver"`
	SSN            string   `json:"ssn" validate:"required,ssn"`
	StartsWith     string   `json:"starts_with" validate:"required,startswith=foo"`
	ULID           string   `json:"ulid" validate:"required,ulid"`
	Unique         []string `json:"unique" validate:"required,unique"`
	Uppercase      string   `json:"uppercase" validate:"required,uppercase"`
	URI            string   `json:"uri" validate:"required,uri"`
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
