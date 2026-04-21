# Contributing to openapier

Thank you for your interest in contributing to `openapier`! We welcome all types
of contributions, from bug reports and documentation updates to new features and
performance improvements.

## Development Setup

### Prerequisites

- **Go:** 1.26 or later (due to dependencies on modern Go features).
- **Mockery:** Used for generating test mocks. Install it via:
  ```bash
  go install github.com/vektra/mockery/v2@latest
  ```
- **Lefthook (Optional):** Used to run pre-commit and pre-push hooks (linting, formatting, tests). Install it via:
  ```bash
  go install github.com/evilmartians/lefthook@latest
  ```
  After installing, run `lefthook install` to set up the hooks in your local repository.


### Building and Testing

We use a `Makefile` to manage common development tasks:

- **Run all tests:** `make test`
- **Update test snapshots:** `make snapshots` (Run this if you change the output
  format or logic and need to update the `expected.yaml` files in `testdata`).
- **Update documentation:** `make docs` (Regenerates `TAGS.md` from the source
  code).
- **Generate mocks:** `make mocks`

## Contribution Workflow

1. **Fork the repository** and create your branch from `main`.
2. **Implement your changes.** If you're adding a feature or fixing a bug,
   please include relevant tests.
3. **Verify your changes:**
   - Run `make test` to ensure all tests pass.
   - Run `go mod tidy` to keep dependencies clean.
   - Ensure your code is formatted with `gofmt`.
4. **Update Documentation:** If your changes add or modify supported tags, run
  `make docs` to update `TAGS.md`.
6. **Commit your changes.** Use clear and descriptive commit messages.
7. **Push to your fork** and submit a Pull Request.

## Coding Guidelines

- Follow standard Go idioms and naming conventions.
- Keep functions focused and well-documented.
- Ensure all exported symbols have appropriate comments.
- Avoid using `init()` functions unless absolutely necessary.
- You'll find some register patterns through the codebase, please do use
  `init()` for those.

## Reporting Issues

- Use the GitHub Issue tracker to report bugs or suggest enhancements.
- Provide a clear description of the issue and, if possible, a minimal Go code
  snippet to reproduce it.

## License

By contributing to this project, you agree that your contributions will be
licensed under the [MIT License](./LICENSE).
