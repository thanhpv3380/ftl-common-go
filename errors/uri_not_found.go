package errors

import (
	"github.com/thanhpv3380/ftl-common-go/common"
)

type UriNotFound struct {
	*common.GeneralError
}

func NewUriNotFound(source string) *UriNotFound {
	return &UriNotFound{
		GeneralError: NewGeneralError(string(common.URI_NOT_FOUND), nil, source, nil),
	}
}
