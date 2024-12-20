package models

import (
	"context"
	"fmt"
	"log"

	"github.com/kittituchdev/pos-guide/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          string  `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string  `bson:"name" json:"name"`
	Description string  `bson:"description" json:"description"`
	Price       float64 `bson:"price" json:"price"`
	Stock       int     `bson:"stock" json:"stock"`
	CreatedAt   int64   `bson:"createdAt" json:"createdAt"` // Milliseconds since epoch
	UpdatedAt   int64   `bson:"updatedAt" json:"updatedAt"` // Milliseconds since epoch
}

var collectionName = "products"

func InsertOneProduct(product Product) error {
	collection := config.MongoClient.Database(config.DatabaseName).Collection(collectionName)
	result, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insert a record with id: ", result.InsertedID)
	return err
}

func InsertManyProducts(products []Product) error {
	// Convert to slice of interface{}
	newProducts := make([]interface{}, len(products))
	for i, product := range products {
		newProducts[i] = product
	}

	collection := config.MongoClient.Database(config.DatabaseName).Collection(collectionName)
	result, err := collection.InsertMany(context.TODO(), newProducts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	return err
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
		"createdAt":   product.CreatedAt,
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

func FindProductByName(productName string) []Product {
	var results []Product

	filter := bson.D{{Key: "name", Value: productName}}

	collection := config.MongoClient.Database(config.DatabaseName).Collection(collectionName)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func FindAllProduct() []Product {
	var results []Product

	collection := config.MongoClient.Database(config.DatabaseName).Collection(collectionName)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}
