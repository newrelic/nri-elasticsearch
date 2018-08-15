package main

import (
	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
)

// populateMetrics wrapper to call each of the individual populate functions
func populateMetrics(i *integration.Integration, client Client) {
	err := populateNodesMetrics(i, client)
	if err != nil {
		log.Error("there was an error populating metrics for nodes", err)
	}
	err = populateClusterMetrics(i, client)
	if err != nil {
		log.Error("there was an error populating metrics for clusters", err)
	}
	err = populateCommonMetrics(i, client)
	if err != nil {
		log.Error("there was an error populating metrics for common metrics", err)
	}
	err = populateIndicesMetrics(i, client)
	if err != nil {
		log.Error("there was an error populating metrics for indices", err)
	}
}

func populateNodesMetrics(i *integration.Integration, client Client) error {
	logger.Infof("Collecting node metrics")
	nodeResponse := new(NodeResponse)
	err := client.Request(nodeStatsEndpoint, nodeResponse)
	if err != nil {
		log.Error("there was an error creating request for node metrics", err)
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
			log.Error("there was an error setting metrics for node metrics on: %s", node, err)
		}
	}
}

func populateClusterMetrics(i *integration.Integration, client Client) error {
	logger.Infof("Collecting cluster metrics.")
	clusterResponse := new(ClusterResponse)
	err := client.Request(clusterEndpoint, clusterResponse)
	if err != nil {
		log.Error("there was an error creating request for cluster metrics", err)
		return err
	}

	err = setMetricsResponse(i, clusterResponse, *clusterResponse.Name, "cluster")
	if err != nil {
		log.Error("there was an error setting metrics for cluster metrics", err)
	}
	return err
}

func populateCommonMetrics(i *integration.Integration, client Client) error {
	logger.Infof("Collecting common metrics.")
	commonResponse := new(CommonMetrics)
	err := client.Request(commonStatsEndpoint, commonResponse)
	if err != nil {
		log.Error("there was an error creating request for common metrics", err)
		return err
	}

	err = setMetricsResponse(i, commonResponse.All, "commonMetrics", "common")
	if err != nil {
		log.Error("there was an error setting metrics for common metrics", err)
	}
	return nil
}

func populateIndicesMetrics(i *integration.Integration, client Client) error {
	logger.Infof("Collecting indices metrics")
	indicesStats := make([]*IndexStats, 0)
	err := client.Request(indicesStatsEndpoint, &indicesStats)
	if err != nil {
		log.Error("there was an error creating request for indices stats", err)
		return err
	}
	setIndicesStatsMetricsResponse(i, indicesStats)
	return nil
}

func setIndicesStatsMetricsResponse(integration *integration.Integration, resp []*IndexStats) {
	for _, object := range resp {
		err := setMetricsResponse(integration, object, *object.UUID, "indices")
		if err != nil {
			log.Error("there was an error setting metrics for indices metrics", err)
		}
	}
}

// setMetricsResponse creates an entity and a metric set for the
// type of response and calls MarshalMetrics using that response
func setMetricsResponse(integration *integration.Integration, resp interface{}, name string, namespace string) error {
	entity, err := integration.Entity(name, namespace)
	if err != nil {
		log.Error("there was an error creating new entity for %s: %v", namespace, err)
		return err
	}

	metricSet := entity.NewMetricSet(namespace+"MetricSet",
		metric.Attribute{Key: "displayName", Value: entity.Metadata.Name},
		metric.Attribute{Key: "entityName", Value: entity.Metadata.Namespace + ":" + entity.Metadata.Name},
	)

	return metricSet.MarshalMetrics(resp)
}
