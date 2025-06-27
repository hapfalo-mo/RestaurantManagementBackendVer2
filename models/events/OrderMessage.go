package events

type OrderMessage struct {
	OrderId  string `json:"orderId"`
	Username string `json:"foodName"`
	Amount   int    `json:"amount"`
	Money    int    `json:"money"`
	Message  string `json:"message"`
}
