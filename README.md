# New Relic Infrastructure Integration for elasticsearch

The New Relic Infrastructure Integration for Elasticsearch captures critical performance metrics and inventory reported by Elasticsearch clusters. Data on the cluster, nodes, shards, and indices is collected.

Inventory data is obtained from the elasticsearch.yml file, and metrics and additional inventory data is obtained from the REST API.

## Requirements

No additional requirements

## Installation

- download an archive file for the `Elasticsearch` Integration
- extract `elasticsearch-definition.yml` and `/bin` directory into `/var/db/newrelic-infra/newrelic-integrations`
- add execute permissions for the binary file `nr-elasticsearch` (if required)
- extract `elasticsearch-config.yml.sample` into `/etc/newrelic-infra/integrations.d`

## Usage

This is the description about how to run the ElasticSearch Integration with New Relic Infrastructure agent, so it is required to have the agent installed (see [agent installation](https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure/installation/install-infrastructure-linux)).

In order to use the ElasticSearch Integration it is required to configure `elasticsearch-config.yml.sample` file. Firstly, rename the file to `elasticsearch-config.yml`. Then, depending on your needs, specify all instances that you want to monitor. Once this is done, restart the Infrastructure agent.

You can view your data in Insights by creating your own custom NRQL queries. To do so use the **ElasticsearchClusterSample**, **ElasticsearchIndexSample**, **ElasticsearchNodeSample** event type.

## Compatibility

* Supported OS: No limitations
* elasticsearch versions: 5.0+

## Integration Development usage

Assuming that you have source code you can build and run the Elasticsearch Integration locally.

* Go to directory of the Elasticsearch Integration and build it
```bash
$ make
```
* The command above will execute tests for the Elasticsearch Integration and build an executable file called `nr-elasticsearch` in `bin` directory.
```bash
$ ./bin/nr-elasticsearch
```
* If you want to know more about usage of `./nr-elasticsearch` check
```bash
$ ./bin/nr-elasticsearch -help
```

For managing external dependencies [govendor tool](https://github.com/kardianos/govendor) is used. It is required to lock all external dependencies to specific version (if possible) into vendor directory.
