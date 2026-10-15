package main

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Create Order พร้อม Transaction (ตรวจสอบสต็อก คำนวณราคารวม และหักสต็อก)
func createOrderDB(db *gorm.DB, customerID uint, productID uint, quantity int) (*Order, error) {
	var order Order

	err := db.Transaction(func(tx *gorm.DB) error {
		// 1. ตรวจสอบลูกค้า
		var customer Customer
		if err := tx.First(&customer, customerID).Error; err != nil {
			return fmt.Errorf("customer not found: %v", err)
		}

		// 2. ตรวจสอบสินค้า
		var product Product
		if err := tx.First(&product, productID).Error; err != nil {
			return fmt.Errorf("product not found: %v", err)
		}

		// 3. ตรวจสอบสต็อก
		if product.Quantity < quantity {
			return fmt.Errorf("insufficient stock for product %s (available: %d, requested: %d)", product.Name, product.Quantity, quantity)
		}

		// 4. คำนวณ TotalPrice
		totalPrice := product.Price * float64(quantity)

		// 5. ตัดสต็อกสินค้า
		product.Quantity -= quantity
		if err := tx.Save(&product).Error; err != nil {
			return fmt.Errorf("failed to update product stock: %v", err)
		}

		// 6. บันทึกคำสั่งซื้อ
		order = Order{
			CustomerID: customerID,
			ProductID:  productID,
			Quantity:   quantity,
			TotalPrice: totalPrice,
			Status:     "Pending",
		}
		if err := tx.Create(&order).Error; err != nil {
			return fmt.Errorf("failed to create order: %v", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return &order, nil
}

// Query Orders
func getOrdersByCustomerDB(db *gorm.DB, customerID uint) ([]Order, error) {
	var orders []Order
	err := db.Where("customer_id = ?", customerID).Find(&orders).Error
	return orders, err
}

func getOrdersByStatusDB(db *gorm.DB, status string) ([]Order, error) {
	var orders []Order
	err := db.Where("status = ?", status).Find(&orders).Error
	return orders, err
}

func getOrdersByDateRangeDB(db *gorm.DB, startDate, endDate time.Time) ([]Order, error) {
	var orders []Order
	err := db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&orders).Error
	return orders, err
}

func getOrdersWithDetailsDB(db *gorm.DB) ([]Order, error) {
	var orders []Order
	err := db.Preload("Customer").Preload("Product").Preload("Product.Category").Find(&orders).Error
	return orders, err
}

// Update Order Status
func updateOrderStatusDB(db *gorm.DB, id uint, status string) error {
	validStatuses := map[string]bool{"Pending": true, "Completed": true, "Cancelled": true}
	if !validStatuses[status] {
		return fmt.Errorf("invalid status '%s'. Status must be Pending, Completed, or Cancelled", status)
	}
	return db.Model(&Order{}).Where("id = ?", id).Update("status", status).Error
}
