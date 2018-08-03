package main

import (
	"testing"
	"flag"

	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/mock"
)

var (
	update = flag.Bool("update", false, "update .golden files")
)

func setupTestArgs() {
	args = argumentList{}
	logger = log.NewStdErr(true)
}

type mockedLogger struct {
	mock.Mock
}
type testLogger struct {
	f func(format string, args ...interface{})
}

func (l *mockedLogger) Debugf(format string, args ...interface{}) {
	args = append([]interface{}{format}, args...)
	l.Called(args...)
}
func (l *mockedLogger) Infof(format string, args ...interface{}) {
	args = append([]interface{}{format}, args...)
	l.Called(args...)
}
func (l *mockedLogger) Errorf(format string, args ...interface{}) {
	args = append([]interface{}{format}, args...)
	l.Called(args...)
}
func (l *mockedLogger) Warnf(format string, args ...interface{}) {
	args = append([]interface{}{format}, args...)
	l.Called(args...)
}

func (l *testLogger) Debugf(format string, args ...interface{}) {
	l.f("DEBUG: "+format, args...)
}
func (l *testLogger) Infof(format string, args ...interface{}) {
	l.f("INFO : "+format, args...)
}
func (l *testLogger) Errorf(format string, args ...interface{}) {
	l.f("ERROR: "+format, args...)
}
func (l *testLogger) Warnf(format string, args ...interface{}) {
	l.f("WARN: "+format, args...)
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
	payload, err := integration.New("Test", "0.0.1", integration.Logger(&testLogger{t.Logf}))
	require.NoError(t, err)
	require.NotNil(t, payload)
	logger = payload.Logger()
	return
}