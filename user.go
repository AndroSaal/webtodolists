package todo

//поля полностью совпадают со структурой базы данных
//Тэги используем для того, чтобы корректно принимать и выводить данные
//на HTTP запросах
type User struct {
	ID 			int    `json:"id"`
	Title 		string `json:"title"`
	Description string `json:"description"`
	Password 	string `json:"password"`
}