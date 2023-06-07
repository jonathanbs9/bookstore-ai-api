package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonathanbs9/bookstore-ai-api/models"
	"gorm.io/gorm"
)

/* type BookController struct {
	db *sql.DB
} */

type BookController struct {
	db *gorm.DB
}

// NewBookController crea una nueva instancia de BookController
/*func NewBookController(db *sql.DB) *BookController {
	return &BookController{
		db: db,
	}
}*/

func NewBookControllerWithGorm(db *gorm.DB) *BookController {
	return &BookController{
		db: db,
	}
}

func (bc *BookController) ListBooks(c *gin.Context) {
	// Consultar la base de datos para obtener los libros y devolver
	// la lista como respuesta JSON
	//books, err := bc.getBooksFromDatabase()
	books, err := bc.getBooksFromDatabaseWithGorm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la lista de libros"})
		return
	}
	c.JSON(http.StatusOK, books)
}

/*func (bc *BookController) getBooksFromDatabase() ([]models.Book, error) {
	query := "SELECT id, title, author, isbn FROM books"

	rows, err := bc.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}*/

func (bc *BookController) getBooksFromDatabaseWithGorm() ([]models.Book, error) {
	var books []models.Book

	result := bc.db.Select("id, title, author, isbn").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

// GetBookByName busca un libro por nombre
/*func (bc *BookController) GetBookByName(c *gin.Context) {
	name := c.Param("name")

	// Consultar la base de datos para obtener el libro por nombre
	book, err := bc.GetBookByNameFromDatabase(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el libro"})
	}

	// Verificar si se encontró un libro con el nombre dado
	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Libro no encontrado"})
		return
	}
	//Devolver el libro como respuesta JSON
	c.JSON(http.StatusOK, book)
}

func (bc *BookController) GetBookByNameFromDatabase(name string) (*models.Book, error) {
	query := "SELECT id, title, author, isbn FROM books WHERE title = ?"

	row := bc.db.QueryRow(query, name)

	var book models.Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No se encontró ningún libro
		}
		return nil, err
	}
	return &book, nil
}

// GetBookByAuthor busca un libro por autor
func (bc *BookController) GetBookByAuthor(c *gin.Context) {
	//author := c.Param("author")

	// Consultar la base de datos para obtener el libro por autor
	// y devolver el libro como respuesta JSON
}

// GetBookByISBN busca un libro por ISBN
func (bc *BookController) GetBookByISBN(c *gin.Context) {
	//isbn := c.Param("isbn")

	// Consultar la base de datos para obtener el libro por ISBN
	// y devolver el libro como respuesta JSON
}

// CreateBook da de alta un nuevo libro
func (bc *BookController) CreateBook(c *gin.Context) {
	//var book models.Book

	// Bind the JSON request body to the book struct

	// Insert the book into the database

	// Return the created book as the JSON response
}

func (bc *BookController) GetBookById(c *gin.Context) {
	id := c.Param("id")

	// Consulto BD
	book, err := bc.GetBookByIdFromDatabase(id)
	log.Println("Book: ", book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el libro"})
	}
	// Verificar si encontró un libro por ID
	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Libro por ID no encontrado"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (bc *BookController) GetBookByIdFromDatabase(id string) (*models.Book, error) {
	query := "SELECT id, title, author, isbn FROM books WHERE id = ?"

	row := bc.db.QueryRow(query, id)

	var book models.Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("NO ENCUENTRA LIBRO -> ASIGNA NIL")
			return nil, nil // No se encontró ningún libro
		}
		return nil, err
	}
	//log.Println(&book)
	return &book, nil
}

// UpdateBook modifica un libro existente
func (bc *BookController) UpdateBook(c *gin.Context) {
	//idParam := c.Param("id")
	//id, err := strconv.Atoi(idParam)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
	//	return
	//}

	//var book models.Book

	// Bind the JSON request body to the book struct

	// Update the book in the database

	// Return the updated book as the JSON response
}*/
