BIN := ./bin/app
MAIN := ./cmd/app/main.go

.PHONY: all
all: gen fmt lint test build

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run: ## Run the application
	go run $(MAIN)

.PHONY: build
build: ## Build the binary
	go build -o $(BIN) $(MAIN)

.PHONY: test
test: ## Run tests with race detector
	go test -race -v -cover ./...

.PHONY: lint
lint: ## Run golangci-lint
	golangci-lint run ./...

.PHONY: fmt
fmt: ## Format code
	go fmt ./...
	go mod tidy

.PHONY: gen
gen: ## Run go generate (mockgen, etc.)
	go generate ./...

.PHONY: proto
proto: ## Generate protobuf and gRPC code
	rm -rf pb/*
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		proto/*.proto

.PHONY: clean
clean: ## Clean binaries and coverage
	rm -rf bin/
	rm -f coverage.out coverage.html
	go clean

.PHONY: evans
evans: ## Start Evans REPL for gRPC debugging
	evans --port 9092 --host localhost --path proto --proto verify.proto repl