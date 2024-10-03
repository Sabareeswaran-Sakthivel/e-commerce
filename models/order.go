package models

type Order struct {
	Id            int     `json:"id" db:"id"`
	CustomerId    string  `json:"customer_id" db:"customer_id"`
	DateOfSale    string  `json:"date_of_sale" db:"date_of_sale"`
	PaymentMethod string  `json:"payment_method" db:"payment_method"`
	ShippingCost  float64 `json:"shipping_cost" db:"shipping_cost"`
	RegionId      int     `json:"region_id" db:"region_id"`
	CampaignId    *int    `json:"campaign_id" db:"campaign_id"`
}
