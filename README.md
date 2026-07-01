# E-Commerce API Project using Golang 

## Goals
- Design a public API for an ecommerce website using golang to server our buyers
- Server to be scalable and handle concurrent users

## ARCHITECT AND DESIGN

    user -> API -> DB (pg or mongo)
- ###  SCHEMA 

    product table ( id , name , price , qty)

    orders (order_id, Cust_id, created_at, status)

    order_items (product_id, Order_id, qty, total_price)

- ### API (v1)

    - GET /health 
    - GET /product?name
    - POST /orders 
