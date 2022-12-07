package dto

type AddProductToBasketRequest struct {
	BaketId string   `json:"basketId"`
	Product *Product `json:"product"`
}

type Product struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	UnitPrice int    `json:"unitPrice"`
}
