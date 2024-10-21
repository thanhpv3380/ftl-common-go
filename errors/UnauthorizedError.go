package errors

import "ftl/kafi-common/common"

type UnauthorizedError struct {
	*GeneralError
}

func NewUnauthorizedError(params map[string][]common.ParamError, source string, messageParams map[string]interface{}) *UnauthorizedError {
	return &UnauthorizedError{
		GeneralError: NewGeneralError(string(common.UNAUTHORIZED), params, source, messageParams),
	}
}
