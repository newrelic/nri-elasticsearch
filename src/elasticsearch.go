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

	// // Create Entity, entities name must be unique
	// e1, err := i.Entity("instance-1", "custom")
	// panicOnErr(err)

	// // Add Event
	// if args.All() || args.Events {
	// 	err = e1.AddEvent(event.New("restart", "status"))
	// 	panicOnErr(err)
	// }

	// // Add Inventory item
	// if args.All() || args.Inventory {
	// 	err = e1.SetInventoryItem("instance", "version", "3.0.1")
	// 	panicOnErr(err)
	// }

	// // Add Metric
	// if args.All() || args.Metrics {
	// 	m1, err := e1.NewMetricSet("CustomSample")
	// 	panicOnErr(err)
	// 	err = m1.SetMetric("some-data", 1000, metric.GAUGE)
	// 	panicOnErr(err)
	// }

	// // Create another Entity
	// e2, err := i.Entity("instance-2", "custom")
	// panicOnErr(err)

	// if args.All() || args.Inventory {
	// 	err = e2.SetInventoryItem("instance", "version", "3.0.4")
	// 	panicOnErr(err)
	// }

	// if args.All() || args.Metrics {
	// 	m2, err := e2.NewMetricSet("CustomSample")
	// 	panicOnErr(err)
	// 	err = m2.SetMetric("some-data", 2000, metric.GAUGE)
	// 	panicOnErr(err)
	// }

	panicOnErr(i.Publish())
	client, err := NewClient(nil)
	panicOnErr(err)
	client.BaseURL = args.Hostname
	endpoint := nodeMetricDefs.Endpoint
	stringResponse, err := getDataFromEndpoint(client, endpoint)
	panicOnErr(err)
	responseObject, err := objx.FromJSON(stringResponse)
	logger.Infof("Collecting metrics.")
	collectNodesMetrics(i, &responseObject)
}

func getDataFromEndpoint(client *Client, endpoint string) (string, error) {
	url := client.BaseURL + endpoint
	response, err := client.client.Get(url)
	if err != nil {
		logger.Errorf("there was an error when getting response from endpoint %v: %v", url, err)
		return "", err
	}
	jsonData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Errorf("there was an error when reading the response body: %v", err)
		return "", err
	}
	err = response.Body.Close()
	//ask hullah about the workaround this
	if err != nil {
		logger.Errorf("there was an error when closing the response body")
	}
	jsonString := string(jsonData)
	return jsonString, err
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
