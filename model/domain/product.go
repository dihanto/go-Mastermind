package domain

type Product struct {
	Id         int
	Name       string
	Price      int
	Quantity   int
	CategoryId int
}

type Category struct {
	Id   int
	Name string
}

type Cart struct {
	ProductId int
	Quantity  int
}
