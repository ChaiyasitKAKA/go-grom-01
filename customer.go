package main

import (
	"fmt"

	"gorm.io/gorm"
)

func createCustomerDB(db *gorm.DB, name string, email string, phone string) uint {
	customer := Customer{
		Name:  name,
		Email: email,
		Phone: phone,
	}

	result := db.Create(&customer)
	if result.Error != nil {
		fmt.Printf("Failed to create customer: %v\n", result.Error)
		return 0
	}

	fmt.Printf("Customer created: %s (%s)\n", customer.Name, customer.Email)
	return customer.ID
}
