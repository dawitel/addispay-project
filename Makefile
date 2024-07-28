# Define variables
PROTOC = protoc
PROTOC_GEN_GO = protoc-gen-go
PROTOC_GEN_GO_GRPC = protoc-gen-go-grpc
PROTO_DIR = internal/proto
PROTO_FILES = $(wildcard $(PROTO_DIR)/*.proto)
OUT_DIR = internal/grpc
GO_OUT = $(OUT_DIR)
GO_GRPC_OUT = $(OUT_DIR)
GO_FILES = $(wildcard $(GO_OUT)/*.go)

# Ensure Go binaries for protobuf are installed
.PHONY: tools
tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate Go code from proto files
.PHONY: proto
proto: $(PROTO_FILES)
	$(PROTOC) --proto_path=$(PROTO_DIR) --go_out=$(GO_OUT) --go-grpc_out=$(GO_GRPC_OUT) $(PROTO_FILES)

# Build the Go project
.PHONY: build
build: proto
	go build -o bin/server ./cmd/grpc_server

# Clean up generated files and build artifacts
.PHONY: clean
clean:
	rm -rf $(GO_OUT)/*.go bin/

# Run tests
.PHONY: test
test:
	go test ./...

# Full build process
.PHONY: all
all: tools proto build
