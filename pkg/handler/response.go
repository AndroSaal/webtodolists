package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct {
	Message string `json:"message"`
}

//обработчик ошибки
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	//возвращение ошибки внутри логгера (чтобы мы увидели)
	logrus.Error(message)
	//возварщение ошибки в качестве ответа (чтобы увидел клиент)
	c.AbortWithStatusJSON(statusCode, error{message})
}
