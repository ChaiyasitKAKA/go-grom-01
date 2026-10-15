package main

import (
	"fmt"

	"gorm.io/gorm"
)

// --- Create ---
func createProductDB(db *gorm.DB, product *Product) error {
	if product.Price < 0 || product.Quantity < 0 {
		return fmt.Errorf("price and quantity must not be negative")
	}
	return db.Create(product).Error
}

// --- Read / Queries ---
func getAllProductsDB(db *gorm.DB) ([]Product, error) {
	var products []Product
	err := db.Preload("Category").Find(&products).Error
	return products, err
}

func getProductByIDDB(db *gorm.DB, id uint) (*Product, error) {
	var product Product
	err := db.Preload("Category").First(&product, id).Error
	return &product, err
}

func getProductByNameDB(db *gorm.DB, name string) ([]Product, error) {
	var products []Product
	err := db.Where("name LIKE ?", "%"+name+"%").Find(&products).Error
	return products, err
}

func getProductsByCategoryDB(db *gorm.DB, categoryID uint) ([]Product, error) {
	var products []Product
	err := db.Where("category_id = ?", categoryID).Find(&products).Error
	return products, err
}

func getProductsByPriceRangeDB(db *gorm.DB, minPrice, maxPrice float64) ([]Product, error) {
	var products []Product
	err := db.Where("price BETWEEN ? AND ?", minPrice, maxPrice).Find(&products).Error
	return products, err
}

func getProductsLowStockDB(db *gorm.DB, maxStock int) ([]Product, error) {
	var products []Product
	err := db.Where("quantity < ?", maxStock).Find(&products).Error
	return products, err
}

func getProductsByCategoryAndPriceDB(db *gorm.DB, categoryID uint, minPrice, maxPrice float64) ([]Product, error) {
	var products []Product
	err := db.Where("category_id = ? AND price BETWEEN ? AND ?", categoryID, minPrice, maxPrice).Find(&products).Error
	return products, err
}

func getProductsSortedByPriceDB(db *gorm.DB, order string) ([]Product, error) {
	var products []Product
	if order != "ASC" && order != "DESC" {
		order = "ASC"
	}
	err := db.Order("price " + order).Find(&products).Error
	return products, err
}

func getTopExpensiveProductsDB(db *gorm.DB, limit int) ([]Product, error) {
	var products []Product
	err := db.Order("price DESC").Limit(limit).Find(&products).Error
	return products, err
}

// --- Update ---
func updateProductByIDDB(db *gorm.DB, id uint, updatedData map[string]interface{}) error {
	if price, ok := updatedData["price"].(float64); ok && price < 0 {
		return fmt.Errorf("price cannot be negative")
	}
	if qty, ok := updatedData["quantity"].(int); ok && qty < 0 {
		return fmt.Errorf("quantity cannot be negative")
	}
	return db.Model(&Product{}).Where("id = ?", id).Updates(updatedData).Error
}

func updateProductPriceDB(db *gorm.DB, id uint, newPrice float64) error {
	if newPrice < 0 {
		return fmt.Errorf("price cannot be negative")
	}
	return db.Model(&Product{}).Where("id = ?", id).Update("price", newPrice).Error
}

func updateProductStockDB(db *gorm.DB, id uint, amount int) error {
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		return err
	}
	newStock := product.Quantity + amount
	if newStock < 0 {
		return fmt.Errorf("resulting stock cannot be negative")
	}
	return db.Model(&product).Update("quantity", newStock).Error
}

// --- Delete ---
func softDeleteProductDB(db *gorm.DB, id uint) error {
	return db.Delete(&Product{}, id).Error
}

func hardDeleteProductDB(db *gorm.DB, id uint) error {
	return db.Unscoped().Delete(&Product{}, id).Error
}

func getDeletedProductsDB(db *gorm.DB) ([]Product, error) {
	var products []Product
	err := db.Unscoped().Where("deleted_at IS NOT NULL").Find(&products).Error
	return products, err
}

func restoreProductDB(db *gorm.DB, id uint) error {
	return db.Unscoped().Model(&Product{}).Where("id = ?", id).Update("deleted_at", nil).Error
}
