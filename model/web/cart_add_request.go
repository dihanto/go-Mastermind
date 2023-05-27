package web

type CartAddRequest struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}
