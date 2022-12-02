package dto

type ProductResponse struct {
	Id          int                `json:"id"`
	Name        string             `json:"name"`
	Brand       ProductBrand       `json:"brand" gorm:"embedded;embeddedPrefix:b_"`
	Category    ProductCategory    `json:"category" gorm:"embedded;embeddedPrefix:c_"`
	SubCategory ProductSubCategory `json:"subCategory" gorm:"embedded;embeddedPrefix:sc_"`
	Quantity    int                `json:"quantity"`
	UnitPrice   int                `json:"unitPrice"`
}

type ProductBrand struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProductCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProductSubCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
