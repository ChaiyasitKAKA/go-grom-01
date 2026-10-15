package main

import (
	"fmt"

	"gorm.io/gorm"
)

func createCategoryDB(db *gorm.DB, name string, description string) uint {
	category := Category{
		Name:        name,
		Description: description,
	}

	result := db.Create(&category)
	if result.Error != nil {
		fmt.Printf("Failed to create category: %v\n", result.Error)
		return 0
	}

	fmt.Printf("Category created: %s (ID: %d)\n", category.Name, category.ID)
	return category.ID
}
