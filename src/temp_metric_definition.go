package main

// All metrics on /_stats
type All struct {
	Primaries *Primaries `json:"primaries"`
}

type Primaries struct {
	Docs     *PrimariesDocs     `json:"docs"`
	Flush    *PrimariesFlush    `json:"flush"`
	Get      *PrimariesGet      `json:"get"`
	Indexing *PrimariesIndexing `json:"indexing"`
	Merges   *PrimariesMerges   `json:"merges"`
	Refresh  *PrimariesRefresh  `json:"refresh"`
	Search   *PrimariesSearch   `json:"search"`
	Store    *PrimariesStore    `json:"store"`
}

type PrimariesDocs struct {
	Count   *int `json:"count" metric_name:"primaries.docsnumber" sourceType:"gauge"`
	Deleted *int `json:"deleted" metric_name:"primaries.docsDeleted" sourceType:"gauge"`
}

type PrimariesFlush struct {
	Total             *int `json:"total" metric_name:"primaries.flushesTotal" sourceType:"gauge"`
	TotalTimeInMillis *int `json:"total_time_in_millis" metric_name:"primaries.flushTotalTimeInMilliseconds" sourceType:"gauge"`
}

type PrimariesGet struct {
	Current             *int `json:"current" metric_name:"primaries.get.requestsCurrent" sourceType:"gauge"`
	ExistsTimeInMillis  *int `json:"exists_time_in_millis" metric_name:"primaries.get.documentsExistInMiliseconds" sourceType:"gauge"`
	ExistsTotal         *int `json:"exists_total" metric_name:"primaries.get.documentsExist" sourceType:"gauge"`
	MissingTimeInMillis *int `json:"missing_time_in_millis" metric_name:"primaries.get.documentsMissingInMiliseconds" sourceType:"gauge"`
	MissingTotal        *int `json:"missing_total" metric_name:"primaries.get.documentsMissing" sourceType:"gauge"`
	TimeInMillis        *int `json:"time_in_millis" metric_name:"primaries.get.requestsInMiliseconds" sourceType:"gauge"`
	Total               *int `json:"total" metric_name:"primaries.get.requests" sourceType:"gauge"`
}

type PrimariesIndexing struct {
	DeleteCurrent      *int `json:"delete_current" metric_name:"primaries.index.docsCurrentlyDeleted" sourceType:"gauge"`
	DeleteTimeInMillis *int `json:"delete_time_in_millis" metric_name:"primaries.index.docsCurrentlyDeletedInMiliseconds" sourceType:"gauge"`
	DeleteTotal        *int `json:"delete_total" metric_name:"primaries.index.docsDeleted" sourceType:"gauge"`
	IndexCurrent       *int `json:"index_current" metric_name:"primaries.index.docsCurrentlyIndexing" sourceType:"gauge"`
	IndexTimeInMillis  *int `json:"index_time_in_millis" metric_name:"primaries.index.docsCurrentlyIndexingInMiliseconds" sourceType:"gauge"`
	IndexTotal         *int `json:"index_total" metric_name:"primaries.index.docsTotal" sourceType:"gauge"`
}

type PrimariesMerges struct {
	Current            *int `json:"current" metric_name:"primaries.merges.current" sourceType:"gauge"`
	CurrentDocs        *int `json:"current_docs" metric_name:"primaries.merges.docsSegementsCurrentlyMerged" sourceType:"gauge"`
	CurrentSizeInBytes *int `json:"current_size_in_bytes" metric_name:"primaries.merges.segementsCurrentlyMergedInBytes" sourceType:"gauge"`
	Total              *int `json:"total" metric_name:"primaries.merges.segementsTotal" sourceType:"gauge"`
	TotalDocs          *int `json:"total_docs" metric_name:"primaries.merges.docsTotal" sourceType:"gauge"`
	TotalSizeInBytes   *int `json:"total_size_in_bytes" metric_name:"primaries.merges.segmentsTotalInBytes" sourceType:"gauge"`
	TotalTimeInMillis  *int `json:"total_time_in_millis" metric_name:"primaries.merges.segmentsTotalInMiliseconds" sourceType:"gauge"`
}

