package request

import "github.com/google/uuid"

type AddProduct struct {
	Id       int       `json:"id"`
	IdSeller uuid.UUID `json:"id_seller"`
	Name     string    `json:"name"`
	Price    int       `json:"price"`
}

type UpdateProduct struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
