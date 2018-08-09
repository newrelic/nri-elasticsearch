package main

type clusterMetric struct {
	jsonKey string
}

type ClusterResponse struct {
	ClusterName                 string   `json:"cluster_name"`
	Status                      *string  `json:"status"`
	TimedOut                    *bool    `json:"timed_out"`
	NumberOfNodes               *int     `json:"number_of_nodes"`
	NumberOfDataNodes           *int     `json:"number_of_data_nodes"`
	ActivePrimaryShards         *int     `json:"active_primary_shards"`
	ActiveShards                *int     `json:"active_shards" nr_name:"activeShardsCluster" sourceType:"Gauge"`
	RelocatingShards            *int     `json:"relocating_shards"`
	InitializingShards          *int     `json:"initializing_shards"`
	UnassignedShards            *int     `json:"unassigned_shards"`
	DelayedUnassignedShards     *int     `json:"delayed_unassigned_shards"`
	NumberOfPendingTasks        *int     `json:"number_of_pending_tasks"`
	NumberOfInFlightFetch       *int     `json:"number_of_in_flight_fetch"`
	TaskMaxWaitingInQueueMillis *int     `json:"task_max_waiting_in_queue_millis"`
	ActiveShardsPercentAsNumber *float64 `json:"active_shards_percent_as_number"`
}

type NodeResponse struct {
	NodeStats    *NodeCounts      `json:"_nodes"`
	ClusterNamne string           `json:"cluster_name"`
	Nodes        map[string]*Node `json:"nodes"`
}

type NodeCounts struct {
	Total      *int `json:"total"`
	Successful *int `json:"successful"`
	Failed     *int `json:"failed"`
}

type Node struct {
	Name      *string        `json:"name"`
	Host      *string        `json:"host"`
	IP        *string        `json:"ip"`
	Indices   *NodeIndices   `json:"indices"`
	Breakers  *NodeBreakers  `json:"breakers"`
	Process   *NodeProcess   `json:"process"`
	Transport *NodeTransport `json:"transport"`
}

type NodeIndices struct {
	Docs         *IndicesDocs         `json:"docs"`
	Store        *IndicesStore        `json:"store"`
	Indexing     *IndicesIndexing     `json:"indexing"`
	Get          *IndicesGet          `json:"get"`
	Search       *IndicesSearch       `json:"search"`
	Merges       *IndicesMerges       `json:"merges"`
	Refresh      *IndicesRefresh      `json:"refresh"`
	Flush        *IndicesFlush        `json:"flush"`
	QueryCache   *IndicesQueryCache   `json:"query_cache"`
	Segments     *IndicesSegments     `json:"segments"`
	Translog     *IndicesTranslog     `json:"translog"`
	RequestCache *IndicesRequestCache `json:"request_cache"`
	Recovery     *IndicesRecovery     `json:"recovery"`
	IDCache      *IndicesIDCache      `json:"id_cache"`
	NodeFs       *IndicesFs           `json:"fs"`
}

type IndicesDocs struct {
	Count *int `json:"count"`
}

type IndicesStore struct {
	SizeInBytes *int `json:"size_in_bytes"`
}

type IndicesIndexing struct {
	IndexTotal           *int `json:"index_total"`
	IndexTimeInMillis    *int `json:"index_time_in_millis"`
	DeleteCurrent        *int `json:"delete_current"`
	DeleteTimeInMillis   *int `json:"delete_time_in_millis"`
	DeleteTotal          *int `json:"delete_total"`
	IndexCurrent         *int `json:"index_current"`
	IndexFailed          *int `json:"index_failed"`
	ThrottleTimeInMillis *int `json:"throttle_time_in_millis"`
}

type IndicesGet struct {
	Current             *int `json:"current"`
	ExistsTimeInMillis  *int `json:"exists_time_in_millis"`
	ExistsTotal         *int `json:"exists_total"`
	MissingTimeInMillis *int `json:"missing_time_in_millis"`
	MissingTotal        *int `json:"missing_total"`
	TimeInMillis        *int `json:"time_in_millis"`
	Total               *int `json:"total"`
}

