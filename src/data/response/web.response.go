package response

type Response struct {
	Code  int    `json:"status"`
	Status string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}