package main

import (
	"reflect"
	"testing"
	"io/ioutil"
	"bytes"
	
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
)

func TestReadConfigFile(t *testing.T) {
	testCases := []struct{
		filePath    string
		expectedMap map[string]interface{}
	}{
		{
			"testdata/elasticsearch_sample.yml",
			map[string]interface{}{
				"path.data": "/var/lib/elasticsearch",
				"path.logs": "/var/log/elasticsearch",
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
	testCases := []struct{
		filePath    string
	}{
		{
			"testdata/elasticsearch_doesntexist.yml",
		},
		{
			"testdata/elasticsearch_bad.yml",
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

	dataPath := "testdata/elasticsearch_sample.yml"
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

	if !bytes.Equal(expected, actual) {
		t.Errorf("Actual JSON results do not match expected .golden file")
	}
}

func TestParsePluginsAndModules(t *testing.T) {
	i, e := getTestingEntity(t)

	dataPath := "testdata/good-node.json"
	goldenPath := dataPath + ".golden"

	statsData, err := ioutil.ReadFile(dataPath)
	assert.NoError(t, err)

	statsJSON, err := objx.FromJSON(string(statsData))
	assert.NoError(t, err)

	populateNodeStatInventory(e, statsJSON)

	actualJSON, err := i.MarshalJSON()
	assert.NoError(t, err)

	if *update {
		t.Log("Writing .golden file")
		err := ioutil.WriteFile(goldenPath, actualJSON, 0644)
		assert.NoError(t, err)
	}

	expectedJSON, _ := ioutil.ReadFile(goldenPath)

	if !bytes.Equal(expectedJSON, actualJSON) {
		t.Errorf("Actual JSON results do not match expected .golden file")
	}
}

func TestParseLocalNode(t *testing.T) {
	dataPath := "testdata/good-nodes-local.json"
	goldenPath := dataPath + ".golden"

	statsData, err := ioutil.ReadFile(dataPath)
	assert.NoError(t, err)
	
	statsJSON, err := objx.FromJSON(string(statsData))
	assert.NoError(t, err)

	actual, err := parseLocalNode(statsJSON)
	assert.NoError(t, err)

	actualString, _ := actual.JSON()
	if *update {
		t.Log("Writing .golden file")
		err := ioutil.WriteFile(goldenPath, []byte(actualString), 0644)
		assert.NoError(t, err)
	}

	expectedJSON, _ := ioutil.ReadFile(goldenPath)

	if !bytes.Equal(expectedJSON, []byte(actualString)) {
		t.Errorf("Actual JSON results do not match expected .golden file")
	}


}