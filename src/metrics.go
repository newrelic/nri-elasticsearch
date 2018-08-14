package main

import (
	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
)

func populateMetrics(i *integration.Integration, client Client) {
	populateNodesMetrics(i, client)
	populateClusterMetrics(i, client)
	populateCommonMetrics(i, client)
	populateIndicesMetrics(i, client)
}

func populateNodesMetrics(i *integration.Integration, client Client) {
	logger.Infof("Collecting node metrics")
	nodeResponse := new(NodeResponse)
	err := client.Request(nodeStatsEndpoint, nodeResponse)
	if err != nil {
		logger.Errorf("there was an error creating request for node metrics", err)
	}

	setNodesMetricsResponse(i, nodeResponse)
}

func setNodesMetricsResponse(integration *integration.Integration, resp *NodeResponse) {
	for node := range resp.Nodes {
		entity, err := integration.Entity(node, "node")
		if err != nil {
			logger.Errorf("there was an error creating new entity for nodes: %v", err)
			continue
		}

		metricSet := entity.NewMetricSet("nodesMetricSet")
		if err != nil {
			logger.Errorf("there was an error creating new metric set for nodes: %v", err)
			continue
		}
		err = metricSet.MarshalMetrics(resp.Nodes[node])
		if err != nil {
			logger.Errorf("there was an error marshaling metrics for node %s: ", node, err)
		}
	}
}

func populateClusterMetrics(i *integration.Integration, client Client) {
	logger.Infof("Collecting cluster metrics.")
	clusterResponse := new(ClusterResponse)
	err := client.Request(clusterEndpoint, clusterResponse)
	if err != nil {
		logger.Errorf("there was an error creating request for cluster metrics", err)
	}

	err = setClusterMetricsResponse(i, clusterResponse)
	if err != nil {
		logger.Errorf("there was an error setting metrics for cluster metrics", err)
	}
}

func setClusterMetricsResponse(integration *integration.Integration, resp *ClusterResponse) error {
	clusterName := *resp.Name
	entity, err := integration.Entity(clusterName, "cluster")
	if err != nil {
		logger.Errorf("there was an error creating new entity for clusters: %v", err)
	}

	metricSet := entity.NewMetricSet("clusterMetricSet")
	if err != nil {
		logger.Errorf("there was an error creating new metric set for clusters: %v", err)
	}

	return metricSet.MarshalMetrics(resp)
}

func populateCommonMetrics(i *integration.Integration, client Client) {
	logger.Infof("Collecting common metrics.")
	commonResponse := new(All)
	err := client.Request(commonStatsEndpoint, commonResponse)
	if err != nil {
		logger.Errorf("there was an error creating request for common metrics", err)
	}

	err = setCommonMetricsResponse(i, commonResponse)
	if err != nil {
		logger.Errorf("there was an error setting metrics for common metrics", err)
	}
}

func setCommonMetricsResponse(integration *integration.Integration, resp *All) error {
	entity, err := integration.Entity("commonMetrics", "common")
	if err != nil {
		logger.Errorf("there was an error creating new entity for common metrics: %v", err)
		return err
	}

	metricSet := entity.NewMetricSet("clusterMetricSet")
	if err != nil {
		logger.Errorf("there was an error creating new metric set for commmon metrics: %v", err)
		return err
	}

	return metricSet.MarshalMetrics(resp)
}

func populateIndicesMetrics(i *integration.Integration, client Client) {
	logger.Infof("Collecting indices metrics")
	indicesStats := make([]*IndexStats, 0)
	err := client.Request(indicesStatsEndpoint, &indicesStats)
	if err != nil {
		logger.Errorf("there was an error creating request for indices stats", err)
	}
	setIndicesStatsMetricsResponse(i, indicesStats)
}

func setIndicesStatsMetricsResponse(integration *integration.Integration, resp []*IndexStats) {
	for _, object := range resp {
		entity, err := integration.Entity(*object.UUID, "indices")
		if err != nil {
			logger.Errorf("there was an error creating new entity for indices stats metrics: %v", err)
		}
		metricSet := entity.NewMetricSet("indicesMetricsSet",
			metric.Attribute{Key: "displayName", Value: entity.Metadata.Name},
			metric.Attribute{Key: "entityName", Value: entity.Metadata.Namespace + ":" + entity.Metadata.Name})

		err = metricSet.MarshalMetrics(object)
		if err != nil {
			logger.Errorf("there was an error marshaling new metric set for index %s: %v", *object.UUID, err)
		}
	}
}
