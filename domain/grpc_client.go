package domain

import (
	"context"
	"log"
	pb "ports_microservice/domain/proto"
	constants "ports_microservice/utils"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// grpcServerAddress that must be dialed by client
var grpcServerAddress = "localhost:" + strconv.Itoa(constants.GrpcServerPort)

func StartClient() {
	conn, err := grpc.Dial(grpcServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Print("Domain gRPC client dialed on: ", grpcServerAddress)
	defer conn.Close()

	client := pb.NewPortDomainClient(conn)
	portsList, err := client.GetPortList(context.Background(), &pb.GetPortListRequest{})
	if err != nil {
		log.Fatalf("failed to get ports list: %v", err)
	}
	log.Printf("ports list: %v", portsList)
}
