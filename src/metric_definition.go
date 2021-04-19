package main

// MasterNodeIdResponse struct (/_cat/master?h=id endpoint)
type MasterNodeIdResponse struct {
	ID string `json:"id"`
}

// CommonMetrics struct
type CommonMetrics struct {
	All     *All `json:"_all"`
	Indices map[string]*Index
}

// All struct
type All struct {
	Primaries *Primaries `json:"primaries"`
}

// Primaries struct
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

// PrimariesDocs struct
type PrimariesDocs struct {
	Count   *int `json:"count" metric_name:"primaries.docsnumber" source_type:"gauge"`
	Deleted *int `json:"deleted" metric_name:"primaries.docsDeleted" source_type:"gauge"`
}

// PrimariesFlush struct
type PrimariesFlush struct {
	Total             *int `json:"total" metric_name:"primaries.flushesTotal" source_type:"gauge"`
	TotalTimeInMillis *int `json:"total_time_in_millis" metric_name:"primaries.flushTotalTimeInMilliseconds" source_type:"gauge"`
}

// PrimariesGet struct
type PrimariesGet struct {
	Current             *int `json:"current" metric_name:"primaries.get.requestsCurrent" source_type:"gauge"`
	ExistsTimeInMillis  *int `json:"exists_time_in_millis" metric_name:"primaries.get.documentsExistInMilliseconds" source_type:"gauge"`
	ExistsTotal         *int `json:"exists_total" metric_name:"primaries.get.documentsExist" source_type:"gauge"`
	MissingTimeInMillis *int `json:"missing_time_in_millis" metric_name:"primaries.get.documentsMissingInMilliseconds" source_type:"gauge"`
	MissingTotal        *int `json:"missing_total" metric_name:"primaries.get.documentsMissing" source_type:"gauge"`
	TimeInMillis        *int `json:"time_in_millis" metric_name:"primaries.get.requestsInMilliseconds" source_type:"gauge"`
	Total               *int `json:"total" metric_name:"primaries.get.requests" source_type:"gauge"`
}

// PrimariesIndexing struct
type PrimariesIndexing struct {
	DeleteCurrent      *int `json:"delete_current" metric_name:"primaries.index.docsCurrentlyDeleted" source_type:"gauge"`
	DeleteTimeInMillis *int `json:"delete_time_in_millis" metric_name:"primaries.index.docsCurrentlyDeletedInMilliseconds" source_type:"gauge"`
	DeleteTotal        *int `json:"delete_total" metric_name:"primaries.index.docsDeleted" source_type:"gauge"`
	IndexCurrent       *int `json:"index_current" metric_name:"primaries.index.docsCurrentlyIndexing" source_type:"gauge"`
	IndexTimeInMillis  *int `json:"index_time_in_millis" metric_name:"primaries.index.docsCurrentlyIndexingInMilliseconds" source_type:"gauge"`
	IndexTotal         *int `json:"index_total" metric_name:"primaries.index.docsTotal" source_type:"gauge"`
}

// PrimariesMerges struct
type PrimariesMerges struct {
	Current            *int `json:"current" metric_name:"primaries.merges.current" source_type:"gauge"`
	CurrentDocs        *int `json:"current_docs" metric_name:"primaries.merges.docsSegmentsCurrentlyMerged" source_type:"gauge"`
	CurrentSizeInBytes *int `json:"current_size_in_bytes" metric_name:"primaries.merges.segmentsCurrentlyMergedInBytes" source_type:"gauge"`
	Total              *int `json:"total" metric_name:"primaries.merges.segmentsTotal" source_type:"gauge"`
	TotalDocs          *int `json:"total_docs" metric_name:"primaries.merges.docsTotal" source_type:"gauge"`
	TotalSizeInBytes   *int `json:"total_size_in_bytes" metric_name:"primaries.merges.segmentsTotalInBytes" source_type:"gauge"`
	TotalTimeInMillis  *int `json:"total_time_in_millis" metric_name:"primaries.merges.segmentsTotalInMilliseconds" source_type:"gauge"`
}

