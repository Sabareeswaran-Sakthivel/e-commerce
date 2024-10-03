# SetUp
- Install Golang version 1.21.0
- Run go mod vendor
- Run go mod tidy
- Build go build -o \<app-name\>
- Start the application by ./\<app-name\>

# Table Schema

## Orders

This table stores the order information.

- id
- customer_id
- date_of_sale
- payment_method
- shipping_cost
- region_id
- campaign_id

### Constraints
- Foreign Key - customer_id refers id from customers table
- Foreign Key - region_id refers id from regions table
- Foreign Key - campaign_id refers id from campaigns table


## Products

This table stores the product information.

- id
- name
- description
- category_id
- unit_price

### Constraints
- Foreign Key - category_id refers id from categories table


## Customers

This table stores the customer information.

- id
- name
- email
- address
- age
- gender
- is_active
- created_at
- updated_at

### Constraints
- Unique Column - email


## Order Details

This table stores the detail information of orders. A order may contain multiple products. For many to many mapping, we are using this table

- id
- order_id
- product_id
- quantity_sold
- discount

### Constraints
- Unique Column - (order_id, product_id)
- FK - order_id refers id from orders table
- FK - product_id refers id from products table

## Categories

This table stores all the categories.

- id
- name

### Constraints
- Unique Column - name


## Region

This table stores all the regions

- id
- name

### Constraints
- Unique Column - name

# Data Refresh Mechanism

Here, I have scheduled a cron every day at 12. 
Another approach, we can create jenkins pipeline and schedule a cron over there to hit the api's.

### API - http://localhost:7075/uploadCSVData
- **Query Parameters**:
    - fileName: csv file to upload the data. It should be in applications route.


# Revenue and Order Management API

This API provides endpoints to retrieve various metrics related to revenue, orders, customers, and products within specified date ranges. 

## Base URL
http://localhost:7075/api


## Endpoints

### Total Revenue

#### Get Total Revenue for a Date Range
- **URL**: `/revenue/total`
- **Method**: `GET`
- **Query Parameters**:
  - `startDate`: Start date for the date range (format: YYYY-MM-DD)
  - `endDate`: End date for the date range (format: YYYY-MM-DD)
- **Response**:
  ```json
  {
    "total_revenue": 143.976
  }
  
#### Get Total Revenue by Product: (For a date range)
- **URL**: `/revenue/products`
- **Method**: `GET`
- **Query Parameters**:
  - `startDate`: Start date for the date range (format: YYYY-MM-DD)
  - `endDate`: End date for the date range (format: YYYY-MM-DD)
- **Response**:
  ```json
  [
    {
        "product_id": "P123",
        "product_name": "UltraBoost Running Shoes",
        "total_revenue": 324
    }
  ]

#### Get Total Revenue by Category: (For a date range)
- **URL**: `/revenue/categories`
- **Method**: `GET`
- **Query Parameters**:
  - `startDate`: Start date for the date range (format: YYYY-MM-DD)
  - `endDate`: End date for the date range (format: YYYY-MM-DD)
- **Response**:
  ```json
  [
    {
        "category_id": "73da15d4-1692-452e-b0f8-fb26313be6ce",
        "category_name": "Clothing",
        "total_revenue": 143.976
    },
    {
        "category_id": "d9dd6ffa-2d1e-4426-93c2-1ae63cfddbe8",
        "category_name": "Shoes",
        "total_revenue": 504
    },
    {
        "category_id": "f6fe0f21-e6e7-481e-864c-1337ed3b6021",
        "category_name": "Electronics",
        "total_revenue": 4064.5915
    }
  ]

#### Get Total Revenue by Region: (For a date range)
- **URL**: `/revenue/region`
- **Method**: `GET`
- **Query Parameters**:
  - `startDate`: Start date for the date range (format: YYYY-MM-DD)
  - `endDate`: End date for the date range (format: YYYY-MM-DD)
- **Response**:
  ```json
  [
    {
        "region_id": 1,
        "region_name": "North America",
        "total_revenue": 297.4915
    },
    {
        "region_id": 2,
        "region_name": "Europe",
        "total_revenue": 1299
    },
    {
        "region_id": 3,
        "region_name": "Asia",
        "total_revenue": 2612.076
    },
    {
        "region_id": 4,
        "region_name": "South America",
        "total_revenue": 180
    }
  ]


### Top N Products:

