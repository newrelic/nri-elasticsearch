package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var NodeTestFile = filepath.Join("testdata", "good-nodes-local.json")

type mockClient struct{
	mock.Mock
}

func (mc mockClient) Request(endpoint string) (objx.Map, error) {
	return getObjxMapFromFile(mc.Called(endpoint).String(0)), nil
}

func TestReadConfigFile(t *testing.T) {
	testCases := []struct {
		filePath    string
		expectedMap map[string]interface{}
	}{
		{
			filepath.Join("testdata", "elasticsearch_sample.yml"),
			map[string]interface{}{
				"path.data":    "/var/lib/elasticsearch",
				"path.logs":    "/var/log/elasticsearch",
				"network.host": "0.0.0.0",
			},
		},
	}

	for _, tc := range testCases {
		setupTestArgs()
		resultMap, err := readConfigFile(tc.filePath)
		if err != nil {
			t.Errorf("couldn't read config file: %v", err)
		} else {
			if expected := reflect.DeepEqual(tc.expectedMap, resultMap); !expected {
				t.Errorf("maps didn't match")
			}
		}
	}
}

func TestConfigErrors(t *testing.T) {
	testCases := []struct {
		filePath string
	}{
		{
			filepath.Join("testdata", "elasticsearch_doesntexist.yml"),
		},
		{
			filepath.Join("testdata", "elasticsearch_bad.yml"),
		},
	}

	for _, tc := range testCases {
		setupTestArgs()
		_, err := readConfigFile(tc.filePath)
		if err == nil {
			t.Errorf("was not expecting a result")
		}
	}
}

func TestPopulateConfigInventory(t *testing.T) {
	i, e := getTestingEntity(t)

	dataPath := filepath.Join("testdata", "elasticsearch_sample.yml")
	goldenPath := dataPath + ".golden"

	args.ConfigPath = dataPath

	populateConfigInventory(e)

	actual, _ := i.MarshalJSON()

	if *update {
		t.Log("Writing .golden file")
		err := ioutil.WriteFile(goldenPath, actual, 0644)
		assert.NoError(t, err)
	}

	expected, _ := ioutil.ReadFile(goldenPath)

	assert.Equal(t, expected, actual)
}

func TestParsePluginsAndModules(t *testing.T) {
	i, e := getTestingEntity(t)

	dataPath := filepath.Join("testdata", "good-node.json")
	goldenPath := dataPath + ".golden"

	statsJSON := getObjxMapFromFile(dataPath)

	populateNodeStatInventory(e, statsJSON)

	actualJSON, err := i.MarshalJSON()
	assert.NoError(t, err)

	if *update {
		t.Log("Writing .golden file")
		err := ioutil.WriteFile(goldenPath, actualJSON, 0644)
		assert.NoError(t, err)
	}

	expectedJSON, _ := ioutil.ReadFile(goldenPath)

	assert.Equal(t, expectedJSON, actualJSON)
}

func TestParseLocalNode(t *testing.T) {
	dataPath := filepath.Join("testdata", "good-nodes-local.json")
	goldenPath := dataPath + ".golden"

	statsJSON := getObjxMapFromFile(dataPath)

	_, actualStats, err := parseLocalNode(statsJSON)
	assert.NoError(t, err)

	actualString, _ := actualStats.JSON()
	if *update {
		t.Log("Writing .golden file")
		err := ioutil.WriteFile(goldenPath, []byte(actualString), 0644)
		assert.NoError(t, err)
	}

	expectedJSON, _ := ioutil.ReadFile(goldenPath)

	assert.Equal(t, string(expectedJSON), actualString)
}

func TestGetLocalNode(t *testing.T) {
	goldenPath := filepath.Join("testdata", "good-nodes-local.json.golden")

	fakeClient := mockClient{}
	mockedReturnVal := filepath.Join("testdata", "good-nodes-local.json")
	fakeClient.On("Request", "/_nodes/_local").Return(mockedReturnVal, nil).Once()

	resultName, resultStats, _ := getLocalNode(fakeClient)
	assert.Equal(t, "z9ZPp87vT92qG1cRVRIcMQ", resultName)

	actualString, _ := resultStats.JSON()
	if *update {
		t.Log("Writing .golden file")
		err := ioutil.WriteFile(goldenPath, []byte(actualString), 0644)
		assert.NoError(t, err)
	}

	expectedJSON, _ := ioutil.ReadFile(goldenPath)

	assert.Equal(t, string(expectedJSON), actualString)
	fakeClient.AssertExpectations(t)
}

func TestPopulateInventory(t *testing.T) {
	setupTestArgs()
	args.ConfigPath = filepath.Join("testdata", "elasticsearch_sample.yml")

	goldenPath := filepath.Join("testdata", "good-inventory.json.golden")

	fakeClient := mockClient{}
	mockedReturnVal := filepath.Join("testdata", "good-nodes-local.json")
	fakeClient.On("Request", "/_nodes/_local").Return(mockedReturnVal, nil).Once()

	i := getTestingIntegration(t)
	populateInventory(i, fakeClient)

	actualJSON, _ := i.MarshalJSON()
	if *update {
		t.Log("Writing .golden file")
		err := ioutil.WriteFile(goldenPath, actualJSON, 0644)
		assert.NoError(t, err)
	}

	expectedJSON, _ := ioutil.ReadFile(goldenPath)

	assert.Equal(t, expectedJSON, actualJSON)
	fakeClient.AssertExpectations(t)
}

func getObjxMapFromFile(fileName string) objx.Map {
	fileBytes, _ := ioutil.ReadFile(fileName)

	var resultMap map[string]interface{}

	_ = json.Unmarshal(fileBytes, &resultMap)

	return objx.New(resultMap)
}