// PrimariesRefresh struct
type PrimariesRefresh struct {
	Total             *int `json:"total" metric_name:"primaries.indexRefreshesTotal" source_type:"gauge"`
	TotalTimeInMillis *int `json:"total_time_in_millis" metric_name:"primaries.indexRefreshesTotalInMilliseconds" source_type:"gauge"`
}

// PrimariesSearch struct
type PrimariesSearch struct {
	FetchCurrent      *int `json:"fetch_current" metric_name:"primaries.queryFetches" source_type:"gauge"`
	FetchTimeInMillis *int `json:"fetch_time_in_millis" metric_name:"primaries.queryFetchesInMilliseconds" source_type:"gauge"`
	FetchTotal        *int `json:"fetch_total" metric_name:"primaries.queryFetchesTotal" source_type:"gauge"`
	QueryCurrent      *int `json:"query_current" metric_name:"primaries.queryActive" source_type:"gauge"`
	QueryTimeInMillis *int `json:"query_time_in_millis" metric_name:"primaries.queriesInMilliseconds" source_type:"gauge"`
	QueryTotal        *int `json:"query_total" metric_name:"primaries.queriesTotal" source_type:"gauge"`
}

// PrimariesStore struct
type PrimariesStore struct {
	SizeInBytes *int `json:"size_in_bytes" metric_name:"primaries.sizeInBytes" source_type:"gauge"`
}

// Index struct
type Index struct {
	Primaries *IndexPrimaryStats `json:"primaries"`
	Totals    *IndexTotalStats   `json:"total"`
}

// IndexPrimaryStats struct
type IndexPrimaryStats struct {
	Store *IndexPrimaryStore `json:"store"`
}

// IndexPrimaryStore struct
type IndexPrimaryStore struct {
	Size *int `json:"size_in_bytes" metric_name:"index.primaryStoreSizeInBytes" source_type:"gauge"`
}

// IndexTotalStats struct
type IndexTotalStats struct {
	Store *IndexTotalStore `json:"store"`
}

// IndexTotalStore struct
type IndexTotalStore struct {
	Size *int `json:"size_in_bytes" metric_name:"index.storeSizeInBytes" source_type:"gauge"`
}

// ClusterResponse struct
type ClusterResponse struct {
	Name                *string `json:"cluster_name"`
	Status              *string `json:"status" metric_name:"cluster.status" source_type:"attribute"`
	NumberOfNodes       *int    `json:"number_of_nodes" metric_name:"cluster.nodes" source_type:"gauge"`
	NumberOfDataNodes   *int    `json:"number_of_data_nodes" metric_name:"cluster.dataNodes" source_type:"gauge"`
	ActivePrimaryShards *int    `json:"active_primary_shards" metric_name:"shards.primaryActive" source_type:"gauge"`
	ActiveShards        *int    `json:"active_shards" metric_name:"shards.active" source_type:"gauge"`
	RelocatingShards    *int    `json:"relocating_shards" metric_name:"shards.relocating" source_type:"gauge"`
	InitializingShards  *int    `json:"initializing_shards" metric_name:"shards.initializing" source_type:"gauge"`
	UnassignedShards    *int    `json:"unassigned_shards" metric_name:"shards.unassigned" source_type:"gauge"`
}

// IndexStats struct
type IndexStats struct {
	Health           *string `json:"health" metric_name:"index.health" source_type:"attribute"`
	DocsCount        *string `json:"docs.count" metric_name:"index.docs" source_type:"gauge"`
	DocsDeleted      *string `json:"docs.deleted" metric_name:"index.docsDeleted" source_type:"gauge"`
	PrimaryShards    *string `json:"pri" metric_name:"index.primaryShards" source_type:"gauge"`
	ReplicaShards    *string `json:"rep" metric_name:"index.replicaShards" source_type:"gauge"`
	PrimaryStoreSize *int    `metric_name:"index.primaryStoreSizeInBytes" source_type:"gauge"`
	StoreSize        *int    `metric_name:"index.storeSizeInBytes" source_type:"gauge"`
	Name             *string `json:"index"`
}

// NodeResponse struct
type NodeResponse struct {
	NodeStats   *NodeCounts      `json:"_nodes"`
	ClusterName string           `json:"cluster_name"`
	Nodes       map[string]*Node `json:"nodes"`
}

