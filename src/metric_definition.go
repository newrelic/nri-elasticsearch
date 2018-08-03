package main

import "github.com/newrelic/infra-integrations-sdk/data/metric"

type metricDefinition struct {
	Name       string
	SourceType metric.SourceType
	APIKey     string
}

type metricSet struct {
	Endpoint   string
	MetricDefs []*metricDefinition
}

var nodeMetricDefs = &metricSet{
	Endpoint: nodeStatsEndpoint,
	MetricDefs: []*metricDefinition{
		{
			Name:       "breakers.estimatedSizeFieldDataCircuitBreakerInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "breakers.fielddata.estimated_size_in_bytes",
		},
		{
			Name:       "breakers.fieldDataCircuitBreakerTripped",
			SourceType: metric.GAUGE,
			APIKey:     "breakers.fielddata.tripped",
		},
		{
			Name:       "breakers.estimatedSizeParentCircuitBreakerInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "breakers.parent.estimated_size_in_bytes",
		},
		{
			Name:       "breakers.parentCircuitBreakerTripped",
			SourceType: metric.GAUGE,
			APIKey:     "breakers.parent.tripped",
		},
		{
			Name:       "breakers.estimatedSizeRequestCircuitBreakerInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "breakers.request.estimated_size_in_bytes",
		},
		{
			Name:       "breakers.requestCircuitBreakerTripped",
			SourceType: metric.GAUGE,
			APIKey:     "breakers.request.tripped",
		},
		{
			Name:       "flush.indexFlushDisk",
			SourceType: metric.GAUGE,
			APIKey:     "indices.indexing.index_total",
		},
		{
			Name:       "flush.timeFlushIndexDiskInSeconds",
			SourceType: metric.GAUGE,
			APIKey:     "indices.indexing.index_time_in_millis",
		},
		{
			Name:       "fs.bytesAvailableJVMInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.fs.total.available_in_bytes",
		},
		{
			Name:       "fs.iOOperations",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.fs.io_stats.total.operations",
		},
		{
			Name:       "fs.bytesUserIoOperationsInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.fs.total.total_in_bytes",
		},
		{
			Name:       "fs.bytesReadsInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.fs.io_stats.total.read_kilobytes",
		},
		{
			Name:       "fs.reads",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.fs.io_stats.total.read_operations",
		},
		{
			Name:       "fs.writesInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.fs.io_stats.total.write_kilobytes",
		},
		{
			Name:       "fs.writesInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.fs.io_stats.total.write_operations",
		},
		{
			Name:       "fs.unallocatedBytesInBYtes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.fs.total.free_in_bytes",
		},
		{
			Name:       "fs.totalSizeInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.fs.total.total_in_bytes",
		},
		{
			Name:       "get.currentRequestsRunning",
			SourceType: metric.GAUGE,
			APIKey:     "indices.get.current",
		},
		{
			Name:       "get.requestsDocumentExistsInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "indices.get.exists_time_in_millis",
		},
		{
			Name:       "get.requestsDcoumentExists",
			SourceType: metric.GAUGE,
			APIKey:     "indices.get.exists_total",
		},
		{
			Name:       "get.requestsDocumentMissingInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "indices.get.missing_time_in_millis",
		},
		{
			Name:       "get.requestsDcoumentMissing",
			SourceType: metric.GAUGE,
			APIKey:     "indices.get.missing_total",
		},
		{
			Name:       "get.timeGetRequestsInMilisecon",
			SourceType: metric.GAUGE,
			APIKey:     "indices.get.time_in_millis",
		},
		{
			Name:       "get.totalGetReqeuests",
			SourceType: metric.GAUGE,
			APIKey:     "indices.get.total",
		},
		{
			Name:       "http.currentOpenConnections",
			SourceType: metric.GAUGE,
			APIKey:     "http.current_open",
		},
		{
			Name:       "http.openedConnections",
			SourceType: metric.GAUGE,
			APIKey:     "http.total_opened",
		},
		{
			Name:       "cache.cacheSizeIDInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.id_cache.memory_size_in_bytes",
		},
		{
			Name:       "indexing.docsCurrentlyDeleted",
			SourceType: metric.GAUGE,
			APIKey:     "indices.indexing.delete_current",
		},
		{
			Name:       "indexing.timeDeletingDocumentsInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "indices.indexing.delete_time_in_millis",
		},
		{
			Name:       "indexing.totalDocumentsDeleted",
			SourceType: metric.GAUGE,
			APIKey:     "indices.indexing.delete_total",
		},
		{
			Name:       "indexing.documentsCurrentlyIndexing",
			SourceType: metric.GAUGE,
			APIKey:     "indices.indexing.index_current",
		},
		{
			Name:       "indexing.timeIndexingDocumentsInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "indices.indexing.index_time_in_millis",
		},
		{
			Name:       "indexing.documentsIndexed",
			SourceType: metric.GAUGE,
			APIKey:     "indices.indexing.index_total",
		},
		{
			Name:       "indices.numberIndices",
			SourceType: metric.GAUGE,
			APIKey:     "indices.docs.count",
		},
		{
			Name:       "indices.indexingOperationsFailed",
			SourceType: metric.GAUGE,
			APIKey:     "indices.indexing.index_failed",
		},
		{
			Name:       "indices.indexingWaitedThrottlingInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "indices.indexing.throttle_time_in_millis",
		},
		{
			Name:       "indices.queryCacheEvictions",
			SourceType: metric.GAUGE,
			APIKey:     "indices.query_cache.evictions",
		},
		{
			Name:       "indices.queryCacheHits",
			SourceType: metric.GAUGE,
			APIKey:     "indices.query_cache.hit_count",
		},
		{
			Name:       "indices.memoryQueryCacheInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.query_cache.memory_size_in_bytes",
		},
		{
			Name:       "indices.queryCacheMisses",
			SourceType: metric.GAUGE,
			APIKey:     "indices.query_cache.miss_count",
		},
		{
			Name:       "indices.recoveryOngoingShardSource",
			SourceType: metric.GAUGE,
			APIKey:     "indices.recovery.current_as_source",
		},
		{
			Name:       "indices.recoveryOngoingShardTarget",
			SourceType: metric.GAUGE,
			APIKey:     "indices.recovery.current_as_target",
		},
		{
			Name:       "indices.recoveryWaitedThrottlingInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "indices.recovery.throttle_time_in_millis",
		},
		{
			Name:       "indices.requestCacheEvicitons",
			SourceType: metric.GAUGE,
			APIKey:     "indices.request_cache.evictions",
		},
		{
			Name:       "indices.requestCacheHits",
			SourceType: metric.GAUGE,
			APIKey:     "indices.request_cache.hit_count",
		},
		{
			Name:       "indices.requestCacheMemoryInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.request_cache.memory_size_in_bytes",
		},
		{
			Name:       "indices.requestCacheMisses",
			SourceType: metric.GAUGE,
			APIKey:     "indices.request_cache.miss_count",
		},
		{
			Name:       "indices.segmentsIndexShard",
			SourceType: metric.GAUGE,
			APIKey:     "indices.segments.count",
		},
		{
			Name:       "indices.segmentsMemoryUsedDocValuesInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.segments.doc_values_memory_in_bytes",
		},
		{
			Name:       "indices.segmentsMemoryUsedFixedBitSetInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.segments.fixed_bit_set_memory_in_bytes",
		},
		{
			Name:       "indices.segmentsMaxMemoryIndexWriterInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.segments.index_writer_max_memory_in_bytes",
		},
		{
			Name:       "indices.segmentsMemoryUsedIndexWriterInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.segments.index_writer_memory_in_bytes",
		},
		{
			Name:       "indices.segmentsMemoryUsedIndexSegmentsInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.segments.memory_in_bytes",
		},
		{
			Name:       "indices.segmentsMemoryUsedNormsInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.segments.norms_memory_in_bytes",
		},
		{
			Name:       "indices.segmentsMemoryUsedStoredFieldsInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.segments.stored_fields_memory_in_bytes",
		},
		{
			Name:       "indices.segmentsMemoryUsedTermVectorsInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.segments.term_vectors_memory_in_bytes",
		},
		{
			Name:       "indices.segmentsMemoryUsedTermsInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.segments.terms_memory_in_bytes",
		},
		{
			Name:       "indices.segmentsMemoryUsedSegmentVersionMapInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.segments.version_map_memory_in_bytes",
		},
		{
			Name:       "indices.translogOperations",
			SourceType: metric.GAUGE,
			APIKey:     "indices.translog.operations",
		},
		{
			Name:       "indices.translogOperationsInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.translog.size_in_bytes",
		},
		{
			Name:       "merges.currentActive",
			SourceType: metric.GAUGE,
			APIKey:     "indices.merges.current",
		},
		{
			Name:       "merges.docsSegementsMerging",
			SourceType: metric.GAUGE,
			APIKey:     "indices.merges.current_docs",
		},
		{
			Name:       "merges.sizeSegementsMergingInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.merges.current_size_in_bytes",
		},
		{
			Name:       "merges.segmentMerges",
			SourceType: metric.GAUGE,
			APIKey:     "indices.merges.total",
		},
		{
			Name:       "merges.docsSegmentMerges",
			SourceType: metric.GAUGE,
			APIKey:     "indices.merges.total_docs",
		},
		{
			Name:       "merges.mergedSegmentsInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.merges.total_size_in_bytes",
		},
		{
			Name:       "merges.totalSegmentMergingInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "indices.merges.total_time_in_millis",
		},
		{
			Name:       "openFD",
			SourceType: metric.GAUGE,
			APIKey:     "process.open_file_descriptors",
		},
		{
			Name:       "refresh.total",
			SourceType: metric.GAUGE,
			APIKey:     "indices.refresh.total",
		},
		{
			Name:       "refresh.totalInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "indices.refresh.total_time_in_millis",
		},
		{
			Name:       "searchFetchCurrentlyRunning",
			SourceType: metric.GAUGE,
			APIKey:     "indices.search.fetch_current",
		},
		{
			Name:       "activeSearches",
			SourceType: metric.GAUGE,
			APIKey:     "indices.search.open_contexts",
		},
		{
			Name:       "activeSearchesInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "indices.search.fetch_time_in_millis",
		},
		{
			Name:       "searchFetches",
			SourceType: metric.GAUGE,
			APIKey:     "indices.search.fetch_total",
		},
		{
			Name:       "currentActiveQueries",
			SourceType: metric.GAUGE,
			APIKey:     "indices.search.query_current",
		},
		{
			Name:       "currentActiveQueriesInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "indices.search.query_time_in_millis",
		},
		{
			Name:       "queriesTotal",
			SourceType: metric.GAUGE,
			APIKey:     "indices.search.query_total",
		},
		{
			Name:       "sizeStoreInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "indices.store.size_in_bytes",
		},
		{
			Name:       "transport.packetsReceived",
			SourceType: metric.GAUGE,
			APIKey:     "transport.rx_count",
		},
		{
			Name:       "transport.packetsReceivedInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "transport.rx_size_in_bytes",
		},
		{
			Name:       "transport.connectionsOpened",
			SourceType: metric.GAUGE,
			APIKey:     "transport.server_open",
		},
		{
			Name:       "transport.packetsSent",
			SourceType: metric.GAUGE,
			APIKey:     "transport.tx_count",
		},
		{
			Name:       "transport.packetsSentInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "transport.tx_size_in_bytes",
		},
		{
			Name:       "vm.gc.collections",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.gc.collection_count",
		},
		{
			Name:       "jvm.gc.collectionsInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.gc.collection_time",
		},
		{
			Name:       "jvm.gc.majorCollectionsOldGenerationObjectsInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.gc.collectors.old.collection_time_in_millis",
		},
		{
			Name:       "jvm.gc.majorCollectionsOldGenerationObjects",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.gc.collectors.old.collection_count",
		},
		{
			Name:       "vm.gc.majorCollectionsYoungGenerationObjectsInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.gc.collectors.young.collection_time_in_millis",
		},
		{
			Name:       "jvm.gc.majorCollectionsYoungGenerationObjects",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.gc.collectors.young.collection_count",
		},
		{
			Name:       "jvm.gc.concurrentMarkSweepInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.gc.concurrent_mark_sweep_collection_time",
		},
		{
			Name:       "jvm.gc.concurrentMarkSweep",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.gc.concurrent_mark_sweep_count",
		},
		{
			Name:       "vm.gc.parallelNewCollectionsInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.gc.par_new_collection_time",
		},
		{
			Name:       "jvm.gc.parallelNewCollections",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.gc.par_new_count",
		},
		{
			Name:       "jvm.mem.heapCommittedInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.mem.heap_committed_in_bytes",
		},
		{
			Name:       "jvm.mem.heapUsed",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.mem.heap_in_use",
		},
		{
			Name:       "jvm.mem.heapMaxInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.mem.heap_max_in_bytes",
		},
		{
			Name:       "jvm.mem.heapUsedInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.mem.heap_used_in_bytes",
		},
		{
			Name:       "jvm.mem.nonHeapCommittedInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.mem.non_heap_committed_in_bytes",
		},
		{
			Name:       "jvm.mem.nonHeapUsedInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.mem.non_heap_used_in_bytes",
		},
		{
			Name:       "jvm.mem.usedYoungGenerationHeapInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.mem.pools.young.used_in_bytes",
		},
		{
			Name:       "jvm.mem.maxYoungGenerationHeapInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.mem.pools.young.max_in_bytes",
		},
		{
			Name:       "jvm.mem.usedOldGenerationHeapInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.mem.pools.old.used_in_bytes",
		},
		{
			Name:       "jvm.mem.maxOldGenerationHeapInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.mem.pools.old.max_in_bytes",
		},
		{
			Name:       "jvm.mem.usedSurvivorSpaceInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.mem.pools.survivor.used_in_bytes",
		},
		{
			Name:       "jvm.mem.maxSurvivorSpaceInBYtes",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.mem.pools.survivor.max_in_bytes",
		},
		{
			Name:       "jvm.ThreadsActive",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.threads.count",
		},
		{
			Name:       "jvm.ThreadsPeak",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.threads.peak_count",
		},
		{
			Name:       "threadpool.bulkActive",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.thread_pool.bulk.active",
		},
		{
			Name:       "threadpool.bulk.Aueue",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.thread_pool.bulk.queue",
		},
		{
			Name:       "threadpool.bulkThreads",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.thread_pool.bulk.threads",
		},
		{
			Name:       "threadpool.bulkRejected",
			SourceType: metric.GAUGE,
			APIKey:     "jvm.thread_pool.bulk.rejected",
		},
		{
			Name:       "threadpoolActivefetchShardStarted",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.fetch_shard_started.active",
		},
		{
			Name:       "threadpool.fetchShardStartedThreads",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.fetch_shard_started.threads",
		},
		{
			Name:       "threadpool.fetchShardStartedQueue",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.fetch_shard_started.queue",
		},
		{
			Name:       "threadpool.fetchShardStartedRejected",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.fetch_shard_started.rejected",
		},
		{
			Name:       "threadpool.fetchShardStoreActive",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.fetch_shard_store.active",
		},
		{
			Name:       "threadpool.fetchShardStoreThreads",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.fetch_shard_store.threads",
		},
		{
			Name:       "threadpool.fetchShardStoreQueue",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.fetch_shard_store.queue",
		},
		{
			Name:       "threadpool.fetchShardStoreRejected",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.fetch_shard_store.rejected",
		},
		{
			Name:       "threadpool.flushActive",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.flush.active",
		},
		{
			Name:       "threadpool.flushQueue",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.flush.queue",
		},
		{
			Name:       "threadpool.flushThreads",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.flush.threads",
		},
		{
			Name:       "threadpool.flushRejected",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.flush.rejected",
		},
		{
			Name:       "threadpool.forceMergeActive",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.force_merge.active",
		},
		{
			Name:       "threadpool.forceMergeThreads",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.force_merge.threads",
		},
		{
			Name:       "threadpool.forceMergeQueue",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.force_merge.queue",
		},
		{
			Name:       "threadpool.forceMergeRejected",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.force_merge.rejected",
		},
		{
			Name:       "threadpool.genericActive",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.generic.active",
		},
		{
			Name:       "threadpool.genericQueue",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.generic.queue",
		},
		{
			Name:       "threadpool.genericThreads",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.generic.threads",
		},
		{
			Name:       "threadpool.genericRejected",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.generic.rejected",
		},
		{
			Name:       "threadpool.getActive",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.get.active",
		},
		{
			Name:       "threadpool.getQueue",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.get.queue",
		},
		{
			Name:       "threadpool.getThreads",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.get.threads",
		},
		{
			Name:       "threadpool.getRejected",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.get.rejected",
		},
		{
			Name:       "threadpool.indexActive",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.index.active",
		},
		{
			Name:       "threadpool.indexQueue",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.index.queue",
		},
		{
			Name:       "threadpool.indexThreads",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.index.threads",
		},
		{
			Name:       "threadpool.indexRejected",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.index.rejected",
		},
		{
			Name:       "threadpool.listenerActive",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.listener.active",
		},
		{
			Name:       "threadpool.listenerQueue",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.listener.queue",
		},
		{
			Name:       "threadpool.listenerThreads",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.listener.threads",
		},
		{
			Name:       "threadpool.listenerRejected",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.listener.rejected",
		},
		{
			Name:       "threadpool.managementActive",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.management.active",
		},
		{
			Name:       "threadpool.managementQueue",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.management.queue",
		},
		{
			Name:       "threadpool.managementThreads",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.management.threads",
		},
		{
			Name:       "threadpool.managementRejected",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.management.rejected",
		},
		{
			Name:       "threadpool.mergeActive",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.merge.active",
		},
		{
			Name:       "threadpool.mergeQueue",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.merge.queue",
		},
		{
			Name:       "threadpool.mergeThreads",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.merge.threads",
		},
		{
			Name:       "threadpool.mergeRejected",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.merge.rejected",
		},
		{
			Name:       "threadpool.percolateActive",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.percolate.active",
		},
		{
			Name:       "threadpool.percolateQueue",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.percolate.queue",
		},
		{
			Name:       "threadpool.percolateThreads",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.percolate.threads",
		},
		{
			Name:       "threadpool.percolateRejected",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.percolate.rejected",
		},
		{
			Name:       "threadpool.refreshActive",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.refresh.active",
		},
		{
			Name:       "threadpool.refreshQueue",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.refresh.queue",
		},
		{
			Name:       "threadpool.refreshThreads",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.refresh.threads",
		},
		{
			Name:       "threadpool.refreshRejected",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.refresh.rejected",
		},
		{
			Name:       "threadpool.searchActive",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.search.active",
		},
		{
			Name:       "threadpool.searchQueue",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.search.queue",
		},
		{
			Name:       "threadpool.searchThreads",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.search.threads",
		},
		{
			Name:       "threadpool.searchRejected",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.search.rejected",
		},
		{
			Name:       "threadpool.snapshotActive",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.snapshot.active",
		},
		{
			Name:       "threadpool.snapshotQueue",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.snapshot.queue",
		},
		{
			Name:       "threadpool.snapshotThreads",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.snapshot.threads",
		},
		{
			Name:       "threadpool.snapshotRejected",
			SourceType: metric.GAUGE,
			APIKey:     "thread_pool.snapshot.rejected",
		},
	},
}

