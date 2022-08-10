package pgsqlclient

import (
	ai "anomaly_identifier/proto/anomaly_identifier"
	"context"
	"google.golang.org/grpc"
)

func getAnomalyIds(field string, restrictiveQuery string) ([]int32, error) {
	conn, _ := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())

	client := ai.NewAnomalyIdentifierServiceClient(conn)
	resp, err := client.Generate(
		context.Background(),
		&ai.AnomalyIdentifierRequest{
			Service:          ai.Services_PGSQL,
			Field:            field,
			RestrictiveQuery: restrictiveQuery,
		},
	)

	if err != nil {
		return []int32{}, err
	}

	return resp.AnomalyIds, nil
}
