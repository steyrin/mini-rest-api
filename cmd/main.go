package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/steyrin/mini-rest-api/config"
	"github.com/steyrin/mini-rest-api/internal/handler"
	"github.com/steyrin/mini-rest-api/internal/repository"
	"github.com/steyrin/mini-rest-api/internal/service"
	"github.com/steyrin/mini-rest-api/internal/tracer"
	"log"
)

func main() {

	shutdown := tracer.InitTracer(false)
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Fatalf("Error shutting down tracer: %v", err)
		}
	}()

	db := config.InitDB()

	bookRepo := repository.NewBookRepository(db)

	bookService := service.NewBookService(bookRepo)

	r := gin.Default()

	err := repository.InitializeBooks(db, context.Background())
	if err != nil {
		log.Fatalf("Ошибка при инициализации книг: %v", err)
	}

	handler.NewBookHandler(r, bookService)

	//router.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "Test!",
	//	})
	//})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
