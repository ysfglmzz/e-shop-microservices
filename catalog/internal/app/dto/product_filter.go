package dto

type ProductFilter struct {
	Limit         *int `form:"limit" binding:"required"`
	Offset        *int `form:"offset" binding:"required"`
	CategoryId    *int `form:"categoryId"`
	SubCategoryId *int `form:"subCategoryId"`
	BrandId       *int `form:"brandId"`
}
