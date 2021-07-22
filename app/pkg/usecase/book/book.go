package book

import (
	"context"
	"time"

	"github.com/bagus-aulia/dot-test/app/pkg/models"
	book_repo "github.com/bagus-aulia/dot-test/app/pkg/repository/book"
)

// Book interface for Book Usecase
type Book interface {
	GetBookList(c context.Context) ([]models.Book, error)
	GetBookByUUID(c context.Context, UUID string) (models.Book, error)
	CreateBook(c context.Context, title string, writer string) (models.Book, error)
	UpdateBook(c context.Context, UUID string, title string, writer string) (models.Book, error)
	DeleteBook(c context.Context, UUID string) (models.Book, error)
}

type bookUsecase struct {
	bookRepo       book_repo.Book
	contextTimeout time.Duration
}

// NewBookUsecase will create new an bookUsecase object representation of book interface
func NewBookUsecase(book book_repo.Book, timeout time.Duration) Book {
	return &bookUsecase{
		bookRepo:       book,
		contextTimeout: timeout,
	}
}

func (u *bookUsecase) GetBookList(c context.Context) ([]models.Book, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.bookRepo.GetBookList(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *bookUsecase) GetBookByUUID(c context.Context, UUID string) (models.Book, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.bookRepo.GetBookByUUID(ctx, UUID)

	return res, err
}

func (u *bookUsecase) CreateBook(c context.Context, title string, writer string) (models.Book, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.bookRepo.CreateBook(ctx, title, writer)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *bookUsecase) UpdateBook(c context.Context, UUID string, title string, writer string) (models.Book, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.bookRepo.UpdateBook(ctx, UUID, title, writer)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *bookUsecase) DeleteBook(c context.Context, UUID string) (models.Book, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.bookRepo.DeleteBook(ctx, UUID)
	if err != nil {
		return res, err
	}

	return res, nil
}
