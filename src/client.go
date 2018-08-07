package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	nrHttp "github.com/newrelic/infra-integrations-sdk/http"
	"github.com/stretchr/objx"
)

const (
	nodeStatsEndpoint          = "/_nodes/stats"
	localNodeInventoryEndpoint = "/_nodes/_local"
	commonStatsEndpoint        = "/_stats"
	clusterEndpoint            = "/_cluster/health"
)

// HTTPClient represents a single connection to an Elasticsearch host
type HTTPClient struct {
	baseURL string
	client  *http.Client
}

// Client represents an object that can make a request for a map
type Client interface {
	Request(string) (objx.Map, error)
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
func (c *HTTPClient) Request(endpoint string) (objx.Map, error) {
	response, err := c.client.Get(c.baseURL + endpoint)
	if err != nil {
		return nil, err
	}
	defer checkErr(response.Body.Close)

	var resultMap map[string]interface{}

	err = json.NewDecoder(response.Body).Decode(&resultMap)
	if err != nil {
		return nil, err
	}

	return objx.New(resultMap), nil
}
