package daoerror

import (
	"strings"

	"github.com/GuillaumeDeconinck/todos-go/pkg/tools"
)

const INTERNAL_ERROR_CODE = "INTERNAL_ERROR"
const NOT_FOUND_CODE = "NOT_FOUND"

var nbErrors = 0

type DaoError struct {
	Code string
}

func (err DaoError) Error() string {
	return err.Code
}

func ConvertToDaoError(err error) *DaoError {
	tools.SugaredLogger.Errorf("Error received %s\n", err)

	if strings.Contains(err.Error(), "FATAL") {
		nbErrors += 1
		// Need a way to clean nbErrors like setTimeout in NodeJS
		if nbErrors > 5 {
			// Close server ?
		}
	}

	Code := INTERNAL_ERROR_CODE
	return &DaoError{
		Code: Code,
	}
}

func New(errorCode string) *DaoError {
	return &DaoError{
		Code: errorCode,
	}
}
