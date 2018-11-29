package main

// LocalNodeResponse struct (/_nodes/_local endpoint)
type LocalNodeResponse struct {
	Nodes map[string]*LocalNode `json:"nodes"`
}

// LocalNode is the node API object
type LocalNode struct {
	Host    *string           `json:"host"`
	Process *LocalNodeProcess `json:"process"`
	Ingest  *LocalNodeIngest  `json:"ingest"`
	Plugins []*LocalNodeAddon `json:"plugins"`
	Modules []*LocalNodeAddon `json:"modules"`
}

// LocalNodeProcess (node/process)
type LocalNodeProcess struct {
	RefreshInterval *int  `json:"refresh_interval_in_millis"`
	ID              *int  `json:"id"`
	MLockAll        *bool `json:"mlockall"`
}

// LocalNodeIngest (node/ingest)
type LocalNodeIngest struct {
	Processors []*IngestProcessor `json:"processors"`
}

// IngestProcessor (node/ingest contains a json array of these processors)
type IngestProcessor struct {
	Type *string `json:"type"`
}

// LocalNodeAddon (modules and plugins)
type LocalNodeAddon struct {
	Name                 *string `json:"name"`
	Version              *string `json:"version"`
	ElasticsearchVersion *string `json:"elasticsearch_version"`
	JavaVersion          *string `json:"java_version"`
	Description          *string `json:"description"`
	ClassName            *string `json:"classname"`
}
