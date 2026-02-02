# Makefile for Averroes Backend

.PHONY: run build clean test db-migrate db-seed

# Build the application
build:
	go build -o bin\api.exe cmd\api\main.go

# Run the application
run: build
	bin\api.exe

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	del /Q bin\api.exe

# Run database migrations (placeholder)
db-migrate:
	@echo Running database migrations...

# Seed database with initial data (placeholder)
db-seed:
	@echo Seeding database with initial data...

# Install dependencies
deps:
	go mod tidy