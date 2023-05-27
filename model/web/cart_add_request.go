package web

type CartAddRequest struct {
	CartId    int `json:"cartId"`
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type CartUpdateRequest struct {
	CartItemId int `json:"cartItemId"`
	CartId     int `json:"cartId"`
	ProductId  int `json:"productId"`
	Quantity   int `json:"quantity"`
}
