package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/newrelic/infra-integrations-sdk/integration"
)

func BenchmarkIndicesStatsMetrics1(b *testing.B) {
	b.StopTimer()
	i, _ := integration.New("Test", "0.0.1")
	benchmarkIndicesStatsMetrics(i, b, 1)
}

func BenchmarkIndicesStatsMetrics10(b *testing.B) {
	b.StopTimer()
	i, _ := integration.New("Test", "0.0.1")
	benchmarkIndicesStatsMetrics(i, b, 10)
}

func BenchmarkIndicesStatsMetrics100(b *testing.B) {
	b.StopTimer()
	i, _ := integration.New("Test", "0.0.1")
	benchmarkIndicesStatsMetrics(i, b, 100)
}

func BenchmarkIndicesStatsMetrics1000(b *testing.B) {
	b.StopTimer()
	i, _ := integration.New("Test", "0.0.1")
	benchmarkIndicesStatsMetrics(i, b, 1000)
}

func BenchmarkIndicesStatsMetrics10000(b *testing.B) {
	b.StopTimer()
	i, _ := integration.New("Test", "0.0.1")
	benchmarkIndicesStatsMetrics(i, b, 10000)
}

func benchmarkIndicesStatsMetrics(i *integration.Integration, b *testing.B, numIndices int) {
	commonMetricsStruct := generateCommonMetricsStruct(numIndices)
	commonMetricsJSON, _ := json.Marshal(commonMetricsStruct)

	indicesStatsMetricsStruct := generateIndicesStatsMetricsStruct(numIndices)
	indicesStatsMetricsJSON, _ := json.Marshal(indicesStatsMetricsStruct)

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client, _ := NewClient("")
	client.baseURL = server.URL

	mux.HandleFunc("/_stats", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(commonMetricsJSON))
	})
	mux.HandleFunc("/_cat/indices", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(indicesStatsMetricsJSON))
	})

	commonMetricsResponse, _ := populateCommonMetrics(i, client)

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		populateIndicesMetrics(i, client, commonMetricsResponse)
	}
	server.Close()
}
func generateIndicesStatsMetricsStruct(numIndices int) []IndexStats {
	indicesStats := make([]IndexStats, numIndices)
	for j := 0; j < numIndices; j++ {
		jString := strconv.Itoa(j)
		index := IndexStats{
			Health:           &jString,
			DocsCount:        &jString,
			DocsDeleted:      &jString,
			PrimaryShards:    &jString,
			ReplicaShards:    &jString,
			PrimaryStoreSize: getRandomIntPointer(),
			StoreSize:        getRandomIntPointer(),
			Name:             &jString,
		}

		indicesStats[j] = index
	}

	return indicesStats
}
