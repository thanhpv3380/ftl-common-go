package errors

import (
	"ftl/kafi-common/common/interfaces"
	"ftl/kafi-common/common/errorCodes"
)

type UriNotFound struct {
	*GeneralError
}

func NewUriNotFound(source string) *UriNotFound {
	return &UriNotFound{
		GeneralError: NewGeneralError(URI_NOT_FOUND, nil, source, nil),
	}
}
