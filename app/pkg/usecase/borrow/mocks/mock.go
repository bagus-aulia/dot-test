package mocks

import (
	"context"

	"github.com/bagus-aulia/dot-test/app/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// BorrowUsecase is an autogenerated mock type for the BorrowUsecase type
type BorrowUsecase struct {
	mock.Mock
}

// GetBorrowList provides a mock function with given fields: ctx
func (_m *BorrowUsecase) GetBorrowList(ctx context.Context) ([]models.Borrow, error) {
	ret := _m.Called(ctx)

	var r0 []models.Borrow
	if rf, ok := ret.Get(0).(func(context.Context) []models.Borrow); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).([]models.Borrow)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBorrowByUUID provides a mock function with given fields: ctx, uuid
func (_m *BorrowUsecase) GetBorrowByUUID(ctx context.Context, UUID string) (models.Borrow, error) {
	ret := _m.Called(ctx, UUID)

	var r0 models.Borrow
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Borrow); ok {
		r0 = rf(ctx, UUID)
	} else {
		r0 = ret.Get(0).(models.Borrow)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, UUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateBorrow provides a mock function with given fields: ctx, userUUID, bookUUIDs
func (_m *BorrowUsecase) CreateBorrow(ctx context.Context, userUUID string, bookUUIDs []string) (*models.Borrow, error) {
	ret := _m.Called(ctx, userUUID, bookUUIDs)

	var r0 *models.Borrow
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) *models.Borrow); ok {
		r0 = rf(ctx, userUUID, bookUUIDs)
	} else {
		r0 = ret.Get(0).(*models.Borrow)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, userUUID, bookUUIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReturnBorrow provides a mock function with given fields: ctx, uuid
func (_m *BorrowUsecase) ReturnBorrow(ctx context.Context, UUID string) (models.Borrow, error) {
	ret := _m.Called(ctx, UUID)

	var r0 models.Borrow
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Borrow); ok {
		r0 = rf(ctx, UUID)
	} else {
		r0 = ret.Get(0).(models.Borrow)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, UUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}