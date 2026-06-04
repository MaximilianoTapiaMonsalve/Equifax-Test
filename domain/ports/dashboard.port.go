package ports

import (
	"context"
	"golang-interview/domain/models"
)

type IDashboardRepository interface {
	FetchUser(ctx context.Context, id int32) (*models.User, error)
	FetchTodos(ctx context.Context, id int32) (*models.Todos, error)
}
