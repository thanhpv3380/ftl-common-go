package errors

import (
	"ftl/kafi-common/common/interfaces"
)

type ForwardError struct {
	Status         Status
	IsForwardError bool
	Err            error
}

func NewForwardError(status Status) *ForwardError {
	return &ForwardError{
		Status:         status,
		IsForwardError: true,
		Err:            fmt.Errorf("forward error"),
	}
}
