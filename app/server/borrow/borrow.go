package borrow

import (
	"time"

	borrow_handler "github.com/bagus-aulia/dot-test/app/pkg/delivery/borrow"
	borrow_repo "github.com/bagus-aulia/dot-test/app/pkg/repository/borrow"
	borrow_usecase "github.com/bagus-aulia/dot-test/app/pkg/usecase/borrow"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Server for Borrw
func Server(e *echo.Echo, DB *gorm.DB, redis *redis.Client, timeout time.Duration) {
	borrowRepo := borrow_repo.NewConBorrowRepository(DB, redis)
	borrowUsecase := borrow_usecase.NewBorrowUsecase(borrowRepo, timeout)

	borrow_handler.NewBorrowHandler(e, borrowUsecase)
}
