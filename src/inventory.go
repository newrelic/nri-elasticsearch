package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"

	"github.com/newrelic/infra-integrations-sdk/integration"
)

func populateInventory(entity *integration.Entity) error {
	return populateConfigInventory(entity)
}

func readConfigFile(filePath string) (map[string]interface{}, error) {
	rawYaml, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Errorf("could not open specified config file: %v", err)
		return nil, err
	}
	
	parsedYaml := make(map[string]interface{})

	err = yaml.Unmarshal(rawYaml, parsedYaml)
	if err != nil {
		logger.Errorf("could not parse configuration yaml: %v", err)
		return nil, err
	}

	return parsedYaml, nil
}

func populateConfigInventory(entity *integration.Entity) error {
	configYaml, err := readConfigFile(args.ConfigPath)
	if err != nil {
		return err
	}
	
	for key, value := range configYaml {
		err = entity.SetInventoryItem("config", key, value)
		if err != nil {
			logger.Errorf("could not set inventory item: %v", err)
		}
	}
	return nil
}