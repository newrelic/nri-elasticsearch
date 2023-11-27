//go:generate goversioninfo
package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
)

type argumentList struct {
	sdkArgs.DefaultArgumentList
	Hostname               string `default:"localhost" help:"Hostname or IP where Elasticsearch Node is running."`
	LocalHostname          string `default:"localhost" help:"Hostname or IP of the Elasticsearch node from which to collect inventory."`
	ClusterEnvironment     string `default:"" help:"A way to further specify which cluster we are gathering data for, example: 'staging'"`
	Port                   int    `default:"9200" help:"Port on which Elasticsearch Node is listening."`
	Username               string `default:"" help:"Username for accessing Elasticsearch Node"`
	Password               string `default:"" help:"Password for the given user."`
	UseSSL                 bool   `default:"false" help:"Signals whether to use SSL or not. Certificate bundle must be supplied"`
	CABundleFile           string `default:"" help:"Alternative Certificate Authority bundle file"`
	CABundleDir            string `default:"" help:"Alternative Certificate Authority bundle directory"`
	SSLAlternativeHostname string `default:"" help:"Alternative hostname to accept certificates from during SSL negotiation"`
	Timeout                int    `default:"30" help:"Timeout for an API call"`
	ConfigPath             string `default:"/etc/elasticsearch/elasticsearch.yml" help:"Path to the ElasticSearch configuration .yml file."`
	CollectIndices         bool   `default:"true" help:"Signals whether to collect indices metrics or not"`
	CollectPrimaries       bool   `default:"true" help:"Signals whether to collect primaries metrics or not"`
	IndicesRegex           string `default:"" help:"A regex pattern that matches the index names to collect. Collects all if unspecified"`
	ShowVersion            bool   `default:"false" help:"Print build information and exit"`
	MasterOnly             bool   `default:"false" help:"Collect cluster metrics on the elected master only"`
	LocalOnly              bool   `default:"false" help:"Collect node metrics on the local node only"`
	TLSInsecureSkipVerify  bool   `default:"false" help:"Enabled TLS insecure skip verify"`
}

const (
	integrationName = "com.newrelic.elasticsearch"
)

var (
	args               argumentList
	integrationVersion = "0.0.0"
	gitCommit          = ""
	buildDate          = ""
)

func main() {
	// Create Integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	logErrorAndExit(err)

	if args.ShowVersion {
		fmt.Printf(
			"New Relic %s integration Version: %s, Platform: %s, GoVersion: %s, GitCommit: %s, BuildDate: %s\n",
			strings.Title(strings.Replace(integrationName, "com.newrelic.", "", 1)),
			integrationVersion,
			fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
			runtime.Version(),
			gitCommit,
			buildDate)
		os.Exit(0)
	}
	if args.MasterOnly && args.LocalOnly {
		logErrorAndExit(fmt.Errorf("Select argument -master_only or -local_only, not both"))
	}

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
