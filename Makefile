.PHONY: proto update-proto clean all

all: update-proto proto

proto: 
	@echo "Generating gRPC Go files..."
	@mkdir -p internal/pb
	protoc -I ./proto \
		--go_out=./internal/pb --go_opt=paths=source_relative \
		--go-grpc_out=./internal/pb --go-grpc_opt=paths=source_relative \
		./proto/cipher.proto
	@echo "Generation complete!"

update-proto:
	@echo "Fetching latest proto contract from Github..."
	git submodule update --remote --merge

clean:
	@echo "Cleaning internal/pb directory..."
	rm -rf internal/pb/*