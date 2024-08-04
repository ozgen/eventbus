build:
	@go build -o bin/app cmd/main.go

.PHONY: build test clean

# Project specific settings
BINARY_NAME=eventbus
PACKAGE=github.com/ozgen/eventbus

# Build compiles the binary for the EventBus project
build:
	@echo "Building ${BINARY_NAME}..."
	@go build -o ${BINARY_NAME} ${PACKAGE}

# Test runs all the tests in the EventBus package
test:
	@echo "Running tests..."
	@go test -v ${PACKAGE}/...

# Clean removes the built binary and other artifacts
clean:
	@echo "Cleaning..."
	@go clean
	@rm -f ${BINARY_NAME}

