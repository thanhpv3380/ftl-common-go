package main

type ParamError struct {
	Code          string                 `json:"code"`
	MessageParams map[string]interface{} `json:"messageParams,omitempty"`
}

type Status struct {
	Code          string                  `json:"code"`
	MessageParams map[string]interface{}  `json:"messageParams,omitempty"`
	Params        map[string][]ParamError `json:"params,omitempty"`
}

type Response struct {
	Status *Status     `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
