package errors

import "github.com/thanhpv3380/ftl-common-go/common"

type InvalidParameterError struct {
	*GeneralError
}

func NewInvalidParameterError(params map[string][]common.ParamError, source string, messageParams map[string]interface{}) *InvalidParameterError {
	return &InvalidParameterError{
		GeneralError: NewGeneralError(string(common.INVALID_PARAMETER), params, source, messageParams),
	}
}
