.PHONY: build test test-fast conformance clean help

# Default target
all: build

# Build the database
build:
	go build -o sql-challenge ./cmd/sql-challenge

# Run full SQLLogicTest suite
test: build
	@echo "Running SQLLogicTest suite..."
	@python3 scripts/run_tests.py

# Run tests, stop on first failure (recommended for development)
test-fast: build
	@echo "Running SQLLogicTest suite (fail-fast mode)..."
	@python3 scripts/run_tests.py -x

# Run a specific test file (e.g., make test-file FILE=select1)
test-file: build
	@python3 scripts/run_tests.py -x -f $(FILE)

# Show conformance summary only
conformance: build
	@echo "Checking conformance..."
	@python3 scripts/run_tests.py --summary

# Clean build artifacts
clean:
	rm -f sql-challenge
	go clean

# Show help
help:
	@echo "SQL Vibe Coding Challenge - Go Seed"
	@echo ""
	@echo "Targets:"
	@echo "  build       - Build database"
	@echo "  test        - Run full SQLLogicTest suite"
	@echo "  test-fast   - Run tests, stop on first failure"
	@echo "  test-file   - Run specific test (FILE=select1)"
	@echo "  conformance - Show pass rate summary"
	@echo "  clean       - Remove build artifacts"
	@echo ""
	@echo "Quick start:"
	@echo "  make build"
	@echo "  make test-fast"
