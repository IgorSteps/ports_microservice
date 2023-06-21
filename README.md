# Ports Microservice

## API Spec
 ### Port Creator gRPC service
 A gRPC service which exposes `ports.v1.PortsService` to create a port. It follows protobuf definitions in `external/proto/ports`

Example, run the following from project root:
```
grpcurl --plaintext -proto external/proto/ports/service.proto -d '{"id": "1", "name": "Great Port", "city": "St. Petersburg","country": "Russia"}' 127.0.0.1:5001 ports.v1.PortsService.CreatePort
```
which should output:
```
{
  "name": "Great Port",
  "city": "St. Petersburg",
  "country": "Russia"
}
```


## Running locally

### Starting the service:
1. Run `make build`.
2. Run `make run`.
3. If changes were made to wire.go files, run `make wire` to regen wire_gen.go file.