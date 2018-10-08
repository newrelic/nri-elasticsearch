package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/newrelic/infra-integrations-sdk/integration"
)

func BenchmarkCommonMetrics10(b *testing.B) {
	b.StopTimer()
	i, _ := integration.New("Test", "0.0.1")
	benchmarkCommonMetrics(i, b, 10)
}

func BenchmarkCommonMetrics100(b *testing.B) {
	b.StopTimer()
	i, _ := integration.New("Test", "0.0.1")
	benchmarkCommonMetrics(i, b, 100)
}

func BenchmarkCommonMetrics1000(b *testing.B) {
	b.StopTimer()
	i, _ := integration.New("Test", "0.0.1")
	benchmarkCommonMetrics(i, b, 1000)
}

func BenchmarkCommonMetrics10000(b *testing.B) {
	b.StopTimer()
	i, _ := integration.New("Test", "0.0.1")
	benchmarkCommonMetrics(i, b, 10000)

}

func BenchmarkCommonMetrics100000(b *testing.B) {
	b.StopTimer()
	i, _ := integration.New("Test", "0.0.1")
	benchmarkCommonMetrics(i, b, 1000000)
}

func benchmarkCommonMetrics(i *integration.Integration, b *testing.B, numIndices int) {
	commonMetricsStruct := generateCommonMetricsStruct(numIndices)

	commonMetricsJSON, _ := json.Marshal(commonMetricsStruct)

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client, _ := NewClient("")
	client.baseURL = server.URL

	mux.HandleFunc("/_stats", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(commonMetricsJSON))
	})

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		populateCommonMetrics(i, client)
	}
	server.Close()
}

func generateCommonMetricsStruct(numIndices int) *CommonMetrics {
	mockedCommonMetricsStruct := &CommonMetrics{
		All: &All{
			Primaries: &Primaries{
				Docs: &PrimariesDocs{
					Count:   getRandomIntPointer(),
					Deleted: getRandomIntPointer(),
				},
				Flush: &PrimariesFlush{
					Total:             getRandomIntPointer(),
					TotalTimeInMillis: getRandomIntPointer(),
				},
				Get: &PrimariesGet{
					Current:             getRandomIntPointer(),
					ExistsTimeInMillis:  getRandomIntPointer(),
					ExistsTotal:         getRandomIntPointer(),
					MissingTimeInMillis: getRandomIntPointer(),
					MissingTotal:        getRandomIntPointer(),
					TimeInMillis:        getRandomIntPointer(),
					Total:               getRandomIntPointer(),
				},
				Indexing: &PrimariesIndexing{
					DeleteCurrent:      getRandomIntPointer(),
					DeleteTimeInMillis: getRandomIntPointer(),
					DeleteTotal:        getRandomIntPointer(),
					IndexCurrent:       getRandomIntPointer(),
					IndexTimeInMillis:  getRandomIntPointer(),
					IndexTotal:         getRandomIntPointer(),
				},
				Merges: &PrimariesMerges{
					Current:            getRandomIntPointer(),
					CurrentDocs:        getRandomIntPointer(),
					CurrentSizeInBytes: getRandomIntPointer(),
					Total:              getRandomIntPointer(),
					TotalDocs:          getRandomIntPointer(),
					TotalSizeInBytes:   getRandomIntPointer(),
					TotalTimeInMillis:  getRandomIntPointer(),
				},
				Refresh: &PrimariesRefresh{
					Total:             getRandomIntPointer(),
					TotalTimeInMillis: getRandomIntPointer(),
				},
				Search: &PrimariesSearch{
					FetchCurrent:      getRandomIntPointer(),
					FetchTimeInMillis: getRandomIntPointer(),
					FetchTotal:        getRandomIntPointer(),
					QueryCurrent:      getRandomIntPointer(),
					QueryTimeInMillis: getRandomIntPointer(),
					QueryTotal:        getRandomIntPointer(),
				},
				Store: &PrimariesStore{
					SizeInBytes: getRandomIntPointer(),
				},
			},
		},
		Indices: make(map[string]*Index),
	}

	for i := 0; i < numIndices; i++ {
		index := &Index{
			Primaries: &IndexPrimaryStats{
				Store: &IndexPrimaryStore{
					Size: getRandomIntPointer(),
				},
			},
			Totals: &IndexTotalStats{
				Store: &IndexTotalStore{
					Size: getRandomIntPointer(),
				},
			},
		}
		temp := strconv.Itoa(i)

		mockedCommonMetricsStruct.Indices[temp] = index
	}

	return mockedCommonMetricsStruct
}

func getRandomIntPointer() *int {
	randomInt := rand.Intn(1000)
	return &randomInt
}