// NodeCounts struct
type NodeCounts struct {
	Total      *int `json:"total"`
	Successful *int `json:"successful"`
	Failed     *int `json:"failed"`
}

// Node struct from /_api/nodes
type Node struct {
	Name       *string         `json:"name"`
	Host       *string         `json:"host" metric_name:"node.hostname" source_type:"attribute"`
	RawIP      interface{}     `json:"ip"`
	IP         string          `metric_name:"node.ipAddress" source_type:"attribute"`
	Indices    *NodeIndices    `json:"indices"`
	Breakers   *NodeBreakers   `json:"breakers"`
	Process    *NodeProcess    `json:"process"`
	Transport  *NodeTransport  `json:"transport"`
	Jvm        *NodeJvm        `json:"jvm"`
	Fs         *NodeFs         `json:"fs"`
	ThreadPool *NodeThreadPool `json:"thread_pool"`
	HTTP       *NodeHTTP       `json:"http"`
}

// NodeIndices struct
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
}

// IndicesDocs struct
type IndicesDocs struct {
	Count *int `json:"count" metric_name:"indices.numberIndices" source_type:"gauge"`
}

// IndicesStore struct
type IndicesStore struct {
	SizeInBytes *int `json:"size_in_bytes" metric_name:"sizeStoreInBytes" source_type:"gauge"`
}

// IndicesIndexing struct
type IndicesIndexing struct {
	IndexTotal           *int `json:"index_total" metric_name:"indexing.documentsIndexed" source_type:"gauge"`
	IndexTimeInMillis    *int `json:"index_time_in_millis" metric_name:"indexing.timeIndexingDocumentsInMilliseconds" source_type:"gauge"`
	DeleteCurrent        *int `json:"delete_current" metric_name:"indexing.docsCurrentlyDeleted" source_type:"gauge"`
	DeleteTimeInMillis   *int `json:"delete_time_in_millis" metric_name:"indexing.timeDeletingDocumentsInMilliseconds" source_type:"gauge"`
	DeleteTotal          *int `json:"delete_total" metric_name:"indexing.totalDocumentsDeleted" source_type:"gauge"`
	IndexCurrent         *int `json:"index_current" metric_name:"indexing.documentsCurrentlyIndexing" source_type:"gauge"`
	IndexFailed          *int `json:"index_failed" metric_name:"indices.indexingOperationsFailed" source_type:"gauge"`
	ThrottleTimeInMillis *int `json:"throttle_time_in_millis" metric_name:"indices.indexingWaitedThrottlingInMilliseconds" source_type:"gauge"`
}

// IndicesGet struct
type IndicesGet struct {
	Current             *int `json:"current" metric_name:"get.currentRequestsRunning" source_type:"gauge"`
	ExistsTimeInMillis  *int `json:"exists_time_in_millis" metric_name:"get.requestsDocumentExistsInMilliseconds" source_type:"gauge"`
	ExistsTotal         *int `json:"exists_total" metric_name:"get.requestsDocumentExists" source_type:"gauge"`
	MissingTimeInMillis *int `json:"missing_time_in_millis" metric_name:"get.requestsDocumentMissingInMilliseconds" source_type:"gauge"`
	MissingTotal        *int `json:"missing_total" metric_name:"get.requestsDocumentMissing" source_type:"gauge"`
	TimeInMillis        *int `json:"time_in_millis" metric_name:"get.timeGetRequestsInMilliseconds" source_type:"gauge"`
	Total               *int `json:"total" metric_name:"get.totalGetRequests" source_type:"gauge"`
}

// IndicesSearch struct
type IndicesSearch struct {
	FetchCurrent *int `json:"fetch_current" metric_name:"searchFetchCurrentlyRunning" source_type:"gauge"`
	OpenContexts *int `json:"open_contexts" metric_name:"activeSearches" source_type:"gauge"`
	FetchTotal   *int `json:"fetch_total" metric_name:"searchFetches" source_type:"gauge"`
	QueryTotal   *int `json:"query_total" metric_name:"queriesTotal" source_type:"gauge"`
	QueryTime    *int `json:"query_time_in_millis" metric_name:"activeSearchesInMilliseconds" source_type:"gauge"`
}

