package errors

import (
	"github/thanhpv3380/ftl-common-go/common"
)

type ObjectNotFoundError struct {
	*GeneralError
}

func NewObjectNotFoundError(source string) *ObjectNotFoundError {
	return &ObjectNotFoundError{
		GeneralError: NewGeneralError(string(common.OBJECT_NOT_FOUND), nil, source, nil),
	}
}
