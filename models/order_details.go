package models

type OrderDetails struct {
	Id           int     `json:"id" db:"id"`
	OrderId      int     `json:"order_id" db:"order_id"`
	ProductId    string  `json:"product_id" db:"product_id"`
	QuantitySold int     `json:"quantity_sold" db:"quantity_sold"`
	Discount     float64 `json:"discount" db:"discount"`
}
