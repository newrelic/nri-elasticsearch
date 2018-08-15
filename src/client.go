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
<<<<<<< HEAD
	BaseURL string
	client  *http.Client
}

type Client interface {
	Request(string, interface{}) error
=======
	baseURL string
	client  *http.Client
}

// Client represents an object that can make a request for a map
type Client interface {
	Request(string) (objx.Map, error)
>>>>>>> master
}

// NewClient creates a new Elasticsearch http client.
// httpClient passed in variable should be nil for non-test usage. It is
// available to enable easier mocking of http calls during tests.
func NewClient(httpClient *http.Client) (*HTTPClient, error) {
	if httpClient == nil {
		var err error
		httpClient, err = nrHttp.New(args.CABundleFile, args.CABundleDir, time.Duration(args.Timeout)*time.Second)
		if err != nil {
			return nil, err
		}
	}

	return &HTTPClient{
		client: httpClient,
		baseURL: func() string {
			if args.UseSSL {
				return fmt.Sprintf("https://%s:%d", args.Hostname, args.Port)
			}

			return fmt.Sprintf("http://%s:%d", args.Hostname, args.Port)
		}(),
	}, nil
}

// Request takes an endpoint, makes a GET request to that endpoint,
// and parses the response JSON into a map, which it returns.
<<<<<<< HEAD
func (c *HTTPClient) Request(endpoint string, v interface{}) error {
	response, err := c.client.Get(c.BaseURL + endpoint)
=======
func (c *HTTPClient) Request(endpoint string) (objx.Map, error) {
	response, err := c.client.Get(c.baseURL + endpoint)
>>>>>>> master
	if err != nil {
		return err
	}
	defer checkErr(response.Body.Close)

	err = json.NewDecoder(response.Body).Decode(&v)
	if err != nil {
		return err
	}

	return nil
}
