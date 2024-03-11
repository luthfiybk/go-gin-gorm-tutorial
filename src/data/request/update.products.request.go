package request

type UpdateProductsRequest struct {
	Id int `validate:"required" json:"id"`
	Product_Name string `validate:"required,min=1,max=255" json:"product_name"`
	Description string `validate:"required,min=1,max=255" json:"description"`
}