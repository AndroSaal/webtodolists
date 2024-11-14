package hendler

import (
	"github.com/gin-gonic/gin"
)

type Hendler struct{}

//инициализация энд-поинтов
/*
Функция возвращает объект *gin.Engine, который реализует интерфейс
Hendler из пакета net/http
*/
func (h *Hendler) InitRoutes() *gin.Engine {
	router := gin.New()

	//инициализация группы /auth
	auth := router.Group("/auth")
	{
		auth.POST("sing-up")
		auth.POST("sing-in")
	}

	//инициализация группы /api
	api := router.Group("/api")
	{
		//группа листс для работы со списками
		lists := api.Group("/lists")
		{
			lists.POST("/")
			lists.GET("/")
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")
		}

		items := lists.Group(":id/items")
		{
			items.POST("/")
			items.GET("/")
			items.GET("/:item_id")
			items.PUT("/:item_id")
			items.DELETE("/:item_id")
		}

	}
	return router
}
