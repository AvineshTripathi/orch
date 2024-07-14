package handlers

import (
	"encoding/json"
	"net/http"
	"orch/models"
	"orch/provisioner-client"
	"orch/utils"
)

// ExampleHandler handles the /example endpoint
func ApiServerStatusHandler(w http.ResponseWriter, r *http.Request) {
	response := models.ExampleResponse{Data: "Server is up"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ProvisionerStatusHandler(w http.ResponseWriter, r *http.Request) {
	response, err := provisioner.GetProvisionerStatus()
	if err != nil {
		utils.HandleError(w, "provision server not up", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ExampleResponse{Data: response})
}
