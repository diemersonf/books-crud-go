package routes

import (
	"github.com/diemersonf/books-crud-go/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.ShowBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.EditeBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}

	return router
}
