package handler

import (
	todo "ToDoApp"
	"net/http"

	"github.com/gin-gonic/gin"
)

// обработчик регистрации
func (h *Handler) singUp(c *gin.Context) {

	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		//возвращение статус кода и сообщение ответа
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

// структура аутентификации пользователя
type userAuth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// обработчик аутентификации
func (h *Handler) singIn(c *gin.Context) {
	var input userAuth

	// с хранит в себе информацию о поступившем запросе,
	// метод BindJSON парсит тело запроса и записывает
	// информацию из тела запроса в переменную input
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
