package errors

import (
	"ftl/kafi-common/common/interfaces"
	"ftl/kafi-common/common/errorCodes"
)

type SystemError struct {
	*GeneralError
}

func NewSystemError(source string) *SystemError {
	return &SystemError{
		GeneralError: NewGeneralError(INTERNAL_SERVER_ERROR, nil, source, nil),
	}
}