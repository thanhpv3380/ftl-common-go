package errors

import (
	"ftl/kafi-common/common/interfaces"
	"ftl/kafi-common/common/errorCodes"
)

type UnauthorizedError struct {
	*GeneralError
}

func NewUnauthorizedError(params map[string][]ParamError, source string, messageParams map[string]interface{}) *UnauthorizedError {
	return &UnauthorizedError{
		GeneralError: NewGeneralError(UNAUTHORIZED, params, source, messageParams),
	}
}
