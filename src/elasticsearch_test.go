package main

import (
	"flag"
	"testing"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-elasticsearch/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	update = flag.Bool("update", false, "update .golden files")
)

func setupTestArgs() {
	args = argumentList{}
	logger = log.NewStdErr(true)
}

func getTestingEntity(t *testing.T, entityArgs ...string) (payload *integration.Integration, entity *integration.Entity) {
	payload = getTestingIntegration(t)
	var err error
	if len(entityArgs) > 1 {
		entity, err = payload.Entity(entityArgs[0], entityArgs[1])
		assert.NoError(t, err)
	} else {
		entity = payload.LocalEntity()
	}
	require.NotNil(t, entity)
	return
}

func getTestingIntegration(t *testing.T) (payload *integration.Integration) {
	payload, err := integration.New("Test", "0.0.1", integration.Logger(&testutils.TestLogger{F: t.Logf}))
	require.NoError(t, err)
	require.NotNil(t, payload)
	logger = payload.Logger()
	return
}
