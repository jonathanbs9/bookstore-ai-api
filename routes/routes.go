package routes

import (
	"github.com/gin-gonic/gin"

	"dev.azure.com/jbsorg/segundo_proyecto/_git/bookstore-ai-api/controllers"
)

// SetupRoutes configura las rutas de la API
func SetupRoutes(router *gin.Engine, bookController *controllers.BookController) {
	// Grupo de rutas para los endpoints relacionados con libros
	booksGroup := router.Group("/books")
	{
		booksGroup.GET("", bookController.ListBooks)
		booksGroup.GET("/name/:name", bookController.GetBookByName)
		booksGroup.GET("/author/:author", bookController.GetBookByAuthor)
		booksGroup.GET("/isbn/:isbn", bookController.GetBookByISBN)
		booksGroup.GET("/id/:id", bookController.GetBookById)
		booksGroup.POST("", bookController.CreateBook)
		booksGroup.PUT("/:id", bookController.UpdateBook)
	}

	// Otras rutas y controladores si es necesario
}
