package test

import (
	"context"
	"errors"
	"golang-interview/domain/models"
	"golang-interview/domain/services"
	"testing"
)

type mockRepo struct {
	user     *models.User
	todos    *models.Todos
	userErr  error
	todosErr error
}

func (m *mockRepo) FetchUser(ctx context.Context, id int32) (*models.User, error) {
	return m.user, m.userErr
}

func (m *mockRepo) FetchTodos(ctx context.Context, id int32) (*models.Todos, error) {
	return m.todos, m.todosErr
}

func TestGetDashboardStatusVeteran(t *testing.T) {
	repo := &mockRepo{
		user: &models.User{
			ID:        420,
			FirstName: "Maximiliano",
			LastName:  "Tapia",
			Age:       69,
		},
		todos: &models.Todos{
			Todos: []models.Todo{},
		},
	}

	svc := services.New(repo)

	resp, err := svc.GetDashboard(context.Background(), 2)
	if err != nil {
		t.Errorf("must not fail: %v", err)
	}

	if resp == nil {
		t.Fatal("response can't be null")
	}

	if resp.Status != "Veteran" {
		t.Errorf("expected veteran, but responded %s", resp.Status)
	}
}

func TestGetDashboardAllTodosCompleted(t *testing.T) {
	repo := &mockRepo{
		user: &models.User{
			ID:        9,
			FirstName: "Tax",
			LastName:  "Mapia",
			Age:       11,
		},
		todos: &models.Todos{
			Todos: []models.Todo{
				{
					ID:        1,
					Todo:      "todo 1",
					Completed: true,
				},
				{
					ID:        2,
					Todo:      "todo 2",
					Completed: true,
				},
			},
		},
	}

	svc := services.New(repo)
	resp, err := svc.GetDashboard(context.Background(), 1)
	if err != nil {
		t.Errorf("must not fail: %v", err)
	}

	if resp.PendingTaskCount != 0 {
		t.Errorf("expected 0 pending todos but have %d", resp.PendingTaskCount)
	}

	if resp.NextUrgentTask != nil {
		t.Errorf("NextUrgentTask must be null but have %s", *resp.NextUrgentTask)
	}
}

func TestGetDashboardTodosUnavailable(t *testing.T) {
	repo := &mockRepo{
		user:     &models.User{ID: 1, FirstName: "A", LastName: "B", Age: 25},
		todosErr: errors.New("timeout al llamar a la api de todos"),
	}

	svc := services.New(repo)
	resp, err := svc.GetDashboard(context.Background(), 1)

	if err != nil {
		t.Errorf("must not fail: %v", err)
	}

	if resp.ErrorWarning == nil {
		t.Fatal("ErrorWarning must not be null")
	}

	warning := *resp.ErrorWarning
	if warning != "Todos Unavailable" {
		t.Errorf("wrong warning, must be 'Todos Unavailable' but have '%s'", warning)
	}
}

func TestGetDashboardUserFails(t *testing.T) {
	repo := &mockRepo{
		userErr: errors.New("usuario not found"),
	}

	svc := services.New(repo)
	resp, err := svc.GetDashboard(context.Background(), 99)

	if err == nil {
		t.Error("se esperaba un error pero no se obtuvo ninguno")
	}

	if resp != nil {
		t.Errorf("se esperaba nil en la respuesta pero se obtuvo %v", resp)
	}
}