// IndicesMerges struct
type IndicesMerges struct {
	Current            *int `json:"current" metric_name:"merges.currentActive" source_type:"gauge"`
	CurrentDocs        *int `json:"current_docs" metric_name:"merges.docsSegmentsMerging" source_type:"gauge"`
	CurrentSizeInBytes *int `json:"current_size_in_bytes" metric_name:"merges.sizeSegmentsMergingInBytes" source_type:"gauge"`
	Total              *int `json:"total" metric_name:"merges.segmentMerges" source_type:"gauge"`
	TotalDocs          *int `json:"total_docs" metric_name:"merges.docsSegmentMerges" source_type:"gauge"`
	TotalSizeInBytes   *int `json:"total_size_in_bytes" metric_name:"merges.mergedSegmentsInBytes" source_type:"gauge"`
	TotalTimeInMillis  *int `json:"total_time_in_millis" metric_name:"merges.totalSegmentMergingInMilliseconds" source_type:"gauge"`
}

// IndicesRefresh struct
type IndicesRefresh struct {
	Total             *int `json:"total" metric_name:"refresh.total" source_type:"gauge"`
	TotalTimeInMillis *int `json:"total_time_in_millis" metric_name:"refresh.totalInMilliseconds" source_type:"gauge"`
}

// IndicesFlush struct
type IndicesFlush struct {
	Total             *int `json:"total" metric_name:"flush.indexRefreshesTotal" source_type:"gauge"`
	TotalTimeInMillis *int `json:"total_time_in_millis" metric_name:"flush.indexRefreshesTotalInMilliseconds" source_type:"gauge"`
}

// IndicesQueryCache struct
type IndicesQueryCache struct {
	Evictions         *int `json:"evictions" metric_name:"indices.queryCacheEvictions" source_type:"gauge"`
	HitCount          *int `json:"hit_count" metric_name:"indices.queryCacheHits" source_type:"gauge"`
	MemorySizeInBytes *int `json:"memory_size_in_bytes" metric_name:"indices.memoryQueryCacheInBytes" source_type:"gauge"`
	MissCount         *int `json:"miss_count" metric_name:"indices.queryCacheMisses" source_type:"gauge"`
}

// IndicesSegments struct
type IndicesSegments struct {
	Count                       *int `json:"count" metric_name:"indices.segmentsIndexShard" source_type:"gauge"`
	DocValuesMemoryInBytes      *int `json:"doc_values_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedDocValuesInBytes" source_type:"gauge"`
	FixedBitSetMemoryInBytes    *int `json:"fixed_bit_set_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedFixedBitSetInBytes" source_type:"gauge"`
	IndexWriterMaxMemoryInBytes *int `json:"index_writer_max_memory_in_bytes" metric_name:"indices.segmentsMaxMemoryIndexWriterInBytes" source_type:"gauge"`
	IndexWriterMemoryInBytes    *int `json:"index_writer_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedIndexWriterInBytes" source_type:"gauge"`
	MemoryInBytes               *int `json:"memory_in_bytes" metric_name:"indices.segmentsMemoryUsedIndexSegmentsInBytes" source_type:"gauge"`
	NormsMemoryInBytes          *int `json:"norms_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedNormsInBytes" source_type:"gauge"`
	StoredFieldsMemoryInBytes   *int `json:"stored_fields_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedStoredFieldsInBytes" source_type:"gauge"`
	TermVectorsMemoryInBytes    *int `json:"term_vectors_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedTermVectorsInBytes" source_type:"gauge"`
	TermsMemoryInBytes          *int `json:"terms_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedTermsInBytes" source_type:"gauge"`
	VersionMapMemoryInBytes     *int `json:"version_map_memory_in_bytes" metric_name:"indices.segmentsMemoryUsedSegmentVersionMapInBytes" source_type:"gauge"`
}

// IndicesTranslog struct
type IndicesTranslog struct {
	Operations  *int `json:"operations" metric_name:"indices.translogOperations" source_type:"gauge"`
	SizeInBytes *int `json:"size_in_bytes" metric_name:"indices.translogOperationsInBytes" source_type:"gauge"`
}

