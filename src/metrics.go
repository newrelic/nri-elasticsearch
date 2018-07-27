package main

import (
	"fmt"

	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/stretchr/objx"
)

func collectNodesMetrics(integration *integration.Integration, response *objx.Map) {
	notFoundMetrics := make([]string, 0)
	for _, metricInfo := range nodeMetricDefs.MetricDefs {
		nodesResponse := response.Get("nodes")
		nodes := nodesResponse.Data().(objx.Map)
		for node := range nodes {

			entity, err := integration.Entity(node, "node")
			if err != nil {
				logger.Errorf("there was an error creating new entity: %v", err)
			}

			metricSet, err := entity.NewMetricSet("nodesMetricSet")
			if err != nil {
				logger.Errorf("there was an error creating new metric set: %v", err)
			}

			nodeData := nodes.Get(node).Data().(objx.Map)

			metricInfoValue, err := parseJSON(nodeData, metricInfo.APIKey)
			if err != nil {
				notFoundMetrics = append(notFoundMetrics, metricInfo.APIKey)
			}
			if metricInfoValue != nil {
				setMetric(metricSet, node, metricInfoValue, metricInfo.SourceType)
			}

		}
	}
}

func collectClusterMetrics(integration *integration.Integration, response *objx.Map) {
	println("collect cluster metrics here")
}

func collectCommonMetrics(integration *integration.Integration, response *objx.Map) {
	println("collect common metrics here")
}

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
	} else if value.IsInt() {
		return value.Int(), nil
	} else {
		return nil, fmt.Errorf("could not parse json for value for key: [%v]: ", key)
	}
}

func convertBoolToInt(val bool) (returnval int) {
	returnval = 0
	if val {
		returnval = 1
	}
	return
}
