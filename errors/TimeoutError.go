package errors

import "github.com/thanhpv3380/ftl-common-go/common"

type TimeoutError struct {
	*GeneralError
}

func NewTimeoutError(params map[string][]common.ParamError, source string, messageParams map[string]interface{}) *TimeoutError {
	return &TimeoutError{
		GeneralError: NewGeneralError(string(common.TIMEOUT_ERROR), params, source, messageParams),
	}
}
