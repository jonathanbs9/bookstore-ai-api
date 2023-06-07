package main

import (
	//"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/jonathanbs9/bookstore-ai-api/controllers"
	"github.com/jonathanbs9/bookstore-ai-api/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	ID     int    `gorm:"primaryKey;autoIncrement"`
	Title  string `gorm:"not null"`
	Author string `gorm:"not null"`
	ISBN   string `gorm:"not null"`
}

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

	//db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/book_inventory")
	//if err != nil {
	//	log.Fatal("Error al conectar a la base de datos: ", err)
	//}
	//defer db.Close()

	//dsn := "jonathanbs:ihmQFPAYETnygjZodt49EA@tcp(azure-mule-888.g8x.cockroachlabs.cloud:26257)/book_inventory"
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	dsn := "postgresql://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=verify-full"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	// Ejecutar script para crear la tabla si no existe
	err = createTableIfNotExists(db)
	if err != nil {
		log.Fatal(err)
	}

	//err = createData(db)
	err = createDataWithGorm(db)
	if err != nil {
		log.Fatal(err)
	}

	// Crear una instancia del enrutador Gin
	router := gin.Default()

	// Inicializar los controladores
	//bookController := controllers.NewBookController(db)
	bookController := controllers.NewBookControllerWithGorm(db)

	// Definir las rutas
	routes.SetupRoutes(router, bookController)

	// Iniciar el servidor HTTP
	err = router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal("Error al iniciar el servidor HTTP: ", err)
	}

}

// Create Table if NOT Exist
func createTableIfNotExists(db *gorm.DB) error {

	err := db.AutoMigrate(&Book{})
	if err != nil {
		return fmt.Errorf("error al crear la tabla: %v", err)
	}

	fmt.Println("Tabla 'books' creada correctamente")

	return nil
}

// Create Data
/*func createData(db *sql.DB) error {
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
}*/

func createDataWithGorm(db *gorm.DB) error {
	// Crea la tabla "books" si no existe
	err := db.AutoMigrate(&Book{})
	if err != nil {
		return fmt.Errorf("error al migrar la tabla 'books': %v", err)
	}

	// Verifica si ya hay registros en la tabla
	var count int64
	if err := db.Model(&Book{}).Count(&count).Error; err != nil {
		return fmt.Errorf("error al contar registros en la tabla 'books': %v", err)
	}

	if count == 0 {
		// Inserta los registros en la tabla "books"
		books := []Book{
			{Title: "Cien años de soledad", Author: "Gabriel García Márquez", ISBN: "9788437604947"},
			{Title: "Harry Potter y la piedra filosofal", Author: "J.K. Rowling", ISBN: "9788478886456"},
			{Title: "El código Da Vinci", Author: "Dan Brown", ISBN: "9780307474278"},
			{Title: "1984", Author: "George Orwell", ISBN: "9788420676778"},
			{Title: "El Alquimista", Author: "Paulo Coelho", ISBN: "9780062315007"},
			{Title: "To Kill a Mockingbird", Author: "Harper Lee", ISBN: "9780060935467"},
			{Title: "The Catcher in the Rye", Author: "J.D. Salinger", ISBN: "9780316769488"},
			{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", ISBN: "9780743273565"},
			{Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", ISBN: "9780618640157"},
			{Title: "Pride and Prejudice", Author: "Jane Austen", ISBN: "9780141439518"},
			{Title: "The Chronicles of Narnia", Author: "C.S. Lewis", ISBN: "9780064404990"},
			{Title: "The Hunger Games", Author: "Suzanne Collins", ISBN: "9780439023528"},
			{Title: "Gone Girl", Author: "Gillian Flynn", ISBN: "9780307588371"},
			{Title: "The Girl on the Train", Author: "Paula Hawkins", ISBN: "9781594634024"},
			{Title: "The Da Vinci Code", Author: "Dan Brown", ISBN: "9780307474278"},
			{Title: "The Fault in Our Stars", Author: "John Green", ISBN: "9780525478812"},
			{Title: "The Help", Author: "Kathryn Stockett", ISBN: "9780399155345"},
			{Title: "The Kite Runner", Author: "Khaled Hosseini", ISBN: "9781594631931"},
			{Title: "The Girl with the Dragon Tattoo", Author: "Stieg Larsson", ISBN: "9780307949486"},
			{Title: "The Hobbit", Author: "J.R.R. Tolkien", ISBN: "9780547928227"},
			{Title: "The Maze Runner", Author: "James Dashner", ISBN: "9780385737951"},
			{Title: "The Book Thief", Author: "Markus Zusak", ISBN: "9780375842207"},
			{Title: "The Giver", Author: "Lois Lowry", ISBN: "9780544340688"},
			{Title: "The Lovely Bones", Author: "Alice Sebold", ISBN: "9780316044936"},
			{Title: "The Secret Life of Bees", Author: "Sue Monk Kidd", ISBN: "9780142001745"},
			{Title: "The Time Traveler's Wife", Author: "Audrey Niffenegger", ISBN: "9781476764832"},
			{Title: "The Twilight Saga", Author: "Stephenie Meyer", ISBN: "9780316067935"},
			{Title: "The Help", Author: "Kathryn Stockett", ISBN: "9780399155345"},
			{Title: "The Pillars of the Earth", Author: "Ken Follett", ISBN: "9780451166890"},
			{Title: "The Power of Now", Author: "Eckhart Tolle", ISBN: "9781577314806"},
		}

		err = db.Create(&books).Error
		if err != nil {
			return fmt.Errorf("error al insertar registros en la tabla 'books': %v", err)
		}

		fmt.Println("Tabla 'books' poblada correctamente")
	} else {
		fmt.Println("La tabla 'books' ya contiene registros")
	}

	return nil
}
