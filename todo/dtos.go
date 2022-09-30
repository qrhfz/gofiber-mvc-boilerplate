package todo

type NewTodoDto struct {
	Todo string
}

type UpdateTodoDto struct {
	Todo string
	Done bool
}
