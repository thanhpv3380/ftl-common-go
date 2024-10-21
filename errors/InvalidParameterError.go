package errors

import (
	"ftl/kafi-common/common/interfaces"
	"ftl/kafi-common/common/errorCodes"
)

type InvalidParameterError struct {
	*GeneralError
}

func NewInvalidParameterError(params map[string][]ParamError, source string, messageParams map[string]interface{}) *InvalidParameterError {
	return &InvalidParameterError{
		GeneralError: NewGeneralError(INVALID_PARAMETER, params, source, messageParams),
	}
}
