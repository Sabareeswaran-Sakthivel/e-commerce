package models

type Product struct {
	Id          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	CategoryId  string  `json:"category_id" db:"category_id"`
	UnitPrice   float64 `json:"unit_price" db:"unit_price"`
}
