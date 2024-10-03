package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sabareeswaran-sakthivel/e-commerce/services"
)

type CronController struct {
	service services.UploadCSVDataService
}

func (cc *CronController) UploadCSVData(c *gin.Context) {
	res, err := cc.service.UploadCSVData(c.Request.URL.Query().Get("fileName"))
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "message": res})
}
