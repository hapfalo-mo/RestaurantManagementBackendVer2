package models

type OrderItemRequest struct {
	OrderId  int     `json:"OrderId"`
	FoodId   int     `json: "FoodId"`
	Quantity int     `json:"Quantity"`
	Price    float64 `json:"Price"`
}
