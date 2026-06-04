package entities

type UserEntity struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

type TodoEntity struct {
	ID        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserID    int    `json:"userId"`
}

type TodosResponseEntity struct {
	Todos []TodoEntity `json:"todos"`
	Total int          `json:"total"`
	Skip  int          `json:"skip"`
	Limit int          `json:"limit"`
}
