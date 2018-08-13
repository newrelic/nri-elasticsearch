package main

import (
	"github.com/newrelic/infra-integrations-sdk/integration"
)

func populateMetrics(i *integration.Integration, client *Client) {
	populateNodesMetrics(i, client)
	populateClusterMetrics(i, client)
	populateCommonMetrics(i, client)
	populateIndicesMetrics(i, client)
}

func populateNodesMetrics(i *integration.Integration, client *Client) {
	logger.Infof("Collecting node metrics")
	nodeResponse := new(NodeResponse)
	err := client.Request(nodeMetricDefs.Endpoint, nodeResponse)
	if err != nil {
		logger.Errorf("there was an error creating request for node metrics", err)
	}
	// curt wrote a thing to gracefully exit, put that here
	setNodesMetricsResponse(i, nodeResponse)
}

func populateClusterMetrics(i *integration.Integration, client *Client) {
	logger.Infof("Collecting cluster metrics.")
	clusterResponse := new(ClusterResponse)
	err := client.Request(clusterEndpoint, clusterResponse)
	if err != nil {
		logger.Errorf("there was an error creating request for cluster metrics", err)
	}
	// curt wrote a thing to gracefully exit, put that here
	err = setClusterMetricsResponse(i, clusterResponse)
	if err != nil {
		logger.Errorf("there was an error setting metrics for cluster metrics", err)
	}
}

func populateCommonMetrics(i *integration.Integration, client *Client) {
	logger.Infof("Collecting common metrics.")
	commonResponse := new(All)
	err := client.Request(commonStatsEndpoint, commonResponse)
	if err != nil {
		logger.Errorf("there was an error creating request for common metrics", err)
	}
	// curt wrote a thing to gracefully exit, put that here
	err = setCommonMetricsResponse(i, commonResponse)
	if err != nil {
		logger.Errorf("there was an error setting metrics for common metrics", err)
	}
}

func populateIndicesMetrics(i *integration.Integration, client *Client) {
	logger.Infof("Collecting indices metrics")
	indicesStatsResponse := new(IndicesStatsResponse)
	err := client.Request(indicesStatsEndpoint, indicesStatsResponse)
	if err != nil {
		logger.Errorf("there was an error creating request for indices stats", err)
	}
	// curt wrote a thing to gracefully exit, put that here
	err = setIndicesStatsMetricsResponse(i, indicesStatsResponse)
	if err != nil {
		logger.Errorf("there was an error setting metrics for indices metrics", err)
	}
}

//TODO make sure this works
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

//TODO make sure this works
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

// TODO make sure this works
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

// TODO make sure this works
func setIndicesStatsMetricsResponse(integration *integration.Integration, resp *IndicesStatsResponse) error {
	entity, err := integration.Entity("indicesStatsMetrics", "indices")
	if err != nil {
		logger.Errorf("there was an error creating new entity for indices stats metrics: %v", err)
		return err
	}

	metricSet := entity.NewMetricSet("indicesMetricsSet")
	if err != nil {
		logger.Errorf("there was an error creating new metric set for indicies stats metrics: %v", err)
		return err
	}

	return metricSet.MarshalMetrics(resp)
}
