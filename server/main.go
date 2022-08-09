package main

import (
	"context"
	"google.golang.org/grpc"
	ai "grpc_anomaly_identifier/proto/anomaly_identifier"
	"log"
	"net"
)

type PasswordGeneratorServiceServer struct {
	ai.AnomalyIdentifierServiceServer
}

func (s *PasswordGeneratorServiceServer) Generate(ctx context.Context,
	req *ai.AnomalyIdentifierRequest) (*ai.AnomalyIdentifierResponse, error) {

	var err error
	response := new(ai.AnomalyIdentifierResponse)

	response.AnomalyIds = make([]int32, 10)

	return response, err
}

func main() {
	server := grpc.NewServer()

	instance := new(PasswordGeneratorServiceServer)

	ai.RegisterAnomalyIdentifierServiceServer(server, instance)

	listener, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatal("Unable to create grpc listener:", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
