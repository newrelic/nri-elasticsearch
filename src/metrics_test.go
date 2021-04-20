package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/newrelic/infra-integrations-sdk/integration"

	"github.com/stretchr/testify/assert"
)

var testClusterName = "goTestCluster"

type testClient struct {
	endpointMapping    map[string]string
	ReturnRequestError bool
}

func (c *testClient) init(filename string, endpoint string, t *testing.T) {
	c.endpointMapping = map[string]string{
		endpoint: filepath.Join("testdata", filename),
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
		_ = ioutil.WriteFile(goldenFile, actualContents, 0644)
	}
	return goldenFile, actualContents
}

func TestGetLocalNodeID(t *testing.T) {
	fakeClient := mockClient{}
	mockedReturnVal := filepath.Join("testdata", "good-nodes-local.json")
	fakeClient.On("Request", localNodeInventoryEndpoint).Return(mockedReturnVal, nil).Once()

	nodeID, err := getLocalNodeID(&fakeClient)
	assert.NoError(t, err)
	assert.Equal(t, "z9ZPp87vT92qG1cRVRIcMQ", nodeID)
}

func TestGetLocalNodeID_Error(t *testing.T) {
	fakeClient := mockClient{}
	mockedReturnVal := filepath.Join("testdata", "bad-nodes-local.json")
	fakeClient.On("Request", localNodeInventoryEndpoint).Return(mockedReturnVal, nil).Once()

	_, err := getLocalNodeID(&fakeClient)
	assert.Error(t, err)
	assert.Equal(t, errLocalNodeID, err)
}

func TestGetMasterNodeID(t *testing.T) {
	fakeClient := mockClient{}
	mockedReturnVal := filepath.Join("testdata", "good-master.json")
	fakeClient.On("Request", electedMasterNodeEndpoint).Return(mockedReturnVal, nil).Once()

	nodeID, err := getMasterNodeID(&fakeClient)
	assert.NoError(t, err)
	assert.Equal(t, "z9ZPp87vT92qG1cRVRIcMQ", nodeID)
}

func TestGetMasterNodeID_Error(t *testing.T) {
	fakeClient := mockClient{}
	fakeClient.On("Request", electedMasterNodeEndpoint).Return("", nil).Once()

	_, err := getMasterNodeID(&fakeClient)
	assert.Error(t, err)
	assert.Equal(t, errMasterNodeID, err)
}

func TestPopulateNodesMetrics(t *testing.T) {
	i := getTestingIntegration(t)
	client := createNewTestClient()
	client.init("nodeStatsMetricsResult.json", nodeStatsEndpoint, t)

	err := populateNodesMetrics(i, client, testClusterName)
	assert.NoError(t, err)

	sourceFile := filepath.Join("testdata", "nodeStatsMetricsResult.json")
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
	err := populateNodesMetrics(i, mockClient, testClusterName)
	assert.Error(t, err, "should be an error")
}

func TestPopulateClusterMetrics(t *testing.T) {
	i := getTestingIntegration(t)
	client := createNewTestClient()
	client.init("clusterStatsMetricsResult.json", clusterEndpoint, t)

	name, err := populateClusterMetrics(i, client, "")
	assert.NotEmpty(t, name)
	assert.NoError(t, err)

	sourceFile := filepath.Join("testdata", "clusterStatsMetricsResult.json")

	goldenFile, actualContents := createGoldenFile(i, sourceFile)
	expectedContents, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Errorf("Failed to load golden file '%s': %s", goldenFile, err.Error())
		t.FailNow()
	}

	actualLength := len(i.Entities[0].Metrics[0].Metrics)
	expectedLength := 12

	assert.Equal(t, expectedContents, actualContents)
	assert.Equal(t, expectedLength, actualLength)
}

func TestPopulateClusterMetrics_Error(t *testing.T) {
	mockClient := createNewTestClient()
	mockClient.ReturnRequestError = true

	i := getTestingIntegration(t)
	_, err := populateClusterMetrics(i, mockClient, "")
	assert.Error(t, err, "should be an error")
}

func TestPopulateCommonMetrics(t *testing.T) {
	i := getTestingIntegration(t)
	client := createNewTestClient()
	args.CollectIndices = true
	args.CollectPrimaries = true
	client.init("commonMetricsResult.json", commonStatsEndpoint, t)

	name, err := populateCommonMetrics(i, client, testClusterName)
	assert.NotEmpty(t, name)
	assert.NoError(t, err)

	sourceFile := filepath.Join("testdata", "commonMetricsResult.json")
	goldenFile, actualContents := createGoldenFile(i, sourceFile)
	expectedContents, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Errorf("Failed to load golden file '%s': %s", goldenFile, err.Error())
		t.FailNow()
	}

	actualLength := len(i.Entities[0].Metrics[0].Metrics)
	expectedLength := 37

	assert.Equal(t, expectedContents, actualContents)
	assert.Equal(t, expectedLength, actualLength)
}

func TestPopulateCommonMetrics_Error(t *testing.T) {
	mockClient := createNewTestClient()
	mockClient.ReturnRequestError = true

	i := getTestingIntegration(t)
	_, err := populateCommonMetrics(i, mockClient, testClusterName)
	assert.Error(t, err, "should be an error")
}

func TestPopulateIndicesMetrics(t *testing.T) {
	i := getTestingIntegration(t)
	client := createNewTestClient()
	client.init("indicesMetricsResult.json", indicesStatsEndpoint, t)

	commonStruct := new(CommonMetrics)
	commonData, _ := ioutil.ReadFile(filepath.Join("testdata", "indicesMetricsResult_Common.json"))
	err := json.Unmarshal(commonData, commonStruct)
	assert.NoError(t, err)

	err = populateIndicesMetrics(i, client, commonStruct, testClusterName)
	assert.NoError(t, err)

	sourceFile := filepath.Join("testdata", "indicesMetricsResult.json")
	goldenFile, actualContents := createGoldenFile(i, sourceFile)

	for j := range i.Entities {
		resultStruct := i.Entities[j].Metrics[0].Metrics
		actualLength := len(resultStruct)
		expectedLength := 11
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
	err := populateIndicesMetrics(i, mockClient, new(CommonMetrics), testClusterName)
	assert.Error(t, err, "should be an error")
}

func TestPopulateIndicesMetrics_NilStats(t *testing.T) {
	// Given a nil commonStats metrics and not nil indicesMetrics
	client := createNewTestClient()
	client.init("indicesMetricsResult.json", indicesStatsEndpoint, t)
	i := getTestingIntegration(t)

	// When executing populateIndicesMetrics
	err := populateIndicesMetrics(i, client, nil, testClusterName)

	// Then should be an error and not panic
	assert.Error(t, err, "should be an error")
}

func TestIndicesRegex(t *testing.T) {
	args.IndicesRegex = "twitter"
	i := getTestingIntegration(t)
	client := createNewTestClient()
	client.init("indicesMetricsResult.json", indicesStatsEndpoint, t)

	commonStruct := new(CommonMetrics)
	commonData, _ := ioutil.ReadFile(filepath.Join("testdata", "indicesMetricsResult_Common.json"))
	err := json.Unmarshal(commonData, commonStruct)
	assert.NoError(t, err)

	err = populateIndicesMetrics(i, client, commonStruct, testClusterName)
	assert.NoError(t, err)

	actualLength := len(i.Entities)
	expectedLength := 1
	actualName := i.Entities[0].Metadata.Name
	expectedName := "twitter"
	assert.Equal(t, expectedLength, actualLength)
	assert.Equal(t, expectedName, actualName)

	args.IndicesRegex = ""
}
