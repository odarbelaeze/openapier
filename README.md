# openapier

`openapier` is a tool for Go that generates OpenAPI v3.1 specifications from code comments. Heavily inspired by [swaggo/swag][swag], it aims to provide a modern, spec-compliant alternative for documenting your Go APIs.

## Features

- **OpenAPI v3.1 Support:** Generates modern OpenAPI specifications.
- **Go Type Analysis:** Automatically resolves request and response types from your Go structs.
- **Validation Support:** Automatically translates Go `validate` tags into OpenAPI schema constraints (e.g., `min`, `max`, `pattern`).
- **Generics Support:** Full support for Go 1.18+ generics in request and response types.
- **Rich Annotations:** Extensive support for spec-level and operation-level metadata.
- **Easy Integration:** Simple CLI tool that fits into any CI/CD pipeline.

## Installation

You can install `openapier` using `go install`:

```bash
go install github.com/odarbelaeze/openapier/cmd/openapier@latest
```

## Usage

Run `openapier` from your project's root directory:

```bash
openapier --main cmd/server/main.go --root . > openapi.yaml
```

### Flags

- `--main`: Path to the file containing the root spec definition (e.g., `main.go`). Default: `main.go`.
- `--root`: Path to the root directory of the Go code to parse. Default: `./`.
- `--format`: Output format (`yaml` or `json`). Default: `yaml`.
- `--debug`: Enable debug logging for troubleshooting.

## Example

### 1. General API Info

In your `main.go` file, add annotations to describe your API:

```go
package main

// @info.title Swagger Example API
// @info.version 1.0
// @info.description This is a sample server Petstore server.
// @info.termsOfService http://swagger.io/terms/
// @info.contact.name API Support
// @info.license.name Apache 2.0
// @server.url https://petstore.swagger.io/v2
func main() {
    // ...
}
```

### 2. Operation Annotations

Annotate your handler functions to define endpoints:

```go
// @summary Add a new pet to the store
// @description get string by ID
// @id get-string-by-int
// @param   some_id int path the ID of the string to return
// @requestBody application/json Pet The pet to add
// @response 200 application/json Pet The added pet
// @router /testapi/get-string-by-int/{some_id} [get]
func GetStringByInt(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

## Supported Tags

For a comprehensive list of all supported tags and their usage, see [TAGS.md](./TAGS.md).

## Development

To regenerate the documentation or run tests:

```bash
make docs
make test
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

Distributed under the MIT License. See `LICENSE` for more information.

[swag]: https://github.com/swaggo/swag