var clusterMetricDefs = &metricSet{
	Endpoint: clusterEndpoint,
	MetricDefs: []*metricDefinition{
		{
			Name:       "activePrimaryShardsCluster",
			SourceType: metric.GAUGE,
			APIKey:     "active_primary_shards",
		},
		{
			Name:       "activeShardsCluster",
			SourceType: metric.GAUGE,
			APIKey:     "active_shards",
		},
		{
			Name:       "index.health",
			SourceType: metric.GAUGE,
			APIKey:     "health",
		},
		{
			Name:       "index.docs",
			SourceType: metric.GAUGE,
			APIKey:     "docs_count",
		},
		{
			Name:       "index.docsDeleted",
			SourceType: metric.GAUGE,
			APIKey:     "docs_deleted",
		},
		{
			Name:       "index.primaryShards",
			SourceType: metric.GAUGE,
			APIKey:     "primary_shards",
		},
		{
			Name:       "index.replicaShards",
			SourceType: metric.GAUGE,
			APIKey:     "replica_shards",
		},
		{
			Name:       "index.primaryStoreSizeInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "primary_store_size",
		},
		{
			Name:       "index.storeSizeInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "store_size",
		},
		{
			Name:       "index.primaryShards",
			SourceType: metric.GAUGE,
			APIKey:     "primary_shards",
		},
		{
			Name:       "dex.replicaShards",
			SourceType: metric.GAUGE,
			APIKey:     "replica_shards",
		},
		{
			Name:       "dex.primaryStoreSizeInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "primary_store_size",
		},
		{
			Name:       "dex.storeSizeInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "store_size",
		},

		{
			Name:       "shards.relocating",
			SourceType: metric.GAUGE,
			APIKey:     "relocating_shards",
		},
		{
			Name:       "shards.Initializing",
			SourceType: metric.GAUGE,
			APIKey:     "initializing_shards",
		},
		{
			Name:       "shards.unassigned",
			SourceType: metric.GAUGE,
			APIKey:     "unassigned_shards",
		},

		{
			Name:       "cluster.dataNodes",
			SourceType: metric.GAUGE,
			APIKey:     "number_of_data_nodes",
		},
		{
			Name:       "cluster.nodes",
			SourceType: metric.GAUGE,
			APIKey:     "number_of_nodes",
		},
		{
			Name:       "cluster.status",
			SourceType: metric.GAUGE,
			APIKey:     "status",
		},
	},
}

