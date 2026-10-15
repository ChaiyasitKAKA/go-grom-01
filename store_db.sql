-- Active: 1792079349380@@127.0.0.1@3306@store_db
-- สร้างและเลือกฐานข้อมูล
CREATE DATABASE IF NOT EXISTS store_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE store_db;

-- เคลียร์ตารางเดิม (ถ้ามี) ตามลำดับ Foreign Key ป้องกัน Error
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS categories;

-- 1. สร้างตาราง Categories
CREATE TABLE categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

-- 2. สร้างตาราง Products
CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    name VARCHAR(255) NOT NULL,
    price DOUBLE NOT NULL,
    quantity INT NOT NULL,
    category_id INT NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- 3. สร้างตาราง Customers
CREATE TABLE customers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(50)
);

-- 4. สร้างตาราง Orders
CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    customer_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    total_price DOUBLE NOT NULL,
    status VARCHAR(50) DEFAULT 'Pending',
    FOREIGN KEY (customer_id) REFERENCES customers(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

-- ==========================================
-- ข้อมูลทดสอบขั้นต่ำ 
-- ==========================================

-- แทรก Category (4 รายการ)
INSERT INTO categories (id, name, description) VALUES
(1, 'Laptops & Computers', 'Notebooks and Desktop PCs'),
(2, 'Smartphones & Tablets', 'Mobile devices and tablets'),
(3, 'Accessories', 'Mouse, Keyboard, Hubs, etc.'),
(4, 'Monitors & Displays', 'Screens and displays');

-- แทรก Product (15 รายการ)
INSERT INTO products (id, name, price, quantity, category_id) VALUES
(1, 'MacBook Air M2', 39900.0, 9, 1),       -- โดนสั่งซื้อไป 1
(2, 'ASUS Vivobook 15', 18900.0, 14, 1),    -- โดนสั่งซื้อไป 1
(3, 'Lenovo IdeaPad Slim 3', 14500.0, 18, 1),-- โดนสั่งซื้อไป 2
(4, 'Dell XPS 13', 45900.0, 5, 1),
(5, 'MSI Gaming Laptop', 32900.0, 7, 1),     -- โดนสั่งซื้อไป 1
(6, 'iPhone 16', 32900.0, 24, 2),           -- โดนสั่งซื้อไป 1
(7, 'Samsung Galaxy S24', 29900.0, 17, 2),  -- โดนสั่งซื้อไป 1
(8, 'Xiaomi 14', 22900.0, 12, 2),
(9, 'iPad Air', 23900.0, 9, 2),             -- โดนสั่งซื้อไป 1
(10, 'Logitech Mouse MX Master 3S', 3900.0, 28, 3), -- โดนสั่งซื้อไป 2
(11, 'Mechanical Keyboard', 2500.0, 24, 3), -- โดนสั่งซื้อไป 1
(12, 'USB-C Hub 7-in-1', 990.0, 47, 3),     -- โดนสั่งซื้อไป 3
(13, 'Gaming Headset', 1800.0, 18, 3),      -- โดนสั่งซื้อไป 2
(14, 'Dell 27\" 4K Monitor', 12900.0, 6, 4), -- โดนสั่งซื้อไป 1
(15, 'LG UltraGear Gaming Monitor', 8900.0, 10, 4);

-- แทรก Customer (8 รายการ)
INSERT INTO customers (id, name, email, phone) VALUES
(1, 'Somchai Jaidee', 'somchai@email.com', '0811111111'),
(2, 'Suda Raksa', 'suda@email.com', '0822222222'),
(3, 'Anonda Manee', 'anonda@email.com', '0833333333'),
(4, 'Kanya Wong', 'kanya@email.com', '0844444444'),
(5, 'Prasert Chai', 'prasert@email.com', '0855555555'),
(6, 'Nicha Suk', 'nicha@email.com', '0866666666'),
(7, 'Arthit Daeng', 'arthit@email.com', '0877777777'),
(8, 'Malee Boon', 'malee@email.com', '0888888888');

-- แทรก Order (12 รายการ - คำนวณ TotalPrice และตัดสต็อกใน Product เรียบร้อย)
INSERT INTO orders (customer_id, product_id, quantity, total_price, status) VALUES
(1, 1, 1, 39900.0, 'Completed'),
(1, 10, 2, 7800.0, 'Completed'),
(2, 6, 1, 32900.0, 'Completed'),
(2, 12, 3, 2970.0, 'Pending'),
(3, 2, 1, 18900.0, 'Completed'),
(4, 7, 1, 29900.0, 'Pending'),
(5, 11, 1, 2500.0, 'Completed'),
(6, 14, 1, 12900.0, 'Completed'),
(7, 3, 2, 29000.0, 'Completed'),
(8, 9, 1, 23900.0, 'Pending'),
(3, 13, 2, 3600.0, 'Completed'),
(5, 5, 1, 32900.0, 'Completed');