type PrimariesRefresh struct {
	Total             *int `json:"total" metric_name:"primaries.indexRefreshesTotal" sourceType:"gauge"`
	TotalTimeInMillis *int `json:"total_time_in_millis" metric_name:"primaries.indexRefreshesTotalInMiliseconds" sourceType:"gauge"`
}

type PrimariesSearch struct {
	FetchCurrent      *int `json:"fetch_current" metric_name:"primaries.queryFetches" sourceType:"gauge"`
	FetchTimeInMillis *int `json:"fetch_time_in_millis" metric_name:"primaries.queryFetchesInMiliseconds" sourceType:"gauge"`
	FetchTotal        *int `json:"fetch_total" metric_name:"primaries.queryFetchesTotal" sourceType:"gauge"`
	QueryCurrent      *int `json:"query_current" metric_name:"primaries.queryActive" sourceType:"gauge"`
	QueryTimeInMillis *int `json:"query_time_in_millis" metric_name:"primaries.queriesInMiliseconds" sourceType:"gauge"`
	QueryTotal        *int `json:"query_total" metric_name:"primaries.queriesTotal" sourceType:"gauge"`
}

type PrimariesStore struct {
	SizeInBytes *int `json:"size_in_bytes" metric_name:"primaries.sizeInBytes" sourceType:"gauge"`
}

type clusterMetric struct {
	jsonKey string
}

// ClusterResponse metrics on /_cluster/health
type ClusterResponse struct {
	Status              *string `json:"status" metric_name:"cluster.status" sourceType:"gauge"`
	NumberOfNodes       *int    `json:"number_of_nodes" metric_name:"cluster.nodes" sourceType:"gauge"`
	NumberOfDataNodes   *int    `json:"number_of_data_nodes" metric_name:"cluster.dataNodes" sourceType:"gauge"`
	ActivePrimaryShards *int    `json:"active_primary_shards" metric_name:"activePrimaryShardsCluster" sourceType:"gauge"`
	ActiveShards        *int    `json:"active_shards" metric_name:"activeShardsCluster" sourceType:"gauge"`
	RelocatingShards    *int    `json:"relocating_shards" metric_name:"shards.relocating" sourceType:"gauge"`
	InitializingShards  *int    `json:"initializing_shards" metric_name:"shards.initializing" sourceType:"gauge"`
	UnassignedShards    *int    `json:"unassigned_shards" metric_name:"shards.unassigned" sourceType:"gauge"`
}

// IndicesStatsResponse on /_cat/indices?format=json
type IndicesStatsResponse struct {
	Incides *[]IndexStats
}

type IndexStats struct {
	Health           *string `json:"health" metric_name:"index.health" sourceType:"gauge"`
	DocsCount        *string `json:"docs.count" metric_name:"index.docs" sourceType:"gauge"`
	DocsDeleted      *string `json:"docs.deleted" metric_name:"index.docsDeleted" sourceType:"gauge"`
	PrimaryShards    *string `json:"pri" metric_name:"index.primaryShareds" sourceType:"gauge"`
	ReplicaShards    *string `json:"rep" metric_name:"index.replicaShards" sourceType:"gauge"`
	PrimaryStoreSize *string `json:"pri.store.size" metric_name:"index.primaryStoreSizeInBytes" sourceType:"gauge"`
	StoreSize        *string `json:"store.size" metric_name:"index.storeSizeInBytes" sourceType:"gauge"`
}

// NodeResponse metrics on /_nodes/stats
type NodeResponse struct {
	NodeStats   *NodeCounts      `json:"_nodes"`
	ClusterName string           `json:"cluster_name"`
	Nodes       map[string]*Node `json:"nodes"`
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
	Count *int `json:"count" metric_name:"indices.numberIndices" sourceType:"gauge"`
}

type IndicesStore struct {
	SizeInBytes *int `json:"size_in_bytes" metric_name:"primaries.sizeInBytes" sourceType:"gauge"`
}

