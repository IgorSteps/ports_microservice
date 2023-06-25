# Ports Microservice Exercise

A micrservice written in GoLang that follows clean architecture. Uses wire for DI, gRPC, REST and GORM with PostgreSQL.

## API Spec

### Port Creator gRPC service

A gRPC service which exposes `ports.v1.PortsService` to create a port. It follows protobuf definitions in `external/proto/ports`.

Example, run the following from project root:

```zsh
grpcurl --plaintext -proto external/proto/ports/service.proto -d '{"id": "1", "name": "Great Port", "city": "St. Petersburg","country": "Russia"}' 127.0.0.1:5001 ports.v1.PortsService.CreatePort
```

which should output:

```json
{
  "name": "Great Port",
  "city": "St. Petersburg",
  "country": "Russia"
}
```

### Port Getter REST service

A REST service which exposes `/ports` endpoint to get ports.

Example, run the following from project root:

```zsh
curl localhost:3000/ports
```

which should output:

```json
{
  [
    {
        "Name": "Great Port",
        "City": "St. Petersburg",
        "Country": "Russia",
        "Alias": null,
        "Regions": null,
        "Coordinates": null,
        "Province": "",
        "Timezone": "",
        "Unlocs": null,
        "Code": ""
    },
    {
        "Name": "Great Port",
        "City": "St. Petersburg",
        "Country": "Russia",
        "Alias": null,
        "Regions": null,
        "Coordinates": null,
        "Province": "",
        "Timezone": "",
        "Unlocs": null,
        "Code": ""
    },
  ] 
}
```

## Running locally

### Prerequisites

- Docker
- Go version > 1.20.5

### Setting up environment

1. Run `docker-compose up -d` to create required PostgreSQL image.

### Starting the service

1. Run `make build`.
2. Run `make run-grpc` to run ports grpc.
3. Run `make run-rest` to run ports rest.
4. If changes were made to wire.go files, run `make wire` to regen `wire_gen.go` file.
5. If changes were made to proto definition files, run `make proto` to regen protobuf files.

## Testing

### Unit

Run `make unit` to run unit tests for the whole project.
