package main

import (
	"fmt"

	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/stretchr/objx"
)

func populateMetrics(i *integration.Integration, client Client) {
	collectNodesMetrics(i, client)
	collectClusterMetrics(i, client)
	collectCommonMetrics(i, client)
}

func collectNodesMetrics(integration *integration.Integration, client Client) {
	log.Info("Collecting node metrics.")
	responseObjectNode, err := client.Request(nodeStatsEndpoint)
	if err != nil {
		log.Error("Could not get node stats from API: %v", err)
		return
	}

	nodes := responseObjectNode.Get("nodes").ObjxMap()
	// endpoint has multiple nodes so we need to collect for all of them
	for node := range nodes {
		entity, err := integration.Entity(node, "node")
		if err != nil {
			log.Error("Could not create new entity for node [%s]: %v", node, err)
			continue
		}

		metricSet := entity.NewMetricSet("nodesMetricSet")

		nodesData := nodes.Get(node).ObjxMap()
		collectMetrics(nodesData, node, metricSet, nodeMetricDefs)
	}
}

func collectClusterMetrics(integration *integration.Integration, client Client) {
	log.Info("Collecting cluster metrics.")
	responseObjectCluster, err := client.Request(clusterEndpoint)
	if err != nil {
		log.Error("Could not get cluster stats from API: %v", err)
		return
	}

	clusterName := responseObjectCluster.Get("cluster_name").Str()
	entity, err := integration.Entity(clusterName, "cluster")
	if err != nil {
		log.Error("Could not create new entity for cluster: %v", err)
		return
	}
	metricSet := entity.NewMetricSet("clusterMetricSet")

	collectMetrics(responseObjectCluster, clusterName, metricSet, clusterMetricDefs)
}

func collectCommonMetrics(integration *integration.Integration, client Client) {
	log.Info("Collecting common metrics.")
	responseObjectCommon, err := client.Request(commonStatsEndpoint)
	if err != nil {
		log.Error("Could not get common stats from API: %v", err)
		return
	}

	entity, err := integration.Entity("commonMetrics", "common")
	if err != nil {
		log.Error("Could not create new entity for common metrics: %v", err)
		return
	}

	metricSet := entity.NewMetricSet("clusterMetricSet")

	collectMetrics(responseObjectCommon, "commonMetrics", metricSet, commonStatsMetricDefs)
}

// generic function that sets metrics in SDK
func collectMetrics(data objx.Map, metricKey string, metricSet *metric.Set, metricDefs *metricSet) {
	notFoundMetrics := make([]string, 0)
	foundMetrics := make([]string, 0)
	for _, metricInfo := range metricDefs.MetricDefs {
		metricInfoValue, err := parseJSON(data, metricInfo.APIKey)
		if err != nil {
			notFoundMetrics = append(notFoundMetrics, metricInfo.APIKey)
		}
		if metricInfoValue != nil {
			setMetric(metricSet, metricInfo.Name, metricInfoValue, metricInfo.SourceType)
			foundMetrics = append(foundMetrics, metricInfo.APIKey)
		}
	}
}

func setMetric(metricSet *metric.Set, metricName string, metricValue interface{}, metricType metric.SourceType) {
	if err := metricSet.SetMetric(metricName, metricValue, metricType); err != nil {
		log.Error("Could not set metric value: %v", err)
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
