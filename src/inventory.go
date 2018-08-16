package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
)

func populateInventory(i *integration.Integration, client Client) {
	// all inventory should be collected on the local node entity so we need to look that up
	localNodeName, localNode, err := getLocalNode(client)
	if err != nil {
		log.Error("Couldn't get local node stats: %v", err)
		return
	}

	localNodeEntity, err := i.Entity(localNodeName, "node")
	if err != nil {
		log.Error("Couldn't get local node entity: %v", err)
		return
	}

	err = populateConfigInventory(localNodeEntity)
	if err != nil {
		log.Error("Couldn't populate config inventory: %v", err)
	}

	populateNodeStatInventory(localNodeEntity, localNode)
}

func readConfigFile(filePath string) (map[string]interface{}, error) {
	rawYaml, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error("Could not open specified config file: %v", err)
		return nil, err
	}

	parsedYaml := make(map[string]interface{})

	err = yaml.Unmarshal(rawYaml, parsedYaml)
	if err != nil {
		log.Error("Could not parse configuration yaml: %v", err)
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
		err = entity.SetInventoryItem("config/"+key, "value", value)
		if err != nil {
			log.Error("Could not set inventory item: %v", err)
		}
	}
	return nil
}

func populateNodeStatInventory(entity *integration.Entity, localNode *LocalNode) {
	parseProcessStats(entity, localNode)
	parsePluginsAndModules(entity, localNode)
	parseNodeIngests(entity, localNode)
}

func getLocalNode(client Client) (localNodeName string, localNodeStats *LocalNode, err error) {
	nodeResponseObject := new(LocalNodeResponse)
	err = client.Request(localNodeInventoryEndpoint, &nodeResponseObject)
	if err != nil {
		return "", nil, err
	}

	localNodeName, localNodeStats, err = parseLocalNode(nodeResponseObject)
	if err != nil {
		return "", nil, err
	}
	return
}

func parseLocalNode(nodeStats *LocalNodeResponse) (string, *LocalNode, error) {
	nodes := nodeStats.Nodes
	if len(nodes) == 1 {
		for k := range nodes {
			return k, nodes[k], nil
		}
	}
	return "", nil, fmt.Errorf("could not identify local node")
}

func parseNodeIngests(entity *integration.Entity, stats *LocalNode) []string {
	processorList := stats.Ingest.Processors
	typeList := []string{}

	for _, processor := range processorList {
		ingestType := processor.Type
		typeList = append(typeList, *ingestType)
	}

	err := entity.SetInventoryItem("config/ingest", "value", strings.Join(typeList, ","))
	if err != nil {
		log.Error("Error setting ingest types: %v", err)
	}

	return typeList
}

func parseProcessStats(entity *integration.Entity, stats *LocalNode) {
	statsFields := reflect.TypeOf(*stats.Process)
	statsValues := reflect.ValueOf(*stats.Process)
	for i := 0; i < statsFields.NumField(); i++ {
		field := statsFields.Field(i)
		jsonKey, ok := field.Tag.Lookup("json")
		if !ok {
			continue
		}
		value := statsValues.Field(i).Interface()
		err := entity.SetInventoryItem("config/process/"+jsonKey, "value", value)
		if err != nil {
			log.Error("Error setting inventory item [%s -> %s]: %v", jsonKey, value, err)
		}
	}
}

func parsePluginsAndModules(entity *integration.Entity, stats *LocalNode) {

	fieldNames := []string{
		"Version",
		"ElasticsearchVersion",
		"JavaVersion",
		"Description",
		"ClassName",
	}

	for _, addonSet := range []string{"plugins", "modules"} {
		var addonStats []*LocalNodeAddon
		if addonSet == "plugins" {
			addonStats = stats.Plugins
		} else {
			addonStats = stats.Modules
		}

		for _, addon := range addonStats {
			addonName := *addon.Name
			addonFields := reflect.ValueOf(*addon)
			for _, field := range fieldNames {
				addonFieldValue := addonFields.FieldByName(field).Interface()
				inventoryKey := fmt.Sprintf("%s/%s/%s", addonSet, addonName, field)
				err := entity.SetInventoryItem("config/"+inventoryKey, "value", addonFieldValue)
				if err != nil {
					log.Error("Error setting inventory item [%s -> %s]: %v", inventoryKey, addonFieldValue, err)
				}
			}
		}
	}
}
