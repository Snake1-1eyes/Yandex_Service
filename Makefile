PROTO_DIR = ./api
PROTO_FILE = order.proto
GO_OUT = ./pkg/api/test
SERVICE_BIN = server

all: generate build run

generate:
	protoc --go_out=$(GO_OUT) --go_opt paths=source_relative \
			--go-grpc_out=$(GO_OUT) --go-grpc_opt paths=source_relative \
			$(PROTO_DIR)/$(PROTO_FILE)

build:
	go build -o $(SERVICE_BIN) ./cmd

run:
	./$(SERVICE_BIN)

clean:
	@rm -f $(SERVICE_BIN)
