# Define the binary name
BINARY_NAME=myapp

# Load env
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

## build: Build the application binary
# build:
# 	go build -o bin/${BINARY_NAME} main.go

## run: Build and run the application
# run: build
# 	./bin/${BINARY_NAME}

## test: Run all tests with coverage
# test:
# 	go test -v -cover ./...

## run graphql 
run-graphql:
	go run ./graphql/cmd

## migrate rest schema
graphql-migrate:
	atlas schema apply --url "$(DATABASE_URL)" --dev-url "$(DEV_URL)" --to "file://./graphql/internal/repository/sql/schema.sql"

## generate graphql repository
graphql-generate-repository:
	cd graphql/internal && sqlc generate

## generate graphql code
graphql-generate-graphql:
	cd graphql/internal && go tool gqlgen generate

## ----

## run rest
run-rest:
	go run ./rest/cmd

## migrate rest schema
rest-migrate:
	atlas schema apply --url "$(DATABASE_URL)" --dev-url "$(DEV_URL)" --to "file://./rest/internal/repository/sql/schema.sql"

## generate rest repository
rest-generate:
	cd rest/internal && sqlc generate

## tidy: Format code and clean up dependencies
tidy:
	go fmt ./...
	go mod tidy

## clean: Remove build artifacts
clean:
	go clean
	rm -f bin/${BINARY_NAME}

# Declare these as "phony" so Make doesn't confuse them with actual files
# .PHONY: build run test migrate tidy clean
