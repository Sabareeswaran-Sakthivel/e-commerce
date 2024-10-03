package repository

import (
	"github.com/sabareeswaran-sakthivel/e-commerce/database/connection"
	"github.com/sabareeswaran-sakthivel/e-commerce/models"
)

type RegionsRepository struct {
}

func (reg *RegionsRepository) GetRegionByName(name string) (models.Region, error) {
	region := models.Region{}
	db := connection.GetDB()

	rows, err := db.Query("SELECT * FROM Regions WHERE name = ?", name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.Region
		)
		err = rows.Scan(
			&ro.Id,
			&ro.Name,
		)
		if err != nil {
			return region, err
		}
		region = ro
	}
	return region, nil
}

func (reg *RegionsRepository) Create(region models.Region) error {
	db := connection.GetDB()

	_, err := db.Exec(
		`INSERT INTO regions (name) VALUES (?);`,
		region.Name,
	)
	if err != nil {
		return err
	}
	return nil
}

func (reg *RegionsRepository) CreateRegionsTable() error {
	db := connection.GetDB()

	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS regions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
		UNIQUE(name)
	);`)

	if err != nil {
		return err
	}
	return nil
}
