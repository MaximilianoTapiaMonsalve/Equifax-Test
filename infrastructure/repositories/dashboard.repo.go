package repositories

import (
	"context"
	"fmt"
	"golang-interview/domain/models"
	"golang-interview/domain/ports"
	"golang-interview/infrastructure/entities"
	"golang-interview/infrastructure/https"
)

type DashboardRepository struct {
	h *https.HTTPClient
}

func New(http *https.HTTPClient) ports.IDashboardRepository {
	return &DashboardRepository{
		h: http,
	}
}

func (r *DashboardRepository) FetchUser(ctx context.Context, id int32) (*models.UserResponse, error) {
	var entity entities.UserEntity
	if err := r.h.Fetch(ctx, fmt.Sprintf("/users/%d", id), &entity); err != nil {
		return nil, err
	}
	return toUserModel(entity), nil
}

func (r *DashboardRepository) FetchTodos(ctx context.Context, id int32) (*models.TodosResponse, error) {
	var entity entities.TodosResponseEntity
	if err := r.h.Fetch(ctx, fmt.Sprintf("/todos/user/%d", id), &entity); err != nil {
		return nil, err
	}
	return toTodoModel(entity), nil
}

func toUserModel(e entities.UserEntity) *models.UserResponse {
	return &models.UserResponse{
		ID:        e.ID,
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Age:       e.Age,
	}
}

func toTodoModel(e entities.TodosResponseEntity) *models.TodosResponse {
	todos := make([]models.Todo, len(e.Todos))
	for i, t := range e.Todos {
		todos[i] = models.Todo{
			ID:        t.ID,
			Todo:      t.Todo,
			Completed: t.Completed,
			UserID:    t.UserID,
		}
	}
	return &models.TodosResponse{
		Todos: todos,
		Total: e.Total,
		Skip:  e.Skip,
		Limit: e.Limit,
	}
}
