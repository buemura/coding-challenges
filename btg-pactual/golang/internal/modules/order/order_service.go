package order

import (
	"context"
	"fmt"

	"github.com/buemura/btg-challenge/internal/infra/database"
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
		fmt.Println(order)
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

func (s *OrderService) getCustomerOrdersCount(customerID int) (*OrdersQuantityOut, error) {
	count, err := s.OrderColl.CountDocuments(context.Background(), bson.M{"customerId": customerID})
	if err != nil {
		return nil, err
	}

	return &OrdersQuantityOut{
		Quantity: int(count),
	}, nil
}

func (s *OrderService) getOrderTotalPrice(orderID int) (*OrderTotalPriceOut, error) {
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "orderid", Value: orderID},
		}},
	}

	unwindStage := bson.D{
		{Key: "$unwind", Value: "$items"},
	}

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$_id"},
			{Key: "totalPrice", Value: bson.D{
				{Key: "$sum", Value: "$items.price"},
			}},
		}},
	}

	cursor, err := s.OrderColl.Aggregate(context.Background(), mongo.Pipeline{matchStage, unwindStage, groupStage})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var result struct {
		ID         interface{} `bson:"_id"`
		TotalPrice float64     `bson:"totalPrice"`
	}

	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&result); err != nil {
			return nil, err
		}
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}
	return &OrderTotalPriceOut{
		TotalPrice: float32(result.TotalPrice),
	}, nil
}
