# Define variables
GO = go
BUILD_DIR = bin
BINARY_NAME_API = orch
BINARY_NAME_PROVISIONER = provisioner-server
PORT = 8080
GRPC_PORT = 50051

# Default target
all: build

# Build the project
build: fmt vet
	@echo "Building the project..."
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME_API) .
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME_PROVISIONER) ./provisioner

# Run the project (default)
run: run-dev

# Run API server in development mode
run-api: build
	@echo "Running REST API server on port $(PORT)..."
	$(BUILD_DIR)/$(BINARY_NAME_API)

# Run Provisioner server in development mode
run-provisioner: build
	@echo "Running gRPC server on port $(GRPC_PORT)..."
	$(BUILD_DIR)/$(BINARY_NAME_PROVISIONER)
	
# Development run target
run-dev: build
	@echo "Running gRPC server on port $(GRPC_PORT)..."
	$(BUILD_DIR)/$(BINARY_NAME_PROVISIONER) &
	@echo "Running REST API server on port $(PORT)..."
	$(BUILD_DIR)/$(BINARY_NAME_API)

# Production run target (empty for now)
run-prod:
	@echo "Production build (empty for now)..."

# Test the project
test:
	@echo "Running tests..."
	$(GO) test ./...

# Format Go code
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# Vet Go code
vet:
	@echo "Vetting code..."
	$(GO) vet ./...

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

# Run static analysis
lint:
	@echo "Linting code..."
	golangci-lint run

# Generate gRPC code
proto:
	@echo "Generating gRPC code..."
	protoc --go_out=paths=source_relative:. \
       --go-grpc_out=paths=source_relative:. \
       --go_opt=Mproto-provisioner/service.proto=orch/proto-provisioner \
       --go-grpc_opt=Mproto-provisioner/service.proto=orch/proto-provisioner \
       proto-provisioner/service.proto

infra:
	@echo "Provisioning infra..."
	docker compose up -d 

# Display available commands
help:
	@echo "Available commands:"
	@echo "  make build     - Build the project"
	@echo "  make run       - Build and run the project (default: development mode)"
	@echo "  make run-dev   - Run the project in development mode"
	@echo "  make run-prod  - Run the project in production mode (empty for now)"
	@echo "  make test      - Run tests"
	@echo "  make fmt       - Format Go code"
	@echo "  make vet       - Vet Go code"
	@echo "  make clean     - Remove build artifacts"
	@echo "  make lint      - Run static analysis"
	@echo "  make proto     - Generate gRPC code"
	@echo "  make help      - Display this help message"
