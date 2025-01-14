package service

import (
	"context"
	"testing"

	"github.com/steyrin/mini-rest-api/internal/model"
	"github.com/steyrin/mini-rest-api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBookService_GetBooks(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.BookRepository)
	mockRepo.On("GetAllBooks", mock.Anything).Return([]model.Book{
		{ID: 1, Name: "Test Book", Genre: "Fiction"},
	}, nil)

	service := NewBookService(mockRepo)

	// Act
	ctx := context.Background()
	books, err := service.GetBooks(ctx)

	// Assert
	assert.NoError(t, err)
	assert.Len(t, books, 1)
	assert.Equal(t, "Test Book", books[0].Name)

	mockRepo.AssertExpectations(t)
}

//func TestBookService_AddBook(t *testing.T) {
//	tests := []struct {
//		name    string
//		input   model.Book
//		wantErr bool
//	}{
//		{name: "Valid Book", input: model.Book{Name: "Valid", Genre: "Fiction"}, wantErr: false},
//		{name: "Missing Name", input: model.Book{Genre: "Fiction"}, wantErr: true},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mockRepo := new(mocks.BookRepository)
//			if !tt.wantErr {
//				mockRepo.On("SaveBook", mock.Anything, &tt.input).Return(&tt.input, nil)
//			} else {
//				mockRepo.On("SaveBook", mock.Anything, &tt.input).Return(nil, errors.New("validation error"))
//			}
//
//			mockRepo.AssertExpectations(t)
//		})
//	}
//}
