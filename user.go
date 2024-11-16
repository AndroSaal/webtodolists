package todo

// поля полностью совпадают со структурой базы данных
// Тэги используем для того, чтобы корректно принимать и выводить данные
// на HTTP запросах
type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"	 binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
	