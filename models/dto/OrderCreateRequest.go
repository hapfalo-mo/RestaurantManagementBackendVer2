package models

type OrderCreateRequest struct {
	UserId     int     `json:"UserId" binding:"required"`
	TotalPrice float64 `json:"TotalPrice" binding:"required"`
}
