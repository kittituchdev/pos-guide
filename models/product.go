package models

import (
	"context"
	"fmt"
	"time"

	"github.com/kittituchdev/pos-guide/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	CreatedAt   int64                `bson:"createdAt" json:"createdAt"` // Milliseconds since epoch
	CreatedBy   string               `bson:"createdBy" json:"createdBy"`
	UpdatedAt   int64                `bson:"updatedAt" json:"updatedAt"` // Milliseconds since epoch
	UpdatedBy   string               `bson:"updatedBy" json:"updatedBy"`
}

var collectionName = "products"

func InsertOneProduct(product Product) error {
	// Context with timeout to avoid long waits
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the collection
	collection := config.MongoClient.Database(config.DatabaseName).Collection(collectionName)

	// Insert the product
	result, err := collection.InsertOne(ctx, product)
	if err != nil {
		return fmt.Errorf("failed to insert product: %v", err) // Return error instead of log.Fatal
	}

	// Log inserted ID
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

func UpdateProductPrice(productId string, price float64) error {
	id, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"price": price}}

	collection := config.MongoClient.Database(config.DatabaseName).Collection(collectionName)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
	return nil
}

func UpdateProduct(productId string, product Product) error {
	id, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"stock":       product.Stock,
		"updatedAt":   product.UpdatedAt,
	}}

	collection := config.MongoClient.Database(config.DatabaseName).Collection(collectionName)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
	return nil
}

// func InsertManyProducts(products []Product) error {
// 	// Convert to slice of interface{}
// 	newProducts := make([]interface{}, len(products))
// 	for i, product := range products {
// 		newProducts[i] = product
// 	}

// 	collection := config.MongoClient.Database(config.DatabaseName).Collection(collectionName)
// 	result, err := collection.InsertMany(context.TODO(), newProducts)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(result)
// 	return err
// }

// func UpdateProduct(productId string, product Product) error {
// 	id, err := primitive.ObjectIDFromHex(productId)
// 	if err != nil {
// 		return err
// 	}
// 	filter := bson.M{"_id": id}
// 	update := bson.M{"$set": bson.M{
// 		"name":        product.Name,
// 		"description": product.Description,
// 		"price":       product.Price,
// 		"stock":       product.Stock,
// 		"createdAt":   product.CreatedAt,
// 		"updatedAt":   product.UpdatedAt,
// 	}}

// 	collection := config.MongoClient.Database(config.DatabaseName).Collection(collectionName)
// 	result, err := collection.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("Modified count: ", result.ModifiedCount)
// 	return nil
// }

// func FindProductByName(productName string) []Product {
// 	var results []Product

// 	filter := bson.D{{Key: "name", Value: productName}}

// 	collection := config.MongoClient.Database(config.DatabaseName).Collection(collectionName)
// 	cursor, err := collection.Find(context.TODO(), filter)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = cursor.All(context.TODO(), &results)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return results
// }
