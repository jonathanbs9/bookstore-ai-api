package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"dev.azure.com/jbsorg/segundo_proyecto/_git/bookstore-ai-api/models"
)

type BookController struct {
	db *sql.DB
}

// NewBookController crea una nueva instancia de BookController
type NewBookController(db *sql.DB) *BookController {
	return &BookController{
		db: db,
	}
}

func (bc *BookController) ListBooks(c *gin.Context){
	// Consultar la base de datos para obtener los libros y devolver
	// la lista como respuesta JSON
}

// GetBookByName busca un libro por nombre
func (bc *BookController) GetBookByName(c *gin.Context) {
	name := c.Param("name")

	// Consultar la base de datos para obtener el libro por nombre
	// y devolver el libro como respuesta JSON
}

// GetBookByAuthor busca un libro por autor
func (bc *BookController) GetBookByAuthor(c *gin.Context) {
	author := c.Param("author")

	// Consultar la base de datos para obtener el libro por autor
	// y devolver el libro como respuesta JSON
}

// GetBookByISBN busca un libro por ISBN
func (bc *BookController) GetBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")

	// Consultar la base de datos para obtener el libro por ISBN
	// y devolver el libro como respuesta JSON
}

// CreateBook da de alta un nuevo libro
func (bc *BookController) CreateBook(c *gin.Context) {
	var book models.Book

	// Bind the JSON request body to the book struct

	// Insert the book into the database

	// Return the created book as the JSON response
}

// UpdateBook modifica un libro existente
func (bc *BookController) UpdateBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var book models.Book

	// Bind the JSON request body to the book struct

	// Update the book in the database

	// Return the updated book as the JSON response
}