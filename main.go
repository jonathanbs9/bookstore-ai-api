package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jonathanbs9/bookstore-ai-api/controllers"
	"github.com/jonathanbs9/bookstore-ai-api/routes"
)

func main() {
	// Configurar la conexión a la base de datos MySQL
	// //dbConfig := mysql.Config{
	// 	User:   "root",
	// 	Passwd: "",
	// 	Net:    "tcp",
	// 	Addr:   "localhost:3306",
	// 	DBName: "book_inventory",
	// }
	//db, err := models.SetupDB(dbConfig)

	/* db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/book_inventory")
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}
	defer db.Close()*/

	dsn := "jonathanbs:ihmQFPAYETnygjZodt49EA@tcp(azure-mule-888.g8x.cockroachlabs.cloud:26257)/book_inventory"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping database:", err)
	}

	// Ejecutar script para crear la tabla si no existe
	err = createTableIfNotExists(db)
	if err != nil {
		log.Fatal(err)
	}

	err = createData(db)
	if err != nil {
		log.Fatal(err)
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

func createTableIfNotExists(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS books (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			author VARCHAR(255) NOT NULL,
			isbn VARCHAR(13) NOT NULL
		);
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error al crear la tabla: %v", err)
	}

	fmt.Println("Tabla 'books' creada correctamente")

	return nil
}

func createData(db *sql.DB) error {
	query := `SELECT COUNT(*) INTO @count FROM books`

	// querytwo := `INSERT INTO books (title, author, isbn)
	// SELECT * FROM(
	// 	SELECT 'Cien años de soledad', 'Gabriel García Márquez', '9788437604947' UNION ALL
	// 	SELECT 'Harry Potter y la piedra filosofal', 'J.K. Rowling', '9788478886456' UNION ALL
	// 	SELECT 'El código Da Vinci', 'Dan Brown', '9780307474278') UNION ALL
	// 	SELECT '1984', 'George Orwell', '9788420676778') UNION ALL
	// 	SELECT 'El Alquimista', 'Paulo Coelho', '9780062315007') UNION ALL
	// 	SELECT 'To Kill a Mockingbird', 'Harper Lee', '9780060935467') UNION ALL
	// 	SELECT 'The Catcher in the Rye', 'J.D. Salinger', '9780316769488') UNION ALL
	// 	SELECT 'The Great Gatsby', 'F. Scott Fitzgerald', '9780743273565') UNION ALL
	// 	SELECT 'The Lord of the Rings', 'J.R.R. Tolkien', '9780618640157') UNION ALL
	// 	SELECT 'Pride and Prejudice', 'Jane Austen', '9780141439518') UNION ALL
	// 	SELECT 'The Chronicles of Narnia', 'C.S. Lewis', '9780064404990') UNION ALL
	// 	SELECT 'The Hunger Games', 'Suzanne Collins', '9780439023528') UNION ALL
	// 	SELECT 'Gone Girl', 'Gillian Flynn', '9780307588371') UNION ALL
	// 	SELECT 'The Girl on the Train', 'Paula Hawkins', '9781594634024') UNION ALL
	// 	SELECT 'The Da Vinci Code', 'Dan Brown', '9780307474278'  UNION ALL
	// 	SELECT 'The Fault in Our Stars', 'John Green', '9780525478812'  UNION ALL
	// 	SELECT 'The Help', 'Kathryn Stockett', '9780399155345'  UNION ALL
	// 	SELECT 'The Kite Runner', 'Khaled Hosseini', '9781594631931'  UNION ALL
	// 	SELECT 'The Girl with the Dragon Tattoo', 'Stieg Larsson', '9780307949486'  UNION ALL
	// 	SELECT 'The Hobbit', 'J.R.R. Tolkien', '9780547928227'  UNION ALL
	// 	SELECT 'The Maze Runner', 'James Dashner', '9780385737951'  UNION ALL
	// 	SELECT 'The Book Thief', 'Markus Zusak', '9780375842207'  UNION ALL
	// 	SELECT 'The Giver', 'Lois Lowry', '9780544340688'  UNION ALL
	// 	SELECT 'The Lovely Bones', 'Alice Sebold', '9780316044936'  UNION ALL
	// 	SELECT 'The Secret Life of Bees', 'Sue Monk Kidd', '9780142001745'  UNION ALL
	// 	SELECT 'The Time Traveler''s Wife', 'Audrey Niffenegger', '9781476764832'  UNION ALL
	// 	SELECT 'The Twilight Saga', 'Stephenie Meyer', '9780316067935'  UNION ALL
	// 	SELECT 'The Help', 'Kathryn Stockett', '9780399155345'  UNION ALL
	// 	SELECT 'The Pillars of the Earth', 'Ken Follett', '9780451166890'  UNION ALL
	// 	SELECT 'The Power of Now', 'Eckhart Tolle', '9781577314806'
	// ) AS data
	// WHERE @count = 0;
	// `

	result, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error al poblar la base de datos en la tabla: %v", err)
	}
	log.Println(result)

	fmt.Println("Tabla 'books' poblada correctamente")

	return nil
}
