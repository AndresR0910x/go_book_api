package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Crea una variable DB de manera global
var DB *gorm.DB

func InitDB() {
	// Verifica que las variables de entorno se encuentren cargadas !
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Inicializa la cadena de conexion
	dsn := os.Getenv("DB_URL")

	// Inicializa la variable DB usando Gorm y pasando el dns usando postgres.open , ademas se pasa la configuracion de gorm

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Si existe un error
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = DB.AutoMigrate(&Book{}); err != nil {
		log.Fatal("Failed to migrate schema: ", err)
	}
}

// C
func CreateBook(c *gin.Context) {
	var book Book

	// bind the request body

	if err := c.ShouldBindJSON(&book); err != nil {
		ResponseJSON(c, http.StatusBadRequest, "Invalid data", nil)
		return 
	}

	DB.Create(&book)
	ResponseJSON(c, http.StatusCreated, "Book created !", &book)
}

// R
// Get the list of available books inside the DB

func GetBooks(c *gin.Context) {
	var books []Book
	DB.Find(&books)

	ResponseJSON(c, http.StatusOK, "Books retrieved successfuly", books)
}

// Get a single book

func GetBook(c *gin.Context) {
	var book Book

	if err := DB.First(&book, c.Param("id")).Error; err != nil {
		ResponseJSON(c, http.StatusNotFound, "Book not found", nil)

		return
	}
	ResponseJSON(c, http.StatusOK, "Book retrieved successfully", book)
}


// Update a book 

func UpdateBook (c *gin.Context) {
	var book Book

	if err := DB.First(&book, c.Param("id")).Error; err != nil {
		ResponseJSON(c, http.StatusNotFound, "Book not found", nil)

		return
	}

	// bind the request body 

	if err := c.ShouldBindJSON(&book); err != nil {
		ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
	}

	DB.Save(&book)

	ResponseJSON(c, http.StatusOK, "Book updated successfully", nil)
}


// Delete a book

func DeleteBook(c *gin.Context) {
	var book Book

	if err := DB.Delete(&book, c.Param("id")).Error; err != nil {
		ResponseJSON(c, http.StatusNotFound, "Book not found", nil)
		return 
	}

	ResponseJSON(c, http.StatusOK, "Book deleted successfully", nil)
}