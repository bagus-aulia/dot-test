package borrow

import (
	"context"
	"time"

	"github.com/bagus-aulia/dot-test/app/pkg/models"
	borrow_repo "github.com/bagus-aulia/dot-test/app/pkg/repository/borrow"
)

// Borrow interface for Borrow Usecase
type Borrow interface {
	GetBorrowList(c context.Context) ([]models.Borrow, error)
	GetBorrowByUUID(ctx context.Context, UUID string) (models.Borrow, error)
	CreateBorrow(ctx context.Context, userUUID string, bookUUIDs []string) (*models.Borrow, error)
	ReturnBorrow(ctx context.Context, UUID string) (models.Borrow, error)
}

type borrowUsecase struct {
	borrowRepo     borrow_repo.Borrow
	contextTimeout time.Duration
}

// NewBorrowUsecase will create new an borrowUsecase object representation of borrow interface
func NewBorrowUsecase(borrow borrow_repo.Borrow, timeout time.Duration) Borrow {
	return &borrowUsecase{
		borrowRepo:     borrow,
		contextTimeout: timeout,
	}
}

func (u *borrowUsecase) GetBorrowList(c context.Context) ([]models.Borrow, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.borrowRepo.GetBorrowList(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *borrowUsecase) GetBorrowByUUID(c context.Context, UUID string) (models.Borrow, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.borrowRepo.GetBorrowByUUID(ctx, UUID)

	return res, err
}

func (u *borrowUsecase) CreateBorrow(c context.Context, userUUID string, bookUUIDs []string) (*models.Borrow, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.borrowRepo.CreateBorrow(ctx, userUUID, bookUUIDs)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (u *borrowUsecase) ReturnBorrow(c context.Context, UUID string) (models.Borrow, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.borrowRepo.ReturnBorrow(ctx, UUID)

	return res, err
}