type IndicesSearch struct {
	FetchCurrent      *int `json:"fetch_current"`
	OpenContexts      *int `json:"open_contexts"`
	FetchTimeInMillis *int `json:"fetch_time_in_millis"`
	FetchTotal        *int `json:"fetch_total"`
	QueryCurrent      *int `json:"query_current"`
	QueryTimeInMillis *int `json:"query_time_in_millis"`
	QueryTotal        *int `json:"query_total"`
}

type IndicesMerges struct {
	Current            *int `json:"current"`
	CurrentDocs        *int `json:"current_docs"`
	CurrentSizeInBytes *int `json:"current_size_in_bytes"`
	Total              *int `json:"total"`
	TotalDocs          *int `json:"total_docs"`
	TotalSizeInBytes   *int `json:"total_size_in_bytes"`
	TotalTimeInMillis  *int `json:"total_time_in_millis"`
}

type IndicesRefresh struct {
	Total             *int `json:"total"`
	TotalTimeInMillis *int `json:"total_time_in_millis"`
}

type IndicesFlush struct {
	Total             *int `json:"total"`
	TotalTimeInMillis *int `json:"total_time_in_millis"`
}

type IndicesQueryCache struct {
	Evictions         *int `json:"evictions"`
	HitCount          *int `json:"hit_count"`
	MemorySizeInBytes *int `json:"memory_size_in_bytes"`
	MissCount         *int `json:"miss_count"`
}

type IndicesSegments struct {
	Count                       *int `json:"count"`
	DocValuesMemoryInBytes      *int `json:"doc_values_memory_in_bytes"`
	FixedBitSetMemoryInBytes    *int `json:"fixed_bit_set_memory_in_bytes"`
	IndexWriterMaxMemoryInBytes *int `json:"index_writer_max_memory_in_bytes"`
	IndexWriterMemoryInBytes    *int `json:"index_writer_memory_in_bytes"`
	MemoryInBytes               *int `json:"memory_in_bytes"`
	NormsMemoryInBytes          *int `json:"norms_memory_in_bytes"`
	StoredFieldsMemoryInBytes   *int `json:"stored_fields_memory_in_bytes"`
	TermVectorsMemoryInBytes    *int `json:"term_vectors_memory_in_bytes"`
	TermsMemoryInBytes          *int `json:"terms_memory_in_bytes"`
	VersionMapMemoryInBytes     *int `json:"version_map_memory_in_bytes"`
}

type IndicesTranslog struct {
	Operations  *int `json:"operations"`
	SizeInBytes *int `json:"size_in_bytes"`
}

type IndicesRequestCache struct {
	Evictions         *int `json:"evictions"`
	HitCount          *int `json:"hit_count"`
	MemorySizeInBytes *int `json:"memory_size_in_bytes"`
	MissCount         *int `json:"miss_count"`
}

type IndicesRecovery struct {
	CurrentAsSource      *int `json:"current_as_source"`
	CurrentAsTarget      *int `json:"current_as_target"`
	ThrottleTimeInMillis *int `json:"throttle_time_in_millis"`
}

type IndicesIDCache struct {
	MemorySizeInBytes *int `json:"memory_size_in_bytes"`
}

type IndicesJvm struct {
	Gc         *JvmGc         `json:"gc"`
	Mem        *JvmMem        `json:"mem"`
	Threads    *JvmThreads    `json:"threads"`
	ThreadPool *JvmThreadPool `json:"thread_pool"`
}

type IndicesFs struct {
	Total   *NodeTotal   `json:"total"`
	IoStats *NodeIoStats `json:"io_stats"`
}

type NodeBreakers struct {
	Fielddata *BreakersFielddata `json:"fielddata"`
	Parent    *BreakersParent    `json:"parent"`
	Request   *BreakersRequest   `json:"request"`
}

type BreakersFielddata struct {
	EstimatedSizeInBytes *int `json:"estimated_size_in_bytes"`
	Tripped              *int `json:"tripped"`
}

type BreakersParent struct {
	EstimatedSizeInBytes *int `json:"estimated_size_in_bytes"`
	Tripped              *int `json:"tripped"`
}

type BreakersRequest struct {
	EstimatedSizeInBytes *int `json:"estimated_size_in_bytes"`
	Tripped              *int `json:"tripped"`
}

type NodeProcess struct {
	OpenFileDescriptors *int `json:"open_file_descriptors"`
}

