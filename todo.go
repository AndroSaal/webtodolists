package todo

// получаем из запроса
type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

// получаем из запроса
type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// для связи сущностей
type UserList struct {
	Id     int
	UserId int
	ListId int
}

// для связи сущностей
type LIstsItem struct {
	Id     int
	ListId int
	ItemId int
}
