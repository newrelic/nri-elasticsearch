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
	},
}

var clusterMetricDefs = &metricSet{
	Endpoint: clusterEndpoint,
	MetricDefs: []*metricDefinition{
		{},
	},
}

var commonStatsMetricDefs = &metricSet{
	Endpoint: statsEndpoint,
	MetricDefs: []*metricDefinition{
		{},
	},
}
