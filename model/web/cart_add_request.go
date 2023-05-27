package web

type CartAddOrUpdateRequest struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type CartUpdateRequest struct {
	Id        int `json:"id"`
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}
