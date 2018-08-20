package main

import (
	"fmt"

	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
)

// populateMetrics wrapper to call each of the individual populate functions
func populateMetrics(i *integration.Integration, client Client) {
	err := populateNodesMetrics(i, client)
	if err != nil {
		log.Error("There was an error populating metrics for nodes: %v", err)
	}
	err = populateClusterMetrics(i, client)
	if err != nil {
		log.Error("There was an error populating metrics for clusters: %v", err)
	}
	err = populateCommonMetrics(i, client)
	if err != nil {
		log.Error("There was an error populating metrics for common metrics: %v", err)
	}
	err = populateIndicesMetrics(i, client)
	if err != nil {
		log.Error("There was an error populating metrics for indices: %v", err)
	}
}

func populateNodesMetrics(i *integration.Integration, client Client) error {
	log.Info("Collecting node metrics")
	nodeResponse := new(NodeResponse)
	err := client.Request(nodeStatsEndpoint, &nodeResponse)
	if err != nil {
		return err
	}

	setNodesMetricsResponse(i, nodeResponse)
	return nil
}

// setNodesMetricsResponse calls setMetricsResponse for each node in the response
func setNodesMetricsResponse(integration *integration.Integration, resp *NodeResponse) {
	for node := range resp.Nodes {
		err := setMetricsResponse(integration, resp.Nodes[node], node, "node")
		if err != nil {
			log.Error("There was an error setting metrics for node metrics on %s: %v", node, err)
		}
	}
}

func populateClusterMetrics(i *integration.Integration, client Client) error {
	log.Info("Collecting cluster metrics.")
	clusterResponse := new(ClusterResponse)
	err := client.Request(clusterEndpoint, &clusterResponse)
	if err != nil {
		return err
	}

	if clusterResponse.Name == nil {
		return fmt.Errorf("cannot set metric response, missing cluster name")
	}
	return setMetricsResponse(i, clusterResponse, *clusterResponse.Name, "cluster")
}

func populateCommonMetrics(i *integration.Integration, client Client) error {
	log.Info("Collecting common metrics.")
	commonResponse := new(CommonMetrics)
	err := client.Request(commonStatsEndpoint, &commonResponse)
	if err != nil {
		return err
	}

	return setMetricsResponse(i, commonResponse.All, "commonMetrics", "common")
}

func populateIndicesMetrics(i *integration.Integration, client Client) error {
	log.Info("Collecting indices metrics")
	indicesStats := make([]*IndexStats, 0)
	err := client.Request(indicesStatsEndpoint, &indicesStats)
	if err != nil {
		return err
	}
	setIndicesStatsMetricsResponse(i, indicesStats)
	return nil
}

func setIndicesStatsMetricsResponse(integration *integration.Integration, resp []*IndexStats) {
	for _, object := range resp {
		if object.UUID == nil {
			log.Error("cannot set metric response, missing UUID")
			continue
		}

		if err := setMetricsResponse(integration, object, *object.UUID, "indicies"); err != nil {
			log.Error("There was an error setting metrics for indices metrics: %v", err)
		}
	}
}

// setMetricsResponse creates an entity and a metric set for the
// type of response and calls MarshalMetrics using that response
func setMetricsResponse(integration *integration.Integration, resp interface{}, name string, namespace string) error {
	entity, err := integration.Entity(name, namespace)
	if err != nil {
		return err
	}

	metricSet := entity.NewMetricSet(namespace+"MetricSet",
		metric.Attribute{Key: "displayName", Value: entity.Metadata.Name},
		metric.Attribute{Key: "entityName", Value: entity.Metadata.Namespace + ":" + entity.Metadata.Name},
	)

	return metricSet.MarshalMetrics(resp)
}
