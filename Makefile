.PHONY: all test

# Build binary
main:
	go build -ldflags="-s -w" -o elebs main.go

# Default target
all: test

# Run all tests with race detection and coverage
test:
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

# Run tests in parallel
test-parallel:
	go test -v -race -coverprofile=coverage.out -covermode=atomic -parallel=4 ./...

# Clean up coverage files
clean:
	rm -f coverage.out
