package main
import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"https://dev.azure.com/jbsorg/Segundo%20proyecto/_git/bookstore-ai-api/?path=/controllers&version=GBmain&_a=contents"
	"https://dev.azure.com/jbsorg/Segundo%20proyecto/_git/bookstore-ai-api/?path=/models&version=GBmain&_a=contents"
	"https://dev.azure.com/jbsorg/Segundo%20proyecto/_git/bookstore-ai-api/?path=/routes&version=GBmain&_a=contents"
	
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