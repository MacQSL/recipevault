package response

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"recipevault/util"
)

type errorResponse struct {
	Error *string `json:"error"`
}

// Standard response struct
type response struct {
	// status message ie: 'success' / 'error' etc...
	Status string `json:"status"`
	// http status code as integer ie: 400 / 200 etc...
	HttpCode int `json:"httpCode"`
	// http status as string ie: "Success" / "Not Found"
	HttpStatus string `json:"httpStatus"`
	// either success payload or error
	Data any `json:"data"`
	// additional response details / message
	// for errors this should be something the UI can ingest
	Message *string `json:"message"`
}

// Generate a success response struct
func newSuccessResponse(code int, data any) *response {
	return &response{
		Status:     "success",
		HttpCode:   code,
		HttpStatus: http.StatusText(code),
		Data:       data,
		Message:    nil, // this can be included if needed in future
	}
}

// Generate an error response struct
func newErrorResponse(code int, err error, message *string) *response {
	var errRes *errorResponse

	// fallback to message if error nil
	if err != nil {
		errMsg := err.Error()
		errRes = &errorResponse{Error: &errMsg}
	} else {
		errRes = &errorResponse{Error: message}
	}

	return &response{
		Status:     "error",
		HttpCode:   code,
		HttpStatus: http.StatusText(code),
		Data:       errRes,
		Message:    message,
	}
}

// Send the response back to client - trasforms data to json
func Send(w http.ResponseWriter, res *response) {
	jData, err := json.Marshal(res)

	// Failed to parse the data to JSON, 422
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Write the status code and response data
	w.WriteHeader(res.HttpCode)
	_, err = w.Write(jData)

	// Handling the edge case where response Write fails
	if err != nil {
		util.NewLogger().Error("Write() failed in response.Send()", err)
	}
}

// Send OK (200) status with data
func Send200(w http.ResponseWriter, data any) {
	Send(w, newSuccessResponse(http.StatusOK, data))
}

// Send CREATED (201) status with data
func Send201(w http.ResponseWriter, data any) {
	Send(w, newSuccessResponse(http.StatusCreated, data))
}

// Send BAD REQUEST (400) status with error and message
func Send400(w http.ResponseWriter, err error, message string) {
	Send(w, newErrorResponse(http.StatusBadRequest, err, &message))
}

// Send NOT FOUND (404) status with error and message
func Send404(w http.ResponseWriter, err error, message string) {
	Send(w, newErrorResponse(http.StatusNotFound, err, &message))
}

// Send FORBIDDEN (403) status with error and message
func Send403(w http.ResponseWriter, err error, message string) {
	Send(w, newErrorResponse(http.StatusForbidden, err, &message))
}

// Send INTERNAL SERVER ERROR (500) status with error and message
func Send500(w http.ResponseWriter, err error, message string) {
	Send(w, newErrorResponse(http.StatusInternalServerError, err, &message))
}

// Send correct response type for error
func SendError(w http.ResponseWriter, err error, message string) {
	switch err {
	case sql.ErrNoRows:
		Send404(w, err, message)
	default:
		Send500(w, err, message)
	}
}
