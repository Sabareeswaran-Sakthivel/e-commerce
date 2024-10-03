package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sabareeswaran-sakthivel/e-commerce/services"
)

type OrderDetailsController struct {
	service services.OrderDetailsService
}

func (oc *OrderDetailsController) GetTopSellingProducts(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	limit := c.Query("limit")

	response, err := oc.service.GetTopSellingProductsForDateRange(startDate, endDate, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (oc *OrderDetailsController) GetTopSellingProductsByCategory(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	limit := c.Query("limit")

	response, err := oc.service.GetTopSellingProductsForDateRangeByCategory(startDate, endDate, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (oc *OrderDetailsController) GetTopSellingProductsByRegion(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	limit := c.Query("limit")

	response, err := oc.service.GetTopSellingProductsForDateRangeByRegion(startDate, endDate, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
