package web

type ProductResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Quantity   int    `json:"quantity"`
	CategoryId int    `json:"categoryId"`
}
