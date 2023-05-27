package main

import (
	"log"

	"dev.azure.com/jbsorg/segundo_proyecto/_git/bookstore-ai-api/controllers"
	"dev.azure.com/jbsorg/segundo_proyecto/_git/bookstore-ai-api/models"
	"dev.azure.com/jbsorg/segundo_proyecto/_git/bookstore-ai-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func main() {
	// Configurar la conexión a la base de datos MySQL
	dbConfig := mysql.Config{
		User:   "your-username",
		Passwd: "your-password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "book_inventory",
	}
	db, err := models.SetupDB(dbConfig)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}

	// Crear una instancia del enrutador Gin
	router := gin.Default()

	// Inicializar los controladores
	bookController := controllers.NewBookController(db)

	// Definir las rutas
	routes.SetupRoutes(router, bookController)

	// Iniciar el servidor HTTP
	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Error al iniciar el servidor HTTP: ", err)
	}
}
