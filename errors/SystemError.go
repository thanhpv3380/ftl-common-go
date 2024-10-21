package errors

import "ftl/kafi-common/common"

type SystemError struct {
	*GeneralError
}

func NewSystemError(source string) *SystemError {
	return &SystemError{
		GeneralError: NewGeneralError(string(common.INTERNAL_SERVER_ERROR), nil, source, nil),
	}
}
