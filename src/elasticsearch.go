package main

import (
	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
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

	// Create Entity, entities name must be unique
	e1, err := i.Entity("instance-1", "custom")
	panicOnErr(err)

	// Add Inventory item
	if args.All() || args.Inventory {
		err = populateInventory(e1)
		panicOnErr(err)
	}

	panicOnErr(i.Publish())
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
