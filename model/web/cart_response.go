package web

type CartResponse struct {
	Id        int `json:"id"`
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}
