package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	yaml "gopkg.in/yaml.v2"
)

func populateInventory(i *integration.Integration, client Client) {
	// all inventory should be collected on the local node entity so we need to look that up
	localNodeName, localNode, err := getLocalNode(client)
	if err != nil {
		log.Error("Could not get local node stats: %v", err)
		return
	}

	clusterResponse, err := getClusterResponse(client)
	if err != nil {
		log.Error("Could not get cluster name: %v", err)
		return
	}

	// This should retrive the entity of the node if was already generated during the metrics processing.
	localNodeEntity, err := getEntity(i, localNodeName, "es-node", *clusterResponse.Name)
	if err != nil {
		log.Error("Could not get local node entity: %v", err)
		return
	}

	err = populateConfigInventory(localNodeEntity)
	if err != nil {
		log.Error("Could not populate config inventory: %v", err)
	}

	populateNodeStatInventory(localNodeEntity, localNode)
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
			return *nodes[k].Name, nodes[k], nil
		}
	}
	return "", nil, errors.New("could not identify local node")
}

func populateConfigInventory(entity *integration.Entity) error {
	configYaml, err := readConfigFile(args.ConfigPath)
	if err != nil {
		return err
	}

	for key, value := range configYaml {
		// special case to look for nested types
		switch value.(type) {
		case map[interface{}]interface{}:
			log.Debug("Unsupported data type '%T' for yaml key %s", value, key)
			continue
		case []interface{}:
			log.Debug("Unsupported data type '%T' for yaml key %s", value, key)
			continue
		}

		err = entity.SetInventoryItem("config/"+key, "value", value)
		if err != nil {
			log.Error("Could not set inventory item: %v", err)
		}
	}
	return nil
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

func populateNodeStatInventory(entity *integration.Entity, localNode *LocalNode) {
	parseProcessStats(entity, localNode)
	parsePluginsAndModules(entity, localNode)
	parseNodeIngests(entity, localNode)
}

func parseProcessStats(entity *integration.Entity, stats *LocalNode) {
	if stats.Process == nil {
		return
	}

	statsFields := reflect.TypeOf(*stats.Process)
	statsValues := reflect.ValueOf(*stats.Process)
	for i := 0; i < statsFields.NumField(); i++ {
		field := statsFields.Field(i)
		jsonKey, ok := field.Tag.Lookup("json")
		if !ok {
			continue
		}
		value := statsValues.Field(i).Interface()

		// we don't want to report nil values
		if reflect.ValueOf(value).IsNil() {
			continue
		}

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

func parseNodeIngests(entity *integration.Entity, stats *LocalNode) []string {
	if stats.Ingest == nil || stats.Ingest.Processors == nil {
		log.Error("No ingest stats defined for node. Skipping processor list collection", stats.Name)
		return []string{}
	}

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
