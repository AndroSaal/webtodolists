package handler

import (
	todo "ToDoApp"

	"github.com/gin-gonic/gin"
)

func (h *Handler) singIn(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {

	}
}

func (h *Handler) singUp(c *gin.Context) {

}
