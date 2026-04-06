SOURCES := $(wildcard ./pkg/**/*.go ./cmd/**/*.go)
TEST_DIRS := $(wildcard pkg/parser/testdata/*/)
SNAPSHOTS := $(addsuffix expected.yaml, $(TEST_DIRS))

snapshots: $(SNAPSHOTS)
	@echo "snapshots are up to date"
.PHONY: snapshots

# % matches the name of the test folder
pkg/parser/testdata/%/expected.yaml: $(SOURCES)
	# $* is the name of the test folder, $@ is the name of the whole rule
	go run ./cmd/openapier --root pkg/parser/testdata/$* > $@

TAGS.md: $(SOURCES)
	go run ./cmd/docs > TAGS.md

docs: TAGS.md
	@echo docs are up to date
.PHONY: docs

test: snapshots
	go test ./...
.PHONY: test
