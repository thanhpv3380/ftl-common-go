package errors

import "github.com/thanhpv3380/ftl-common-go/common"

type SystemError struct {
	*GeneralError
}

func NewSystemError(source string) *SystemError {
	return &SystemError{
		GeneralError: NewGeneralError(string(common.INTERNAL_SERVER_ERROR), nil, source, nil),
	}
}
