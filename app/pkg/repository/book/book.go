package book

import (
	"context"
	"encoding/json"

	"github.com/bagus-aulia/dot-test/app/helpers"
	"github.com/bagus-aulia/dot-test/app/pkg/models"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// Book interface for Book Repository
type Book interface {
	GetBookList(ctx context.Context) ([]models.Book, error)
	GetBookByUUID(ctx context.Context, UUID string) (models.Book, error)
	CreateBook(ctx context.Context, title string, writer string) (models.Book, error)
	UpdateBook(ctx context.Context, UUID string, title string, writer string) (models.Book, error)
	DeleteBook(ctx context.Context, UUID string) (models.Book, error)
}

type conBookRepository struct {
	Conn  *gorm.DB
	Redis *redis.Client
}

var (
	redisKey = "dot-book"
)

// NewConBookRepository is a publisher of Book repository
func NewConBookRepository(Conn *gorm.DB, Redis *redis.Client) Book {
	return &conBookRepository{Conn, Redis}
}

// GetBookList to show all book
func (m *conBookRepository) GetBookList(ctx context.Context) ([]models.Book, error) {
	var books []models.Book

	// check redis data
	redisData, _ := helpers.GetRedisData(ctx, m.Redis, redisKey)
	if redisData != "" {
		if err := json.Unmarshal([]byte(redisData), &books); err != nil {
			return books, err
		}

		return books, nil
	}

	// get data from db
	result := m.Conn.WithContext(ctx).Find(&books)

	// set redis data
	dataJSON, _ := json.Marshal(books)
	err := helpers.SetRedisData(ctx, m.Redis, redisKey, string(dataJSON))
	if err != nil {
		return books, err
	}

	return books, result.Error
}

// GetBookByUUID to show single book data by uuid
func (m *conBookRepository) GetBookByUUID(ctx context.Context, UUID string) (models.Book, error) {
	var book models.Book

	// check redis data
	key := redisKey + "_" + UUID
	redisData, _ := helpers.GetRedisData(ctx, m.Redis, key)
	if redisData != "" {
		if err := json.Unmarshal([]byte(redisData), &book); err != nil {
			return book, err
		}

		return book, nil
	}

	// get data from db
	result := m.Conn.WithContext(ctx).Where("uuid = ?", UUID).First(&book)

	// set redis data
	dataJSON, _ := json.Marshal(book)
	err := helpers.SetRedisData(ctx, m.Redis, key, string(dataJSON))
	if err != nil {
		return book, err
	}

	return book, result.Error
}

// CreateBook to insert book data
func (m *conBookRepository) CreateBook(ctx context.Context, title string, writer string) (models.Book, error) {
	// delete redis data
	helpers.DelRedisData(ctx, m.Redis, redisKey)

	book := models.Book{
		Title:  title,
		Writer: &writer,
		UUID:   helpers.GenerateUUID("Edition"),
	}

	result := m.Conn.WithContext(ctx).Create(&book)

	return book, result.Error
}

// UpdateBook to update book data
func (m *conBookRepository) UpdateBook(ctx context.Context, UUID string, title string, writer string) (models.Book, error) {
	// delete redis data
	key := redisKey + "_" + UUID
	helpers.DelRedisData(ctx, m.Redis, redisKey)
	helpers.DelRedisData(ctx, m.Redis, key)

	book := models.Book{}
	result := m.Conn.WithContext(ctx).Where("uuid = ?", UUID).First(&book)
	if result.Error != nil {
		return book, result.Error
	}

	updated := models.Book{
		Title:  title,
		Writer: &writer,
	}
	result = m.Conn.WithContext(ctx).Model(&book).Updates(updated)

	return book, result.Error
}

// DeleteBook to delete book data
func (m *conBookRepository) DeleteBook(ctx context.Context, UUID string) (models.Book, error) {
	// delete redis data
	key := redisKey + "_" + UUID
	helpers.DelRedisData(ctx, m.Redis, redisKey)
	helpers.DelRedisData(ctx, m.Redis, key)

	book := models.Book{}
	result := m.Conn.WithContext(ctx).Where("uuid = ?", UUID).First(&book)
	if result.Error != nil {
		return book, result.Error
	}

	result = m.Conn.WithContext(ctx).Delete(&book)

	return book, result.Error
}
