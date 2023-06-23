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
	$(GO) build -o $(DIST_DIR)/ports-rest $(CMD_DIR)/portrest

## run-grpc: run ports grpc 
.PHONY: run-grpc
run-grpc:
	./dist/ports-grpc

## run-rest: run ports rest 
.PHONY: run-rest
run-rest:
	./dist/ports-rest

## wire: generate DI files
.PHONY: wire
wire:
	$(WIRE) $(CMD_DIR)/portgrpc
	$(WIRE) $(CMD_DIR)/portrest


