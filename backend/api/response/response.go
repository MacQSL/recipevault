package response

import (
	"encoding/json"
	"net/http"
	"recipevault/util"
)

type httpError struct {
	Error string `json:"error"`
}

// Generate new HTTP error structure
func newHttpError(msg string) *httpError {
	return &httpError{Error: msg}
}

// Send the response back to client - trasforms data to json
func Send(w http.ResponseWriter, code int, data any) {
	jData, err := json.Marshal(data)

	// Failed to parse the data to JSON, 422
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Write the status code and response data
	w.WriteHeader(code)
	_, err = w.Write(jData)

	// Handling the edge case where response Write fails
	if err != nil {
		util.NewLogger().Error("Write() failed in response.Send()", err)
	}
}

// Send OK (200) status with data
func Send200(w http.ResponseWriter, data any) {
	Send(w, http.StatusOK, data)
}

// Send CREATED (201) status with data
func Send201(w http.ResponseWriter, data any) {
	Send(w, http.StatusCreated, data)
}

// Send BAD REQUEST (400) status with error and message
func Send400(w http.ResponseWriter, msg string) {
	Send(w, http.StatusBadRequest, newHttpError(msg))
}

// Send NOT FOUND (404) status with error and message
func Send404(w http.ResponseWriter, msg string) {
	Send(w, http.StatusNotFound, newHttpError(msg))
}

// Send FORBIDDEN (403) status with error and message
func Send403(w http.ResponseWriter, msg string) {
	Send(w, http.StatusForbidden, newHttpError(msg))
}

// Send INTERNAL SERVER ERROR (500) status with error and message
func Send500(w http.ResponseWriter, msg string) {
	Send(w, http.StatusInternalServerError, newHttpError(msg))
}
