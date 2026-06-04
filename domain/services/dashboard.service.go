package services

import (
	"context"
	"fmt"
	"golang-interview/api/dtos"
	"golang-interview/domain/models"
	"golang-interview/domain/ports"
	"sync"
)

type DashboardService struct {
	repo ports.IDashboardRepository
}

func New(repo ports.IDashboardRepository) *DashboardService {
	return &DashboardService{repo: repo}
}

func (s *DashboardService) GetDashboard(ctx context.Context, id int32) (*dtos.DashboardDto, error) {

	var user *models.User
	var todos *models.Todos
	var userErr error
	var todosErr error
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		user, userErr = s.repo.FetchUser(ctx, id)
	}()

	go func() {
		defer wg.Done()
		todos, todosErr = s.repo.FetchTodos(ctx, id)
	}()

	wg.Wait()

	if userErr != nil {
		return nil, fmt.Errorf("fetching user: %w", userErr)
	}

	status := "Rookie"
	fmt.Println(user.Age)
	if user.Age > 50 {
		status = "Veteran"
	}

	response := s.toDto(user, status)

	if todosErr != nil {
		warning := "Todos Unavailable"
		response.ErrorWarning = &warning
		return response, nil
	}

	pendingCount, firstTitle := s.summarizeTodos(todos.Todos)
	response.PendingTaskCount = pendingCount
	response.NextUrgentTask = firstTitle

	return response, nil
}

func (s *DashboardService) toDto(model *models.User, status string) *dtos.DashboardDto {
	return &dtos.DashboardDto{
		ID:       model.ID,
		FullName: model.FirstName + " " + model.LastName,
		Status:   status,
	}
}

func (s *DashboardService) summarizeTodos(todos []models.Todo) (int, *string) {
	var count int
	var firstTitle *string

	for _, t := range todos {
		if !t.Completed {
			count++
			if firstTitle == nil {
				title := t.Todo
				firstTitle = &title
			}
		}
	}

	return count, firstTitle
}
