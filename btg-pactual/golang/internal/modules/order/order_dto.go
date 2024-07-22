package order

type OrderCreatedIn struct {
	OrderID    int     `json:"orderId"`
	CustomerID int     `json:"customerId"`
	Items      []*Item `json:"items"`
}

type OrderListIn struct {
	CustomerID int
	Page       int
	Items      int
}

type OrderListOut struct {
	Data []*Order        `json:"data"`
	Meta *PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	Page       int `json:"page"`
	Items      int `json:"items"`
	TotalPage  int `json:"totalPage"`
	TotalItems int `json:"totalItems"`
}
