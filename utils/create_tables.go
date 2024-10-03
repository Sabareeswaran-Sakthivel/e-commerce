package utils

import (
	"github.com/sabareeswaran-sakthivel/e-commerce/models"
	"github.com/sabareeswaran-sakthivel/e-commerce/repository"
)

func CreateTables() {

	campaingRepo := repository.CampaignsRepository{}
	categoryRepo := repository.CategoriesRepository{}
	regionRepo := repository.RegionsRepository{}
	customerRepo := repository.CustomersRepository{}
	productRepo := repository.ProductsRepository{}
	orderRepo := repository.OrdersRepository{}
	orderDetailsRepo := repository.OrderDetailsRepository{}

	err := campaingRepo.CreateCampaignsTable()
	if err != nil {
		panic(err)
	}

	err = categoryRepo.CreateCategoriesTable()
	if err != nil {
		panic(err)
	}

	err = regionRepo.CreateRegionsTable()
	if err != nil {
		panic(err)
	}

	err = customerRepo.CreateCustomersTable()
	if err != nil {
		panic(err)
	}

	err = productRepo.CreateProductsTable()
	if err != nil {
		panic(err)
	}

	err = orderRepo.CreateOrdersTable()
	if err != nil {
		panic(err)
	}

	err = orderDetailsRepo.CreateOrderDetailsTable()
	if err != nil {
		panic(err)
	}
}

func InsertRegions() {
	regions := []string{"North America", "Europe", "Asia", "South America"}
	regionRepo := repository.RegionsRepository{}
	for _, region := range regions {
		regionM := models.Region{
			Name: region,
		}
		regionRepo.Create(regionM)
	}
}
