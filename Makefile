SOURCES := $(wildcard ./pkg/**/*.go ./cmd/**/*.go)

snapshots: pkg/parser/testdata/generics/expected.yaml pkg/parser/testdata/types/expected.yaml pkg/parser/testdata/simple/expected.yaml
	@echo "snapshots are up to date"
.PHONY: snapshots

# % matches the name of the test folder
pkg/parser/testdata/%/expected.yaml: $(SOURCES)
	# $* is the name of the test folder, $@ is the name of the whole rule
	go run ./cmd/openapier --root pkg/parser/testdata/$* > $@

test: snapshots
	go test ./...
.PHONY: test
