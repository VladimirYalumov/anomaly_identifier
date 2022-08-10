package main

import (
	ai "anomaly_identifier/proto/anomaly_identifier"
	"anomaly_identifier/server/services"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type AnomalyIdentifierServiceServer struct {
	ai.AnomalyIdentifierServiceServer
}

func (s *AnomalyIdentifierServiceServer) Generate(ctx context.Context,
	req *ai.AnomalyIdentifierRequest) (*ai.AnomalyIdentifierResponse, error) {

	response := new(ai.AnomalyIdentifierResponse)
	service, err := services.GetService(req)

	if err != nil {
		return response, err
	}

	err = service.InitDB(
		services.ConnectionSettings[0],
		services.ConnectionSettings[1],
		services.ConnectionSettings[2],
		services.ConnectionSettings[3],
		services.ConnectionSettings[4],
	)

	if err != nil {
		return response, err
	}

	restrictiveQuery := service.MakeRestrictiveQuery(req.GetLimitation())
	response.AnomalyIds, err = service.GetResult(req.Field, restrictiveQuery)

	return response, err
}

func main() {
	server := grpc.NewServer()

	instance := new(AnomalyIdentifierServiceServer)

	ai.RegisterAnomalyIdentifierServiceServer(server, instance)

	listener, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatal("Unable to create grpc listener:", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