// IndicesRequestCache struct
type IndicesRequestCache struct {
	Evictions         *int `json:"evictions" metric_name:"indices.requestCacheEvictions" source_type:"gauge"`
	HitCount          *int `json:"hit_count" metric_name:"indices.requestCacheHits" source_type:"gauge"`
	MemorySizeInBytes *int `json:"memory_size_in_bytes" metric_name:"indices.requestCacheMemoryInBytes" source_type:"gauge"`
	MissCount         *int `json:"miss_count" metric_name:"indices.requestCacheMisses" source_type:"gauge"`
}

// IndicesRecovery struct
type IndicesRecovery struct {
	CurrentAsSource      *int `json:"current_as_source" metric_name:"indices.recoveryOngoingShardSource" source_type:"gauge"` //
	CurrentAsTarget      *int `json:"current_as_target" metric_name:"indices.recoveryOngoingShardTarget" source_type:"gauge"`
	ThrottleTimeInMillis *int `json:"throttle_time_in_millis" metric_name:"indices.recoveryWaitedThrottlingInMilliseconds" source_type:"gauge"` //
}

// IndicesIDCache struct
type IndicesIDCache struct {
	MemorySizeInBytes *int `json:"memory_size_in_bytes" metric_name:"cache.cacheSizeIDInBytes" source_type:"gauge"`
}

// NodeFs struct
type NodeFs struct {
	Total             *FsTotal             `json:"total"`
	IoStats           *FsIoStats           `json:"io_stats"`
	MostUsageEstimate *FsMostUsageEstimate `json:"most_usage_estimate"`
}

// NodeBreakers struct
type NodeBreakers struct {
	Fielddata *BreakersFielddata `json:"fielddata"`
	Parent    *BreakersParent    `json:"parent"`
	Request   *BreakersRequest   `json:"request"`
}

// BreakersFielddata struct
type BreakersFielddata struct {
	EstimatedSizeInBytes *int `json:"estimated_size_in_bytes" metric_name:"breakers.estimatedSizeFieldDataCircuitBreakerInBytes" source_type:"gauge"`
	Tripped              *int `json:"tripped" metric_name:"breakers.fieldDataCircuitBreakerTripped" source_type:"gauge"`
}

// BreakersParent struct
type BreakersParent struct {
	EstimatedSizeInBytes *int `json:"estimated_size_in_bytes" metric_name:"breakers.estimatedSizeParentCircuitBreakerInBytes" source_type:"gauge"`
	Tripped              *int `json:"tripped" metric_name:"breakers.parentCircuitBreakerTripped" source_type:"gauge"`
}

// BreakersRequest struct
type BreakersRequest struct {
	EstimatedSizeInBytes *int `json:"estimated_size_in_bytes" metric_name:"breakers.estimatedSizeRequestCircuitBreakerInBytes" source_type:"gauge"`
	Tripped              *int `json:"tripped" metric_name:"breakers.requestCircuitBreakerTripped" source_type:"gauge"`
}

// NodeProcess struct
type NodeProcess struct {
	OpenFileDescriptors *int `json:"open_file_descriptors" metric_name:"openFD" source_type:"gauge"`
}

// NodeTransport struct
type NodeTransport struct {
	RxCount       *int `json:"rx_count" metric_name:"transport.packetsReceived" source_type:"gauge"`
	RxSizeInBytes *int `json:"rx_size_in_bytes" metric_name:"transport.packetsReceivedInBytes" source_type:"gauge"`
	ServerOpen    *int `json:"server_open" metric_name:"transport.connectionsOpened" source_type:"gauge"`
	TxCount       *int `json:"tx_count" metric_name:"transport.packetsSent" source_type:"gauge"`
	TxSizeInBytes *int `json:"tx_size_in_bytes" metric_name:"transport.packetsSentInBytes" source_type:"gauge"`
}