type IndicesIndexing struct {
	IndexTotal           *int `json:"index_total" metric_name:"indexing.documentsIndexed" sourceType:"gauge"`
	IndexTimeInMillis    *int `json:"index_time_in_millis" metric_name:"indexing.timeIndexingDocumentsInMiliseconds" sourceType:"gauge"`
	DeleteCurrent        *int `json:"delete_current" metric_name:"indexing.docsCurrentlyDeleted" sourceType:"gauge"`
	DeleteTimeInMillis   *int `json:"delete_time_in_millis" metric_name:"indexing.timeDeletingDocumentsInMiliseconds" sourceType:"gauge"`
	DeleteTotal          *int `json:"delete_total" metric_name:"indexing.totalDocumentsDeleted" sourceType:"gauge"`
	IndexCurrent         *int `json:"index_current" metric_name:"indexing.documentsCurrentlyIndexing" sourceType:"gauge"`
	IndexFailed          *int `json:"index_failed" metric_name:"indices.indexingOperationsFailed" sourceType:"gauge"`
	ThrottleTimeInMillis *int `json:"throttle_time_in_millis" metric_name:"indices.indexingWaitedThrottlingInMiliseconds" sourceType:"gauge"`
}

type IndicesGet struct {
	Current             *int `json:"current" metric_name:"get.currentRequestsRunning" sourceType:"gauge"`
	ExistsTimeInMillis  *int `json:"exists_time_in_millis" metric_name:"get.requestsDocumentExistsInMiliseconds" sourceType:"gauge"`
	ExistsTotal         *int `json:"exists_total" metric_name:"get.requestsDcoumentExists" sourceType:"gauge"`
	MissingTimeInMillis *int `json:"missing_time_in_millis" metric_name:"get.requestsDocumentMissingInMiliseconds" sourceType:"gauge"`
	MissingTotal        *int `json:"missing_total" metric_name:"get.requestsDcoumentMissing" sourceType:"gauge"`
	TimeInMillis        *int `json:"time_in_millis" metric_name:"get.timeGetRequestsInMiliseconds" sourceType:"gauge"`
	Total               *int `json:"total" metric_name:"get.totalGetReqeuests" sourceType:"gauge"`
}

type IndicesSearch struct {
	FetchCurrent      *int `json:"fetch_current" metric_name:"primaries.queryFetches" sourceType:"gauge"`
	OpenContexts      *int `json:"open_contexts" metric_name:"activeSearches" sourceType:"gauge"`
	FetchTimeInMillis *int `json:"fetch_time_in_millis" metric_name:"primaries.queryFetchesInMiliseconds" sourceType:"gauge"`
	FetchTotal        *int `json:"fetch_total" metric_name:"primaries.queryFetchesTotal" sourceType:"gauge"`
	QueryCurrent      *int `json:"query_current" metric_name:"primaries.queryActive" sourceType:"gauge"`
	QueryTimeInMillis *int `json:"query_time_in_millis" metric_name:"primaries.queriesInMiliseconds" sourceType:"gauge"`
	QueryTotal        *int `json:"query_total" metric_name:"primaries.queriesTotal" sourceType:"gauge"`
}

type IndicesMerges struct {
	Current            *int `json:"current" metric_name:"merges.currentActive" sourceType:"gauge"`
	CurrentDocs        *int `json:"current_docs" metric_name:"merges.docsSegementsMerging" sourceType:"gauge"`
	CurrentSizeInBytes *int `json:"current_size_in_bytes" metric_name:"merges.sizeSegementsMergingInBytes" sourceType:"gauge"`
	Total              *int `json:"total" metric_name:"merges.segmentMerges" sourceType:"gauge"`
	TotalDocs          *int `json:"total_docs" metric_name:"merges.docsSegmentMerges" sourceType:"gauge"`
	TotalSizeInBytes   *int `json:"total_size_in_bytes" metric_name:"merges.mergedSegmentsInBytes" sourceType:"gauge"`
	TotalTimeInMillis  *int `json:"total_time_in_millis" metric_name:"merges.totalSegmentMergingInMiliseconds" sourceType:"gauge"`
}

type IndicesRefresh struct {
	Total             *int `json:"total" metric_name:"primaries.indexRefreshesTotal" sourceType:"gauge"`
	TotalTimeInMillis *int `json:"total_time_in_millis" metric_name:"primaries.indexRefreshesTotalInMiliseconds" sourceType:"gauge"`
}