#### Get Overall: (Based on quantity sold within a date range)
- **URL**: `/top/products`
- **Method**: `GET`
- **Query Parameters**:
  - `startDate`: Start date for the date range (format: YYYY-MM-DD)
  - `endDate`: End date for the date range (format: YYYY-MM-DD)
  - `limit`: Top N products
- **Response**:
  ```json
  [
    {
        "category_id": "d9dd6ffa-2d1e-4426-93c2-1ae63cfddbe8",
        "category_name": "Shoes",
        "product_id": "P123",
        "product_name": "UltraBoost Running Shoes",
        "total_quantity_sold": 3
    },
    {
        "category_id": "f6fe0f21-e6e7-481e-864c-1337ed3b6021",
        "category_name": "Electronics",
        "product_id": "P456",
        "product_name": "iPhone 15 Pro",
        "total_quantity_sold": 3
    }
  ]

#### Get TOP N products by Category: (Based on quantity sold within a date range)
- **URL**: `/top/products/category`
- **Method**: `GET`
- **Query Parameters**:
  - `startDate`: Start date for the date range (format: YYYY-MM-DD)
  - `endDate`: End date for the date range (format: YYYY-MM-DD)
  - `limit`: Top N products
- **Response**:
  ```json
  [
    {
      "category_id": "73da15d4-1692-452e-b0f8-fb26313be6ce",
      "category_name": "Clothing",
      "products": 
        [
          {
            "id": "P789",
            "name": "Levi's 501 Jeans",
            "total_quantity_sold": 3
          }
        ]
    }
  ]

#### Get TOP N products by Region: (Based on quantity sold within a date range)
- **URL**: `/top/products/region`
- **Method**: `GET`
- **Query Parameters**:
  - `startDate`: Start date for the date range (format: YYYY-MM-DD)
  - `endDate`: End date for the date range (format: YYYY-MM-DD)
  - `limit`: Top N products
- **Response**:
  ```json
  [
    {
      "region_id": 1,
      "region_name": "North America",
      "products":
        [
          {
            "id": "P234",
            "name": "Sony WH-1000XM5 Headphones",
            "total_quantity_sold": 1
          }
        ]
    },
    {
      "region_id": 3,
      "region_name": "Asia",
      "products": 
        [
          {
            "id": "P456",
            "name": "iPhone 15 Pro",
            "total_quantity_sold": 2
          }
        ]
    }
  ]


### Customer Analysis:

#### Get Total Number of Orders: (Within a date range)
- **URL**: `/orders`
- **Method**: `GET`
- **Query Parameters**:
  - `startDate`: Start date for the date range (format: YYYY-MM-DD)
  - `endDate`: End date for the date range (format: YYYY-MM-DD)
- **Response**:
  ```json
  {
    "orders": [
      {
        "id": 1003,
        "customer_id": "C456",
        "date_of_sale": "2024-02-28T00:00:00Z",
        "payment_method": "Debit Card",
        "shipping_cost": 5,
        "region_id": 3,
        "campaign_id": null
      },
      {
        "id": 1004,
        "customer_id": "C101",
        "date_of_sale": "2024-03-10T00:00:00Z",
        "payment_method": "Credit Card",
        "shipping_cost": 8,
        "region_id": 4,
        "campaign_id": null
      }
    ],
    "total_orders": 4
  }

#### Get Total Number of Customers: (Within a date range)
- **URL**: `/orders/customers`
- **Method**: `GET`
- **Query Parameters**:
  - `startDate`: Start date for the date range (format: YYYY-MM-DD)
  - `endDate`: End date for the date range (format: YYYY-MM-DD)
- **Response**:
  ```json
  {
    "customers": [
      {
        "id": "C456",
        "name": "John Smith",
        "email": "johnsmith@email.com",
        "address": "",
        "age": null,
        "gender": "",
        "is_active": true,
        "created_at": "",
        "updated_at": ""
      },
      {
        "id": "C101",
        "name": "Sarah Johnson",
        "email": "sarahjohnson@email.com",
        "address": "",
        "age": null,
        "gender": "",
        "is_active": true,
        "created_at": "",
        "updated_at": ""
      }
    ],
    "total_customers": 2
  }


#### Get Average Order Value: (Within a date range)
- **URL**: `/orders/average`
- **Method**: `GET`
- **Query Parameters**:
  - `startDate`: Start date for the date range (format: YYYY-MM-DD)
  - `endDate`: End date for the date range (format: YYYY-MM-DD)
- **Response**:
  ```json
  {
    "TotalRevenue": 4388.5675,
    "TotalOrders": 5,
    "AverageOrderValue": 877.7135000000001
  }