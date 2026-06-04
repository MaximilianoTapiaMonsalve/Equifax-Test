package models

type User struct {
	ID        int32
	FirstName string
	LastName  string
	Age       int32
}

type Todo struct {
	ID        int32
	Todo      string
	Completed bool
	UserID    int32
}

type Todos struct {
	Todos []Todo
	Total int
	Skip  int
	Limit int
}