type IndicesFlush struct {
	Total             *int `json:"total" metric_name:"flush.indexRefreshesTotal" sourceType:"gauge"`
	TotalTimeInMillis *int `json:"total_time_in_millis" metric_name:"flush.indexRefreshesTotalInMiliseconds" sourceType:"gauge"`
}

type IndicesQueryCache struct {
	Evictions         *int `json:"evictions" metric_name:"indices.queryCacheEvictions" sourceType:"gauge"`
	HitCount          *int `json:"hit_count" metric_name:"indices.queryCacheHits" sourceType:"gauge"`
	MemorySizeInBytes *int `json:"memory_size_in_bytes" metric_name:"indices.memoryQueryCacheInBytes" sourceType:"gauge"`
	MissCount         *int `json:"miss_count" metric_name:"indices.queryCacheMisses" sourceType:"gauge"`
}

type IndicesSegments struct {
	Count                       *int `json:"count" metric_name:"indices.segmentsIndexShard" sourceType:"gauge"`
	DocValuesMemoryInBytes      *int `json:"doc_values_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedDocValuesInBytes" sourceType:"gauge"`
	FixedBitSetMemoryInBytes    *int `json:"fixed_bit_set_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedFixedBitSetInBytes" sourceType:"gauge"`
	IndexWriterMaxMemoryInBytes *int `json:"index_writer_max_memory_in_bytes" metric_name:"indices.segmentsMaxMemoryIndexWriterInBytes" sourceType:"gauge"`
	IndexWriterMemoryInBytes    *int `json:"index_writer_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedIndexWriterInBytes" sourceType:"gauge"`
	MemoryInBytes               *int `json:"memory_in_bytes" metric_name:"indices.segmentsMemoryUsedIndexSegmentsInBytes" sourceType:"gauge"`
	NormsMemoryInBytes          *int `json:"norms_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedNormsInBytes" sourceType:"gauge"`
	StoredFieldsMemoryInBytes   *int `json:"stored_fields_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedStoredFieldsInBytes" sourceType:"gauge"`
	TermVectorsMemoryInBytes    *int `json:"term_vectors_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedTermVectorsInBytes" sourceType:"gauge"`
	TermsMemoryInBytes          *int `json:"terms_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedTermsInBytes" sourceType:"gauge"`
	VersionMapMemoryInBytes     *int `json:"version_map_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedSegmentVersionMapInBytes" sourceType:"gauge"`
}

type IndicesTranslog struct {
	Operations  *int `json:"operations" metric_name:"indices.translogOperations" sourceType:"gauge"`
	SizeInBytes *int `json:"size_in_bytes" metric_name:"indices.translogOperationsInBytes" sourceType:"gauge"`
}

type IndicesRequestCache struct {
	Evictions         *int `json:"evictions" metric_name:"indices.requestCacheEvicitons" sourceType:"gauge"`
	HitCount          *int `json:"hit_count" metric_name:"indices.requestCacheHits" sourceType:"gauge"`
	MemorySizeInBytes *int `json:"memory_size_in_bytes" metric_name:"indices.requestCacheMemoryInBytes" sourceType:"gauge"`
	MissCount         *int `json:"miss_count" metric_name:"indices.requestCacheMisses" sourceType:"gauge"`
}

type IndicesRecovery struct {
	CurrentAsSource      *int `json:"current_as_source" metric_name:"indices.recoveryOngoingShardSource" sourceType:"gauge"`
	CurrentAsTarget      *int `json:"current_as_target" metric_name:"indices.recoveryOngoingShardTarget" sourceType:"gauge"`
	ThrottleTimeInMillis *int `json:"throttle_time_in_millis" metric_name:"indices.recoveryWaitedThrottlingInMiliseconds" sourceType:"gauge"`
}

type IndicesIDCache struct {
	MemorySizeInBytes *int `json:"memory_size_in_bytes" metric_name:"cache.cacheSizeIDInBytes" sourceType:"gauge"`
}

