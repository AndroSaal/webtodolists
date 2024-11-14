package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct{}

//инициализация энд-поинтов
/*
Функция возвращает объект *gin.Engine, который реализует интерфейс
Hendler из пакета net/http
*/
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	//инициализация группы /auth
	auth := router.Group("/auth")
	{
		auth.POST("sing-up", h.singUp)
		auth.POST("sing-in", h.singIn)
	}

	//инициализация группы /api
	api := router.Group("/api")
	{
		//группа lists для работы со списками
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
		}

		//группа items методы для работы с задачами
		items := lists.Group(":id/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getAllItem)
			items.GET("/:item_id", h.getItemById)
			items.PUT("/:item_id", h.updateItem)
			items.DELETE("/:item_id", h.deleteItem)
		}

	}
	return router
}
