package errors

import (
	"ftl/kafi-common/common/interfaces"
	"ftl/kafi-common/common/errorCodes"
)

type GeneralError struct {
	Code          string
	MessageParams map[string]interface{}
	Source        string
	Params        map[string][]ParamError
	IsSystemError bool
}

func NewGeneralError(code string, params map[string][]ParamError, source string, messageParams map[string]interface{}) *GeneralError {
	if code == "" {
		code = INTERNAL_SERVER_ERROR
	}
	return &GeneralError{
		Code:          code,
		Source:        source,
		Params:        params,
		MessageParams: messageParams,
		IsSystemError: true,
	}
}

func (e *GeneralError) ToStatus() Status {
	return Status{
		Code:          e.Code,
		Params:        e.Params,
		MessageParams: e.MessageParams,
	}
}

func CreateFromStatus(status Status) *GeneralError {
	return NewGeneralError(status.Code, status.Params, "", status.MessageParams)
}