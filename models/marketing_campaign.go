package models

type Campaign struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	EndDate     string `json:"end_date" db:"end_date"`
	StartDate   string `json:"start_date" db:"start_date"`
	Description string `json:"description" db:"description"`
}
