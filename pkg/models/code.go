package models

// Code ..
type Code struct {
	UserCode        string `json:"user_code"`
	DeviceCode      string `json:"device_code"`
	VerificationURI string `json:"verification_uri"`
	ExpiresIn       string `json:"ExpiresIn"`
	Interval        int    `json:"interval"`
	Message         string `json:"message"`
}
