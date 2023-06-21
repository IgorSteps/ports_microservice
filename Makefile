WIRE := wire
GO := go

CMD_DIR := ./cmd
DIST_DIR := ./dist

## gen: generate Go files from proto
.PHONY: gen
gen:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	external/proto/ports/service.proto \
	external/proto/ports/create.proto

## build: build the project
.PHONY: build
build:
	$(GO) build -o $(DIST_DIR)/ports-grpc $(CMD_DIR)/portgrpc

## run: run the project
.PHONY: run
run:
	./dist/ports-grpc

## wire: generate DI files
.PHONY: wire
wire:
	$(WIRE) $(CMD_DIR)/portgrpc


