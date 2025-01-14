package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/steyrin/mini-rest-api/internal/model"
	"github.com/steyrin/mini-rest-api/internal/service"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
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

//func (h *BookHandler) GetBooks(c *gin.Context) {
//	ctx := c.Request.Context()
//	books, err := h.BookService.GetBooks(ctx)
//	if err != nil {
//		logrus.Errorf("Ошибка при получении книг: %v", err)
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить книги"})
//		return
//	}
//	c.JSON(http.StatusOK, books)
//}

func (h *BookHandler) GetBooks(c *gin.Context) {
	ctx := c.Request.Context()

	// инит трейс
	tracer := otel.Tracer("book-handler")
	ctx, span := tracer.Start(ctx, "GetBooks")
	defer span.End() // завершаем трейс в конце выполнения метода

	books, err := h.BookService.GetBooks(ctx)
	if err != nil {
		logrus.Errorf("Ошибка при получении книг: %v", err)

		// добавляем ошибку в трейс
		span.RecordError(err)
		span.SetAttributes(attribute.String("error.message", err.Error()))
		span.SetStatus(codes.Error, "Не удалось получить книги")

		// Возвращаем ответ с ошибкой
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить книги"})
		return
	}

	span.SetStatus(codes.Ok, "Успешно получили книги")
	c.JSON(http.StatusOK, books)
}

//func (h *BookHandler) AddBook(c *gin.Context) {
//	var book model.Book
//	if err := c.ShouldBindJSON(&book); err != nil {
//		logrus.Errorf("Ошибка при привязке данных: %v", err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
//		return
//	}
//
//	ctx := c.Request.Context()
//	newBook, err := h.BookService.AddBook(ctx, &book)
//	if err != nil {
//		logrus.Errorf("Ошибка при добавлении книги: %v", err)
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось добавить книгу"})
//		return
//	}
//
//	c.JSON(http.StatusCreated, newBook)
//}

func (h *BookHandler) AddBook(c *gin.Context) {
	ctx := c.Request.Context()

	tracer := otel.Tracer("book-handler")
	ctx, span := tracer.Start(ctx, "AddBook")
	defer span.End()

	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {

		// Логируем и записываем ошибку
		logrus.Errorf("Ошибка при парсинге данных книги: %v", err)
		span.RecordError(err)
		span.SetAttributes(attribute.String("error.message", err.Error()))
		span.SetStatus(codes.Error, "Invalid input data")

		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	createdBook, err := h.BookService.AddBook(ctx, &book)
	if err != nil {
		// Логируем и записываем ошибку
		logrus.Errorf("Ошибка при сохранении книги: %v", err)
		span.RecordError(err)
		span.SetAttributes(attribute.String("error.message", err.Error()))
		span.SetStatus(codes.Error, "Failed to save book")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить книгу"})
		return
	}

	span.SetAttributes(
		attribute.Int64("book.id", createdBook.ID),
		attribute.String("book.name", createdBook.Name),
	)
	span.SetStatus(codes.Ok, "Book added successfully")

	c.JSON(http.StatusCreated, createdBook)
}