var commonStatsMetricDefs = &metricSet{
	Endpoint: commonStatsEndpoint,
	MetricDefs: []*metricDefinition{
		{
			Name:       "primaries.docsnumber",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.docs.count",
		},
		{
			Name:       "primaries.docsDeleted",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.docs.deleted",
		},
		{
			Name:       "primaries.flushesTotal",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.flush.total",
		},

		{
			Name:       "primaries.flushTotalTimeInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.flush.total_time_in_millis",
		},
		{
			Name:       "primaries.get.requestsCurrent",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.get.current",
		},
		{
			Name:       "primaries.get.documentsExistInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.get.exists_time_in_millis",
		},
		{
			Name:       "primaries.get.documentsExist",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.get.exists_total",
		},
		{
			Name:       "primaries.get.documentsMissingInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.get.missing_time_in_millis",
		},
		{
			Name:       "primaries.get.documentsMissing",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.get.missing_total",
		},
		{
			Name:       "primaries.get.requestsInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.get.time_in_millis",
		},
		{
			Name:       "primaries.get.requests",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.get.total",
		},
		{
			Name:       "primaries.index.docsCurrentlyDeleted",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.indexing.delete_current",
		},
		{
			Name:       "primaries.index.docsCurrentlyDeletedInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.indexing.delete_time_in_millis",
		},
		{
			Name:       "primaries.index.docsDeleted",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.indexing.delete_total",
		},
		{
			Name:       "primaries.index.docsCurrentlyIndexing",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.indexing.index_current",
		},
		{
			Name:       "primaries.index.docsCurrentlyIndexingInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.indexing.index_time_in_millis",
		},
		{
			Name:       "primaries.index.docsTotal",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.indexing.index_total",
		},
		{
			Name:       "primaries.merges.current",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.merges.current",
		},
		{
			Name:       "primaries.merges.docsSegementsCurrentlyMerged",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.merges.current_docs",
		},
		{
			Name:       "primaries.merges.segementsCurrentlyMergedInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.merges.current_size_in_bytes",
		},
		{
			Name:       "primaries.merges.segementsTotal",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.merges.total",
		},
		{
			Name:       "primaries.merges.docsTotal",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.merges.total_docs",
		},
		{
			Name:       "primaries.merges.segmentsTotalInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.merges.total_size_in_bytes",
		},
		{
			Name:       "primaries.merges.segmentsTotalInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.merges.total_time_in_millis",
		},
		{
			Name:       "primaries.indexRefreshesTotal",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.refresh.total",
		},
		{
			Name:       "primaries.indexRefreshesTotalInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.refresh.total_time_in_millis",
		},
		{
			Name:       "primaries.queryFetches",
			SourceType: metric.GAUGE,
			APIKey:     "all.primaries.search.fetch_current",
		},
		{
			Name:       "primaries.queryFetchesInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.search.fetch_time_in_millis",
		},
		{
			Name:       "primaries.queryFetchesTotal",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.search.fetch_total",
		},
		{
			Name:       "primaries.queryActive",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.search.query_current",
		},
		{
			Name:       "primaries.queriesInMiliseconds",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.search.query_time_in_millis",
		},
		{
			Name:       "primaries.queriesTotal",
			SourceType: metric.GAUGE,
			APIKey:     "elasticsearch.primaries.search.query.time",
		},
		{
			Name:       "primaries.sizeInBytes",
			SourceType: metric.GAUGE,
			APIKey:     "_all.primaries.store.size_in_bytes",
		},
	},
}
