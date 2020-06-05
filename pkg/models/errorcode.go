package models

// ErrorCode ..
type ErrorCode struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// IsError ..
func (e *ErrorCode) IsError() bool {
	return e.Error != ""
}