// FsTotal struct
type FsTotal struct {
	AvailableInBytes  *int `json:"available_in_bytes" metric_name:"fs.bytesAvailableJVMInBytes" source_type:"gauge"`
	TotalInBytes      *int `json:"total_in_bytes" metric_name:"fs.totalSizeInBytes" source_type:"gauge"`
	FreeInBytes       *int `json:"free_in_bytes" metric_name:"fs.unallocatedBytesInBYtes" source_type:"gauge"`
	DiskIoSizeInBytes *int `json:"disk_io_size_in_bytes" metric_name:"fs.bytesUserIoOperationsInBytes" source_type:"gauge"`
}

// FsMostUsageEstimate struct
type FsMostUsageEstimate struct {
	UsedDiskPercent *float64 `json:"used_disk_percent" metric_name:"fs.usedDiskPercent" source_type:"gauge"`
}

// FsIoStats struct
type FsIoStats struct {
	Devices *IoStatsTotal `json:"total"`
}

// IoStatsTotal struct
type IoStatsTotal struct {
	Operations      *int `json:"operations" metric_name:"fs.iOOperations" source_type:"gauge"`
	ReadKilobytes   *int `json:"read_kilobytes" metric_name:"fs.dataWritten" source_type:"gauge"`
	ReadOperations  *int `json:"read_operations" metric_name:"fs.readOperations" source_type:"gauge"`
	WriteKilobytes  *int `json:"write_kilobytes" metric_name:"fs.dataWritten" source_type:"gauge"`
	WriteOperations *int `json:"write_operations" metric_name:"fs.writeOperations" source_type:"gauge"`
}

// NodeJvm struct
type NodeJvm struct {
	Gc      *JvmGc      `json:"gc"`
	Mem     *JvmMem     `json:"mem"`
	Threads *JvmThreads `json:"threads"`
}

// JvmGc struct
type JvmGc struct {
	CollectionCount                   *int          `json:"collection_count" metric_name:"jvm.gc.collections" source_type:"gauge"`
	CollectionTime                    *int          `json:"collection_time" metric_name:"jvm.gc.collectionsInMilliseconds" source_type:"gauge"`
	ConcurrentMarkSweepCollectionTime *int          `json:"concurrent_mark_sweep_collection_time" metric_name:"jvm.gc.concurrentMarkSweepInMilliseconds" source_type:"gauge"`
	ConcurrentMarkSweepCount          *int          `json:"concurrent_mark_sweep_count" metric_name:"jvm.gc.concurrentMarkSweep" source_type:"gauge"`
	ParNewCollectionTime              *int          `json:"par_new_collection_time" metric_name:"jvm.gc.parallelNewCollectionsInMilliseconds" source_type:"gauge"`
	ParNewCount                       *int          `json:"par_new_count" metric_name:"jvm.gc.parallelNewCollections" source_type:"gauge"`
	Collectors                        *GcCollectors `json:"collectors"`
}

// GcCollectors struct
type GcCollectors struct {
	Old   *CollectorsOld   `json:"old"`
	Young *CollectorsYoung `json:"young"`
}

// CollectorsOld struct
type CollectorsOld struct {
	CollectionTimeInMillis *int `json:"collection_time_in_millis" metric_name:"jvm.gc.majorCollectionsOldGenerationObjectsInMilliseconds" source_type:"gauge"`
	CollectionCount        *int `json:"collection_count" metric_name:"jvm.gc.majorCollectionsOldGenerationObjects" source_type:"gauge"`
}

// CollectorsYoung struct
type CollectorsYoung struct {
	CollectionTimeInMillis *int `json:"collection_time_in_millis" metric_name:"jvm.gc.minorCollectionsYoungGenerationObjectsInMilliseconds" source_type:"gauge"`
	CollectionCount        *int `json:"collection_count" metric_name:"jvm.gc.minorCollectionsYoungGenerationObjects" source_type:"gauge"`
}

