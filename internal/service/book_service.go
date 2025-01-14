package service

import (
	"context" // Импортируем контекст
	"github.com/steyrin/mini-rest-api/internal/model"
	"github.com/steyrin/mini-rest-api/internal/repository"
)

type BookService interface {
	GetBooks(ctx context.Context) ([]model.Book, error)
	AddBook(ctx context.Context, book *model.Book) (*model.Book, error)
}

type bookService struct {
	repository repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repository: repo}
}

func (s *bookService) GetBooks(ctx context.Context) ([]model.Book, error) {
	return s.repository.GetAllBooks(ctx)
}

func (s *bookService) AddBook(ctx context.Context, book *model.Book) (*model.Book, error) {
	return s.repository.SaveBook(ctx, book)
}
