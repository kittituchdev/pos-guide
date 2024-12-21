package models

import (
	"context"
	"fmt"
	"time"

	"github.com/kittituchdev/pos-guide/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Option struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Price       float64            `bson:"price" json:"price"`
	OptionGroup primitive.ObjectID `bson:"optionGroup" json:"optionGroup"`
	IsActive    bool               `bson:"isActive" json:"isActive"`
	IsDelete    bool               `bson:"isDelete" json:"isDelete"`
	CreatedAt   int64              `bson:"createdAt" json:"createdAt"` // Milliseconds since epoch
	CreatedBy   string             `bson:"createdBy" json:"createdBy"`
	UpdatedAt   int64              `bson:"updatedAt" json:"updatedAt"` // Milliseconds since epoch
	UpdatedBy   string             `bson:"updatedBy" json:"updatedBy"`
}

var optionCollectionName = "options"

func InsertOneOption(option Option) error {
	// Context with timeout to avoid long waits
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the collection
	collection := config.MongoClient.Database(config.DatabaseName).Collection(optionCollectionName)

	// Insert the option
	result, err := collection.InsertOne(ctx, option)
	if err != nil {
		return fmt.Errorf("failed to insert option: %v", err) // Return error instead of log.Fatal
	}

	// Log inserted ID
	fmt.Println("Inserted a record with ID:", result.InsertedID)
	return nil
}

func FindAllOption() ([]Option, error) {
	var results []Option

	collection := config.MongoClient.Database(config.DatabaseName).Collection(optionCollectionName)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
