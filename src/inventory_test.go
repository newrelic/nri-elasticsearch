package main

import (
	"reflect"
	"testing"
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