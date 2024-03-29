WIRE := wire
GO := go

CMD_DIR := ./cmd
DIST_DIR := ./dist
INTERNAL_DIR := ./internal
FUNCTIONAL_DIR := ./functionaltests

## proto: generate Go files from proto
.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	external/proto/ports/service.proto \
	external/proto/ports/create.proto

## wire: generate DI files
.PHONY: wire
wire:
	$(WIRE) $(CMD_DIR)/portgrpc
	$(WIRE) $(CMD_DIR)/portrest

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

## mocks: generate mocks
.PHONY: mocks
mocks:
	mockery --dir=./internal --output=./mocks

## unit: runs unit tests
.PHONY: unit
unit:
	$(GO) test $(INTERNAL_DIR)/...

## functional: runs functional tests
.PHONY: functional
functional:
	$(GO) test $(FUNCTIONAL_DIR)/...