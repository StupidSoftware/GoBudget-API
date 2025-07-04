.PHONY: run build test lint clean env
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

