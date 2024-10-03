package services

import (
	"errors"
	"fmt"

	"github.com/sabareeswaran-sakthivel/e-commerce/constants"
	"github.com/sabareeswaran-sakthivel/e-commerce/models"
	"github.com/sabareeswaran-sakthivel/e-commerce/repository"
)

type OrdersService struct {
	or repository.OrdersRepository
}

func (os *OrdersService) GetOrders(startDate, endDate string) {
	orders, err := os.or.GetOrdersForDateRange(startDate, endDate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(orders)
}

func (os *OrdersService) GetTotalRevenue(startDate, endDate string) (map[string]float64, error) {
	totalRevenue, err := os.or.GetTotalRevenue(startDate, endDate)
	if err != nil {
		return nil, err
	}
	return map[string]float64{
		"total_revenue": totalRevenue,
	}, nil
}

func (os *OrdersService) GetTotalRevenueByProduct(startDate, endDate string) ([]models.TotalRevenueByProduct, error) {
	totalRevenueByProduct, err := os.or.GetTotalRevenueByProductForDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}
	return totalRevenueByProduct, nil
}

func (os *OrdersService) GetTotalRevenueByRegion(startDate, endDate string) ([]models.TotalRevenueByRegion, error) {
	totalRevenueByRegion, err := os.or.GetTotalRevenueByRegionForDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}
	return totalRevenueByRegion, nil
}

func (os *OrdersService) GetTotalRevenueByCategory(startDate, endDate string) ([]models.TotalRevenueByCategory, error) {
	totalRevenueByCategory, err := os.or.GetTotalRevenueByCategoryForDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}
	return totalRevenueByCategory, nil
}

func (os *OrdersService) GetAverageOrderValue(startDate, endDate string) (models.AverageOrderValue, error) {
	if startDate == "" || endDate == "" {
		return models.AverageOrderValue{}, errors.New(constants.INVALID_DATE_RANGE)
	}
	averageOrderValue, err := os.or.GetAverageOrderValueForDateRange(startDate, endDate)
	if err != nil {
		return models.AverageOrderValue{}, err
	}
	return averageOrderValue, nil
}

func (os *OrdersService) GetTotalCustomers(startDate, endDate string) (map[string]interface{}, error) {
	if startDate == "" || endDate == "" {
		return map[string]interface{}{}, errors.New(constants.INVALID_DATE_RANGE)
	}
	customers, err := os.or.GetTotalCustomersForDateRange(startDate, endDate)
	if err != nil {
		return map[string]interface{}{}, err
	}

	return map[string]interface{}{
		"total_customers": len(customers),
		"customers":       customers,
	}, nil
}

func (os *OrdersService) GetOrdersForDateRange(startDate, endDate string) (map[string]interface{}, error) {
	if startDate == "" || endDate == "" {
		return map[string]interface{}{}, errors.New(constants.INVALID_DATE_RANGE)
	}
	orders, err := os.or.GetOrdersForDateRange(startDate, endDate)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return map[string]interface{}{
		"total_orders": len(orders),
		"orders":       orders,
	}, nil
}
