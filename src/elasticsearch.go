package main

import (
	"io/ioutil"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/stretchr/objx"
)

type argumentList struct {
	sdkArgs.DefaultArgumentList
	Hostname     string `default:"localhost" help:"Hostname or IP where Elasticsearch Node is running."`
	Port         int    `default:"9200" help:"Port on which Elasticsearch Node is listening."`
	Username     string `default:"" help:"Username for accessing Elasticsearch Node"`
	Password     string `default:"" help:"Password for the given user."`
	UseSSL       bool   `default:"false" help:"Signals whether to use SSL or not. Certificate bundle must be supplied"`
	CABundleFile string `default:"" help:"Alternative Certificate Authority bundle file"`
	CABundleDir  string `default:"" help:"Alternative Certificate Authority bundle directory"`
	Timeout      int    `default:"30" help:"Timeout for an API call"`
	ConfigPath   string `default:"/etc/elasticsearch/elasticsearch.yml" help:"Path to the ElasticSearch configuration .yml file."`
}

const (
	integrationName    = "com.newrelic.elasticsearch"
	integrationVersion = "0.1.0"
)

var (
	args   argumentList
	logger log.Logger
)

func main() {
	// Create Integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	panicOnErr(err)
	logger = i.Logger()

	client, err := NewClient(nil)
	panicOnErr(err)

	// Add Inventory item
	if args.All() || args.Inventory {
		populateInventory(i)
	}

	// TODO refactor this
	if args.All() || args.Metrics {
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

	panicOnErr(i.Publish())
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

func checkErr(f func() error) {
	if err := f(); err != nil {
		logger.Errorf("%v", err)
	}
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
