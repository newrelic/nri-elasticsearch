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

func BenchmarkNodeMetrics10(b *testing.B) {
	b.StopTimer()
	i, _ := integration.New("Test", "0.0.1")
	benchmarkNodeMetrics(i, b, 10)
}

func BenchmarkNodeMetrics100(b *testing.B) {
	b.StopTimer()
	i, _ := integration.New("Test", "0.0.1")
	benchmarkNodeMetrics(i, b, 100)
}

func BenchmarkNodeMetrics1000(b *testing.B) {
	b.StopTimer()
	i, _ := integration.New("Test", "0.0.1")
	benchmarkNodeMetrics(i, b, 1000)
}

func BenchmarkNodeMetrics10000(b *testing.B) {
	b.StopTimer()
	i, _ := integration.New("Test", "0.0.1")
	benchmarkNodeMetrics(i, b, 10000)
}

func benchmarkNodeMetrics(i *integration.Integration, b *testing.B, numNodes int) {
	nodesMetricsStruct := generateNodeMetricsStruct(numNodes)

	nodeMetricsJSON, _ := json.Marshal(nodesMetricsStruct)

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client, _ := NewClient("")
	client.baseURL = server.URL

	mux.HandleFunc("/_nodes/stats", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(nodeMetricsJSON))
	})

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		populateNodesMetrics(i, client)
	}
	server.Close()
}

