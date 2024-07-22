package order

import (
	"context"
	"math"

	"github.com/buemura/btg-challenge/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderService struct {
	OrderColl *mongo.Collection
}

func NewOrderService() *OrderService {
	return &OrderService{
		OrderColl: database.OrderColl,
	}
}

func (s *OrderService) InsertOrder(orderIn *OrderCreatedIn) error {
	var totalAmount float64
	for _, i := range orderIn.Items {
		totalAmount += i.Price
	}

	_, err := s.OrderColl.InsertOne(context.Background(), &Order{
		OrderID:     orderIn.OrderID,
		CustomerID:  orderIn.CustomerID,
		TotalAmount: totalAmount,
		Items:       orderIn.Items,
	})

	if err != nil {
		return err
	}
	return nil
}

func (s *OrderService) getCustomerOrderList(in *OrderListIn) (*OrderListOut, error) {
	filter := bson.M{"customerId": in.CustomerID}
	count, err := s.OrderColl.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	limit := int64(in.Items)
	skip := int64((in.Page - 1) * in.Items)
	opts := &options.FindOptions{Limit: &limit, Skip: &skip}
	cursor, err := s.OrderColl.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var orders []*Order
	for cursor.Next(context.Background()) {
		var order Order
		err := cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &OrderListOut{
		Data: orders,
		Meta: &PaginationMeta{
			Page:       in.Page,
			Items:      in.Items,
			TotalPage:  int(math.Ceil(float64(count) / float64(in.Items))),
			TotalItems: int(count),
		},
	}, nil
}
