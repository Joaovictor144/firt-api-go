package routes

import (
	"first-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controllers.ShareBook)
			books.GET("/", controllers.ShowBooks)
			books.POST("/", controllers.CreateBook)
		}

	}

	return router
}

