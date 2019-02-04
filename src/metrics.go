package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
)

const indexLimit = 500

// populateMetrics wrapper to call each of the individual populate functions
func populateMetrics(i *integration.Integration, client Client, env string) {
	clusterName, err := populateClusterMetrics(i, client, env)
	if err != nil {
		log.Error("There was an error populating metrics for clusters: %v", err)
	}

	err = populateNodesMetrics(i, client, clusterName)
	if err != nil {
		log.Error("There was an error populating metrics for nodes: %v", err)
	}

	// we want to use the response from common to populate some index-specific stats.
	commonResponse, err := populateCommonMetrics(i, client, clusterName)
	if err != nil {
		log.Error("There was an error populating metrics for common metrics: %v", err)
	}

	if args.CollectIndices {
		err = populateIndicesMetrics(i, client, commonResponse, clusterName)
		if err != nil {
			log.Error("There was an error populating metrics for indices: %v", err)
		}
	}
}

func populateNodesMetrics(i *integration.Integration, client Client, clusterName *string) error {
	log.Info("Collecting node metrics")
	nodeResponse := new(NodeResponse)
	err := client.Request(nodeStatsEndpoint, &nodeResponse)
	if err != nil {
		return err
	}

	setNodesMetricsResponse(i, nodeResponse, clusterName)
	return nil
}

// setNodesMetricsResponse calls setMetricsResponse for each node in the response
func setNodesMetricsResponse(integration *integration.Integration, resp *NodeResponse, clusterName *string) {
	for node := range resp.Nodes {
		err := setMetricsResponse(integration, resp.Nodes[node], *resp.Nodes[node].Name, "node", clusterName)
		if err != nil {
			log.Error("There was an error setting metrics for node metrics on %s: %v", node, err)
		}
	}
}

func populateClusterMetrics(i *integration.Integration, client Client, env string) (*string, error) {
	log.Info("Collecting cluster metrics.")
	clusterResponse := new(ClusterResponse)
	err := client.Request(clusterEndpoint, &clusterResponse)
	if err != nil {
		return nil, err
	}

	if clusterResponse.Name == nil {
		return nil, errors.New("cannot set metric response, missing cluster name")
	}

	if env != "" {
		*clusterResponse.Name = *clusterResponse.Name + ":" + env
	}

	return clusterResponse.Name, setMetricsResponse(i, clusterResponse, *clusterResponse.Name, "cluster", nil)
}

func populateCommonMetrics(i *integration.Integration, client Client, clusterName *string) (*CommonMetrics, error) {
	log.Info("Collecting common metrics.")
	commonResponse := new(CommonMetrics)
	err := client.Request(commonStatsEndpoint, &commonResponse)
	if err != nil {
		return nil, err
	}

	if args.CollectPrimaries {
		err = setMetricsResponse(i, commonResponse.All, "commonMetrics", "common", clusterName)
	}

	return commonResponse, err
}

func populateIndicesMetrics(i *integration.Integration, client Client, commonStats *CommonMetrics, clusterName *string) error {
	log.Info("Collecting indices metrics")
	indicesStats := make([]*IndexStats, 0)
	err := client.Request(indicesStatsEndpoint, &indicesStats)
	if err != nil {
		return err
	}

	indexRegex, err := buildRegex()
	if err != nil {
		return err
	}

	setIndicesStatsMetricsResponse(i, indicesStats, commonStats, clusterName, indexRegex)
	return nil
}

func buildRegex() (indexRegex *regexp.Regexp, err error) {
	if args.IndicesRegex != "" {
		indexRegex, err = regexp.Compile(args.IndicesRegex)
		if err != nil {
			return indexRegex, err
		}
	}
	return indexRegex, nil
}

func setIndicesStatsMetricsResponse(integration *integration.Integration, indexResponse []*IndexStats, commonResponse *CommonMetrics, clusterName *string, indexRegex *regexp.Regexp) {
	type indexStatsObject struct {
		name  string
		stats *IndexStats
	}
	indicesToCollect := make([]indexStatsObject, 0, len(indexResponse))

	for _, object := range indexResponse {
		if object.Name == nil {
			log.Error("Can't set metric response, missing index name")
			continue
		}

		if indexRegex != nil && !indexRegex.MatchString(*object.Name) {
			log.Debug("Can't set metric response, index does not match regex")
			continue
		}

		// cross reference with common stats
		index, err := getIndexFromCommon(*object.Name, commonResponse.Indices)
		if err != nil {
			log.Error("Couldn't match index name in common index stats response: %v", err)
			continue
		}

		// populate fields from stats
		object.PrimaryStoreSize = index.Primaries.Store.Size
		object.StoreSize = index.Totals.Store.Size

		indicesToCollect = append(indicesToCollect, indexStatsObject{
			*object.Name,
			object,
		})
	}

	// enforce index limit
	if length := len(indicesToCollect); length > indexLimit {
		log.Error("Could not collect index metrics: attempting to collect %d indices which exceeds the maximum of %d. Use the index regex configuration parameter to limit collection size.", length, indexLimit)
		return
	}

	for _, index := range indicesToCollect {
		if err := setMetricsResponse(integration, index.stats, index.name, "index", clusterName); err != nil {
			log.Error("There was an error setting metrics for indices metrics: %v", err)
		}
	}
}

func getIndexFromCommon(indexName string, indexList map[string]*Index) (*Index, error) {
	indexStats, ok := indexList[indexName]
	if !ok {
		return nil, fmt.Errorf("index '%s' not contained in list", indexName)
	}
	return indexStats, nil
}

// setMetricsResponse creates an entity and a metric set for the
// type of response and calls MarshalMetrics using that response
func setMetricsResponse(integration *integration.Integration, resp interface{}, name string, namespace string, clusterName *string) error {
	entity, err := integration.Entity(name, namespace)
	if err != nil {
		return err
	}

	idAttributes := []metric.Attribute{
		{Key: "displayName", Value: entity.Metadata.Name},
		{Key: "entityName", Value: entity.Metadata.Namespace + ":" + entity.Metadata.Name},
	}

	// If clusterName is non empty apply it
	if clusterName != nil {
		idAttributes = append(idAttributes, metric.Attribute{Key: "clusterName", Value: *clusterName})
	}

	metricSet := entity.NewMetricSet(getSampleName(namespace), idAttributes...)

	return metricSet.MarshalMetrics(resp)
}

func getSampleName(entityType string) string {
	return fmt.Sprintf("Elasticsearch%sSample", strings.Title(entityType))
}
