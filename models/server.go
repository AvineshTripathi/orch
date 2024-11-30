package models

type ServerResponse struct {
	Data string `json:"data"`
}

type TaskRequest struct {
	PluginName string
	Config     any
}
