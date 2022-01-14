.DEFAULT_GOAL := lint

NAME := $(shell basename $(CURDIR))

clean:
	@echo "Cleaning ${NAME}..."
	@go clean -i ./...
	@rm -rf bin

test: clean
	@echo "Testing ${NAME}..."
	@go test ./... -cover -race

lint: test
	@echo "Linting ${NAME}..."
	@go vet ./...
	@golangci-lint run #https://golangci-lint.run/usage/install/
