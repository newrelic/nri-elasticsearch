package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/newrelic/infra-integrations-sdk/data/attribute"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
)

var (
	errLocalNodeID  = errors.New("could not identify local node ID")
	errMasterNodeID = errors.New("could not get master node ID")
)

// populateMetrics wrapper to call each of the individual populate functions
func populateMetrics(i *integration.Integration, client Client, env string) {
	if args.MasterOnly {
		nodeID, err := getLocalNodeID(client)
		if err != nil {
			log.Error("There was an error gathering the host node ID: %v", err)
			return
		}
		masterID, err := getMasterNodeID(client)
		if err != nil {
			log.Error("There was an error gathering the elected master node ID: %v", err)
			return
		}
		if nodeID != masterID {
			log.Info("The host is not the elected master, cluster metrics will be skipped")
			return
		}
		log.Info("The host is the elected master, cluster metrics will be collected")
	}

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

func getLocalNodeID(client Client) (nodeId string, err error) {
	var nodeResponseObject LocalNodeResponse
	err = client.Request(localNodeInventoryEndpoint, &nodeResponseObject)
	if err != nil {
		return "", err
	}
	return parseLocalNodeID(nodeResponseObject)
}

func parseLocalNodeID(nodeStats LocalNodeResponse) (string, error) {
	nodes := nodeStats.Nodes
	if len(nodes) == 1 {
		for k := range nodes {
			return k, nil
		}
	}
	return "", errLocalNodeID
}

func getMasterNodeID(client Client) (masterId string, err error) {
	var masterIdResponseObject []MasterNodeIdResponse
	err = client.Request(electedMasterNodeEndpoint, &masterIdResponseObject)
	if err != nil {
		return "", err
	}
	if len(masterIdResponseObject) < 1 {
		return "", errMasterNodeID
	}
	return masterIdResponseObject[0].ID, nil
}

func populateNodesMetrics(i *integration.Integration, client Client, clusterName string) error {
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
func setNodesMetricsResponse(integration *integration.Integration, resp *NodeResponse, clusterName string) {
	for _, node := range resp.Nodes {
		node.IP = ipToString(node.RawIP)
		err := setMetricsResponse(integration, *node, *node.Name, "es-node", clusterName)
		if err != nil {
			log.Error("There was an error setting metrics for node metrics on %s: %v", node, err)
		}
	}
}

func ipToString(ip interface{}) string {
	switch v := ip.(type) {
	case nil:
		return ""
	case string:
		return v
	case *string:
		return *v
	case []string:
		return strings.Join(v, ",")
	case []interface{}:
		strArray := make([]string, len(v))
		for i, elem := range v {
			if str, ok := elem.(string); ok {
				strArray[i] = str
			} else {
				return fmt.Sprintf("%#v", ip)
			}
		}
		return strings.Join(strArray, ",")
	default:
		return fmt.Sprintf("%#v", ip)
	}
}

func populateClusterMetrics(i *integration.Integration, client Client, env string) (string, error) {
	log.Info("Collecting cluster metrics.")

	clusterResponse, err := getClusterResponse(client)
	if err != nil {
		return "", err
	}
	err = setMetricsResponse(i, clusterResponse, *clusterResponse.Name, "es-cluster", *clusterResponse.Name)
	if err != nil {
		return "", err
	}
	return *clusterResponse.Name, nil
}

func getClusterResponse(client Client) (*ClusterResponse, error) {
	clusterResponse := new(ClusterResponse)
	err := client.Request(clusterEndpoint, &clusterResponse)
	if err != nil {
		return nil, err
	}

	if clusterResponse.Name == nil {
		return nil, errors.New("cannot set metric response, missing cluster name")
	}
	return clusterResponse, nil
}

func populateCommonMetrics(i *integration.Integration, client Client, clusterName string) (*CommonMetrics, error) {
	log.Info("Collecting common metrics.")
	commonResponse := new(CommonMetrics)
	err := client.Request(commonStatsEndpoint, &commonResponse)
	if err != nil {
		return nil, err
	}

	if args.CollectPrimaries {
		err = setMetricsResponse(i, commonResponse.All, "commonMetrics", "es-common", clusterName)
	}

	return commonResponse, err
}

func populateIndicesMetrics(i *integration.Integration, client Client, commonStats *CommonMetrics, clusterName string) error {
	log.Info("Collecting indices metrics")

	if commonStats == nil {
		return fmt.Errorf("Common Metrics stats cannot be null to compute Index Metrics")
	}

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

func setIndicesStatsMetricsResponse(integration *integration.Integration, indexResponse []*IndexStats, commonResponse *CommonMetrics, clusterName string, indexRegex *regexp.Regexp) {
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

	for _, index := range indicesToCollect {
		if err := setMetricsResponse(integration, index.stats, index.name, "es-index", clusterName); err != nil {
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
func setMetricsResponse(i *integration.Integration, resp interface{}, name string, namespace string, clusterName string) error {

	entity, err := getEntity(i, name, namespace, clusterName)
	if err != nil {
		return err
	}

	msAttributes := []attribute.Attribute{
		{Key: "displayName", Value: entity.Metadata.Name},
		{Key: "entityName", Value: entity.Metadata.Namespace + ":" + entity.Metadata.Name},
	}

	metricSet := entity.NewMetricSet(getSampleName(namespace), msAttributes...)

	return metricSet.MarshalMetrics(resp)
}

// getEntity generates or retrives an entity if exist
func getEntity(i *integration.Integration, name, namespace, clusterName string) (*integration.Entity, error) {
	entityIDAttrs := []integration.IDAttribute{
		{Key: "clusterName", Value: clusterName},
	}
	if args.ClusterEnvironment != "" {
		entityIDAttrs = append(entityIDAttrs, integration.IDAttribute{Key: "env", Value: args.ClusterEnvironment})
	}

	entity, err := i.EntityReportedVia(fmt.Sprintf("%s:%d", args.Hostname, args.Port), name, namespace, entityIDAttrs...)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func getSampleName(entityType string) string {
	return fmt.Sprintf("Elasticsearch%sSample", strings.Title(strings.TrimPrefix(entityType, "es-")))
}
