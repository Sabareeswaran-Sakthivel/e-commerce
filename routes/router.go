// Package routes implements the routines for creating and handling api routes of the application.
package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/sabareeswaran-sakthivel/e-commerce/constants"
	"github.com/sabareeswaran-sakthivel/e-commerce/http/controllers"
)

// NewRouter returns a new gin router for the application routes
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "message": constants.MSG_NOT_FOUND_ERR})
	})

	// health check api endpoint
	router.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200, "message": "Server is running!"})
	})

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"POST", "GET", "DELETE"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           10 * time.Minute,
	}))
	cronController := new(controllers.CronController)
	ordersController := new(controllers.OrdersController)
	orderDetailsController := new(controllers.OrderDetailsController)

	router.GET("/uploadCSVData", cronController.UploadCSVData)

	router.GET("/api/revenue/total", ordersController.GetTotalRevenue)
	router.GET("/api/revenue/products", ordersController.GetTotalRevenueByProduct)
	router.GET("/api/revenue/region", ordersController.GetTotalRevenueByRegion)
	router.GET("/api/revenue/categories", ordersController.GetTotalRevenueByCategory)

	router.GET("/api/top/products", orderDetailsController.GetTopSellingProducts)
	router.GET("/api/top/products/category", orderDetailsController.GetTopSellingProductsByCategory)
	router.GET("/api/top/products/region", orderDetailsController.GetTopSellingProductsByRegion)

	router.GET("/api/orders/average", ordersController.GetAverageOrderValue)
	router.GET("/api/orders/customers", ordersController.GetTotalCustomers)
	router.GET("/api/orders/", ordersController.GetOrdersForDateRange)

	return router
}
