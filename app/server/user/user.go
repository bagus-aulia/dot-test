package user

import (
	"time"

	user_handler "github.com/bagus-aulia/dot-test/app/pkg/delivery/user"
	user_repo "github.com/bagus-aulia/dot-test/app/pkg/repository/user"
	user_usecase "github.com/bagus-aulia/dot-test/app/pkg/usecase/user"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Server for User
func Server(e *echo.Echo, DB *gorm.DB, redis *redis.Client, timeout time.Duration) {
	userRepo := user_repo.NewConUserRepository(DB, redis)
	userUsecase := user_usecase.NewUserUsecase(userRepo, timeout)

	user_handler.NewUserHandler(e, userUsecase)
}
