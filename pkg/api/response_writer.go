package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gobase.com/base/pkg/registering"
)

type errorResponse struct {
	ErrorCode string      `json:"error_code"`
	Field     string      `json:"field,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

func writeJSON(c echo.Context, body interface{}) {
	c.JSON(http.StatusOK, body)
}

func writeSuccess(c echo.Context) {
	c.Response().Header().Add("Content-Type", "application/json")
	c.NoContent(http.StatusOK)
}

func writeErrors(c echo.Context, errs []error, placeholderData interface{}) {
	var response []errorResponse
	// set the appropriate headers
	c.Response().Header().Add("Content-Type", "application/json")

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
	c.JSON(status, response)
}

func writeError(c echo.Context, err error, placeholderData interface{}) {
	var response []errorResponse
	// set the appropriate headers
	c.Response().Header().Add("Content-Type", "application/json")

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
	c.JSON(status, response)
}
