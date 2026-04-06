SOURCES := $(wildcard ./pkg/**/*.go ./cmd/**/*.go)

snapshots: pkg/parser/testdata/generics/expected.yaml pkg/parser/testdata/types/expected.yaml pkg/parser/testdata/simple/expected.yaml
	@echo "snapshots are up to date"
.PHONY: snapshots

pkg/parser/testdata/generics/expected.yaml: $(SOURCES)
	go run ./cmd/openapier --root pkg/parser/testdata/generics > pkg/parser/testdata/generics/expected.yaml

pkg/parser/testdata/types/expected.yaml: $(SOURCES)
	go run ./cmd/openapier --root pkg/parser/testdata/types > pkg/parser/testdata/types/expected.yaml

pkg/parser/testdata/simple/expected.yaml: $(SOURCES)
	go run ./cmd/openapier --root pkg/parser/testdata/simple > pkg/parser/testdata/simple/expected.yaml

test: snapshots
	go test ./...
