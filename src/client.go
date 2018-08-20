package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	nrHttp "github.com/newrelic/infra-integrations-sdk/http"
)

const (
	nodeStatsEndpoint          = "/_nodes/stats"
	localNodeInventoryEndpoint = "/_nodes/_local"
	commonStatsEndpoint        = "/_stats"
	clusterEndpoint            = "/_cluster/health"
	indicesStatsEndpoint       = "/_cat/indices?format=json"
)

// HTTPClient represents a single connection to an Elasticsearch host
type HTTPClient struct {
	baseURL string
	client  *http.Client
}

// Client interface that assists in mocking for tests
type Client interface {
	Request(string, interface{}) error
}

// NewClient creates a new Elasticsearch http client.
// The hostnameOverride parameter specifies a hostname that the client should connect to.
// Passing in an empty string causes the client to use the hostname specified in the command-line args. (default behavior)
func NewClient(hostnameOverride string) (*HTTPClient, error) {
	httpClient, err := nrHttp.New(args.CABundleFile, args.CABundleDir, time.Duration(args.Timeout)*time.Second)
	if err != nil {
		return nil, err
	}

	return &HTTPClient{
		client: httpClient,
		baseURL: func() string {
			protocol := "http"
			if args.UseSSL {
				protocol = "https"
			}

			hostname := args.Hostname
			if hostnameOverride != "" {
				hostname = hostnameOverride
			}

			return fmt.Sprintf("%s://%s:%d", protocol, hostname, args.Port)
		}(),
	}, nil
}

// Request takes an endpoint, makes a GET request to that endpoint,
// and parses the response JSON into a map, which it returns.
func (c *HTTPClient) Request(endpoint string, v interface{}) error {
	response, err := c.client.Get(c.baseURL + endpoint)
	if err != nil {
		return err
	}
	defer checkErr(response.Body.Close)

	err = json.NewDecoder(response.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}
