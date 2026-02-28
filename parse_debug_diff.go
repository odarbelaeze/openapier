package main

import (
	"fmt"
	"log"

	"github.com/odarbelaeze/openapier/pkg/parser"

	"github.com/google/go-cmp/cmp"
)

func main() {
	p := parser.NewParser()
	spec, err := p.Parse("./pkg/parser/testdata/simple", "main.go")
	if err != nil {
		log.Fatal(err)
	}

	p2 := parser.NewParser()
	spec2, err := p2.Parse("./pkg/parser/testdata/simple", "main.go")
	if err != nil {
		log.Fatal(err)
	}

	// Print actual and expected JSON logic
	diff := cmp.Diff(spec, spec2)
	fmt.Println("Self Diff:", diff)
}
