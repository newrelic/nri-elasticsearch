package main

import (
	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/data/event"
	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
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
	args argumentList
)

func main() {
	// Create Integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	panicOnErr(err)

	// Create Entity, entities name must be unique
	e1, err := i.Entity("instance-1", "custom")
	panicOnErr(err)

	// Add Event
	if args.All() || args.Events {
		err = e1.AddEvent(event.New("restart", "status"))
		panicOnErr(err)
	}

	// Add Inventory item
	if args.All() || args.Inventory {
		err = e1.SetInventoryItem("instance", "version", "3.0.1")
		panicOnErr(err)
	}

	// Add Metric
	if args.All() || args.Metrics {
		m1, err := e1.NewMetricSet("CustomSample")
		panicOnErr(err)
		err = m1.SetMetric("some-data", 1000, metric.GAUGE)
		panicOnErr(err)
	}

	// Create another Entity
	e2, err := i.Entity("instance-2", "custom")
	panicOnErr(err)

	if args.All() || args.Inventory {
		err = e2.SetInventoryItem("instance", "version", "3.0.4")
		panicOnErr(err)
	}

	if args.All() || args.Metrics {
		m2, err := e2.NewMetricSet("CustomSample")
		panicOnErr(err)
		err = m2.SetMetric("some-data", 2000, metric.GAUGE)
		panicOnErr(err)
	}

	panicOnErr(i.Publish())
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
