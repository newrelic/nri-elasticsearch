package main

import (
	"fmt"
	"io/ioutil"

	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/stretchr/objx"
)

func populateMetrics(i *integration.Integration, client *Client) {
	logger.Infof("Collecting node metrics.")
	stringResponseNode, err := getDataFromEndpoint(client, nodeMetricDefs.Endpoint)
	panicOnErr(err)
	responseObjectNode, err := objx.FromJSON(stringResponseNode)
	panicOnErr(err)
	collectNodesMetrics(i, &responseObjectNode)

	logger.Infof("Collecting cluster metrics.")
	stringResponseCluster, err := getDataFromEndpoint(client, clusterEndpoint)
	panicOnErr(err)
	responseObjectCluster, err := objx.FromJSON(stringResponseCluster)
	panicOnErr(err)
	collectClusterMetrics(i, &responseObjectCluster)

	logger.Infof("Collecting common metrics.")
	stringResponseCommon, err := getDataFromEndpoint(client, commonStatsEndpoint)
	panicOnErr(err)
	responseObjectCommon, err := objx.FromJSON(stringResponseCommon)
	panicOnErr(err)
	collectCommonMetrics(i, &responseObjectCommon)
}

func getDataFromEndpoint(client *Client, endpoint string) (string, error) {
	url := client.BaseURL + endpoint

	response, err := client.client.Get(url)
	if err != nil {
		logger.Errorf("there was an error when getting response from endpoint %v: %v", url, err)
		return "", err
	}

	defer checkErr(response.Body.Close)

	jsonData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Errorf("there was an error when reading the response body: %v", err)
		return "", err
	}

	jsonString := string(jsonData)
	return jsonString, err
}

func collectNodesMetrics(integration *integration.Integration, response *objx.Map) {
	nodesResponse := response.Get("nodes")
	nodes := nodesResponse.Data().(objx.Map)
	// endpoint has multiple nodes so we need to collect for all of them
	for node := range nodes {
		entity, err := integration.Entity(node, "node")
		if err != nil {
			logger.Errorf("there was an error creating new entity for nodes: %v", err)
			continue
		}

		metricSet, err := entity.NewMetricSet("nodesMetricSet")
		if err != nil {
			logger.Errorf("there was an error creating new metric set for nodes: %v", err)
			continue
		}

		nodesData := nodes.Get(node).ObjxMap()
		collectMetrics(nodesData, node, metricSet, nodeMetricDefs)
	}
}

func collectClusterMetrics(integration *integration.Integration, response *objx.Map) {
	clusterName := response.Get("cluster_name").Str()
	entity, err := integration.Entity(clusterName, "cluster")
	if err != nil {
		logger.Errorf("there was an error creating new entity for clusters: %v", err)
		return
	}
	metricSet, err := entity.NewMetricSet("clusterMetricSet")
	if err != nil {
		logger.Errorf("there was an error creating new metric set for clusters: %v", err)
		return
	}

	collectMetrics(*response, clusterName, metricSet, clusterMetricDefs)
}

func collectCommonMetrics(integration *integration.Integration, response *objx.Map) {
	entity, err := integration.Entity("commonMetrics", "common")
	if err != nil {
		logger.Errorf("there was an error creating new entity for common metrics: %v", err)
		return
	}

	metricSet, err := entity.NewMetricSet("clusterMetricSet")
	if err != nil {
		logger.Errorf("there was an error creating new metric set for commmon metrics: %v", err)
		return
	}

	collectMetrics(*response, "commonMetrics", metricSet, commonStatsMetricDefs)
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