// JvmMem struct
type JvmMem struct {
	HeapCommittedInBytes    *int      `json:"heap_committed_in_bytes" metric_name:"jvm.mem.heapCommittedInBytes" source_type:"gauge"`
	HeapInUse               *int      `json:"heap_used_percent" metric_name:"jvm.mem.heapUsed" source_type:"gauge"`
	HeapMaxInBytes          *int      `json:"heap_max_in_bytes" metric_name:"jvm.mem.heapMaxInBytes" source_type:"gauge"`
	HeapUsedInBytes         *int      `json:"heap_used_in_bytes" metric_name:"jvm.mem.heapUsedInBytes" source_type:"gauge"`
	NonHeapCommittedInBytes *int      `json:"non_heap_committed_in_bytes" metric_name:"jvm.mem.nonHeapCommittedInBytes" source_type:"gauge"`
	NonHeapUsedInBytes      *int      `json:"non_heap_used_in_bytes" metric_name:"jvm.mem.nonHeapUsedInBytes" source_type:"gauge"`
	Pools                   *MemPools `json:"pools"`
}

// MemPools struct
type MemPools struct {
	Young    *PoolsYoung    `json:"young"`
	Old      *PoolsOld      `json:"old"`
	Survivor *PoolsSurvivor `json:"survivor"`
}

// PoolsYoung struct
type PoolsYoung struct {
	UsedInBytes *int `json:"used_in_bytes" metric_name:"jvm.mem.usedYoungGenerationHeapInBytes" source_type:"gauge"`
	MaxInBytes  *int `json:"max_in_bytes" metric_name:"jvm.mem.maxYoungGenerationHeapInBytes" source_type:"gauge"`
}

// PoolsOld struct
type PoolsOld struct {
	UsedInBytes *int `json:"used_in_bytes" metric_name:"jvm.mem.usedOldGenerationHeapInBytes" source_type:"gauge"`
	MaxInBytes  *int `json:"max_in_bytes" metric_name:"jvm.mem.maxOldGenerationHeapInBytes" source_type:"gauge"`
}

// PoolsSurvivor struct
type PoolsSurvivor struct {
	UsedInBytes *int `json:"used_in_bytes" metric_name:"jvm.mem.usedSurvivorSpaceInBytes" source_type:"gauge"`
	MaxInBytes  *int `json:"max_in_bytes" metric_name:"jvm.mem.maxSurvivorSpaceInBYtes" source_type:"gauge"`
}

// JvmThreads struct
type JvmThreads struct {
	Count     *int `json:"count" metric_name:"jvm.ThreadsActive" source_type:"gauge"`
	PeakCount *int `json:"peak_count" metric_name:"jvm.ThreadsPeak" source_type:"gauge"`
}

// NodeThreadPool struct
type NodeThreadPool struct {
	Bulk              *ThreadPoolBulk              `json:"bulk"`
	FetchShardStarted *ThreadPoolFetchShardStarted `json:"fetch_shard_started"`
	FetchShardStore   *ThreadPoolFetchShardStore   `json:"fetch_shard_store"`
	Flush             *ThreadPoolFlush             `json:"flush"`
	ForceMerge        *ThreadPoolForceMerge        `json:"force_merge"`
	Generic           *ThreadPoolGeneric           `json:"generic"`
	Get               *ThreadPoolGet               `json:"get"`
	Index             *ThreadPoolIndex             `json:"index"`
	Listener          *ThreadPoolListener          `json:"listener"`
	Management        *ThreadPoolManagement        `json:"management"`
	Merge             *ThreadPoolMerge             `json:"merge"`
	Percolate         *ThreadPoolPercolate         `json:"percolate"`
	Refresh           *ThreadPoolRefresh           `json:"refresh"`
	Search            *ThreadPoolSearch            `json:"search"`
	Snapshot          *ThreadPoolSnapshot          `json:"snapshot"`
}

// ThreadPoolBulk struct
type ThreadPoolBulk struct {
	Active   *int `json:"active" metric_name:"threadpool.bulkActive" source_type:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.bulkQueue" source_type:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.bulkThreads" source_type:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.bulkRejected" source_type:"gauge"`
}

// ThreadPoolFetchShardStarted struct
type ThreadPoolFetchShardStarted struct {
	Active   *int `json:"active" metric_name:"threadpool.activeFetchShardStarted" source_type:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.fetchShardStartedThreads" source_type:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.fetchShardStartedQueue" source_type:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.fetchShardStartedRejected" source_type:"gauge"`
}

