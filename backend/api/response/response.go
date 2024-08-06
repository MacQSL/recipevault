package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// 1. Respond with success and data
// 2. Parse data to JSON
// 3. Respond with error and data
// 4. Respond with error formatted as { "error": "aasdfadf" }

type errorResponse struct {
	Error interface{} `json:"error"`
}

// Is the status successfull ie: 200 / 201
//
// Very simple implementation of this function, could be
// extended to strip the first digit from status and check
// appropriately. ie: [2]00 (true) vs [4]00 (false)
func isSuccessStatus(status int) bool {
	return status == 200 || status == 201
}

// Send the response back to client - trasforms data to json
func Send(w http.ResponseWriter, status int, data any) {
	var jData []byte
	var err error

	if isSuccessStatus(status) {
		jData, err = json.Marshal(data)
	} else {
		jData, err = json.Marshal(&errorResponse{Error: data})
	}

	// Failed to parse the data to JSON, 422
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Write the status code and response data
	w.WriteHeader(status)
	_, err = w.Write(jData)

	// Handling the offchance the response writer fails
	if err != nil {
		log.Println("ERROR: response.Send() Write", err)
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

// Send NOT FOUND (404)status with data
func Send404(w http.ResponseWriter, data any) {
	Send(w, http.StatusNotFound, data)
}

// Send BAD REQUEST (404)status with data
func Send400(w http.ResponseWriter, data any) {
	Send(w, http.StatusBadRequest, data)
}

// Send INTERNAL SERVER ERROR (500) status with data
func Send500(w http.ResponseWriter, data any) {
	Send(w, http.StatusInternalServerError, data)
}
