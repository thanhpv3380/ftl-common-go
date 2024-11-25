package errors

import (
	"github.com/thanhpv3380/ftl-common-go/common"
)

func NewError(code string) *common.GeneralError {
	if code == "" {
		code = string(common.INTERNAL_SERVER_ERROR)
	}

	return &common.GeneralError{Code: code}
}

func NewGeneralError(code string, params map[string][]common.ParamError, source string, messageParams map[string]interface{}) *common.GeneralError {
	if code == "" {
		code = string(common.INTERNAL_SERVER_ERROR)
	}

	return &common.GeneralError{
		Code:          code,
		Source:        source,
		Params:        params,
		MessageParams: messageParams,
		IsSystemError: true,
	}
}
