package event

type OrderCompletedEvent struct {
	Products []*ProductInfo `json:"products"`
}

type ProductInfo struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}
