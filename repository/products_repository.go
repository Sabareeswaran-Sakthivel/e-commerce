package repository

import (
	"github.com/sabareeswaran-sakthivel/e-commerce/database/connection"
	"github.com/sabareeswaran-sakthivel/e-commerce/models"
)

type ProductsRepository struct {
}

func (pr *ProductsRepository) GetByName(name string) (models.Product, error) {
	product := models.Product{}
	db := connection.GetDB()

	rows, err := db.Query("SELECT * FROM Products WHERE name = ?", name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.Product
		)
		err = rows.Scan(
			&ro.Id,
			&ro.Name,
			&ro.Description,
			&ro.CategoryId,
			&ro.UnitPrice,
		)
		if err != nil {
			return product, err
		}
		product = ro
	}
	return product, nil
}

func (pr *ProductsRepository) GetProductsForDateRange(startDate, endDate string) ([]models.Product, error) {
	products := make([]models.Product, 0)
	db := connection.GetDB()

	rows, err := db.Query("SELECT * FROM Products WHERE date >= ? AND date <= ?", startDate, endDate)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.Product
		)
		err = rows.Scan(
			&ro.Id,
			&ro.Name,
			&ro.Description,
			&ro.CategoryId,
			&ro.UnitPrice,
		)
		if err != nil {
			return products, err
		}
		products = append(products, ro)
	}
	return products, nil
}

func (pr *ProductsRepository) Create(product models.Product) error {
	db := connection.GetDB()

	_, err := db.Exec(
		`INSERT INTO products (id, name, description, category_id, unit_price) VALUES (?, ?, ?, ?, ?);`,
		product.Id,
		product.Name,
		product.Description,
		product.CategoryId,
		product.UnitPrice,
	)
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProductsRepository) CreateProductsTable() error {
	db := connection.GetDB()

	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    category_id INT,
    unit_price DECIMAL(10, 2) NOT NULL,
		FOREIGN KEY (category_id) REFERENCES Categories(id)
	);`)

	if err != nil {
		return err
	}
	return nil
}
