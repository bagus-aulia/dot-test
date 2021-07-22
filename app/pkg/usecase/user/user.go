package user

import (
	"context"
	"time"

	"github.com/bagus-aulia/dot-test/app/pkg/models"
	user_repo "github.com/bagus-aulia/dot-test/app/pkg/repository/user"
)

// User interface for User Usecase
type User interface {
	GetUserList(c context.Context) ([]models.User, error)
	GetUserByUUID(c context.Context, UUID string) (models.User, error)
	CreateUser(c context.Context, title string, writer string) (models.User, error)
	UpdateUser(c context.Context, UUID string, title string, writer string) (models.User, error)
}

type userUsecase struct {
	userRepo       user_repo.User
	contextTimeout time.Duration
}

// NewUserUsecase will create new an userUsecase object representation of user interface
func NewUserUsecase(user user_repo.User, timeout time.Duration) User {
	return &userUsecase{
		userRepo:       user,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) GetUserList(c context.Context) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.userRepo.GetUserList(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userUsecase) GetUserByUUID(c context.Context, UUID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.userRepo.GetUserByUUID(ctx, UUID)

	return res, err
}

func (u *userUsecase) CreateUser(c context.Context, title string, writer string) (models.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.userRepo.CreateUser(ctx, title, writer)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *userUsecase) UpdateUser(c context.Context, UUID string, title string, writer string) (models.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.userRepo.UpdateUser(ctx, UUID, title, writer)
	if err != nil {
		return res, err
	}

	return res, nil
}
