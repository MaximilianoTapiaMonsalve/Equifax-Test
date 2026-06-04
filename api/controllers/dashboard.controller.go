package controllers

import (
	"golang-interview/domain/services"
	"net/http"

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

func (dc *DashboardController) GetDashboard(gctx *gin.Context) {
	type pathParam struct {
		ID int32 `uri:"id" binding:"required"`
	}

	var params pathParam

	if err := gctx.ShouldBindUri(&params); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	response, err := dc.dashboardService.GetDashboard(gctx, params.ID)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gctx.JSON(http.StatusOK, response)
}
