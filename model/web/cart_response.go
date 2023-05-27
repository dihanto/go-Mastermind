package web

type CartResponse struct {
	CartItemId int `json:"cartItemId"`
	CartId     int `json:"cartId"`
	ProductId  int `json:"productId"`
	Quantity   int `json:"quantity"`
}
