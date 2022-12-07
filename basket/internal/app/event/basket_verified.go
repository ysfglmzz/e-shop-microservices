package event

type BasketVerified struct {
	UserId    int        `json:"userId"`
	Products  []*Product `json:"products"`
	ItemCount int        `json:"itemCount"`
}

type Product struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	UnitPrice int    `json:"unitPrice"`
}
