package handler

import (
	"net/http"
	"strings"

	"errors"

	"github.com/gin-gonic/gin"
)

const (
	autorizationHeader = "Authorization"
	userCtx            = "userId"
)

func (h *Handler) UserIdentity(c *gin.Context) {

	// Получаем токен из хэдера авторизации
	header := c.GetHeader(autorizationHeader)

	// Проверяем, что получили не пустой хэдер
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty autorization Header")
		return
	}

	// Разделяем Хэдер по пробелам, при корректном Хэдере, функция вернет 2 элемента
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid autorization Header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user not found")
		return 0, errors.New("user not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is unvalid type")
		return 0, errors.New("user id type is unvalid")
	}

	return idInt, nil
}
