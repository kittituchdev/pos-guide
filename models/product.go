package models

import (
	"context"
	"fmt"
	"time"

	"github.com/kittituchdev/pos-guide/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product struct for database model
type Product struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string               `bson:"name" json:"name"`
	Description string               `bson:"description" json:"description"`
	Price       float64              `bson:"price" json:"price"`
	Stock       int                  `bson:"stock" json:"stock"`
	Images      []string             `bson:"images" json:"images"`
	Options     []primitive.ObjectID `bson:"options" json:"options"`
	Categories  []primitive.ObjectID `bson:"categories" json:"categories"`
	IsActive    bool                 `bson:"isActive" json:"isActive"`
	IsDelete    bool                 `bson:"isDelete" json:"isDelete"`
	CreatedAt   int64                `bson:"createdAt" json:"createdAt"`
	CreatedBy   string               `bson:"createdBy" json:"createdBy"`
	UpdatedAt   int64                `bson:"updatedAt" json:"updatedAt"`
	UpdatedBy   string               `bson:"updatedBy" json:"updatedBy"`
}

// UpdateProductInput struct for partial updates
type UpdateProductInput struct {
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Price       *float64  `json:"price,omitempty"`
	Stock       *int      `json:"stock,omitempty"`
	Images      *[]string `json:"images,omitempty"`
	Options     *[]primitive.ObjectID `json:"options,omitempty"`
	Categories  *[]primitive.ObjectID `json:"categories,omitempty"`
	IsActive    *bool     `json:"isActive,omitempty"`
	IsDelete    *bool     `json:"isDelete,omitempty"`
	UpdatedBy   *string   `json:"updatedBy,omitempty"`
}

var collectionName = "products"

func InsertOneProduct(product Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.MongoClient.Database(config.DatabaseName).Collection(collectionName)
	result, err := collection.InsertOne(ctx, product)
	if err != nil {
		return fmt.Errorf("failed to insert product: %v", err)
	}

	fmt.Println("Inserted a record with ID:", result.InsertedID)
	return nil
}

func FindAllProduct() ([]Product, error) {
	var results []Product

	collection := config.MongoClient.Database(config.DatabaseName).Collection(collectionName)
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

func UpdateProduct(productId string, updates UpdateProductInput) error {
	id, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	updateFields := bson.M{}

	// Dynamically build update fields based on provided input
	if updates.Name != nil {
		updateFields["name"] = *updates.Name
	}
	if updates.Description != nil {
		updateFields["description"] = *updates.Description
	}
	if updates.Price != nil {
		updateFields["price"] = *updates.Price
	}
	if updates.Stock != nil {
		updateFields["stock"] = *updates.Stock
	}
	if updates.Images != nil {
		updateFields["images"] = *updates.Images
	}
	if updates.Options != nil {
		updateFields["options"] = *updates.Options
	}
	if updates.Categories != nil {
		updateFields["categories"] = *updates.Categories
	}
	if updates.IsActive != nil {
		updateFields["isActive"] = *updates.IsActive
	}
	if updates.IsDelete != nil {
		updateFields["isDelete"] = *updates.IsDelete
	}
	if updates.UpdatedBy != nil {
		updateFields["updatedBy"] = *updates.UpdatedBy
	}
	updateFields["updatedAt"] = time.Now().UnixMilli()

	update := bson.M{"$set": updateFields}

	collection := config.MongoClient.Database(config.DatabaseName).Collection(collectionName)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
	return nil
}