type IndicesJvm struct {
	Gc         *JvmGc         `json:"gc"`
	Mem        *JvmMem        `json:"mem"`
	Threads    *JvmThreads    `json:"threads"`
	ThreadPool *JvmThreadPool `json:"thread_pool"`
}

type IndicesFs struct {
	Total   *FsTotal   `json:"total"`
	IoStats *FsIoStats `json:"io_stats"`
}

type NodeBreakers struct {
	Fielddata *BreakersFielddata `json:"fielddata"`
	Parent    *BreakersParent    `json:"parent"`
	Request   *BreakersRequest   `json:"request"`
}

type BreakersFielddata struct {
	EstimatedSizeInBytes *int `json:"estimated_size_in_bytes" metric_name:"breakers.estimatedSizeFieldDataCircuitBreakerInBytes" sourceType:"gauge"`
	Tripped              *int `json:"tripped" metric_name:"breakers.fieldDataCircuitBreakerTripped" sourceType:"gauge"`
}

type BreakersParent struct {
	EstimatedSizeInBytes *int `json:"estimated_size_in_bytes" metric_name:"breakers.estimatedSizeParentCircuitBreakerInBytes" sourceType:"gauge"`
	Tripped              *int `json:"tripped" metric_name:"breakers.parentCircuitBreakerTripped" sourceType:"gauge"`
}

type BreakersRequest struct {
	EstimatedSizeInBytes *int `json:"estimated_size_in_bytes" metric_name:"breakers.estimatedSizeRequestCircuitBreakerInBytes" sourceType:"gauge"`
	Tripped              *int `json:"tripped" metric_name:"breakers.requestCircuitBreakerTripped" sourceType:"gauge"`
}

type NodeProcess struct {
	OpenFileDescriptors *int `json:"open_file_descriptors" metric_name:"openFD" sourceType:"gauge"`
}

type NodeTransport struct {
	RxCount       *int `json:"rx_count" metric_name:"transport.packetsReceived" sourceType:"gauge"`
	RxSizeInBytes *int `json:"rx_size_in_bytes" metric_name:"transport.packetsReceivedInBytes" sourceType:"gauge"`
	ServerOpen    *int `json:"server_open" metric_name:"transport.connectionsOpened" sourceType:"gauge"`
	TxCount       *int `json:"tx_count" metric_name:"transport.packetsSent" sourceType:"gauge"`
	TxSizeInBytes *int `json:"tx_size_in_bytes" metric_name:"transport.packetsSentInBytes" sourceType:"gauge"`
}

type FsTotal struct {
	AvailableInBytes  *int `json:"available_in_bytes" metric_name:"fs.bytesAvailableJVMInBytes" sourceType:"gauge"`
	TotalInBytes      *int `json:"total_in_bytes" metric_name:"fs.totalSizeInBytes" sourceType:"gauge"`
	FreeInBytes       *int `json:"free_in_bytes" metric_name:"fs.unallocatedBytesInBYtes" sourceType:"gauge"`
	DiskIoSizeInBytes *int `json:"disk_io_size_in_bytes" metric_name:"fs.bytesUserIoOperationsInBytes" sourceType:"gauge"`
}

type FsIoStats struct {
	Operations      *int `json:"operations" metric_name:"fs.iOOperations" sourceType:"gauge"`
	ReadKilobytes   *int `json:"read_kilobytes" metric_name:"fs.bytesReadsInBytes" sourceType:"gauge"`
	ReadOperations  *int `json:"read_operations" metric_name:"fs.reads" sourceType:"gauge"`
	WriteKilobytes  *int `json:"write_kilobytes" metric_name:"fs.writesInBytes" sourceType:"gauge"`
	WriteOperations *int `json:"write_operations" metric_name:"fs.writesInBytes" sourceType:"gauge"`
}

