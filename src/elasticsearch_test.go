package main

import (
	"github.com/newrelic/infra-integrations-sdk/log"
)

func setupTestArgs() {
	args = argumentList{}
	logger = log.NewStdErr(true)
}
