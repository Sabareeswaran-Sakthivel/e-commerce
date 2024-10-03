package repository

import (
	"github.com/sabareeswaran-sakthivel/e-commerce/database/connection"
	"github.com/sabareeswaran-sakthivel/e-commerce/models"
)

type CampaignsRepository struct {
}

func (cmr *CampaignsRepository) GetCampaignByName(name string) (models.Campaign, error) {
	campaign := models.Campaign{}
	db := connection.GetDB()

	rows, err := db.Query("SELECT * FROM Campaigns WHERE name = ?", name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ro models.Campaign
		)
		err = rows.Scan(
			&ro.Id,
			&ro.Name,
		)
		if err != nil {
			return campaign, err
		}
		campaign = ro
	}
	return campaign, nil
}

func (cmr *CampaignsRepository) Create(campaign models.Campaign) error {
	db := connection.GetDB()

	_, err := db.Exec(
		`INSERT INTO campaigns (name, end_date, start_date, description) VALUES (?, ?, ?, ?);`,
		campaign.Name,
		campaign.EndDate,
		campaign.StartDate,
		campaign.Description,
	)
	if err != nil {
		return err
	}
	return nil
}

func (cmr *CampaignsRepository) CreateCampaignsTable() error {
	db := connection.GetDB()

	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS campaigns (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
		UNIQUE (name)
	);`)

	if err != nil {
		return err
	}
	return nil
}
