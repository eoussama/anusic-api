package models

// Response type
type Response struct {
	Status   int         `json:"status"`
	HasError bool        `json:"hasError"`
	Error    *Error      `json:"error,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}
