GO = go
BUILD_DIR = bin
BINARY_NAME_API = orch
BINARY_NAME_PROVISIONER = provisioner
PORT = 8080
GRPC_PORT = 50051
PLUGINS_DIR = provisioner/taskPlugins

all: generate-plugin build

build: fmt vet
	@echo "Building the project..."
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME_API) .
	@echo "Building provisioner..."
	cd provisioner && $(GO) build -o ../$(BUILD_DIR)/$(BINARY_NAME_PROVISIONER) .


generate-plugin:
	@echo "Generating plugins..."
	pwd && cd provisioner/taskPlugins && for i in *.go; do echo "acbd" $i; $(GO) build -buildmode=plugin -o $$i.so $i; done

run: run-dev

run-api: build
	@echo "Running REST API server on port $(PORT)..."
	$(BUILD_DIR)/$(BINARY_NAME_API)

run-provisioner: build
	@echo "Running gRPC server on port $(GRPC_PORT)..."
	$(BUILD_DIR)/$(BINARY_NAME_PROVISIONER)

run-dev: generate-plugin build
	@echo "Running gRPC server on port $(GRPC_PORT)..."
	$(BUILD_DIR)/$(BINARY_NAME_PROVISIONER) &
	# TODO: need to find a better solution
	sleep 12
	@echo "Running REST API server on port $(PORT)..."
	$(BUILD_DIR)/$(BINARY_NAME_API)

run-prod:
	@echo "Production build (empty for now)..."
	@echo "To run in production, use: make run-api && make run-provisioner"

test:
	@echo "Running tests..."
	$(GO) test ./...

fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

vet:
	@echo "Vetting code..."
	$(GO) vet ./...

clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

lint:
	@echo "Linting code..."
	golangci-lint run

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

help:
	@echo "Available commands:"
	@echo "  make all       - Build the project and generate plugins"
	@echo "  make build     - Build the project"
	@echo "  make run       - Build and run the project (default: development mode)"
	@echo "  make run-dev   - Run the project in development mode"
	@echo "  make run-api   - Run only the REST API server"
	@echo "  make run-provisioner - Run only the gRPC server"
	@echo "  make run-prod  - Production build (empty for now)"
	@echo "  make test      - Run tests"
	@echo "  make fmt       - Format Go code"
	@echo "  make vet       - Vet Go code"
	@echo "  make clean     - Remove build artifacts"
	@echo "  make lint      - Run static analysis"
	@echo "  make proto     - Generate gRPC code"
	@echo "  make infra     - Provision infrastructure using Docker"
	@echo "  make help      - Display this help message"
