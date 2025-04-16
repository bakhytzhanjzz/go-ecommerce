# 🛒 Microservices-Based E-Commerce Platform

A simple e-commerce system developed as part of the **Advanced Programming II – Assignment 1**, designed using **Clean Architecture** principles. The platform consists of **three independent microservices** that interact via REST APIs.

> **👤 Student:** Bakhytzhan Abdilmazhit  
> **🏫 Institution:** Astana IT University  
> **📘 Course:** Advanced Programming II  
> **📝 Assignment:** Clean Architecture & Microservices Implementation

---

## 🎯 Project Goal

Create a minimal yet functional e-commerce solution composed of three decoupled services:

1. **API Gateway (Gin)** – Manages routing, request logging, telemetry, and future authentication.
2. **Inventory Service (Gin + DB)** – Handles product data, categories, and stock levels.
3. **Order Service (Gin + DB)** – Manages order creation, status updates, and payment tracking.

---

## 🧱 Architecture Overview

Each microservice adheres to the **Clean Architecture** pattern, organizing code by:

- **Domain logic**
- **Infrastructure**
- **Interfaces**
- **Delivery mechanisms**

---

## 📦 Inventory Service

Responsible for managing **products**, **categories**, and **stock availability**.

### 🔧 Key Features

- Full CRUD support for products and categories
- Product listing with pagination and filters

### 🔌 API Endpoints

| Method | Route                | Description              |
|--------|----------------------|--------------------------|
| POST   | `/products`          | Add a new product        |
| GET    | `/products/:id`      | Get product by ID        |
| PATCH  | `/products/:id`      | Update product details   |
| DELETE | `/products/:id`      | Remove a product         |
| GET    | `/products`          | List available products  |

---

## 🧾 Order Service

Manages the lifecycle of **orders**, including creation, updates, and payments.

### 🔧 Key Features

- Links orders to products and quantities
- Supports order status tracking

### 🔌 API Endpoints

| Method | Route                | Description                   |
|--------|----------------------|-------------------------------|
| POST   | `/orders`            | Place a new order             |
| GET    | `/orders/:id`        | Get order details             |
| PATCH  | `/orders/:id`        | Update order status           |
| GET    | `/orders`            | View user’s order history     |

---

## 🚪 API Gateway

Serves as the unified access point for all services. Responsibilities include:

- Routing traffic to Inventory and Order services
- Centralized logging and telemetry
- Placeholder for authentication middleware

---

## 🛠️ Tech Stack

- **Go + Gin** – Backend framework for all services
- **Database** – Any DB engine (PostgreSQL, MongoDB, etc.)
- **Docker (optional)** – Containerization of services
- **REST** – API communication standard

---

## ▶️ Getting Started

Each service runs independently. For example, to launch the Inventory Service:

```bash
cd inventory-service
go run main.go
```

---

### 🎷 Bonus

Fun fact: I love jazz. So while your orders are processing and logs are flying by—maybe cue up some Coltrane or Miles Davis in the background.