type NodeTransport struct {
	RxCount       *int `json:"rx_count"`
	RxSizeInBytes *int `json:"rx_size_in_bytes"`
	ServerOpen    *int `json:"server_open"`
	TxCount       *int `json:"tx_count"`
	TxSizeInBytes *int `json:"tx_size_in_bytes"`
}

type NodeTotal struct {
	AvailableInBytes  *int `json:"available_in_bytes"`
	TotalInBytes      *int `json:"total_in_bytes"`
	FreeInBytes       *int `json:"free_in_bytes"`
	DiskIoSizeInBytes *int `json:"disk_io_size_in_bytes"`
}

type NodeIoStats struct {
	Operations      *int `json:"operations"`
	ReadKilobytes   *int `json:"read_kilobytes"`
	ReadOperations  *int `json:"read_operations"`
	WriteKilobytes  *int `json:"write_kilobytes"`
	WriteOperations *int `json:"write_operations"`
}

type JvmGc struct {
	CollectionCount                   *int            `json:"collection_count"`
	CollectionTime                    *int            `json:"collection_time"`
	ConcurrentMarkSweepCollectionTime *int            `json:"concurrent_mark_sweep_collection_time"`
	ConcurrentMarkSweepCount          *int            `json:"concurrent_mark_sweep_count"`
	ParNewCollectionTime              *int            `json:"par_new_collection_time"`
	ParNewCount                       *int            `json:"par_new_count"`
	Collectors                        *NodeCollectors `json:"collectors"`
}

type NodeCollectors struct {
	Old   *NodeCollectorsOld   `json:"old"`
	Young *NodeCollectorsYoung `json:"young"`
}

type NodeCollectorsOld struct {
	CollectionTimeInMillis *int `json:"collection_time_in_millis"`
	CollectionCount        *int `json:"collection_count"`
}

type NodeCollectorsYoung struct {
	CollectionTimeInMillis *int `json:"collection_time_in_millis"`
	CollectionCount        *int `json:"collection_count"`
}

type JvmMem struct {
	HeapCommittedInBytes    *int       `json:"heap_committed_in_bytes"`
	HeapInUse               *int       `json:"heap_in_use"`
	HeapMaxInBytes          *int       `json:"heap_max_in_bytes"`
	HeapUsedInBytes         *int       `json:"heap_used_in_bytes"`
	NonHeapCommittedInBytes *int       `json:"non_heap_committed_in_bytes"`
	NonHeapUsedInBytes      *int       `json:"non_heap_used_in_bytes"`
	Pools                   *NodePools `json:pools`
}

type NodePools struct {
	Young    *int `json:"young"`
	Old      *int `json:"old"`
	Survivor *int `json:"survivor`
}

type NodePoolsYoung struct {
	UsedInBytes *int `json"used_in_bytes"`
	MaxInBytes  *int `json"max_in_bytes"`
}
type NodePoolsOld struct {
	UsedInBytes *int `json"used_in_bytes"`
	MaxInBytes  *int `json"max_in_bytes"`
}
type NodePoolsSurvivor struct {
	UsedInBytes *int `json"used_in_bytes"`
	MaxInBytes  *int `json"max_in_bytes"`
}

type JvmThreads struct {
	Count     *int `json:"countt"`
	PeakCount *int `json:"peak_count"`
}

type JvmThreadPool struct {
	Bulk              *ThreadPoolBulk              `json:"bulk"`
	FetchShardStarted *ThreadPoolFetchShardStarted `json:"fetch_shard_started"`
	FetchShardStore   *ThreadPoolFetchShardStore   `json:"fetch_shard_store"`
	Flush             *ThreadPoolFlush             `json:"flush"`
	ForceMerge        *ThreadPoolForceMerge        `json:"	force_merge"`
	Generic           *ThreadPoolGeneric           `json:"	generic"`
	Get               *ThreadPoolGet               `json:"	get"`
	Index             *ThreadPoolIndex             `json:"	index"`
	Listener          *ThreadPoolListener          `json:"	listener"`
	Management        *ThreadPoolManagement        `json:"	management"`
	Merge             *ThreadPoolMerge             `json:"	merge"`
	Percolate         *ThreadPoolPercolate         `json:"	percolate"`
	Refresh           *ThreadPoolRefresh           `json:"	refresh"`
	Search            *ThreadPoolSearch            `json:"	search"`
	Snapshot          *ThreadPoolSnapshot          `json:"	snapshot"`
}

