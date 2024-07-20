package order

import (
	"context"

	"github.com/buemura/btg-challenge/infra/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderService struct {
	OrderColl *mongo.Collection
}

func NewOrderService() *OrderService {
	return &OrderService{
		OrderColl: database.OrderColl,
	}
}

func (s *OrderService) InsertOrder(order *Order) error {
	_, err := s.OrderColl.InsertOne(context.Background(), order)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderService) getCustomerOrderList(customerID int) ([]*Order, error) {
	cursor, err := s.OrderColl.Find(context.Background(), bson.M{"customerId": customerID})
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

	return orders, nil
}
