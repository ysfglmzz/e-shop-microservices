package dto

type CreateProductDTO struct {
	Name          string `json:"name"`
	BrandId       int    `json:"brandId"`
	CategoryId    int    `json:"categoryId"`
	SubCategoryId int    `json:"subCategoryId"`
	Quantity      int    `json:"quantity"`
	UnitPrice     int    `json:"unitPrice"`
}
