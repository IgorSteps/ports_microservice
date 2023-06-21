WIRE := wire

CMD_DIR := ./cmd

## gen: generate Go files from proto
.PHONY: gen
gen:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	external/proto/ports.proto


## wire: generate DI files
.PHONY: wire
wire:
	$(WIRE) $(CMD_DIR)/portgrpc