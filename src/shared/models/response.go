package models

// Response type
type Response struct {
	HasError bool   `json:"hasError"`
	Error    *Error `json:"error,omitempty"`
}
