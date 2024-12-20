// транспортный уровень

package handler

import (
	todo "ToDoApp/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	lists, err := h.services.TodoList.GetAllList(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListResponse{
		Data: lists,
	})

}

func (h *Handler) getListById(c *gin.Context) {
	userId, ok := getUserId(c)
	if ok != nil {
		newErrorResponse(c, http.StatusBadRequest, "User not found")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	{
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

	}

	list, err := h.services.TodoList.GetById(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNonAuthoritativeInfo, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	{
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

	}

	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.UpdateById(userId, listId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, ststusResponse{"ok"})
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, ok := getUserId(c)
	if ok != nil {
		newErrorResponse(c, http.StatusBadRequest, "User not found")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	{
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

	}

	err = h.services.TodoList.DeleteById(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	} else {

	}

	c.JSON(http.StatusOK, ststusResponse{
		Status: "ok",
	})
}
