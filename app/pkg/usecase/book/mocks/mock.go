package mocks

import (
	"context"

	"github.com/bagus-aulia/dot-test/app/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// BookUsecase is an autogenerated mock type for the BookUsecase type
type BookUsecase struct {
	mock.Mock
}

// GetBookList provides a mock function with given fields: ctx
func (_m *BookUsecase) GetBookList(ctx context.Context) ([]models.Book, error) {
	ret := _m.Called(ctx)

	var r0 []models.Book
	if rf, ok := ret.Get(0).(func(context.Context) []models.Book); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).([]models.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBookByUUID provides a mock function with given fields: ctx, uuid
func (_m *BookUsecase) GetBookByUUID(ctx context.Context, UUID string) (models.Book, error) {
	ret := _m.Called(ctx, UUID)

	var r0 models.Book
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Book); ok {
		r0 = rf(ctx, UUID)
	} else {
		r0 = ret.Get(0).(models.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, UUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateBook provides a mock function with given fields: ctx, title, writer
func (_m *BookUsecase) CreateBook(ctx context.Context, title string, writer string) (models.Book, error) {
	ret := _m.Called(ctx, title, writer)

	var r0 models.Book
	if rf, ok := ret.Get(0).(func(context.Context, string, string) models.Book); ok {
		r0 = rf(ctx, title, writer)
	} else {
		r0 = ret.Get(0).(models.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, title, writer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBook provides a mock function with given fields: ctx, uuid, title, writer
func (_m *BookUsecase) UpdateBook(ctx context.Context, UUID string, title string, writer string) (models.Book, error) {
	ret := _m.Called(ctx, UUID, title, writer)

	var r0 models.Book
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) models.Book); ok {
		r0 = rf(ctx, UUID, title, writer)
	} else {
		r0 = ret.Get(0).(models.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, UUID, title, writer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBook provides a mock function with given fields: ctx, uuid
func (_m *BookUsecase) DeleteBook(ctx context.Context, UUID string) (models.Book, error) {
	ret := _m.Called(ctx, UUID)

	var r0 models.Book
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Book); ok {
		r0 = rf(ctx, UUID)
	} else {
		r0 = ret.Get(0).(models.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, UUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
