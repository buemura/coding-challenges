package order

type Item struct {
	Product  string  `json:"product"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Order struct {
	OrderID     int     `json:"orderId" bson:"orderId"`
	CustomerID  int     `json:"customerId" bson:"customerId"`
	TotalAmount float64 `json:"totalAmount" bson:"totalAmount"`
	Items       []*Item `json:"items"`
}