type JvmGc struct {
	CollectionCount                   *int          `json:"collection_count" metric_name:"jvm.gc.collections" sourceType:"gauge"`
	CollectionTime                    *int          `json:"collection_time" metric_name:"jvm.gc.collectionsInMiliseconds" sourceType:"gauge"`
	ConcurrentMarkSweepCollectionTime *int          `json:"concurrent_mark_sweep_collection_time" metric_name:"jvm.gc.concurrentMarkSweepInMiliseconds" sourceType:"gauge"`
	ConcurrentMarkSweepCount          *int          `json:"concurrent_mark_sweep_count" metric_name:"jvm.gc.concurrentMarkSweep" sourceType:"gauge"`
	ParNewCollectionTime              *int          `json:"par_new_collection_time" metric_name:"jvm.gc.parallelNewCollectionsInMiliseconds" sourceType:"gauge"`
	ParNewCount                       *int          `json:"par_new_count" metric_name:"jvm.gc.parallelNewCollections" sourceType:"gauge"`
	Collectors                        *GcCollectors `json:"collectors"`
}

type GcCollectors struct {
	Old   *CollectorsOld   `json:"old"`
	Young *CollectorsYoung `json:"young"`
}

type CollectorsOld struct {
	CollectionTimeInMillis *int `json:"collection_time_in_millis" metric_name:"jvm.gc.majorCollectionsOldGenerationObjectsInMiliseconds" sourceType:"gauge"`
	CollectionCount        *int `json:"collection_count" metric_name:"jvm.gc.majorCollectionsOldGenerationObjects" sourceType:"gauge"`
}

type CollectorsYoung struct {
	CollectionTimeInMillis *int `json:"collection_time_in_millis" metric_name:"jvm.gc.majorCollectionsYoungGenerationObjectsInMiliseconds" sourceType:"gauge"`
	CollectionCount        *int `json:"collection_count" metric_name:"jvm.gc.majorCollectionsYoungGenerationObjects" sourceType:"gauge"`
}

type JvmMem struct {
	HeapCommittedInBytes    *int      `json:"heap_committed_in_bytes" metric_name:"jvm.mem.heapCommittedInBytes" sourceType:"gauge"`
	HeapInUse               *int      `json:"heap_in_use" metric_name:"jvm.mem.heapUsed" sourceType:"gauge"`
	HeapMaxInBytes          *int      `json:"heap_max_in_bytes" metric_name:"jvm.mem.heapMaxInBytes" sourceType:"gauge"`
	HeapUsedInBytes         *int      `json:"heap_used_in_bytes" metric_name:"jvm.mem.heapUsedInBytes" sourceType:"gauge"`
	NonHeapCommittedInBytes *int      `json:"non_heap_committed_in_bytes" metric_name:"jvm.mem.nonHeapCommittedInBytes" sourceType:"gauge"`
	NonHeapUsedInBytes      *int      `json:"non_heap_used_in_bytes" metric_name:"jvm.mem.nonHeapUsedInBytes" sourceType:"gauge"`
	Pools                   *MemPools `json:pools`
}

type MemPools struct {
	Young    *PoolsYoung    `json:"young"`
	Old      *PoolsOld      `json:"old"`
	Survivor *PoolsSurvivor `json:"survivor`
}

type PoolsYoung struct {
	UsedInBytes *int `json"used_in_bytes" metric_name:"jvm.mem.usedYoungGenerationHeapInBytes" sourceType:"gauge"`
	MaxInBytes  *int `json"max_in_bytes" metric_name:"jvm.mem.maxYoungGenerationHeapInBytes" sourceType:"gauge"`
}
type PoolsOld struct {
	UsedInBytes *int `json"used_in_bytes" metric_name:"jvm.mem.usedOldGenerationHeapInBytes" sourceType:"gauge"`
	MaxInBytes  *int `json"max_in_bytes" metric_name:"jvm.mem.maxOldGenerationHeapInBytes" sourceType:"gauge"`
}
type PoolsSurvivor struct {
	UsedInBytes *int `json"used_in_bytes" metric_name:"jvm.mem.usedSurvivorSpaceInBytes" sourceType:"gauge"`
	MaxInBytes  *int `json"max_in_bytes" metric_name:"jvm.mem.maxSurvivorSpaceInBYtes" sourceType:"gauge"`
}

