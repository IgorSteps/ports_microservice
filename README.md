# Ports Microservice
How to GET all ports:
```
curl http://localhost:5080/ports
```

gRPC curl sample ports
```
grpcurl -plaintext localhost:50051 PortDomain/GetPortList
```