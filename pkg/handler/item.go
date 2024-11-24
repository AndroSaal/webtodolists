package handler

import (
	todo "ToDoApp/entities"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {

	//для создания задачи нужно получать id пользователя и id списка
	userId, ok := getUserId(c) //получаем id пользователя из midleware
	if ok != nil {
		newErrorResponse(c, http.StatusBadRequest, "This user not found")
		return
	}

	//получение id списка
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf(err.Error(), " This list not found"))
		return
	}

	//получение input: из json в struct
	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//отправляем на уровень ниже - в сервис(в уровень бизнес-логики)
	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	//отправка ответа, если все закончилось хорошо
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) getAllItem(c *gin.Context) {

}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
