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

func (r *DashboardRepository) FetchUser(ctx context.Context, id int32) (*models.User, error) {
	fmt.Println("FetchingUser")
	var entity entities.UserEntity
	if err := r.h.Fetch(ctx, fmt.Sprintf("/users/%d", id), &entity); err != nil {
		return nil, err
	}
	return toUserModel(entity), nil
}

func (r *DashboardRepository) FetchTodos(ctx context.Context, id int32) (*models.Todos, error) {
	fmt.Println("FetchingTodos")
	var entity entities.TodosResponseEntity
	if err := r.h.Fetch(ctx, fmt.Sprintf("/todos/user/%d", id), &entity); err != nil {
		return nil, err
	}
	return toTodoModel(entity), nil
}

func toUserModel(e entities.UserEntity) *models.User {
	return &models.User{
		ID:        int32(e.ID), //<-- this can give some cutted bits but im guessing the dummy api does not have too many users (assuming that the ids are autoncremental)
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Age:       int32(e.Age),
	}
}

func toTodoModel(e entities.TodosResponseEntity) *models.Todos {
	todos := make([]models.Todo, len(e.Todos)) // <-- i had panics sometimes doing this, but i dont know why :s (i know its index errors but idk what casue it)
	for i, t := range e.Todos {
		todos[i] = models.Todo{
			ID:        int32(t.ID),
			Todo:      t.Todo,
			Completed: t.Completed,
			UserID:    int32(t.UserID),
		}
	}
	return &models.Todos{
		Todos: todos,
		Total: e.Total,
		Skip:  e.Skip,
		Limit: e.Limit,
	}
}
