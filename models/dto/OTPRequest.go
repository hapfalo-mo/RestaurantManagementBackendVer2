package models

type OTPRequest struct {
	OTP       string `json:"otp"`
	UserEmail string `json:"userEmail"`
}
