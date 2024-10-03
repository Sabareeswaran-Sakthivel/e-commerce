package models

type TotalRevenueByProduct struct {
	ProductId    string  `json:"product_id" db:"product_id"`
	ProductName  string  `json:"product_name" db:"product_name"`
	TotalRevenue float64 `json:"total_revenue" db:"total_revenue"`
}

type TotalRevenueByRegion struct {
	RegionId     int     `json:"region_id" db:"region_id"`
	RegionName   string  `json:"region_name" db:"region_name"`
	TotalRevenue float64 `json:"total_revenue" db:"total_revenue"`
}

type TotalRevenueByCategory struct {
	CategoryId   string  `json:"category_id" db:"category_id"`
	CategoryName string  `json:"category_name" db:"category_name"`
	TotalRevenue float64 `json:"total_revenue" db:"total_revenue"`
}