// ThreadPoolFetchShardStore struct
type ThreadPoolFetchShardStore struct {
	Active   *int `json:"active" metric_name:"threadpool.fetchShardStoreActive" source_type:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.fetchShardStoreThreads" source_type:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.fetchShardStoreQueue" source_type:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.fetchShardStoreRejected" source_type:"gauge"`
}

// ThreadPoolFlush struct
type ThreadPoolFlush struct {
	Active   *int `json:"active" metric_name:"threadpool.flushActive" source_type:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.flushQueue" source_type:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.flushThreads" source_type:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.flushRejected" source_type:"gauge"`
}

// ThreadPoolForceMerge struct
type ThreadPoolForceMerge struct {
	Active   *int `json:"active" metric_name:"threadpool.forceMergeActive" source_type:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.forceMergeThreads" source_type:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.forceMergeQueue" source_type:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.forceMergeRejected" source_type:"gauge"`
}

// ThreadPoolGeneric struct
type ThreadPoolGeneric struct {
	Active   *int `json:"active" metric_name:"threadpool.genericActive" source_type:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.genericQueue" source_type:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.genericThreads" source_type:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.genericRejected" source_type:"gauge"`
}

// ThreadPoolGet struct
type ThreadPoolGet struct {
	Active   *int `json:"active" metric_name:"threadpool.getActive" source_type:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.getQueue" source_type:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.getThreads" source_type:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.getRejected" source_type:"gauge"`
}

// ThreadPoolIndex struct
type ThreadPoolIndex struct {
	Active   *int `json:"active" metric_name:"threadpool.indexActive" source_type:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.indexQueue" source_type:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.indexThreads" source_type:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.indexRejected" source_type:"gauge"`
}

// ThreadPoolListener struct
type ThreadPoolListener struct {
	Active   *int `json:"active" metric_name:"threadpool.listenerActive" source_type:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.listenerQueue" source_type:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.listenerThreads" source_type:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.listenerRejected" source_type:"gauge"`
}

// ThreadPoolManagement struct
type ThreadPoolManagement struct {
	Active   *int `json:"active" metric_name:"threadpool.managementActive" source_type:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.managementQueue" source_type:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.managementThreads" source_type:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.managementRejected" source_type:"gauge"`
}

// ThreadPoolMerge struct
type ThreadPoolMerge struct {
	Active   *int `json:"active" metric_name:"threadpool.mergeActive" source_type:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.mergeQueue" source_type:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.mergeThreads" source_type:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.mergeRejected" source_type:"gauge"`
}

// ThreadPoolPercolate struct
type ThreadPoolPercolate struct {
	Active   *int `json:"active" metric_name:"threadpool.percolateActive" source_type:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.percolateQueue" source_type:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.percolateThreads" source_type:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.percolateRejected" source_type:"gauge"`
}

// ThreadPoolRefresh struct
type ThreadPoolRefresh struct {
	Active   *int `json:"active" metric_name:"threadpool.refreshActive" source_type:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.refreshQueue" source_type:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.refreshThreads" source_type:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.refreshRejected" source_type:"gauge"`
}

// ThreadPoolSearch struct
type ThreadPoolSearch struct {
	Active   *int `json:"active" metric_name:"threadpool.searchActive" source_type:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.searchQueue" source_type:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.searchThreads" source_type:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.searchRejected" source_type:"gauge"`
}

// ThreadPoolSnapshot struct
type ThreadPoolSnapshot struct {
	Active   *int `json:"active" metric_name:"threadpool.snapshotActive" source_type:"gauge"`
	Queue    *int `json:"queue" metric_name:"threadpool.snapshotQueue" source_type:"gauge"`
	Threads  *int `json:"threads" metric_name:"threadpool.snapshotThreads" source_type:"gauge"`
	Rejected *int `json:"rejected" metric_name:"threadpool.snapshotRejected" source_type:"gauge"`
}

// NodeHTTP struct
type NodeHTTP struct {
	CurrentOpen *int `json:"current_open" metric_name:"http.currentOpenConnections" source_type:"gauge"`
	TotalOpened *int `json:"total_opened" metric_name:"http.openedConnections" source_type:"gauge"`
}
