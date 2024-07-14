package handlers

import (
	"encoding/json"
	"net/http"
	"orch/models"
)

func CreateNewEnvHandler(w http.ResponseWriter, r *http.Request) {
	response := models.ExampleResponse{Data: "Enviroment was created"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
