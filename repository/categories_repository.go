package repository

import (
	"github.com/sabareeswaran-sakthivel/e-commerce/database/connection"
	"github.com/sabareeswaran-sakthivel/e-commerce/models"
)

type CategoriesRepository struct {
}

func (cat *CategoriesRepository) GetCategoryByName(name string) (models.Category, error) {
	category := models.Category{}
	db := connection.GetDB()

	rows, err := db.Query("SELECT * FROM categories WHERE name = ?", name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.Category
		)
		err = rows.Scan(
			&ro.Id,
			&ro.Name,
		)
		if err != nil {
			return category, err
		}
		category = ro
	}
	return category, nil
}

func (cat *CategoriesRepository) Create(category models.Category) error {
	db := connection.GetDB()

	_, err := db.Exec(
		`INSERT INTO categories (id, name) VALUES (?, ?);`,
		category.Id,
		category.Name,
	)
	if err != nil {
		return err
	}
	return nil
}

func (cat *CategoriesRepository) CreateCategoriesTable() error {
	db := connection.GetDB()

	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS categories (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
	);`)

	if err != nil {
		return err
	}
	return nil
}
