package models

type FoodMenuResponse struct {
	Id          int     `json: "Id"`
	FoodName    string  `json:"FoodName"`
	Price       float64 `json:"Price"`
	Description string  `json: "Description"`
	ImageURL    string  `json :"FoodURL"`
	CreatedAt   string  `json:"CreatedAt"`
	UpdatedAt   string  `json:"UpdatedAt"`
	DeletedAt   *string `json:"DeletedAt"`
}
