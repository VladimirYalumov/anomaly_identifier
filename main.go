package main

import (
	"anomaly_identifier/client"
)

func main() {
	_ = client.AddLimitation("created_at", 0, "15.08.2022")
	_, _ = client.GetPGSqlAnomalyIds("price")
}
