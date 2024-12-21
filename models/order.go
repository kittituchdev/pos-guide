package models

import (
	"context"
	"fmt"
	"time"

	"github.com/kittituchdev/pos-guide/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	OrderNumber string             `bson:"orderNumber" json:"orderNumber"`
	TotalPrice  float64            `bson:"totalPrice" json:"totalPrice"`
	Items       []OrderItem        `bson:"items" json:"items"`
	IsActive    bool               `bson:"isActive" json:"isActive"`
	IsDelete    bool               `bson:"isDelete" json:"isDelete"`
	CreatedAt   int64              `bson:"createdAt" json:"createdAt"` // Milliseconds since epoch
	CreatedBy   string             `bson:"createdBy" json:"createdBy"`
	UpdatedAt   int64              `bson:"updatedAt" json:"updatedAt"` // Milliseconds since epoch
	UpdatedBy   string             `bson:"updatedBy" json:"updatedBy"`
}

type OrderItem struct {
	ProductID primitive.ObjectID `bson:"productID" json:"productID"`
	Quantity  int                `bson:"quantity" json:"quantity"`
	Price     float64            `bson:"price" json:"price"`
	Total     float64            `bson:"total" json:"total"`
}

var orderCollectionName = "orders"

func InsertOneOrder(order Order) error {
	// Context with timeout to avoid long waits
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the collection
	collection := config.MongoClient.Database(config.DatabaseName).Collection(orderCollectionName)

	// Insert the order
	result, err := collection.InsertOne(ctx, order)
	if err != nil {
		return fmt.Errorf("failed to insert order: %v", err) // Return error instead of log.Fatal
	}

	// Log inserted ID
	fmt.Println("Inserted a record with ID:", result.InsertedID)
	return nil
}
