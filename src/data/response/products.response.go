package response

type ProductsResponse struct {
	Id int `json:"id"`
	Product_Name string `json:"product_name"`
	Description string `json:"description"`
}