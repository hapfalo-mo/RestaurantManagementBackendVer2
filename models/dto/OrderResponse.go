package models

type OrderResponse struct {
	Id          int     `json:"OrderId"`
	OrderStatus int     `json:"OrderStatus"`
	TotalPrice  float64 `json:"TotalPrice"`
	OrderedAt   string  `json:"OrderedAt"`
	UpdatedAt   string  `json:"UpdatesAt"`
	DeletedAt   *string `json:"DeletedAt"`
	Note        *string `json:"Note"`
	Feedback    *int    `json:"Feedback"`
}
