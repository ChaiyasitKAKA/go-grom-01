package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Category: ประเภทสินค้า (1:N กับ Product)
type Category struct {
	gorm.Model
	Name        string `gorm:not null`
	Description string
	Products    []Product `gorm:"foreignKey:CategoryID"`
}

// Product: สินค้า (N:1 กับ Category, 1:N กับ Order)
type Product struct {
	gorm.Model
	Name       string  `gorm:not null`
	Price      float64 `gorm:not null`
	Quantity   int     `gorm:not null`
	CategoryID uint    `gorm:not null`
	Category   Category
	Orders     []Order `gorm:"foreignKey:ProductID"`
}

// Customer: ลูกค้า (1:N กับ Order)
type Customer struct {
	gorm.Model
	Name   string `gorm:not null`
	Email  string `gorm:unique;not null`
	Phone  string
	Orders []Order `gorm:"foreignKey:CustomerID"`
}

// Order: คำสั่งซื้อ (N:1 กับ Customer และ Product)
type Order struct {
	gorm.Model
	CustomerID uint `gorm:not null`
	Customer   Customer
	ProductID  uint `gorm:not null`
	Product    Product
	Quantity   int     `gorm:not null`
	TotalPrice float64 `gorm:not null`
	Status     string  `gorm:default:'Pending'` // Pending / Completed / Cancelled
}

func connectDB() *gorm.DB {
	// ปรับ DSN ตาม User, Password และ Host ของคุณ (ตัวอย่าง: root:@tcp(127.0.0.1:3306)/store_db?charset=utf8mb4&parseTime=True&loc=Local)
	dsn := "root:@tcp(127.0.0.1:3306)/store_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connected successfully!")
	return db
}

func main() {
	db := connectDB()

	// ทำการ AutoMigrate สร้างตารางทั้ง 4
	err := db.AutoMigrate(&Category{}, &Product{}, &Customer{}, &Order{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	fmt.Println("Database migration completed successfully!")

}