type ThreadPoolBulk struct {
	Active   *int `json:"active"`
	Queue    *int `json:"queue"`
	Threads  *int `json:"threads"`
	Rejected *int `json:"rejected"`
}

type ThreadPoolFetchShardStarted struct {
	Active   *int `json:"active"`
	Queue    *int `json:"queue"`
	Threads  *int `json:"threads"`
	Rejected *int `json:"rejected"`
}

type ThreadPoolFetchShardStore struct {
	Active   *int `json:"active"`
	Queue    *int `json:"queue"`
	Threads  *int `json:"threads"`
	Rejected *int `json:"rejected"`
}

type ThreadPoolFlush struct {
	Active   *int `json:"active"`
	Queue    *int `json:"queue"`
	Threads  *int `json:"threads"`
	Rejected *int `json:"rejected"`
}

type ThreadPoolForceMerge struct {
	Active   *int `json:"active"`
	Queue    *int `json:"queue"`
	Threads  *int `json:"threads"`
	Rejected *int `json:"rejected"`
}

type ThreadPoolGeneric struct {
	Active   *int `json:"active"`
	Queue    *int `json:"queue"`
	Threads  *int `json:"threads"`
	Rejected *int `json:"rejected"`
}

type ThreadPoolGet struct {
	Active   *int `json:"active"`
	Queue    *int `json:"queue"`
	Threads  *int `json:"threads"`
	Rejected *int `json:"rejected"`
}

type ThreadPoolIndex struct {
	Active   *int `json:"active"`
	Queue    *int `json:"queue"`
	Threads  *int `json:"threads"`
	Rejected *int `json:"rejected"`
}

type ThreadPoolListener struct {
	Active   *int `json:"active"`
	Queue    *int `json:"queue"`
	Threads  *int `json:"threads"`
	Rejected *int `json:"rejected"`
}

type ThreadPoolManagement struct {
	Active   *int `json:"active"`
	Queue    *int `json:"queue"`
	Threads  *int `json:"threads"`
	Rejected *int `json:"rejected"`
}

type ThreadPoolMerge struct {
	Active   *int `json:"active"`
	Queue    *int `json:"queue"`
	Threads  *int `json:"threads"`
	Rejected *int `json:"rejected"`
}

type ThreadPoolPercolate struct {
	Active   *int `json:"active"`
	Queue    *int `json:"queue"`
	Threads  *int `json:"threads"`
	Rejected *int `json:"rejected"`
}

type ThreadPoolRefresh struct {
	Active   *int `json:"active"`
	Queue    *int `json:"queue"`
	Threads  *int `json:"threads"`
	Rejected *int `json:"rejected"`
}

type ThreadPoolSearch struct {
	Active   *int `json:"active"`
	Queue    *int `json:"queue"`
	Threads  *int `json:"threads"`
	Rejected *int `json:"rejected"`
}

type ThreadPoolSnapshot struct {
	Active   *int `json:"active"`
	Queue    *int `json:"queue"`
	Threads  *int `json:"threads"`
	Rejected *int `json:"rejected"`
}

type ThreadPoolHTTP struct {
	CurrentOpen *int `json:"current_open"`
	TotalOpened *int `json:"total_opened"`
}

// {
// 	"cluster_name": "elasticsearch",
// 	"status": "green",
// 	"timed_out": false,
// 	"number_of_nodes": 1,
// 	"number_of_data_nodes": 1,
// 	"active_primary_shards": 0,
// 	"active_shards": 0,
// 	"relocating_shards": 0,
// 	"initializing_shards": 0,
// 	"unassigned_shards": 0,
// 	"delayed_unassigned_shards": 0,
// 	"number_of_pending_tasks": 0,
// 	"number_of_in_flight_fetch": 0,
// 	"task_max_waiting_in_queue_millis": 0,
// 	"active_shards_percent_as_number": 100.0
// }

// type commonStatsMetricDefs struct {
// 	Shards struct {
// 		Total      int `json:"total"`
// 		Successful int `json:"successful"`
// 		Failed     int `json:"failed"`
// 	} `json:"_shards"`
// 	All struct {
// 		Primaries struct {
// 		} `json:"primaries"`
// 		Total struct {
// 		} `json:"total"`
// 	} `json:"_all"`
// 	Indices struct {
// 	} `json:"indices"`
// }
