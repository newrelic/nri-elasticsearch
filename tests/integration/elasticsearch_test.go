// +build integration

package integration

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-elasticsearch/tests/integration/helpers"
	"github.com/newrelic/nri-elasticsearch/tests/integration/jsonschema"
	"github.com/stretchr/testify/assert"
)

var (
	secondsWaited               = 0
	elasticsearchMaxTimeoutWait = 30
	iName                       = "elasticsearch"

	defaultContainer          = "integration_nri-elasticsearch_1"
	defaultBinPath            = "/nri-elasticsearch"
	defaultClusterEnvironment = "staging"

	// cli flags
	container = flag.String("container", defaultContainer, "container where the integration is installed")
	binPath   = flag.String("bin", defaultBinPath, "Integration binary path")

	clusterEnvironment = flag.String("cluster_environment", defaultClusterEnvironment, "default cluster environment")
)

// Returns the standard output, or fails testing if the command returned an error
func runIntegration(t *testing.T, envVars ...string) (string, string, error) {
	t.Helper()

	command := make([]string, 0)
	command = append(command, *binPath)
	command = append(command, "--cluster_environment", *clusterEnvironment)

	stdout, stderr, err := helpers.ExecInContainer(*container, command, envVars...)

	if stderr != "" {
		log.Debug("Integration command Standard Error: ", stderr)
	}

	return stdout, stderr, err
}

func TestMain(m *testing.M) {
	flag.Parse()
	log.Info("Waiting for elasticsearch cluster to be ready")
	ensureElasticsearchClusterReady()
	log.Info("Elasticsearch cluster ready")

	result := m.Run()
	os.Exit(result)
}

func TestElasticsearchIntegration(t *testing.T) {
	stdout, stderr, err := runIntegration(t, "HOSTNAME=elasticsearch", "LOCAL_HOSTNAME=elasticsearch")
	assert.NotNil(t, stderr, "unexpected stderr")
	assert.NoError(t, err, "Unexpected error")

	schemaPath := filepath.Join("json-schema-files", "elasticsearch-schema.json")
	err = jsonschema.Validate(schemaPath, stdout)
	assert.NoError(t, err, "The output of Elasticsearch integration doesn't have expected format.")
}

func TestElasticsearchIntegrationOnlyMetrics(t *testing.T) {
	stdout, stderr, err := runIntegration(t, "METRICS=true", "HOSTNAME=elasticsearch", "LOCAL_HOSTNAME=elasticsearch")
	assert.NotNil(t, stderr, "unexpected stderr")
	assert.NoError(t, err, "Unexpected error")

	schemaPath := filepath.Join("json-schema-files", "elasticsearch-schema-metrics.json")
	err = jsonschema.Validate(schemaPath, stdout)
	assert.NoError(t, err, "The output of Elasticsearch integration doesn't have expected format.")
}

func TestElasticsearchIntegrationOnlyInventory(t *testing.T) {
	stdout, stderr, err := runIntegration(t, "INVENTORY=true", "HOSTNAME=elasticsearch", "LOCAL_HOSTNAME=elasticsearch")
	assert.NotNil(t, stderr, "unexpected stderr")
	assert.NoError(t, err, "Unexpected error")

	schemaPath := filepath.Join("json-schema-files", "elasticsearch-schema-inventory.json")
	err = jsonschema.Validate(schemaPath, stdout)
	assert.NoError(t, err, "The output of Elasticsearch integration doesn't have expected format.")
}

func ensureElasticsearchClusterReady() {
	fmt.Print("...")
	responseMaster, _ := http.Get("http://localhost:9200/_cat/master?h=id")
	responseSlave, _ := http.Get("http://localhost:9202/_cat/master?h=id")
	if (responseMaster == nil || responseSlave == nil) && secondsWaited < elasticsearchMaxTimeoutWait {
		secondsWaited++
		time.Sleep(1 * time.Second)
		ensureElasticsearchClusterReady()
	}
}
