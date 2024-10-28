package handlers

import (
	"encoding/json"
	"github.com/AvineshTripathi/orch/models"
	"github.com/AvineshTripathi/orch/provisioner-client"
	"github.com/AvineshTripathi/orch/utils"
	"net/http"
)

func ApiServerStatusHandler(w http.ResponseWriter, r *http.Request) {
	response := models.ServerResponse{Data: "Server is up"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ProvisionerStatusHandler(w http.ResponseWriter, r *http.Request) {
	response, err := provisioner.GetProvisionerStatus()
	if err != nil {
		utils.HandleError(w, "provision server not up", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ServerResponse{Data: response})
}
