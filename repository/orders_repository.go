package repository

import (
	"github.com/sabareeswaran-sakthivel/e-commerce/database/connection"
	"github.com/sabareeswaran-sakthivel/e-commerce/models"
)

type OrdersRepository struct {
}

func (or *OrdersRepository) GetById(id int) (models.Order, error) {
	order := models.Order{}
	db := connection.GetDB()

	rows, err := db.Query("SELECT * FROM Orders WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.Order
		)
		err = rows.Scan(
			&ro.Id,
			&ro.CustomerId,
			&ro.DateOfSale,
			&ro.ShippingCost,
			&ro.RegionId,
			&ro.CampaignId,
		)
		if err != nil {
			return order, err
		}
		order = ro
	}
	return order, nil
}

func (or *OrdersRepository) GetOrdersForDateRange(startDate, endDate string) ([]models.Order, error) {
	orders := make([]models.Order, 0)
	db := connection.GetDB()

	rows, err := db.Query("SELECT * FROM Orders WHERE date_of_sale BETWEEN ? AND ?", startDate, endDate)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.Order
		)
		err = rows.Scan(
			&ro.Id,
			&ro.CustomerId,
			&ro.DateOfSale,
			&ro.PaymentMethod,
			&ro.ShippingCost,
			&ro.RegionId,
			&ro.CampaignId,
		)
		if err != nil {
			return orders, err
		}
		orders = append(orders, ro)
	}
	return orders, nil
}

func (or *OrdersRepository) GetTotalCustomersForDateRange(startDate, endDate string) ([]models.Customer, error) {
	customers := make([]models.Customer, 0)
	db := connection.GetDB()

	query := `select DISTINCT customers.id,
		customers.name,
		customers.email,
		customers.address,
		customers.age,
		customers.gender,
		customers.is_active
		from orders
		INNER JOIN customers ON orders.customer_id = customers.id
		WHERE orders.date_of_sale BETWEEN ? AND ?;`

	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.Customer
		)
		err = rows.Scan(
			&ro.Id,
			&ro.Name,
			&ro.Email,
			&ro.Address,
			&ro.Age,
			&ro.Gender,
			&ro.IsActive,
		)
		if err != nil {
			return customers, err
		}
		customers = append(customers, ro)
	}
	return customers, nil
}

func (or *OrdersRepository) GetOrdersForDateRangeByCustomerId(startDate, endDate string, customerId int) ([]models.Order, error) {
	orders := make([]models.Order, 0)
	db := connection.GetDB()

	rows, err := db.Query("SELECT * FROM Orders WHERE date_of_sale BETWEEN ? AND ? AND customer_id = ?", startDate, endDate, customerId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.Order
		)
		err = rows.Scan(
			&ro.Id,
			&ro.CustomerId,
			&ro.DateOfSale,
			&ro.PaymentMethod,
			&ro.ShippingCost,
			&ro.RegionId,
			&ro.CampaignId,
		)
		if err != nil {
			return orders, err
		}
		orders = append(orders, ro)
	}
	return orders, nil
}

func (or *OrdersRepository) GetAverageOrderValueForDateRange(startDate, endDate string) (models.AverageOrderValue, error) {
	totalRevenue := models.AverageOrderValue{}
	db := connection.GetDB()

	query := `SELECT 
    IFNULL(SUM(od.quantity_sold * p.unit_price * (1 - IFNULL(od.discount, 0))), 0) AS total_revenue,
    COUNT(DISTINCT o.id) AS total_orders,
    IFNULL(SUM(od.quantity_sold * p.unit_price * (1 - od.discount)) / IFNULL(COUNT(DISTINCT o.id), 0), 0) AS average_order_value
		FROM 
		  orders o
		JOIN 
		  order_details od ON o.id = od.order_id
		JOIN 
		  products p ON od.product_id = p.id
		WHERE 
		    o.date_of_sale BETWEEN ? AND ?;
	`

	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.AverageOrderValue
		)
		err = rows.Scan(
			&ro.TotalRevenue,
			&ro.TotalOrders,
			&ro.AverageOrderValue,
		)
		if err != nil {
			return totalRevenue, err
		}
		totalRevenue = ro
	}
	return totalRevenue, nil
}

func (or *OrdersRepository) GetTotalRevenueByProductForDateRange(startDate, endDate string) ([]models.TotalRevenueByProduct, error) {
	totalRevenueByProduct := make([]models.TotalRevenueByProduct, 0)
	db := connection.GetDB()

	query := `select   
		products.id AS product_id,
		products.name AS product_name,
		SUM(order_details.quantity_sold * products.unit_price * (1 - order_details.discount)) AS total_revenue
		from
		orders 
		INNER JOIN order_details ON orders.id = order_details.order_id 
		INNER join products ON order_details.product_id = products.id
		WHERE orders.date_of_sale BETWEEN ? AND ?
		GROUP BY products.id;
	`

	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.TotalRevenueByProduct
		)
		err = rows.Scan(
			&ro.ProductId,
			&ro.ProductName,
			&ro.TotalRevenue,
		)
		if err != nil {
			return totalRevenueByProduct, err
		}
		totalRevenueByProduct = append(totalRevenueByProduct, ro)
	}
	return totalRevenueByProduct, nil
}

