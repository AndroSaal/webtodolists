package handler

import (
	todo "ToDoApp"
	"net/http"

	"github.com/gin-gonic/gin"
)

// транспортный уровень
func (h *Handler) createList(c *gin.Context) {

	userId, ok := getUserId(c)
	if ok != nil {
		newErrorResponse(c, http.StatusBadRequest, "User not found")
		return
	}

	//инициализируем сущность Лист, из контекста десериализуем в сущность
	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//вызов метода создание списка
	id, err := h.services.TodoList.CreateList(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})

}

type getAllListResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllList(c *gin.Context) {

	//получение id user'a
	userId, ok := getUserId(c)
	if ok != nil {
		newErrorResponse(c, http.StatusBadRequest, "User not found")
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListResponse{
		Data: lists,
	})

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
