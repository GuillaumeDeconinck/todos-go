package daoerror

import (
	"fmt"

	"github.com/GuillaumeDeconinck/todos-go/pkg/tools"
)

const INTERNAL_ERROR_CODE = "INTERNAL_ERROR"
const NOT_FOUND_CODE = "NOT_FOUND"

type DaoError struct {
	Code string
}

func (err DaoError) Error() string {
	fmt.Println("Error() called")
	return err.Code
}

func ConvertToDaoError(err error) *DaoError {
	tools.SugaredLogger.Errorf("Error received %s\n", err)
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
