package errors

import (
	"ftl/kafi-common/common/interfaces"
	"ftl/kafi-common/common/errorCodes"
)

type ObjectNotFoundError struct {
	*GeneralError
}

func NewObjectNotFoundError(source string) *ObjectNotFoundError {
	return &ObjectNotFoundError{
		GeneralError: NewGeneralError(OBJECT_NOT_FOUND, nil, source, nil),
	}
}
