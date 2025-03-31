package models

type OrderDetailResponse struct {
	ID         int     `json:"Id"`
	Name       string  `json:"FoodName"`
	Price      float64 `json:"Price"`
	Quantity   int     `json:"Quantity"`
	TotalPrice float64 `json:"TotalPrice"`
}
