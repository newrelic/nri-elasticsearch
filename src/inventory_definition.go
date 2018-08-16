package main

type LocalNodeResponse struct {
	Nodes       map[string]*LocalNode `json:"nodes"`
}

type LocalNode struct {
	Process *LocalNodeProcess `json:"process"`
	Ingest  *LocalNodeIngest  `json:"ingest"`
	Plugins []*LocalNodeAddon `json:"plugins"`
	Modules []*LocalNodeAddon `json:"modules"`
}

type LocalNodeProcess struct {
	RefreshInterval *int  `json:"refresh_interval_in_millis"`
	ID              *int  `json:"id"`
	MLockAll        *bool `json:"mlockall"`
}

type LocalNodeIngest struct {
	Processors []*IngestProcessor `json:"processors"`
}

type IngestProcessor struct {
	Type *string `json:"type"`
}

type LocalNodeAddon struct {
	Name                 *string `json:"name"`
	Version              *string `json:"version"`
	ElasticsearchVersion *string `json:"elasticsearch_version"`
	JavaVersion          *string `json:"java_version"`
	Description          *string `json:"description"`
	ClassName            *string `json:"classname"`
}