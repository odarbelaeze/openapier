package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/odarbelaeze/openapier/pkg/schema/validator"
)

func main() {
	fmt.Println("# OpenAPI Tags Documentation")
	fmt.Println("\nThis document lists all the supported `@` tags and validation tags for the `openapier` tool.")

	fmt.Println("\n## Spec-level Tags")
	fmt.Println("\nThese tags are used to define general information about the API (Info, Servers, etc.).")
	specComments := spec.Default().Comments()
	printComments(specComments, "@")

	fmt.Println("\n## Operation-level Tags")
	fmt.Println("\nThese tags are used to define individual API operations (Paths, Parameters, Responses, etc.).")
	opComments := operation.Default().Comments()
	printComments(opComments, "@")

	fmt.Println("\n## Validation Tags")
	fmt.Println("\nThese tags are used in Go struct fields within the `validate` tag to define schema constraints.")
	validationTags := validator.Default().Validators()
	printComments(validationTags, "")
}

type comment interface {
	Tag() string
	Usage() string
}

func printComments[T comment](comments []T, prefix string) {
	fmt.Println("\n| Tag | Usage |")
	fmt.Println("| :--- | :--- |")

	sort.Slice(comments, func(i, j int) bool {
		return comments[i].Tag() < comments[j].Tag()
	})

	for _, c := range comments {
		tag := strings.ReplaceAll(c.Tag(), "`", "\\`")
		usage := strings.ReplaceAll(c.Usage(), "`", "\\`")
		fmt.Printf("| `%s%s` | `%s` |\n", prefix, tag, usage)
	}
}
