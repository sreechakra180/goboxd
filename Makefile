<<<<<<< HEAD
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
=======
.PHONY: build run test integration lint

COMPOSE ?= docker compose
TOOLS   := $(COMPOSE) --profile tools run --rm tools

build:
	$(COMPOSE) build goboxd

run:
	$(COMPOSE) up goboxd

test:
	$(TOOLS) go test ./...

integration:
	$(TOOLS) go test -tags=integration ./tests/...

lint:
	$(TOOLS) golangci-lint run ./...
>>>>>>> ea5ab73544b662686ae9797ef36e03b1aced31bc
