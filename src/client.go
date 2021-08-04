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
	nodeStatsEndpoint          = "/_nodes/stats"
	localNodeInventoryEndpoint = "/_nodes/_local"
	commonStatsEndpoint        = "/_stats"
	clusterEndpoint            = "/_cluster/health"
	indicesStatsEndpoint       = "/_cat/indices?format=json"
	electedMasterNodeEndpoint  = "/_cat/master?h=id&format=json"
)

// HTTPClient represents a single connection to an Elasticsearch host
type HTTPClient struct {
	baseURL  string
	useAuth  bool
	username string
	password string
	client   *http.Client
}

// Client interface that assists in mocking for tests
type Client interface {
	Request(string, interface{}) error
}

type errorResponse struct {
	Error *errorBody `json:"error"`
}

type errorBody struct {
	Type   *string `json:"type"`
	Reason *string `json:"reason"`
}

// NewClient creates a new Elasticsearch http client.
// The hostname parameter specifies the hostname that the client should connect to.
// Passing in an empty string causes the client to use the hostname specified in the command-line args. (default behavior)
func NewClient(hostname string) (*HTTPClient, error) {
	clientOptions := getClientOptionsFromArgs()
	httpClient, err := nrHttp.New(clientOptions...)
	if err != nil {
		return nil, err
	}

	return &HTTPClient{
		client:   httpClient,
		useAuth:  args.Username != "" || args.Password != "",
		username: args.Username,
		password: args.Password,
		baseURL: func() string {
			protocol := "http"
			if args.UseSSL {
				protocol = "https"
			}
			return fmt.Sprintf("%s://%s:%d", protocol, hostname, args.Port)
		}(),
	}, nil
}

func getClientOptionsFromArgs() []nrHttp.ClientOption {
	var clientOptions []nrHttp.ClientOption
	if args.CABundleDir != "" {
		clientOptions = append(clientOptions, nrHttp.WithCABundleDir(args.CABundleDir))
	}
	if args.CABundleFile != "" {
		clientOptions = append(clientOptions, nrHttp.WithCABundleFile(args.CABundleFile))
	}
	if args.SSLAlternativeHostname != "" {
		clientOptions = append(clientOptions, nrHttp.WithAcceptInvalidHostname(args.SSLAlternativeHostname))
	}
	if args.TLSInsecureSkipVerify {
		clientOptions = append(clientOptions, nrHttp.WithTLSInsecureSkipVerify())
	}
	clientOptions = append(clientOptions, nrHttp.WithTimeout(time.Duration(args.Timeout)*time.Second))
	return clientOptions
}

// Request takes an endpoint, makes a GET request to that endpoint,
// and parses the response JSON into a map, which it returns.
func (c *HTTPClient) Request(endpoint string, v interface{}) error {
	request, err := http.NewRequest("GET", c.baseURL+endpoint, nil)
	if err != nil {
		return err
	}
	if c.useAuth {
		request.SetBasicAuth(c.username, c.password)
	}

	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer checkErr(response.Body.Close)

	err = c.checkStatusCode(response)
	if err != nil {
		return err
	}

	err = json.NewDecoder(response.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func (c *HTTPClient) checkStatusCode(response *http.Response) error {
	if response.StatusCode == 200 {
		return nil
	}

	// try parsing error in body, otherwise return generic error
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("status code %v - could not parse body from response: %v", response.StatusCode, err)
	}

	var errResponse errorResponse
	err = json.Unmarshal(responseBytes, &errResponse)
	if err != nil {
		return fmt.Errorf("status code %v - could not parse error information from response: %v", response.StatusCode, err)
	}

	return fmt.Errorf("status code %v - received error of type '%s' from Elasticsearch: %s", response.StatusCode, *errResponse.Error.Type, *errResponse.Error.Reason)
}
