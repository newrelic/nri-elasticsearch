package main

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopulateNodesMetrics(t *testing.T) {
	sourceFile := filepath.Join("testdata", "nodeStatsMetricsResult.json")
	i := getTestingIntegration(t)
	data := getObjxMapFromFile(sourceFile)
	populateNodesMetrics(i, &data)

	assert.Equal(t, 1, len(i.Entities))
	assert.Equal(t, 1, len(i.Entities[0].Metrics))
	goldenFile := sourceFile + ".golden"
	actual, _ := i.Entities[0].Metrics[0].MarshalJSON()

	if *update {
		ioutil.WriteFile(goldenFile, actual, 0644)
	}

	expected, _ := ioutil.ReadFile(goldenFile)
	actualLength := len(i.Entities[0].Metrics[0].Metrics)
	expectedLength := len(nodeMetricDefs.MetricDefs)

	assert.Equal(t, expected, actual)
	assert.Equal(t, actualLength, expectedLength)
}
