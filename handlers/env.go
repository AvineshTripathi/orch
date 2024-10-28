package handlers

import (
	"encoding/json"
	"github.com/AvineshTripathi/orch/models"
	"github.com/AvineshTripathi/orch/provisioner/queue"
	"github.com/AvineshTripathi/orch/provisioner/task"
	"github.com/AvineshTripathi/orch/utils"
	"io"
	"net/http"
)

func AddTaskToQueueHandler(client *queue.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			utils.HandleError(w, "cannot read message", http.StatusInternalServerError)
		}

		var req models.TaskRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			utils.HandleError(w, "invalid message", http.StatusInternalServerError)
		}

		tsk := task.NewTask(req.PluginName, "", req.Config)

		_, err = client.AddNewTask(tsk)
		if err != nil {
			utils.HandleError(w, "couldnot add task to the queue", http.StatusInternalServerError)
		}

		response := models.ServerResponse{Data: "Enviroment was created"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
