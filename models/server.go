package models

type ServerResponse struct {
	Data string `json:"data"`
}

type TaskRequest struct {
	URL  string
	Data string
}
