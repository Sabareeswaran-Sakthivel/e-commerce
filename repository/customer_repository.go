package repository

import (
	"github.com/sabareeswaran-sakthivel/e-commerce/database/connection"
	"github.com/sabareeswaran-sakthivel/e-commerce/models"
)

type CustomersRepository struct {
}

func (cus *CustomersRepository) GetCustomerByEmail(email string) (models.Customer, error) {
	customer := models.Customer{}
	db := connection.GetDB()

	rows, err := db.Query("SELECT * FROM customers WHERE email = ?", email)
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
			&ro.CreatedAt,
			&ro.UpdatedAt,
		)
		if err != nil {
			return customer, err
		}
		customer = ro
	}
	return customer, nil
}

func (cus *CustomersRepository) Create(customer models.Customer) error {
	db := connection.GetDB()

	_, err := db.Exec(
		`INSERT INTO customers (id, name, email, address, age, gender, is_active) VALUES (?, ?, ?, ?, ?, ?, ?);`,
		customer.Id,
		customer.Name,
		customer.Email,
		customer.Address,
		customer.Age,
		customer.Gender,
		customer.IsActive,
	)
	if err != nil {
		return err
	}
	return nil
}

func (cus *CustomersRepository) CreateCustomersTable() error {
	db := connection.GetDB()

	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS customers (
    id varchar(36) PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    address TEXT,
    age INTEGER,            
    gender TEXT,             
    is_active BOOLEAN DEFAULT TRUE,   
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE (email)
		);
	`)

	if err != nil {
		return err
	}
	return nil
}
