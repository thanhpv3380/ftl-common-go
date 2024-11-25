package utils

import (
	"encoding/json"

	"github.com/thanhpv3380/ftl-common-go/common"
	"github.com/thanhpv3380/ftl-common-go/errors"
)

func ExtractResponse(data interface{}) (*interface{}, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("")
	}

	var response common.MessageResponse
	if err := json.Unmarshal(jsonData, &response); err != nil {
		return nil, errors.NewError("")
	}

	if response.Status.Code != "" {
		return nil, &response.Status
	}

	return &response.Data, nil
}
