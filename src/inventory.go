package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/stretchr/objx"
)

func populateInventory(i *integration.Integration, client *Client) {
	// all inventory should be collected on the local node entity so we need to look that up
	localNodeName, localNode, err := getLocalNode(client)
	if err != nil {
		logger.Errorf("couldn't get local node stats: %v", err)
		return
	}

	localNodeEntity, err := lookupLocalNode(i, localNodeName)
	if err != nil {
		logger.Errorf("couldn't look up local node entity: %v", err)
	}

	err = populateConfigInventory(localNodeEntity)
	if err != nil {
		logger.Errorf("couldn't populate config inventory: %v", err)
	}

	populateNodeStatInventory(localNodeEntity, localNode)
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

func populateNodeStatInventory(entity *integration.Entity, localNode objx.Map) {
	parseProcessStats(entity, localNode)
	parsePluginsAndModules(entity, localNode)
	parseNodeIngests(entity, localNode)
}

func getLocalNode(client *Client) (localNodeName string, localNodeStats objx.Map, err error) {
	nodeStats, err := client.Request(localNodeInventoryEndpoint)
	if err != nil {
		return "", nil, err
	}

	localNodeName, localNodeStats, err = parseLocalNode(nodeStats)
	if err != nil {
		return "", nil, err
	}
	return
}

func parseLocalNode(nodeStats objx.Map) (string, objx.Map, error) {
	nodes := nodeStats.Get("nodes").ObjxMap()
	if len(nodes) == 1 {
		for k := range nodes {
			return k, nodes.Get(k).ObjxMap(), nil
		}
	}
	return "", nil, fmt.Errorf("could not identify local node")
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
		err := entity.SetInventoryItem("config", "process."+k, v)
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

func lookupLocalNode(i *integration.Integration, nodeName string) (*integration.Entity, error) {
	for _, e := range i.Entities {
		if e.Metadata.Name == nodeName {
			return e, nil
		}
	}
	logger.Infof("entity for local node did not exist after parsing metrics; creating entity")
	localEntity, err := i.Entity(nodeName, "node")
	if err != nil {
		return nil, err
	}
	return localEntity, err
}