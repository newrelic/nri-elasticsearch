package main

import (
	"fmt"

	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/stretchr/objx"
)

func collectNodesMetrics(integration *integration.Integration, response *objx.Map) {
	// notFoundMetrics := make([]string, 0)
	for _, metric := range nodeMetricDefs.MetricDefs {
		nodes := response.Get("nodes")
		for node := range nodes.Data().(objx.Map) {
			entity, err := integration.Entity(node, "node")
			metricSet, err := entity.NewMetricSet("nodesMetricSet")
			if err != nil {
				logger.Errorf("there was an error creating new metric set: %v", err)
			}
			nodesData := nodes.Data().(objx.Map).Get(node).Data().(objx.Map)
			metricInfoValue, err := parseJSON(nodesData, metric.APIKey)
			if err != nil {
				logger.Errorf("there was an error parsing the json:")
			}
		}
	}
}

// func populateMetrics(metricSet *metric.Set, entityType string, response *objx.Map, metricDefs *metricSet) {
// 	// notFoundMetrics := make([]string, 0)
// 	for metricKey := range metricDefs.MetricDefs {
// 		println(metricKey)
// 	}
// }

func setMetric(metricSet *metric.Set, metricName string, metricValue interface{}, metricType metric.SourceType) {
	if err := metricSet.SetMetric(metricName, metricValue, metricType); err != nil {
		logger.Errorf("There was an error when trying to set metric value: %s", err)
	}
}

func parseJSON(jsonData objx.Map, key string) (interface{}, error) {
	value := jsonData.Get(key)
	if value.IsStr() {
		return value.Str(), nil
	} else if value.IsBool() {
		return convertBoolToInt(value.Bool()), nil
	} else if value.IsFloat64() {
		return value.Float64(), nil
	}

	return nil, fmt.Errorf("could not parse json for value for key: [%v]: ", key)
}

func convertBoolToInt(val bool) (returnval int) {
	returnval = 0
	if val {
		returnval = 1
	}
	return
}
