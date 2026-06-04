package controllers

import (
	"golang-interview/domain/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	dashboardService *services.DashboardService
}

func New(service *services.DashboardService) *DashboardController {
	return &DashboardController{
		dashboardService: service,
	}
}

func (dc *DashboardController) GetDashboard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	response, err := dc.dashboardService.GetDashboard(c.Request.Context(), int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
