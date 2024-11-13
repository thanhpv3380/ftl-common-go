package errors

import (
	"fmt"

	"github.com/thanhpv3380/ftl-common-go/common"
)

type GeneralError struct {
	Code          string                         `json:"code"`
	MessageParams map[string]interface{}         `json:"messageParams"`
	Source        string                         `json:"source"`
	Params        map[string][]common.ParamError `json:"params"`
	IsSystemError bool                           `json:"isSystemError"`
}

func (e *GeneralError) Error() string {
	return fmt.Sprintf(e.Code)
}

func NewGeneralError(code string, params map[string][]common.ParamError, source string, messageParams map[string]interface{}) *GeneralError {
	if code == "" {
		code = string(common.INTERNAL_SERVER_ERROR)
	}

	return &GeneralError{
		Code:          code,
		Source:        source,
		Params:        params,
		MessageParams: messageParams,
		IsSystemError: true,
	}
}
