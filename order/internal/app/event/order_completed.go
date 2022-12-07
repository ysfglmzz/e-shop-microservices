package event

type OrderCompleted struct {
	UserId   int            `json:"userId"`
	Products []*ProductInfo `json:"products"`
}

type ProductInfo struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}
