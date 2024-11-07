package errors

import (
	"fmt"

	"github.com/thanhpv3380/ftl-common-go/common"
)

type GeneralError struct {
	Code          string
	MessageParams map[string]interface{}
	Source        string
	Params        map[string][]common.ParamError
	IsSystemError bool
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