func generateNodeMetricsStruct(numNodes int) *NodeResponse {
	mockedNodeMetricsStruct := &NodeResponse{
		NodeStats: &NodeCounts{
			Total:      getRandomIntPointer(),
			Successful: getRandomIntPointer(),
			Failed:     getRandomIntPointer(),
		},
		ClusterName: "BenchmarkCluster",
		Nodes:       make(map[string]*Node),
	}

	for i := 0; i < numNodes; i++ {
		iString := strconv.Itoa(i)
		node := &Node{
			Name: &iString,
			Host: &iString,
			Indices: &NodeIndices{
				Docs: &IndicesDocs{
					Count: getRandomIntPointer(),
				},
				Store: &IndicesStore{
					SizeInBytes: getRandomIntPointer(),
				},
				Indexing: &IndicesIndexing{
					IndexTotal:           getRandomIntPointer(),
					IndexTimeInMillis:    getRandomIntPointer(),
					DeleteCurrent:        getRandomIntPointer(),
					DeleteTimeInMillis:   getRandomIntPointer(),
					DeleteTotal:          getRandomIntPointer(),
					IndexCurrent:         getRandomIntPointer(),
					IndexFailed:          getRandomIntPointer(),
					ThrottleTimeInMillis: getRandomIntPointer(),
				},
				Get: &IndicesGet{
					Current:             getRandomIntPointer(),
					ExistsTimeInMillis:  getRandomIntPointer(),
					ExistsTotal:         getRandomIntPointer(),
					MissingTimeInMillis: getRandomIntPointer(),
					MissingTotal:        getRandomIntPointer(),
					TimeInMillis:        getRandomIntPointer(),
					Total:               getRandomIntPointer(),
				},
				Search: &IndicesSearch{
					FetchCurrent: getRandomIntPointer(),
					OpenContexts: getRandomIntPointer(),
					FetchTotal:   getRandomIntPointer(),
					QueryTotal:   getRandomIntPointer(),
				},
				Merges: &IndicesMerges{
					Current:            getRandomIntPointer(),
					CurrentDocs:        getRandomIntPointer(),
					CurrentSizeInBytes: getRandomIntPointer(),
					Total:              getRandomIntPointer(),
					TotalDocs:          getRandomIntPointer(),
					TotalSizeInBytes:   getRandomIntPointer(),
					TotalTimeInMillis:  getRandomIntPointer(),
				},
				Refresh: &IndicesRefresh{
					Total:             getRandomIntPointer(),
					TotalTimeInMillis: getRandomIntPointer(),
				},
				Flush: &IndicesFlush{
					Total:             getRandomIntPointer(),
					TotalTimeInMillis: getRandomIntPointer(),
				},
				QueryCache: &IndicesQueryCache{
					Evictions:         getRandomIntPointer(),
					HitCount:          getRandomIntPointer(),
					MemorySizeInBytes: getRandomIntPointer(),
					MissCount:         getRandomIntPointer(),
				},
				Segments: &IndicesSegments{
					Count: getRandomIntPointer(),
					DocValuesMemoryInBytes:      getRandomIntPointer(),
					FixedBitSetMemoryInBytes:    getRandomIntPointer(),
					IndexWriterMaxMemoryInBytes: getRandomIntPointer(),
					IndexWriterMemoryInBytes:    getRandomIntPointer(),
					MemoryInBytes:               getRandomIntPointer(),
					NormsMemoryInBytes:          getRandomIntPointer(),
					StoredFieldsMemoryInBytes:   getRandomIntPointer(),
					TermVectorsMemoryInBytes:    getRandomIntPointer(),
					TermsMemoryInBytes:          getRandomIntPointer(),
					VersionMapMemoryInBytes:     getRandomIntPointer(),
				},
				Translog: &IndicesTranslog{
					Operations:  getRandomIntPointer(),
					SizeInBytes: getRandomIntPointer(),
				},
				RequestCache: &IndicesRequestCache{
					Evictions:         getRandomIntPointer(),
					HitCount:          getRandomIntPointer(),
					MemorySizeInBytes: getRandomIntPointer(),
					MissCount:         getRandomIntPointer(),
				},
				Recovery: &IndicesRecovery{
					CurrentAsSource:      getRandomIntPointer(),
					CurrentAsTarget:      getRandomIntPointer(),
					ThrottleTimeInMillis: getRandomIntPointer(),
				},
				IDCache: &IndicesIDCache{
					MemorySizeInBytes: getRandomIntPointer(),
				},
			},
			Breakers: &NodeBreakers{
				Fielddata: &BreakersFielddata{
					EstimatedSizeInBytes: getRandomIntPointer(),
					Tripped:              getRandomIntPointer(),
				},
				Parent: &BreakersParent{
					EstimatedSizeInBytes: getRandomIntPointer(),
					Tripped:              getRandomIntPointer(),
				},
				Request: &BreakersRequest{
					EstimatedSizeInBytes: getRandomIntPointer(),
					Tripped:              getRandomIntPointer(),
				},
			},
			Process: &NodeProcess{
				OpenFileDescriptors: getRandomIntPointer(),
			},
			Transport: &NodeTransport{
				RxCount:       getRandomIntPointer(),
				RxSizeInBytes: getRandomIntPointer(),
				ServerOpen:    getRandomIntPointer(),
				TxCount:       getRandomIntPointer(),
				TxSizeInBytes: getRandomIntPointer(),
			},
			Jvm: &NodeJvm{
				Gc: &JvmGc{
					CollectionCount:                   getRandomIntPointer(),
					CollectionTime:                    getRandomIntPointer(),
					ConcurrentMarkSweepCollectionTime: getRandomIntPointer(),
					ConcurrentMarkSweepCount:          getRandomIntPointer(),
					ParNewCollectionTime:              getRandomIntPointer(),
					ParNewCount:                       getRandomIntPointer(),
					Collectors: &GcCollectors{
						Old: &CollectorsOld{
							CollectionTimeInMillis: getRandomIntPointer(),
							CollectionCount:        getRandomIntPointer(),
						},
						Young: &CollectorsYoung{
							CollectionTimeInMillis: getRandomIntPointer(),
							CollectionCount:        getRandomIntPointer(),
						},
					},
				},
				Mem: &JvmMem{
					HeapCommittedInBytes:    getRandomIntPointer(),
					HeapInUse:               getRandomIntPointer(),
					HeapMaxInBytes:          getRandomIntPointer(),
					HeapUsedInBytes:         getRandomIntPointer(),
					NonHeapCommittedInBytes: getRandomIntPointer(),
					NonHeapUsedInBytes:      getRandomIntPointer(),
					Pools: &MemPools{
						Young: &PoolsYoung{
							UsedInBytes: getRandomIntPointer(),
							MaxInBytes:  getRandomIntPointer(),
						},
						Old: &PoolsOld{
							UsedInBytes: getRandomIntPointer(),
							MaxInBytes:  getRandomIntPointer(),
						},
						Survivor: &PoolsSurvivor{
							UsedInBytes: getRandomIntPointer(),
							MaxInBytes:  getRandomIntPointer(),
						},
					},
				},
				Threads: &JvmThreads{
					Count:     getRandomIntPointer(),
					PeakCount: getRandomIntPointer(),
				},
			},
			Fs: &NodeFs{
				Total: &FsTotal{
					AvailableInBytes:  getRandomIntPointer(),
					TotalInBytes:      getRandomIntPointer(),
					FreeInBytes:       getRandomIntPointer(),
					DiskIoSizeInBytes: getRandomIntPointer(),
				},
				IoStats: &FsIoStats{
					Devices: &IoStatsTotal{
						Operations:      getRandomIntPointer(),
						ReadKilobytes:   getRandomIntPointer(),
						ReadOperations:  getRandomIntPointer(),
						WriteKilobytes:  getRandomIntPointer(),
						WriteOperations: getRandomIntPointer(),
					},
				},
			},
			ThreadPool: &NodeThreadPool{
				Bulk: &ThreadPoolBulk{
					Active:   getRandomIntPointer(),
					Queue:    getRandomIntPointer(),
					Threads:  getRandomIntPointer(),
					Rejected: getRandomIntPointer(),
				},
				FetchShardStarted: &ThreadPoolFetchShardStarted{
					Active:   getRandomIntPointer(),
					Queue:    getRandomIntPointer(),
					Threads:  getRandomIntPointer(),
					Rejected: getRandomIntPointer(),
				},
				FetchShardStore: &ThreadPoolFetchShardStore{
					Active:   getRandomIntPointer(),
					Queue:    getRandomIntPointer(),
					Threads:  getRandomIntPointer(),
					Rejected: getRandomIntPointer(),
				},
				Flush: &ThreadPoolFlush{
					Active:   getRandomIntPointer(),
					Queue:    getRandomIntPointer(),
					Threads:  getRandomIntPointer(),
					Rejected: getRandomIntPointer(),
				},
				ForceMerge: &ThreadPoolForceMerge{
					Active:   getRandomIntPointer(),
					Queue:    getRandomIntPointer(),
					Threads:  getRandomIntPointer(),
					Rejected: getRandomIntPointer(),
				},
				Generic: &ThreadPoolGeneric{
					Active:   getRandomIntPointer(),
					Queue:    getRandomIntPointer(),
					Threads:  getRandomIntPointer(),
					Rejected: getRandomIntPointer(),
				},
				Get: &ThreadPoolGet{
					Active:   getRandomIntPointer(),
					Queue:    getRandomIntPointer(),
					Threads:  getRandomIntPointer(),
					Rejected: getRandomIntPointer(),
				},
				Index: &ThreadPoolIndex{
					Active:   getRandomIntPointer(),
					Queue:    getRandomIntPointer(),
					Threads:  getRandomIntPointer(),
					Rejected: getRandomIntPointer(),
				},
				Listener: &ThreadPoolListener{
					Active:   getRandomIntPointer(),
					Queue:    getRandomIntPointer(),
					Threads:  getRandomIntPointer(),
					Rejected: getRandomIntPointer(),
				},
				Management: &ThreadPoolManagement{
					Active:   getRandomIntPointer(),
					Queue:    getRandomIntPointer(),
					Threads:  getRandomIntPointer(),
					Rejected: getRandomIntPointer(),
				},
				Merge: &ThreadPoolMerge{
					Active:   getRandomIntPointer(),
					Queue:    getRandomIntPointer(),
					Threads:  getRandomIntPointer(),
					Rejected: getRandomIntPointer(),
				},
				Percolate: &ThreadPoolPercolate{
					Active:   getRandomIntPointer(),
					Queue:    getRandomIntPointer(),
					Threads:  getRandomIntPointer(),
					Rejected: getRandomIntPointer(),
				},
				Refresh: &ThreadPoolRefresh{
					Active:   getRandomIntPointer(),
					Queue:    getRandomIntPointer(),
					Threads:  getRandomIntPointer(),
					Rejected: getRandomIntPointer(),
				},
				Search: &ThreadPoolSearch{
					Active:   getRandomIntPointer(),
					Queue:    getRandomIntPointer(),
					Threads:  getRandomIntPointer(),
					Rejected: getRandomIntPointer(),
				},
				Snapshot: &ThreadPoolSnapshot{
					Active:   getRandomIntPointer(),
					Queue:    getRandomIntPointer(),
					Threads:  getRandomIntPointer(),
					Rejected: getRandomIntPointer(),
				},
			},
			HTTP: &NodeHTTP{
				CurrentOpen: getRandomIntPointer(),
				TotalOpened: getRandomIntPointer(),
			},
		}

		temp := strconv.Itoa(i)
		mockedNodeMetricsStruct.Nodes[temp] = node
	}

	return mockedNodeMetricsStruct
}
