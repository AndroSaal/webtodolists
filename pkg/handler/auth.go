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

// обработчик аутентификация
func (h *Handler) singIn(c *gin.Context) {

}
