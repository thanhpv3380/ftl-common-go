package errors

import (
	"ftl/kafi-common/common"
)

type ObjectNotFoundError struct {
	*GeneralError
}

func NewObjectNotFoundError(source string) *ObjectNotFoundError {
	return &ObjectNotFoundError{
		GeneralError: NewGeneralError(string(common.OBJECT_NOT_FOUND), nil, source, nil),
	}
}
