package errors

import (
	"fmt"

	"github.com/thanhpv3380/ftl-common-go/common"
)

type ForwardError struct {
	Status         common.Status
	IsForwardError bool
}

func (e *ForwardError) Error() string {
	return fmt.Sprintf(e.Status.Code)
}

func NewForwardError(status common.Status) *ForwardError {
	return &ForwardError{
		Status:         status,
		IsForwardError: true,
	}
}
