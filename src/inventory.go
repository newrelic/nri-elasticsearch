package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
	"strings"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/stretchr/objx"
)

func populateInventory(entity *integration.Entity) {
	err := populateConfigInventory(entity)
	if err != nil {
		logger.Errorf("couldn't populate config inventory: %v", err)
	}
	err = populateNodeStatInventory(entity)
	if err != nil {
		logger.Errorf("couldn't populate node stat inventory: %v", err)
	}
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

func populateNodeStatInventory(entity *integration.Entity) error {
	
	// we'll get the stats back as part of metrics so this will change
	client, err := NewClient(nil)
	if err != nil {
		return err
	}

	localNodeStats, err := client.Request(localNodeInventoryEndpoint)
	if err != nil {
		return err
	}
	//

	localNode, err := getLocalNode(localNodeStats)
	if err != nil {
		return err
	}

	parseProcessStats(entity, localNode)
	parsePluginsAndModules(entity, localNode)
	parseNodeIngests(entity, localNode)
	return nil
}

func getLocalNode(localNodeStats objx.Map) (objx.Map, error) {
	nodes := localNodeStats.Get("nodes").ObjxMap()
	if len(nodes) == 1 {
		for _, v := range nodes {
			return objx.New(v), nil
		}
	}

	return nil, fmt.Errorf("could not identify local node")
}

func parseNodeIngests(entity *integration.Entity, stats objx.Map) []string {
	processorList := stats.Get("ingest.processors").ObjxMapSlice()

	typeList := []string{}

	for _, processor := range processorList {
		ingestType := processor.Get("type").String()
		typeList = append(typeList, ingestType)
	}

	err := entity.SetInventoryItem("config", "ingest", strings.Join(typeList, ","))
	if err != nil {
		logger.Errorf("error setting ingest types: %v", err)
	}

	return typeList
}

func parseProcessStats(entity *integration.Entity, stats objx.Map) {
	processStats := stats.Get("process").ObjxMap()

	for k, v := range processStats {
		err := entity.SetInventoryItem("config", "process." + k, v)
		if err != nil {
			logger.Errorf("error setting inventory item [%v -> %v]: %v", k, v, err)
		}
	}
}

func parsePluginsAndModules(entity *integration.Entity, stats objx.Map) {

	fieldNames := []string{
		"version",
		"elasticsearch_version",
		"java_version",
		"description",
		"classname",
	}

	for _, addonType := range []string{"plugins", "modules"} {
		addonStats := stats.Get(addonType).ObjxMapSlice()
		for _, addon := range addonStats {
			addonName := addon.Get("name").Str()
			for _, field := range fieldNames {
				inventoryKey := fmt.Sprintf("%v.%v.%v", addonType, addonName, field)
				inventoryValue := addon.Get(field).Str()
				err := entity.SetInventoryItem("config", inventoryKey, inventoryValue)
				if err != nil {
					logger.Errorf("error setting inventory item [%v -> %v]: %v", inventoryKey, inventoryValue, err)
				}
			}
		}
	}
}