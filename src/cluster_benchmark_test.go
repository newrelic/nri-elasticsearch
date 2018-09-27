package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/newrelic/infra-integrations-sdk/integration"
)

func BenchmarkClusterMetrics(b *testing.B) {
	b.StopTimer()
	i, _ := integration.New("Test", "0.0.1")
	benchmarkClusterMetrics(i, b)
}

func benchmarkClusterMetrics(i *integration.Integration, b *testing.B) {
	clusterMetricsStruct := generateClusterMetricsStruct()

	clusterMetricsJSON, _ := json.Marshal(clusterMetricsStruct)

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client, _ := NewClient("")
	client.baseURL = server.URL

	mux.HandleFunc("/_cluster/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(clusterMetricsJSON))
	})

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		populateClusterMetrics(i, client)
	}
	server.Close()
}

func generateClusterMetricsStruct() *ClusterResponse {
	testName := "testName"
	testStatus := "testStatus"
	mockedClusterMetricsStruct := &ClusterResponse{
		Name:                &testName,
		Status:              &testStatus,
		NumberOfNodes:       getRandomIntPointer(),
		NumberOfDataNodes:   getRandomIntPointer(),
		ActivePrimaryShards: getRandomIntPointer(),
		ActiveShards:        getRandomIntPointer(),
		RelocatingShards:    getRandomIntPointer(),
		InitializingShards:  getRandomIntPointer(),
		UnassignedShards:    getRandomIntPointer(),
	}
	return mockedClusterMetricsStruct
}
