.PHONY: proto-tools proto update-proto clean all

GO ?= go
MODULE ?= github.com/basarardabaykal/cipherconverter-microservice
GOBIN := $(shell $(GO) env GOBIN 2>/dev/null)
GOPATH := $(shell $(GO) env GOPATH 2>/dev/null)
ifeq ($(strip $(GOBIN)),)
PROTOC_BIN_DIR := $(GOPATH)/bin
else
PROTOC_BIN_DIR := $(GOBIN)
endif
PROTOC_GEN_GO := $(PROTOC_BIN_DIR)/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(PROTOC_BIN_DIR)/protoc-gen-go-grpc

all: update-proto proto

proto-tools:
	@echo "Installing protoc Go plugins..."
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

proto: proto-tools
	@echo "Generating gRPC Go files..."
	@mkdir -p internal/pb/common internal/pb/symmetric
	@# Remove legacy flat-layout files that cause mixed-package build failures.
	rm -f internal/pb/cipher.pb.go internal/pb/cipher_grpc.pb.go internal/pb/common.pb.go internal/pb/symmetric.pb.go internal/pb/symmetric_grpc.pb.go
	protoc -I ./proto \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
		--go_out=. --go_opt=module=$(MODULE) --go_opt=paths=import \
		--go-grpc_out=. --go-grpc_opt=module=$(MODULE) --go-grpc_opt=paths=import \
		./proto/common.proto \
		./proto/symmetric.proto
	@echo "Generation complete!"

update-proto:
	@echo "Fetching latest proto contract from Github..."
	git submodule update --remote --merge

clean:
	@echo "Cleaning internal/pb directory..."
	rm -rf internal/pb/common internal/pb/symmetric
	rm -f internal/pb/cipher.pb.go internal/pb/cipher_grpc.pb.go internal/pb/common.pb.go internal/pb/symmetric.pb.go internal/pb/symmetric_grpc.pb.go