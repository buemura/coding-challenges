package order

type Item struct {
	Product  string  `json:"product"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
}

type Order struct {
	OrderID    int     `json:"orderId"`
	CustomerID int     `json:"customerId"`
	Items      []*Item `json:"items"`
}
