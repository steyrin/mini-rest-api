package repository

import (
	"context"
	"github.com/steyrin/mini-rest-api/internal/model"
	"github.com/uptrace/bun"
	"log"
)

type BookRepository interface {
	GetAllBooks(ctx context.Context) ([]model.Book, error)
	SaveBook(ctx context.Context, book *model.Book) (*model.Book, error)
}

type bookRepository struct {
	db *bun.DB
}

func NewBookRepository(db *bun.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	var books []model.Book
	err := r.db.NewSelect().Model(&books).Scan(ctx)
	return books, err
}

func (r *bookRepository) SaveBook(ctx context.Context, book *model.Book) (*model.Book, error) {
	_, err := r.db.NewInsert().Model(book).Exec(ctx)
	return book, err
}

func InitializeBooks(db *bun.DB, ctx context.Context) error {
	var books []model.Book
	err := db.NewSelect().Model(&books).Limit(1).Scan(ctx)
	if err != nil {
		log.Printf("Ошибка при проверке книг в базе данных: %v", err)
		return err
	}

	if len(books) == 0 {
		initialBooks := []model.Book{
			{
				Name:        "The Go Programming Language",
				Genre:       "Programming",
				Description: "A comprehensive book about Go programming.",
				Year:        2015,
				Rating:      4.8,
				Price:       39.99,
			},
			{
				Name:        "Learning Python",
				Genre:       "Programming",
				Description: "An in-depth guide to Python.",
				Year:        2020,
				Rating:      4.7,
				Price:       29.99,
			},
		}

		_, err := db.NewInsert().Model(&initialBooks).Exec(ctx)
		if err != nil {
			log.Printf("Ошибка при добавлении книг: %v", err)
			return err
		}

		log.Println("Базы данных книги были успешно добавлены!")
	}

	return nil
}
