package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateOrderNumber generates an order number in YYMMDDHHmmss+2 random digits
func GenerateOrderNumber() string {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	now := time.Now().In(location)
	// Format the current timestamp as YYMMDDHHmmss
	timestamp := now.Format("060102150405") // YYMMDDHHmmss

	// Generate 2 random digits
	randomNumber := rand.Intn(100) // Generates a number between 0-99

	// Combine timestamp and random number
	return fmt.Sprintf("%s%02d", timestamp, randomNumber)
}
