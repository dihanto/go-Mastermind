package domain

type Product struct {
	ProductId  int
	Name       string
	Price      int
	CategoryId int
}

type Category struct {
	Id   int
	Name string
}

type Cart struct {
	CartId int
	UserId string
}

type CartItem struct {
	CartItemId int
	CartId     int
	ProductId  int
	Quantity   int
}

type User struct {
	Id       string
	Username string
}
