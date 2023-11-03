package main

import (
	_ "golang/crud-basic/docs"

	"net/http"

	"errors"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

// @title 		Basic Golang CRUD
// @version 	1.0
// @description A Testing Swagger for My Project

// @host 		localhost:8080
// @BasePath	/

type book struct {
	ID       string `json:"id" example:"1"`
	Title    string `json:"title" example:"Feb's Tutor"`
	Author   string `json:"author" example:"Febs"`
	Quantity int    `json:"quantity" example:"10"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// GetBooks			godoc
// @Summary			Get All Books
// @Description		Return All books Data
// @Tags			Books
// @Success			200 {object} []book
// @Router			/books [get]
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// GetBookByID		godoc
// @Summary			Get Book by ID
// @Description		Return book Data
// @Tags			Books
// @Param			id path string true "search by id"
// @Success			200 {object} book
// @Router			/books/{id} [get]
func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// CheckoutBook		godoc
// @Summary			Checkout Book
// @Description		Checkout book Data
// @Tags			Books
// @Param			id query string true "id for checkout"
// @Success			200 {object} book
// @Router			/checkout [patch]
func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"messsage": "Book not found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"messsage": "Book not available."})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

// ReturnBook		godoc
// @Summary			Return Book
// @Description		Return book Data
// @Tags			Books
// @Param			id query string true "id for checkout"
// @Success			200 {object} book
// @Router			/return [patch]
func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"messsage": "Book not found."})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

var PG_CONN = "postgresql://yomi:yomi123@localhost:5432/main"
var crudDB *gorm.DB

func main() {
	db, err := gorm.Open(postgres.Open(PG_CONN), &gorm.Config{})

	if err != nil {
		return
	}

	crudDB = db

	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/books/:id", bookById)
	router.GET("/books", getBooks)
	router.PATCH("/return", returnBook)
	router.PATCH("/checkout", checkoutBook)
	router.POST("/books", createBook)
	router.Run()
}
