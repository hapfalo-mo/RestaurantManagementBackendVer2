package models

type LoginGoogleRequest struct {
	Email    string `json: "email" binding:"requried"`
	IsVerify bool   `json: "isVerify" binding:"requried"`
}
