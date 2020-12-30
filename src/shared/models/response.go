package models

// Response type
type Response struct {
	HasError bool
	Error    *Error
	Data     interface{}
}
