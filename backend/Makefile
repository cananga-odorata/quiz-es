.PHONY: help build run test migrate lint docker-up docker-down clean

# Default target
help:
	@echo "Quiz Management System — Available Commands"
	@echo ""
	@echo "Development:"
	@echo "  make run          - Run the API server"
	@echo "  make build        - Build the API binary"
	@echo "  make test         - Run tests with coverage"
	@echo "  make test-fast    - Run tests without coverage"
	@echo ""
	@echo "Database:"
	@echo "  make migrate      - Run database migration"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-up    - Start Docker containers"
	@echo "  make docker-down  - Stop Docker containers"
	@echo "  make docker-build - Build Docker image"
	@echo ""
	@echo "Cleanup:"
	@echo "  make clean        - Remove build artifacts"

# Build the application
build:
	@echo "Building..."
	go build -o bin/api cmd/api/main.go
	go build -o bin/migrate cmd/migrate/main.go
	@echo "Build complete: bin/api, bin/migrate"

# Run the application
run:
	go run cmd/api/main.go

# Run tests with coverage
test:
	go test -v -race -coverprofile=coverage.out ./internal/modules/quiz/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# Run tests without coverage (faster)
test-fast:
	go test -v ./internal/modules/quiz/...

# Run database migration
migrate:
	go run cmd/migrate/main.go

# Format code
fmt:
	go fmt ./...

# Tidy dependencies
tidy:
	go mod tidy

# Start Docker containers
docker-up:
	docker-compose up -d

# Stop Docker containers
docker-down:
	docker-compose down

# Build Docker image
docker-build:
	docker build -t quiz-api:latest .

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html
