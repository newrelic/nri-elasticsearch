package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	nrHttp "github.com/newrelic/infra-integrations-sdk/http"
)

const (
	nodeIngestEndpoint         = "/_nodes/ingest"
	nodeProcessEndpoint        = "/_nodes/process"
	nodePluginsEndpoint        = "/_nodes/plugins"
	nodeStatsEndpoint          = "/_nodes/stats"
	localNodeInventoryEndpoint = "/_nodes/_local"
	commonStatsEndpoint        = "/_stats"
	clusterEndpoint            = "/_cluster/health"
	indicesStatsEndpoint       = "/_cat/indices?format=json"
)

// Client represents a single connection to an Elasticsearch host
type Client struct {
	BaseURL string
	client  *http.Client
}

// NewClient creates a new Elasticsearch http client.
// httpClient passed in variable should be nil for non-test usage. It is
// available to enable easier mocking of http calls during tests.
func NewClient(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		var err error
		httpClient, err = nrHttp.New(args.CABundleFile, args.CABundleDir, time.Duration(args.Timeout)*time.Second)
		if err != nil {
			return nil, err
		}
	}

	return &Client{
		client: httpClient,
		BaseURL: func() string {
			if args.UseSSL {
				return fmt.Sprintf("https://%s:%d", args.Hostname, args.Port)
			}

			return fmt.Sprintf("http://%s:%d", args.Hostname, args.Port)
		}(),
	}, nil
}

// Request takes an endpoint, makes a GET request to that endpoint,
// and parses the response JSON into a map, which it returns.
func (c *Client) Request(endpoint string, v interface{}) error {
	response, err := c.client.Get(c.BaseURL + endpoint)
	if err != nil {
		return err
	}
	defer checkErr(response.Body.Close)

	var resultMap map[string]interface{}

	err = json.NewDecoder(response.Body).Decode(&v)
	if err != nil {
		return err
	}

	return nil
}

// Do sends an API request and returns and API response. The API response
// is JSON decoded and stored in the value pointed to by v, or returned
// as an error if an API error occurred.
func (c Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	// err = CheckResponse(resp)
	// if err != nil {
	// 	return resp, err
	// }

	if v != nil {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp, err
		}

		err = json.Unmarshal(data, v)
	}

	return resp, err
}

// func CheckResponse(resp *http.Response) error {
// 	if resp.StatusCode < 400 {
// 		return nil
// 	}

// 	errorResponse := &ErrorResponse{Response: resp}
// 	data, err := ioutil.ReadAll(resp.Body)
// 	if err == nil && data != nil {
// 		json.Unmarshal(data, errorResponse)
// 	}

// 	return errorResponse
// }

// type ErrorResponse struct {
// 	Response *http.Response

// 	Message string `json:"error"`
// }s
