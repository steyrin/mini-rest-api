package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/steyrin/mini-rest-api/internal/model"
	"github.com/steyrin/mini-rest-api/internal/service"
	"net/http"
)

type BookHandler struct {
	BookService service.BookService
}

func NewBookHandler(r *gin.Engine, bs service.BookService) {
	handler := &BookHandler{BookService: bs}

	r.GET("/books", handler.GetBooks)
	r.POST("/books", handler.AddBook)
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	ctx := c.Request.Context()
	books, err := h.BookService.GetBooks(ctx)
	if err != nil {
		logrus.Errorf("Ошибка при получении книг: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить книги"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) AddBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		logrus.Errorf("Ошибка при привязке данных: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	ctx := c.Request.Context()
	newBook, err := h.BookService.AddBook(ctx, &book)
	if err != nil {
		logrus.Errorf("Ошибка при добавлении книги: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось добавить книгу"})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}
