//go:generate goversioninfo
package main

import (
	"os"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
)

type argumentList struct {
	sdkArgs.DefaultArgumentList
	Hostname           string `default:"localhost" help:"Hostname or IP where Elasticsearch Node is running."`
	LocalHostname      string `default:"localhost" help:"Hostname or IP of the Elasticsearch node from which to collect inventory."`
	ClusterEnvironment string `default:"" help:"A way to further specify which cluster we are gathering data for, example: 'staging'"`
	Port               int    `default:"9200" help:"Port on which Elasticsearch Node is listening."`
	Username           string `default:"" help:"Username for accessing Elasticsearch Node"`
	Password           string `default:"" help:"Password for the given user."`
	UseSSL             bool   `default:"false" help:"Signals whether to use SSL or not. Certificate bundle must be supplied"`
	CABundleFile       string `default:"" help:"Alternative Certificate Authority bundle file"`
	CABundleDir        string `default:"" help:"Alternative Certificate Authority bundle directory"`
	Timeout            int    `default:"30" help:"Timeout for an API call"`
	ConfigPath         string `default:"/etc/elasticsearch/elasticsearch.yml" help:"Path to the ElasticSearch configuration .yml file."`
	CollectIndices     bool   `default:"true" help:"Signals whether to collect indices metrics or not"`
	CollectPrimaries   bool   `default:"true" help:"Signals whether to collect primaries metrics or not"`
	IndicesRegex       string `default:"" help:"JSON array of index names from which to collect metrics"`
}

const (
	integrationName    = "com.newrelic.elasticsearch"
	integrationVersion = "4.2.0"
)

var (
	args argumentList
)

func main() {
	// Create Integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	logErrorAndExit(err)

	// Create a client for metrics
	metricsClient, err := NewClient(args.Hostname)
	logErrorAndExit(err)

	if args.All() || args.Metrics {
		populateMetrics(i, metricsClient, args.ClusterEnvironment)
	}

	// Create a client for inventory. Inventory needs to make REST calls against
	// localhost to get information relative to this node only.
	inventoryClient, err := NewClient(args.LocalHostname)
	logErrorAndExit(err)

	if args.All() || args.Inventory {
		populateInventory(i, inventoryClient)
	}

	logErrorAndExit(i.Publish())
}

// checkErr logs an error if it exists
func checkErr(f func() error) {
	if err := f(); err != nil {
		log.Error("%v", err)
	}
}

// logErrorAndExit logs an error if it exits and
// exits with a status code of 1
func logErrorAndExit(err error) {
	if err != nil {
		log.Error("Encountered fatal error: %v", err)
		os.Exit(1)
	}
}
