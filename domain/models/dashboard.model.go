package models

type UserResponse struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

type Todo struct {
	ID        int
	Todo      string
	Completed bool
	UserID    int
}

type TodosResponse struct {
	Todos []Todo
	Total int
	Skip  int
	Limit int
}

type DashboardResponse struct {
	ID               int     `json:"id"`
	FullName         string  `json:"full_name"`
	Status           string  `json:"status"`
	PendingTaskCount int     `json:"pending_task_count"`
	NextUrgentTask   *string `json:"next_urgent_task"`
	ErrorWarning     *string `json:"error_warning"`
}
