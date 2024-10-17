package models

type OrderMessage struct {
	OrderID  string          `json:"orderId"`
	ClientID string          `json:"clientId"`
	Products []ProductDetail `json:"products"`
}

type ProductDetail struct {
	ProductID string  `json:"productId"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
}
