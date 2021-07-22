package book

import (
	"time"

	book_handler "github.com/bagus-aulia/dot-test/app/pkg/delivery/book"
	book_repo "github.com/bagus-aulia/dot-test/app/pkg/repository/book"
	book_usecase "github.com/bagus-aulia/dot-test/app/pkg/usecase/book"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Server for Book
func Server(e *echo.Echo, DB *gorm.DB, redis *redis.Client, timeout time.Duration) {
	bookRepo := book_repo.NewConBookRepository(DB, redis)
	bookUsecase := book_usecase.NewBookUsecase(bookRepo, timeout)

	book_handler.NewBookHandler(e, bookUsecase)
}
