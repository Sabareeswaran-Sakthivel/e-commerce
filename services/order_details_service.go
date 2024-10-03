package services

import (
	"fmt"
	"strconv"

	"github.com/sabareeswaran-sakthivel/e-commerce/models"
	"github.com/sabareeswaran-sakthivel/e-commerce/repository"
)

type OrderDetailsService struct {
	ord repository.OrderDetailsRepository
}

func (ord *OrderDetailsService) Create(orderId int, productId string, quantitySold int, discount float64) {
	orderDetails := models.OrderDetails{
		OrderId:      orderId,
		ProductId:    productId,
		QuantitySold: quantitySold,
		Discount:     discount,
	}

	err := ord.ord.Create(orderDetails)
	if err != nil {
		fmt.Println(err)
	}

	err = ord.ord.Create(orderDetails)
	if err != nil {
		fmt.Println(err)
	}
}

func (ord *OrderDetailsService) GetTopSellingProductsForDateRange(startDate, endDate string, limit string) ([]models.TopSellingProduct, error) {
	if limit == "" {
		limit = "10"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return nil, err
	}

	topSellingProducts, err := ord.ord.GetTopSellingProductsForDateRange(startDate, endDate, limitInt)
	if err != nil {
		return nil, err
	}
	return topSellingProducts, nil
}

func (ord *OrderDetailsService) GetTopSellingProductsForDateRangeByCategory(startDate, endDate string, limit string) ([]models.TopSellingProductByCategoryResponse, error) {

	if limit == "" {
		limit = "10"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return nil, err
	}

	topSellingProducts, err := ord.ord.GetTopSellingProductsForDateRangeByCategory(startDate, endDate, limitInt)
	if err != nil {
		return nil, err
	}
	products := make(map[string][]models.TopSellingProductByCategoryResponse)

	for _, product := range topSellingProducts {
		if _, exists := products[product.CategoryName]; !exists {
			products[product.CategoryName] = append(products[product.CategoryName], models.TopSellingProductByCategoryResponse{
				CategoryID:   product.CategoryID,
				CategoryName: product.CategoryName,
				Products:     []models.ProductResponse{{Id: product.ProductID, Name: product.ProductName, TotalQuantitySold: product.TotalQuantitySold}},
			})
		} else {
			products[product.CategoryName][0].Products = append(products[product.CategoryName][0].Products, models.ProductResponse{Id: product.ProductID, Name: product.ProductName, TotalQuantitySold: product.TotalQuantitySold})
		}
	}

	response := make([]models.TopSellingProductByCategoryResponse, 0)
	for _, v := range products {
		response = append(response, models.TopSellingProductByCategoryResponse{
			CategoryID:   v[0].CategoryID,
			CategoryName: v[0].CategoryName,
			Products:     v[0].Products,
		})
	}
	return response, nil
}

func (ord *OrderDetailsService) GetTopSellingProductsForDateRangeByRegion(startDate, endDate string, limit string) ([]models.TopSellingProductByRegionResponse, error) {
	if limit == "" {
		limit = "10"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return nil, err
	}

	topSellingProducts, err := ord.ord.GetTopSellingProductsForDateRangeByRegion(startDate, endDate, limitInt)
	if err != nil {
		return nil, err
	}

	products := make(map[string][]models.TopSellingProductByRegionResponse)

	for _, product := range topSellingProducts {
		if _, exist := products[product.RegionName]; !exist {
			products[product.RegionName] = append(products[product.RegionName], models.TopSellingProductByRegionResponse{
				RegionID:   product.RegionID,
				RegionName: product.RegionName,
				Products:   []models.ProductResponse{{Id: product.ProductID, Name: product.ProductName, TotalQuantitySold: product.TotalQuantitySold}},
			})
		} else {
			products[product.RegionName][0].Products = append(products[product.RegionName][0].Products, models.ProductResponse{Id: product.ProductID, Name: product.ProductName, TotalQuantitySold: product.TotalQuantitySold})
		}
	}

	response := make([]models.TopSellingProductByRegionResponse, 0)
	for _, v := range products {
		response = append(response, models.TopSellingProductByRegionResponse{
			RegionID:   v[0].RegionID,
			RegionName: v[0].RegionName,
			Products:   v[0].Products,
		})
	}
	return response, nil
}
