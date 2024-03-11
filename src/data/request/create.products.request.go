package request

type CreateProductsRequest struct {
	Product_Name string `validate:"required,min=1,max=255" json:"product_name"`
	Description string `validate:"required,min=1,max=255" json:"description"`
}