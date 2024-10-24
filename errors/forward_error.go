package errors

import (
	"fmt"
	"ftl/kafi-common/common"
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
