package utils

import (
	"net/http"
)

// HandleError sends an error response with the given message and status code
func HandleError(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}