func (or *OrdersRepository) GetTotalRevenue(startDate, endDate string) (float64, error) {
	totalRevenue := 0.0
	db := connection.GetDB()

	query := `SELECT 
    IFNULL(SUM(od.quantity_sold * p.unit_price * (1 - IFNULL(od.discount, 0))), 0) AS total_revenue
		FROM 
		    Orders o
		JOIN 
		    Order_details od ON o.id = od.order_id
		JOIN 
		    Products p ON od.product_id = p.id
		WHERE 
		    o.date_of_sale BETWEEN ? AND ?;
	`

	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro float64
		)
		err = rows.Scan(
			&ro,
		)
		if err != nil {
			return totalRevenue, err
		}
		totalRevenue = ro
	}
	return totalRevenue, nil
}

func (or *OrdersRepository) GetTotalRevenueByProduct(startDate, endDate string) ([]models.TotalRevenueByProduct, error) {
	totalRevenueByProduct := make([]models.TotalRevenueByProduct, 0)
	db := connection.GetDB()

	query := `SELECT 
    p.id AS product_id,
    p.name AS product_name,
		IFNULL(SUM(od.quantity_sold * p.unit_price * (1 - IFNULL(od.discount, 0))), 0) AS total_revenue
		FROM 
		    Orders o
		JOIN 
		    Order_details od ON o.id = od.order_id
		JOIN 
		    Products p ON od.product_id = p.id
		WHERE 
		    o.date_of_sale BETWEEN ? AND ?;
	`

	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.TotalRevenueByProduct
		)
		err = rows.Scan(
			&ro.ProductId,
			&ro.ProductName,
			&ro.TotalRevenue,
		)
		if err != nil {
			return totalRevenueByProduct, err
		}
		totalRevenueByProduct = append(totalRevenueByProduct, ro)
	}
	return totalRevenueByProduct, nil
}

func (or *OrdersRepository) GetTotalRevenueByRegionForDateRange(startDate, endDate string) ([]models.TotalRevenueByRegion, error) {
	totalRevenueByRegion := make([]models.TotalRevenueByRegion, 0)
	db := connection.GetDB()

	query := `SELECT 
    r.id AS region_id,
    r.name AS region_name,
    IFNULL(SUM(od.quantity_sold * p.unit_price * (1 - IFNULL(od.discount, 0))), 0) AS total_revenue
		FROM 
		    orders o
		JOIN 
		    order_details od ON o.id = od.order_id
		JOIN 
		    products p ON od.product_id = p.id
		JOIN 
		    regions r ON o.region_id = r.id
		WHERE 
		    o.date_of_sale BETWEEN ? AND ?
		GROUP BY r.id;
	`

	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.TotalRevenueByRegion
		)
		err = rows.Scan(
			&ro.RegionId,
			&ro.RegionName,
			&ro.TotalRevenue,
		)
		if err != nil {
			return totalRevenueByRegion, err
		}
		totalRevenueByRegion = append(totalRevenueByRegion, ro)
	}
	return totalRevenueByRegion, nil
}

func (or *OrdersRepository) GetTotalRevenueByCategoryForDateRange(startDate, endDate string) ([]models.TotalRevenueByCategory, error) {
	totalRevenueByCategory := make([]models.TotalRevenueByCategory, 0)
	db := connection.GetDB()

	query := `select 
			categories.id as category_id,
			categories.name as category_name,
			SUM(order_details.quantity_sold * products.unit_price * (1 - order_details.discount)) AS total_revenue
			from orders
			INNER JOIN order_details ON orders.id = order_details.order_id
			INNER join products ON order_details.product_id = products.id
			INNER JOIN categories ON products.category_id = categories.id
			where orders.date_of_sale BETWEEN ? AND ?
			GROUP BY categories.id;
	`

	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.TotalRevenueByCategory
		)
		err = rows.Scan(
			&ro.CategoryId,
			&ro.CategoryName,
			&ro.TotalRevenue,
		)
		if err != nil {
			return totalRevenueByCategory, err
		}
		totalRevenueByCategory = append(totalRevenueByCategory, ro)
	}
	return totalRevenueByCategory, nil
}

func (or *OrdersRepository) Create(order models.Order) error {
	db := connection.GetDB()

	_, err := db.Exec(
		`INSERT INTO orders (id, customer_id, date_of_sale, payment_method, shipping_cost, region_id, campaign_id) VALUES (?, ?, ?, ?, ?, ?, ?);`,
		order.Id,
		order.CustomerId,
		order.DateOfSale,
		order.PaymentMethod,
		order.ShippingCost,
		order.RegionId,
		order.CampaignId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (or *OrdersRepository) CreateOrdersTable() error {
	db := connection.GetDB()

	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS orders (
    id INT PRIMARY KEY,
    customer_id VARCHAR(36) NOT NULL,
    date_of_sale DATE NOT NULL,
    payment_method TEXT NOT NULL,
    shipping_cost REAL NOT NULL,
    region_id INTEGER NOT NULL,
    campaign_id INTEGER,
    FOREIGN KEY (customer_id) REFERENCES customers(id),
    FOREIGN KEY (region_id) REFERENCES regions(id),
    FOREIGN KEY (campaign_id) REFERENCES campaigns(id)
	);`)

	if err != nil {
		return err
	}
	return nil
}
