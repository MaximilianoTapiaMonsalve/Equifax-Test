package services

import (
	"context"
	"fmt"
	"golang-interview/domain/models"
	"golang-interview/domain/ports"
)

type DashboardService struct {
	repo ports.IDashboardRepository
}

func New(repo ports.IDashboardRepository) *DashboardService {
	return &DashboardService{repo: repo}
}

func (s *DashboardService) GetDashboard(ctx context.Context, id int32) (*models.DashboardResponse, error) {

	return nil, fmt.Errorf("Not Implemented yet")
}