type JvmThreads struct {
	Count     *int `json:"count" metric_name:"jvm.ThreadsActive" sourceType:"gauge"`
	PeakCount *int `json:"peak_count" metric_name:"jvm.ThreadsPeak" sourceType:"gauge"`
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
	Active   *int `json:"active" metric_name:"threadpool.bulkActive" sourceType:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.bulk.Aueue" sourceType:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.bulkThreads" sourceType:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.bulkRejected" sourceType:"gauge"`
}

type ThreadPoolFetchShardStarted struct {
	Active   *int `json:"active" metric_name:"threadpoolActivefetchShardStarted" sourceType:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.fetchShardStartedThreads" sourceType:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.fetchShardStartedQueue" sourceType:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.fetchShardStartedRejected" sourceType:"gauge"`
}

type ThreadPoolFetchShardStore struct {
	Active   *int `json:"active" metric_name:"threadpool.fetchShardStoreActive" sourceType:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.fetchShardStoreThreads" sourceType:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.fetchShardStoreQueue" sourceType:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.fetchShardStoreRejected" sourceType:"gauge"`
}

type ThreadPoolFlush struct {
	Active   *int `json:"active" metric_name:"threadpool.flushActive" sourceType:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.flushQueue" sourceType:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.flushThreads" sourceType:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.flushRejected" sourceType:"gauge"`
}

type ThreadPoolForceMerge struct {
	Active   *int `json:"active" metric_name:"threadpool.forceMergeActive" sourceType:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.forceMergeThreads" sourceType:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.forceMergeQueue" sourceType:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.forceMergeRejected" sourceType:"gauge"`
}

type ThreadPoolGeneric struct {
	Active   *int `json:"active" metric_name:"threadpool.genericActive" sourceType:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.genericQueue" sourceType:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.genericThreads" sourceType:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.genericRejected" sourceType:"gauge"`
}

type ThreadPoolGet struct {
	Active   *int `json:"active" metric_name:"threadpool.getActive" sourceType:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.getQueue" sourceType:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.getThreads" sourceType:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.getRejected" sourceType:"gauge"`
}

type ThreadPoolIndex struct {
	Active   *int `json:"active" metric_name:"threadpool.indexActive" sourceType:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.indexQueue" sourceType:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.indexThreads" sourceType:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.indexRejected" sourceType:"gauge"`
}

type ThreadPoolListener struct {
	Active   *int `json:"active" metric_name:"threadpool.listenerActive" sourceType:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.listenerQueue" sourceType:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.listenerThreads" sourceType:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.listenerRejected" sourceType:"gauge"`
}

type ThreadPoolManagement struct {
	Active   *int `json:"active" metric_name:"threadpool.managementActive" sourceType:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.managementQueue" sourceType:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.managementThreads" sourceType:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.managementRejected" sourceType:"gauge"`
}

type ThreadPoolMerge struct {
	Active   *int `json:"active" metric_name:"threadpool.mergeActive" sourceType:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.mergeQueue" sourceType:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.mergeThreads" sourceType:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.mergeRejected" sourceType:"gauge"`
}

type ThreadPoolPercolate struct {
	Active   *int `json:"active" metric_name:"threadpool.percolateActive" sourceType:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.percolateQueue" sourceType:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.percolateThreads" sourceType:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.percolateRejected" sourceType:"gauge"`
}

type ThreadPoolRefresh struct {
	Active   *int `json:"active" metric_name:"threadpool.refreshActive" sourceType:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.refreshQueue" sourceType:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.refreshThreads" sourceType:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.refreshRejected" sourceType:"gauge"`
}

type ThreadPoolSearch struct {
	Active   *int `json:"active" metric_name:"threadpool.searchActive" sourceType:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.searchQueue" sourceType:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.searchThreads" sourceType:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.searchRejected" sourceType:"gauge"`
}

type ThreadPoolSnapshot struct {
	Active   *int `json:"active" metric_name:"threadpool.snapshotActive" sourceType:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.snapshotQueue" sourceType:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.snapshotThreads" sourceType:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.snapshotRejected" sourceType:"gauge"`
}

type ThreadPoolHTTP struct {
	CurrentOpen *int `json:"current_open" metric_name:"http.currentOpenConnections" sourceType:"gauge"`
	TotalOpened *int `json:"total_opened" metric_name:"http.openedConnections" sourceType:"gauge"`
}
