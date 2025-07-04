.PHONY: run build test lint clean env dev logs logs-api logs-db
.ALL: run

APP_NAME=GoBudget
CMD_DIR=cmd/server
ENV_FILE=.env
OUT_DIR=bin

run:
	@echo ">> Running $(APP_NAME)..."
	@ENV_FILE=$(ENV_FILE) go run $(CMD_DIR)/main.go

build:
	@echo ">> Building $(APP_NAME)..."
	@go build -o $(OUT_DIR)/$(APP_NAME) $(CMD_DIR)/main.go

test:
	@echo ">> Running tests..."
	@go test ./...

lint:
	@echo ">> Linting code..."
	@go vet ./...

clean:
	@echo ">> Cleaning build artifacts..."
	@rm -f $(APP_NAME)

env:
	@echo ">> Showing environment variables from $(ENV_FILE)"
	@cat $(ENV_FILE)

dev:
	@echo ">> Starting development environment with Docker Compose..."
	@docker-compose up --build --remove-orphans --force-recreate -d

logs:
	@echo ">> Tailing logs from docker-compose (press Ctrl+C to exit)"
	@docker-compose logs -f

logs-api:
	@echo ">> Tailing logs from API service (Ctrl+C to exit)"
	@docker-compose logs -f api

logs-db:
	@echo ">> Tailing logs from DB service (Ctrl+C to exit)"
	@docker-compose logs -f postgres