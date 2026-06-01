.PHONY: run run-docker build test clean

run:
	@echo "Starting GoShield (Local Mode)..."
	go run cmd/goshield/main.go

run-docker:
	@echo "Starting GoShield (Docker Isolated Mode)..."
	USE_DOCKER=true go run cmd/goshield/main.go

build:
	@echo "Building GoShield..."
	go build -o bin/goshield cmd/goshield/main.go

test:
	@echo "Running tests..."
	go test ./...

clean:
	@echo "Cleaning up..."
	rm -rf bin/
