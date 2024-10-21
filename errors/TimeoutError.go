package errors

import "ftl/kafi-common/common"

type TimeoutError struct {
	*GeneralError
}

func NewTimeoutError(params map[string][]common.ParamError, source string, messageParams map[string]interface{}) *TimeoutError {
	return &TimeoutError{
		GeneralError: NewGeneralError(string(common.TIMEOUT_ERROR), params, source, messageParams),
	}
}
