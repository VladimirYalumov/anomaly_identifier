package client

import (
	ai "anomaly_identifier/proto/anomaly_identifier"
	"context"
	"google.golang.org/grpc"
)

var limitations []*ai.Limitation

func GetPGSqlAnomalyIds(field string) ([]int64, error) {
	conn, _ := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())

	client := ai.NewAnomalyIdentifierServiceClient(conn)
	resp, err := client.Generate(
		context.Background(),
		&ai.AnomalyIdentifierRequest{
			Service:    ai.Service_PGSQL,
			Field:      field,
			Limitation: limitations,
		},
	)

	if err != nil {
		return []int64{}, err
	}

	return resp.AnomalyIds, nil
}

func AddLimitation(field string, comparisonSign int, value string) (err error) {
	limitation := ai.Limitation{
		Field:          field,
		ComparisonSign: ai.ComparisonSign(comparisonSign),
		Value:          value,
	}
	limitations = append(limitations, &limitation)
	return
}
