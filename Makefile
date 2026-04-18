SOURCES := $(filter %.go, $(shell git ls-files --cached --others --exclude-standard pkg cmd))
TEST_DIRS := $(wildcard pkg/parser/testdata/*/)
SNAPSHOTS := $(addsuffix expected.yaml, $(TEST_DIRS))

snapshots: $(SNAPSHOTS)
	@echo "snapshots are up to date"
.PHONY: snapshots

# % matches the name of the test folder
pkg/parser/testdata/%/expected.yaml: $(SOURCES)
	go run ./cmd/openapier --root pkg/parser/testdata/$* > $@

TAGS.md: $(SOURCES)
	go run ./cmd/docs > TAGS.md

docs: TAGS.md
	@echo docs are up to date
.PHONY: docs

test: mocks snapshots
	go test ./...
.PHONY: test

mocks:
	mockery --log-level ERROR
.PHONY: mocks
