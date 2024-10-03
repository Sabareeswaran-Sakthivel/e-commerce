package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sabareeswaran-sakthivel/e-commerce/services"
)

type OrdersController struct {
	service services.OrdersService
}

func (oc *OrdersController) GetTotalRevenue(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	response, err := oc.service.GetTotalRevenue(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (oc *OrdersController) GetTotalRevenueByProduct(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	response, err := oc.service.GetTotalRevenueByProduct(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (oc *OrdersController) GetTotalRevenueByRegion(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	response, err := oc.service.GetTotalRevenueByRegion(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (oc *OrdersController) GetTotalRevenueByCategory(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	response, err := oc.service.GetTotalRevenueByCategory(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (oc *OrdersController) GetAverageOrderValue(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	response, err := oc.service.GetAverageOrderValue(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (oc *OrdersController) GetTotalCustomers(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	response, err := oc.service.GetTotalCustomers(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (oc *OrdersController) GetOrdersForDateRange(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	response, err := oc.service.GetOrdersForDateRange(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
