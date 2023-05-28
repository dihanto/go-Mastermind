package web

type CartAddRequest struct {
	CartId    int `validate:"required,numeric" json:"cartId"`
	ProductId int `validate:"required,numeric" json:"productId"`
	Quantity  int `validate:"required,numeric" json:"quantity"`
}

type CartUpdateRequest struct {
	CartItemId int `validate:"required,numeric" json:"cartItemId"`
	CartId     int `validate:"required,numeric" json:"cartId"`
	ProductId  int `json:"productId"`
	Quantity   int `json:"quantity"`
}
