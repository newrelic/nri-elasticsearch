//go:build integration

package integration

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/newrelic/infra-integrations-sdk/v3/log"
	"github.com/newrelic/nri-elasticsearch/tests/integration/helpers"
	"github.com/newrelic/nri-elasticsearch/tests/integration/jsonschema"
	"github.com/stretchr/testify/assert"
)

var (
	secondsWaited               = 0
	elasticsearchMaxTimeoutWait = 60
	iName                       = "elasticsearch"

	defaultContainer          = "integration_nri-elasticsearch_1"
	defaultBinPath            = "/nri-elasticsearch"
	defaultClusterEnvironment = "staging"
	defaultUsername           = "elastic"
	defaultPassword           = "elastic"

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
	command = append(command, "--use_ssl")
	command = append(command, "--tls_insecure_skip_verify=true")
	command = append(command, "--username", defaultUsername)
	command = append(command, "--password", defaultPassword)

	stdout, stderr, err := helpers.ExecInContainer(*container, command, envVars...)

	if stderr != "" {
		log.Debug("Integration command Standard Error: ", stderr)
	}

	return stdout, stderr, err
}

func TestMain(m *testing.M) {
	flag.Parse()
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transCfg}
	log.Info("Waiting for elasticsearch cluster to be ready")
	ensureElasticsearchClusterReady(client)
	log.Info("Elasticsearch cluster ready")

	result := m.Run()
	os.Exit(result)
}

func TestElasticsearchIntegrationAll(t *testing.T) {
	stdout, stderr, err := runIntegration(t, "HOSTNAME=elasticsearch", "LOCAL_HOSTNAME=elasticsearch")
	assert.NotNil(t, stderr, "unexpected stderr")
	assert.NoError(t, err, "Unexpected error")

	errorExpected := false
	schemaPath := filepath.Join("json-schema-files", "elasticsearch-schema.json")
	err = jsonschema.Validate(schemaPath, stdout, errorExpected)
	assert.NoError(t, err, "The output of Elasticsearch integration doesn't have expected format.")
}

func TestElasticsearchIntegrationOnlyMetrics(t *testing.T) {
	stdout, stderr, err := runIntegration(t, "METRICS=true", "HOSTNAME=elasticsearch", "LOCAL_HOSTNAME=elasticsearch")
	assert.NotNil(t, stderr, "unexpected stderr")
	assert.NoError(t, err, "Unexpected error")

	errorExpected := false
	schemaPath := filepath.Join("json-schema-files", "elasticsearch-schema-metrics.json")
	err = jsonschema.Validate(schemaPath, stdout, errorExpected)
	assert.NoError(t, err, "The output of Elasticsearch integration doesn't have expected format.")
}

func TestElasticsearchIntegrationOnlyInventory(t *testing.T) {
	stdout, stderr, err := runIntegration(t, "INVENTORY=true", "HOSTNAME=elasticsearch", "LOCAL_HOSTNAME=elasticsearch")
	assert.NotNil(t, stderr, "unexpected stderr")
	assert.NoError(t, err, "Unexpected error")

	errorExpected := false
	schemaPath := filepath.Join("json-schema-files", "elasticsearch-schema-inventory.json")
	err = jsonschema.Validate(schemaPath, stdout, errorExpected)
	assert.NoError(t, err, "The output of Elasticsearch integration doesn't have expected format.")
}

func TestElasticsearchIntegrationAllOnSlave_OnlyMasterFlagTrue(t *testing.T) {
	stdout, stderr, err := runIntegration(t, "MASTER_ONLY=true", "PORT=9200", "HOSTNAME=elasticsearch-replica", "LOCAL_HOSTNAME=elasticsearch-replica")
	assert.NotNil(t, stderr, "unexpected stderr")
	assert.NoError(t, err, "Unexpected error")

	errorExpected := false
	schemaPath := filepath.Join("json-schema-files", "elasticsearch-schema-inventory.json")
	err = jsonschema.Validate(schemaPath, stdout, errorExpected)
	assert.NoError(t, err, "The output of Elasticsearch integration doesn't have expected format.")
}

func TestElasticsearchIntegrationAllOnSlave_OnlyMasterFlagFalse(t *testing.T) {
	stdout, stderr, err := runIntegration(t, "MASTER_ONLY=false", "PORT=9200", "HOSTNAME=elasticsearch-replica", "LOCAL_HOSTNAME=elasticsearch-replica")
	assert.NotNil(t, stderr, "unexpected stderr")
	assert.NoError(t, err, "Unexpected error")

	errorExpected := true
	schemaPath := filepath.Join("json-schema-files", "elasticsearch-schema-inventory.json")
	err = jsonschema.Validate(schemaPath, stdout, errorExpected)
	assert.Error(t, err)
}

func ensureElasticsearchClusterReady(client *http.Client) {
	fmt.Print(".")
	requestMaster, _ := http.NewRequest("GET", "https://localhost:9200/_nodes/stats", nil)
	requestMaster.SetBasicAuth(defaultUsername, defaultPassword)
	requestSlave, _ := http.NewRequest("GET", "https://localhost:9202/_nodes/stats", nil)
	requestSlave.SetBasicAuth(defaultUsername, defaultPassword)
	responseMaster, _ := client.Do(requestMaster)
	responseSlave, _ := client.Do(requestSlave)
	if (responseMaster == nil || responseSlave == nil) && secondsWaited < elasticsearchMaxTimeoutWait {
		secondsWaited++
		time.Sleep(1 * time.Second)
		ensureElasticsearchClusterReady(client)
	}
}
