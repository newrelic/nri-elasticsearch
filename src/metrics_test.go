package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/newrelic/infra-integrations-sdk/integration"

	"github.com/stretchr/testify/assert"
)

var testWorkDir string

func getWorkDir(t *testing.T) string {
	var err error
	if testWorkDir == "" {
		testWorkDir, err = os.Getwd()
	}

	if err != nil {
		t.Errorf("Unable to get working directory: %s", err.Error())
	}

	return testWorkDir
}

type testClient struct {
	endpointMapping    map[string]string
	ReturnRequestError bool
}

func (c *testClient) init(filename string, endpoint string, t *testing.T) {
	c.endpointMapping = map[string]string{
		endpoint: filepath.Join(getWorkDir(t), "testdata", filename),
	}
}

func (c *testClient) Request(endpoint string, v interface{}) error {
	if c.ReturnRequestError {
		return errors.New("error")
	}

	jsonPath := c.endpointMapping[endpoint]

	jsonData, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, v)
}

func createNewTestClient() *testClient {
	return new(testClient)
}

func createGoldenFile(i *integration.Integration, sourceFile string) (string, []byte) {
	goldenFile := sourceFile + ".golden"
	actualContents, _ := i.Entities[0].Metrics[0].MarshalJSON()

	if *update {
		ioutil.WriteFile(goldenFile, actualContents, 0644)
	}
	return goldenFile, actualContents
}

func TestPopulateNodesMetrics(t *testing.T) {
	i := getTestingIntegration(t)
	client := createNewTestClient()
	client.init("nodeStatsMetricsResult.json", nodeStatsEndpoint, t)

	populateNodesMetrics(i, client)

	sourceFile := filepath.Join(getWorkDir(t), "testdata", "nodeStatsMetricsResult.json")
	goldenFile, actualContents := createGoldenFile(i, sourceFile)
	expectedContents, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Errorf("Failed to load golden file '%s': %s", goldenFile, err.Error())
		t.FailNow()
	}

	assert.Equal(t, 1, len(i.Entities))
	assert.Equal(t, 1, len(i.Entities[0].Metrics))
	assert.Equal(t, expectedContents, actualContents)
}

func TestPopulateNodesMetrics_Error(t *testing.T) {
	mockClient := createNewTestClient()
	mockClient.ReturnRequestError = true

	i := getTestingIntegration(t)
	err := populateNodesMetrics(i, mockClient)
	assert.Error(t, err, "should be an error")
}

func TestPopulateClusterMetrics(t *testing.T) {
	i := getTestingIntegration(t)
	client := createNewTestClient()
	client.init("clusterStatsMetricsResult.json", clusterEndpoint, t)

	populateClusterMetrics(i, client)

	sourceFile := filepath.Join(getWorkDir(t), "testData", "clusterStatsMetricsResult.json")

	goldenFile, actualContents := createGoldenFile(i, sourceFile)
	expectedContents, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Errorf("Failed to load golden file '%s': %s", goldenFile, err.Error())
		t.FailNow()
	}

	actualLength := len(i.Entities[0].Metrics[0].Metrics)
	expectedLength := 11

	assert.Equal(t, expectedContents, actualContents)
	assert.Equal(t, expectedLength, actualLength)
}

func TestPopulateClusterMetrics_Error(t *testing.T) {
	mockClient := createNewTestClient()
	mockClient.ReturnRequestError = true

	i := getTestingIntegration(t)
	err := populateClusterMetrics(i, mockClient)
	assert.Error(t, err, "should be an error")
}

func TestPopulateCommonMetrics(t *testing.T) {
	i := getTestingIntegration(t)
	client := createNewTestClient()
	client.init("commonMetricsResult.json", commonStatsEndpoint, t)

	populateCommonMetrics(i, client)

	sourceFile := filepath.Join(getWorkDir(t), "testData", "commonMetricsResult.json")
	goldenFile, actualContents := createGoldenFile(i, sourceFile)
	expectedContents, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Errorf("Failed to load golden file '%s': %s", goldenFile, err.Error())
		t.FailNow()
	}

	actualLength := len(i.Entities[0].Metrics[0].Metrics)
	expectedLength := 36

	assert.Equal(t, expectedContents, actualContents)
	assert.Equal(t, expectedLength, actualLength)
}

func TestPopulateCommonMetrics_Error(t *testing.T) {
	mockClient := createNewTestClient()
	mockClient.ReturnRequestError = true

	i := getTestingIntegration(t)
	_, err := populateCommonMetrics(i, mockClient)
	assert.Error(t, err, "should be an error")
}

func TestPopulateIndicesMetrics(t *testing.T) {
	i := getTestingIntegration(t)
	client := createNewTestClient()
	client.init("indicesMetricsResult.json", indicesStatsEndpoint, t)

	commonStruct := new(CommonMetrics)
	commonData, _ := ioutil.ReadFile(filepath.Join(getWorkDir(t), "testdata", "indicesMetricsResult_Common.json"))
	json.Unmarshal(commonData, commonStruct)

	populateIndicesMetrics(i, client, commonStruct)

	sourceFile := filepath.Join(getWorkDir(t), "testData", "indicesMetricsResult.json")
	goldenFile, actualContents := createGoldenFile(i, sourceFile)

	for j := range i.Entities {
		resultStruct := i.Entities[j].Metrics[0].Metrics
		actualLength := len(resultStruct)
		expectedLength := 10
		assert.Equal(t, expectedLength, actualLength)
	}

	expectedContents, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Errorf("Failed to load golden file '%s': %s", goldenFile, err.Error())
		t.FailNow()
	}
	assert.Equal(t, expectedContents, actualContents)
}

func TestPopulateIndicesMetrics_Error(t *testing.T) {
	mockClient := createNewTestClient()
	mockClient.ReturnRequestError = true

	i := getTestingIntegration(t)
	err := populateIndicesMetrics(i, mockClient, new(CommonMetrics))
	assert.Error(t, err, "should be an error")
}
