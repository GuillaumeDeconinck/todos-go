package routes

import (
	"fmt"
	"reflect"

	daoerror "github.com/GuillaumeDeconinck/todos-go/internal/api/dao/daoError"
)

type HttpError struct {
	StatusCode int    `json:"statusCode"`
	ErrorName  string `json:"error"`
	// Message    string `json:"message"`
}

func (err HttpError) Error() string {
	return err.ErrorName
}

func ConvertToHttpError(err error) *HttpError {
	var statusCode int
	var errorName string

	switch reflect.TypeOf(err).String() {
	// If it's a DaoError coming from the layer below
	case "*daoerror.DaoError":
		fmt.Printf("Dao error detected\n")
		switch err.Error() {
		case daoerror.INTERNAL_ERROR_CODE:
			statusCode = 500
			errorName = "Internal server error"
		case daoerror.NOT_FOUND_CODE:
			statusCode = 404
			errorName = "Not found"
		default:
			statusCode = 500
			errorName = "Internal server error"
		}
	// If it's an "error" (base class from Golang)
	default:
		fmt.Printf("Default error detected")
		statusCode = 500
		errorName = "Internal server error"
	}

	return &HttpError{
		StatusCode: statusCode,
		ErrorName:  errorName,
	}
}
