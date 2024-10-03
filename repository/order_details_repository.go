package repository

import (
	"github.com/sabareeswaran-sakthivel/e-commerce/database/connection"
	"github.com/sabareeswaran-sakthivel/e-commerce/models"
)

type OrderDetailsRepository struct {
}

func (ord *OrderDetailsRepository) GetOrderDetailsByOrderId(orderId int) ([]models.OrderDetails, error) {
	orderDetails := make([]models.OrderDetails, 0)
	db := connection.GetDB()

	rows, err := db.Query("SELECT * FROM order_details WHERE order_id = ?", orderId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.OrderDetails
		)
		err = rows.Scan(
			&ro.Id,
			&ro.OrderId,
			&ro.ProductId,
			&ro.QuantitySold,
			&ro.Discount,
		)
		if err != nil {
			return orderDetails, err
		}
		orderDetails = append(orderDetails, ro)
	}
	return orderDetails, nil
}

func (ord *OrderDetailsRepository) GetTopSellingProductsForDateRange(startDate, endDate string, limit int) ([]models.TopSellingProduct, error) {
	orderDetails := make([]models.TopSellingProduct, 0)
	db := connection.GetDB()

	query := `SELECT
    p.id AS product_id,
    p.name AS product_name,
		c.id AS category_id,
		c.name AS category_name,
		SUM(od.quantity_sold) AS total_quantity_sold
		FROM 
		  order_details od
		INNER JOIN 
		  orders o ON od.order_id = o.id
		INNER JOIN 
		  products p ON od.product_id = p.id
		INNER JOIN
			categories c ON p.category_id = c.id
		WHERE 
		    o.date_of_sale BETWEEN ? AND ?
		GROUP BY 
		    p.id, p.name
		ORDER BY 
		    total_quantity_sold DESC
		LIMIT ?;`

	rows, err := db.Query(query, startDate, endDate, limit)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.TopSellingProduct
		)
		err = rows.Scan(
			&ro.ProductID,
			&ro.ProductName,
			&ro.CategoryID,
			&ro.CategoryName,
			&ro.TotalQuantitySold,
		)
		if err != nil {
			return orderDetails, err
		}
		orderDetails = append(orderDetails, ro)
	}
	return orderDetails, nil
}

func (ord *OrderDetailsRepository) GetTopSellingProductsForDateRangeByCategory(startDate, endDate string, limit int) ([]models.TopSellingProductByCategory, error) {
	topSellingProducts := make([]models.TopSellingProductByCategory, 0)
	db := connection.GetDB()

	query := prepareQureyForTopSellingProductsByCategory()

	rows, err := db.Query(query, startDate, endDate, limit)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.TopSellingProductByCategory
		)
		err = rows.Scan(
			&ro.CategoryID,
			&ro.CategoryName,
			&ro.ProductID,
			&ro.ProductName,
			&ro.TotalQuantitySold,
		)
		if err != nil {
			return topSellingProducts, err
		}
		topSellingProducts = append(topSellingProducts, ro)
	}
	return topSellingProducts, nil
}

func (ord *OrderDetailsRepository) GetTopSellingProductsForDateRangeByRegion(startDate, endDate string, limit int) ([]models.TopSellingProductByRegion, error) {
	topSellingProducts := make([]models.TopSellingProductByRegion, 0)
	db := connection.GetDB()

	query := prepareQureyForTopSellingProductsByRegion()

	rows, err := db.Query(query, startDate, endDate, limit)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.TopSellingProductByRegion
		)
		err = rows.Scan(
			&ro.ProductID,
			&ro.ProductName,
			&ro.RegionID,
			&ro.RegionName,
			&ro.TotalQuantitySold,
		)
		if err != nil {
			return topSellingProducts, err
		}
		topSellingProducts = append(topSellingProducts, ro)
	}
	return topSellingProducts, nil
}

func (ord *OrderDetailsRepository) CreateOrderDetailsTable() error {
	db := connection.GetDB()

	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS order_details (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    order_id INT NOT NULL,
    product_id VARCHAR(36) NOT NULL,
    quantity_sold INTEGER NOT NULL,
    discount REAL,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id),
		UNIQUE (order_id, product_id)
	);`)

	if err != nil {
		return err
	}
	return nil
}

func (ord *OrderDetailsRepository) Create(orderDetails models.OrderDetails) error {
	db := connection.GetDB()

	_, err := db.Exec(
		`INSERT INTO order_details (order_id, product_id, quantity_sold, discount) VALUES (?, ?, ?, ?);`,
		orderDetails.OrderId,
		orderDetails.ProductId,
		orderDetails.QuantitySold,
		orderDetails.Discount,
	)
	if err != nil {
		return err
	}
	return nil
}

func prepareQureyForTopSellingProductsByCategory() string {
	return `WITH filteredProducts AS (
    SELECT 
        c.id AS category_id,
        c.name AS category_name,
        p.id AS product_id,
        p.name AS product_name,
        SUM(od.quantity_sold) AS total_quantity_sold,
        ROW_NUMBER() OVER (PARTITION BY c.id ORDER BY SUM(od.quantity_sold) DESC) AS rank
    FROM 
        order_details od
    JOIN 
        orders o ON od.order_id = o.id
    JOIN 
        products p ON od.product_id = p.id
    JOIN 
        categories c ON p.category_id = c.id
    WHERE 
        o.date_of_sale BETWEEN ? AND ?
    GROUP BY 
        c.id, c.name, p.id, p.name
		)
		SELECT
		    category_id,
		    category_name,
		    product_id,
		    product_name,
		    total_quantity_sold
		FROM 
		    filteredProducts
		WHERE 
		    rank <= ?;
		`
}

func prepareQureyForTopSellingProductsByRegion() string {
	return `WITH filteredProducts AS (
    SELECT 
        p.id AS product_id,
        p.name AS product_name,
				r.id AS region_id,
				r.name AS region_name,
        SUM(od.quantity_sold) AS total_quantity_sold,
        ROW_NUMBER() OVER (PARTITION BY r.id ORDER BY SUM(od.quantity_sold) DESC) AS rank
    FROM 
        order_details od
    INNER JOIN
        orders o ON od.order_id = o.id
    INNER JOIN 
        products p ON od.product_id = p.id
    INNER JOIN 
        regions r ON o.region_id = r.id
    WHERE 
        o.date_of_sale BETWEEN ? AND ?
    GROUP BY 
        r.id, r.name, p.id, p.name
		)
		SELECT
		    product_id,
		    product_name,
				region_id,
				region_name,
		    total_quantity_sold
		FROM 
		    filteredProducts
		WHERE 
		    rank <= ?;
		`
}
