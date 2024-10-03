package models

type TopSellingProduct struct {
	CategoryID        string `json:"category_id"`
	CategoryName      string `json:"category_name"`
	ProductID         string `json:"product_id"`
	ProductName       string `json:"product_name"`
	TotalQuantitySold int    `json:"total_quantity_sold"`
}

type TopSellingProductByCategory struct {
	CategoryID        string `json:"category_id"`
	CategoryName      string `json:"category_name"`
	ProductID         string `json:"product_id"`
	ProductName       string `json:"product_name"`
	TotalQuantitySold int    `json:"total_quantity_sold"`
}

type TopSellingProductByCategoryResponse struct {
	CategoryID   string            `json:"category_id"`
	CategoryName string            `json:"category_name"`
	Products     []ProductResponse `json:"products"`
}

type ProductResponse struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	TotalQuantitySold int    `json:"total_quantity_sold"`
}

type TopSellingProductByRegion struct {
	RegionID          int    `json:"region_id"`
	RegionName        string `json:"region_name"`
	ProductID         string `json:"product_id"`
	ProductName       string `json:"product_name"`
	TotalQuantitySold int    `json:"total_quantity_sold"`
}

type TopSellingProductByRegionResponse struct {
	RegionID   int               `json:"region_id"`
	RegionName string            `json:"region_name"`
	Products   []ProductResponse `json:"products"`
}
