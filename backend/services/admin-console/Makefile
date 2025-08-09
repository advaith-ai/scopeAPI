# Admin Console Microservice Makefile

.PHONY: help build run test clean docker-build docker-run dev

# Default target
help:
	@echo "Admin Console Microservice - Available commands:"
	@echo "  build        - Build the Go application"
	@echo "  run          - Run the application locally"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  dev          - Build and run in development mode"
	@echo "  full-build   - Build Angular and Go application"

# Build the Go application
build:
	@echo "Building Go application..."
	go mod tidy
	go build -o admin-console ./cmd/main.go

# Run the application locally
run: build
	@echo "Running admin console service..."
	./admin-console

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f admin-console
	rm -rf dist

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t scopeapi-admin-console .

# Run Docker container
docker-run: docker-build
	@echo "Running Docker container..."
	docker run -p 8080:8080 scopeapi-admin-console

# Development mode - build and run
dev: build
	@echo "Starting development server..."
	./admin-console

# Full build - Angular + Go
full-build:
	@echo "Building Angular application..."
	cd ../../../adminConsole && npm install && npm run build
	@echo "Copying Angular build..."
	mkdir -p dist
	cp -r ../../../adminConsole/dist/admin-console/* dist/
	@echo "Building Go application..."
	go mod tidy
	go build -o admin-console ./cmd/main.go
	@echo "Build completed successfully!"

# Install dependencies
deps:
	@echo "Installing Go dependencies..."
	go mod tidy
	@echo "Installing Angular dependencies..."
	cd ../../../adminConsole && npm install

# Format code
fmt:
	@echo "Formatting Go code..."
	go fmt ./...
	@echo "Formatting Angular code..."
	cd ../../../adminConsole && npm run format

# Lint code
lint:
	@echo "Linting Go code..."
	golangci-lint run
	@echo "Linting Angular code..."
	cd ../../../adminConsole && npm run lint 