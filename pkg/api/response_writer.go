package api

import (
	"encoding/json"
	"net/http"

	"gobase.com/base/pkg/registering"
)

type errorResponse struct {
	ErrorCode string      `json:"error_code"`
	Field     string      `json:"field,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

//
func writeJSON(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// pack the response struct into a JSON response body and send it
	responseBytes, _ := json.Marshal(&body)
	w.Write(responseBytes)
}

func writeSuccess(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func writeErrors(w http.ResponseWriter, errs []error, placeholderData interface{}) {
	var response []errorResponse
	// set the appropriate headers
	w.Header().Add("Content-Type", "application/json")

	status := http.StatusBadRequest

	for _, err := range errs {
		// depending on the type of the error...
		switch err.(type) {
		case *registering.Error:
			rErr := err.(*registering.Error)
			//log warning
			response = append(response, errorResponse{ErrorCode: string(rErr.Code), Field: rErr.Field, Data: rErr.Data})
			status = http.StatusBadRequest
		default:
			//log error
			status = http.StatusInternalServerError
		}
	}

	// pack the response struct into a JSON response body and send it
	w.WriteHeader(status)
	responseBytes, _ := json.Marshal(&response)
	w.Write(responseBytes)

}

func writeError(w http.ResponseWriter, err error, placeholderData interface{}) {
	var response []errorResponse
	// set the appropriate headers
	w.Header().Add("Content-Type", "application/json")

	status := http.StatusBadRequest

	// depending on the type of the error...
	switch err.(type) {
	case *registering.Error:
		rErr := err.(*registering.Error)
		//log warning
		response = append(response, errorResponse{ErrorCode: string(rErr.Code), Field: rErr.Field, Data: rErr.Data})
		status = http.StatusBadRequest
	default:
		//log error
		status = http.StatusInternalServerError
	}

	// pack the response struct into a JSON response body and send it
	w.WriteHeader(status)
	responseBytes, _ := json.Marshal(&response)
	w.Write(responseBytes)

}
