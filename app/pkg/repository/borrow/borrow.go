package borrow

import (
	"context"
	"encoding/json"
	"time"

	"github.com/bagus-aulia/dot-test/app/helpers"
	"github.com/bagus-aulia/dot-test/app/pkg/models"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// Borrow interface for Borrow Repository
type Borrow interface {
	GetBorrowList(ctx context.Context) ([]models.Borrow, error)
	GetBorrowByUUID(ctx context.Context, UUID string) (models.Borrow, error)
	CreateBorrow(ctx context.Context, userUUID string, bookUUIDs []string) (models.Borrow, error)
	ReturnBorrow(ctx context.Context, UUID string) (models.Borrow, error)
}

type conBorrowRepository struct {
	Conn  *gorm.DB
	Redis *redis.Client
}

var (
	redisKey = "dot-borrow"
)

// NewConBorrowRepository is a publisher of Borrow repository
func NewConBorrowRepository(Conn *gorm.DB, Redis *redis.Client) Borrow {
	return &conBorrowRepository{Conn, Redis}
}

// GetBorrowList to show all Borrow
func (m *conBorrowRepository) GetBorrowList(ctx context.Context) ([]models.Borrow, error) {
	var borrowRes []models.Borrow

	// check redis data
	redisData, _ := helpers.GetRedisData(ctx, m.Redis, redisKey)
	if redisData != "" {
		if err := json.Unmarshal([]byte(redisData), &borrowRes); err != nil {
			return borrowRes, err
		}

		return borrowRes, nil
	}

	// transaction to get data from db
	err := m.Conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// borrow list
		var borrows []models.Borrow
		if err := tx.Find(&borrows).Error; err != nil {
			return err
		}

		for _, borrow := range borrows {
			// collect user data
			var user models.User
			if err := tx.First(&user, borrow.UserID).Error; err != nil {
				return err
			}
			borrow.User = &user

			// collect borrow detail datas
			var details []models.BorrowDetail
			var detailRes []models.BorrowDetail
			if err := tx.Where("borrow_id = ?", borrow.ID).Find(&details).Error; err != nil {
				return err
			}

			for _, detail := range details {
				// collect book data
				var book models.Book
				if err := tx.First(&book, detail.BookID).Error; err != nil {
					return err
				}
				detail.Book = &book
				detailRes = append(detailRes, detail)
			}
			borrow.BorrowDetail = detailRes

			borrowRes = append(borrowRes, borrow)
		}

		return nil
	})

	// set redis data
	dataJSON, _ := json.Marshal(borrowRes)
	err = helpers.SetRedisData(ctx, m.Redis, redisKey, string(dataJSON))
	if err != nil {
		return borrowRes, err
	}

	return borrowRes, err
}

// GetBorrowByUUID to show single Borrow data by uuid
func (m *conBorrowRepository) GetBorrowByUUID(ctx context.Context, UUID string) (models.Borrow, error) {
	var borrowRes models.Borrow

	// check redis data
	key := redisKey + "_" + UUID
	redisData, _ := helpers.GetRedisData(ctx, m.Redis, key)
	if redisData != "" {
		if err := json.Unmarshal([]byte(redisData), &borrowRes); err != nil {
			return borrowRes, err
		}

		return borrowRes, nil
	}

	// transaction to get data from db
	err := m.Conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// borrow data
		var borrow models.Borrow
		if err := tx.Where("uuid = ?", UUID).First(&borrow).Error; err != nil {
			return err
		}

		// collect user data
		var user models.User
		if err := tx.First(&user, borrow.UserID).Error; err != nil {
			return err
		}
		borrow.User = &user

		// collect borrow detail datas
		var details []models.BorrowDetail
		var detailRes []models.BorrowDetail
		if err := tx.Where("borrow_id = ?", borrow.ID).Find(&details).Error; err != nil {
			return err
		}

		for _, detail := range details {
			// collect book data
			var book models.Book
			if err := tx.First(&book, detail.BookID).Error; err != nil {
				return err
			}
			detail.Book = &book
			detailRes = append(detailRes, detail)
		}
		borrow.BorrowDetail = detailRes

		borrowRes = borrow

		return nil
	})

	// set redis data
	dataJSON, _ := json.Marshal(borrowRes)
	err = helpers.SetRedisData(ctx, m.Redis, key, string(dataJSON))
	if err != nil {
		return borrowRes, err
	}

	return borrowRes, err
}

// CreateBorrow to store borrow data
func (m *conBorrowRepository) CreateBorrow(ctx context.Context, userUUID string, bookUUIDs []string) (models.Borrow, error) {
	// delete redis data
	helpers.DelRedisData(ctx, m.Redis, redisKey)

	// transaction
	borrowUUID := helpers.GenerateUUID("Borrow")
	var borrowData models.Borrow

	err := m.Conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// collect user data
		var user models.User
		if err := tx.Where("uuid = ?", userUUID).First(&user).Error; err != nil {
			return err
		}

		// collect book datas
		var books []models.Book
		if err := tx.Where("uuid IN ? AND borrowed = ?", bookUUIDs, false).Find(&books).Error; err != nil {
			return err
		}
		if len(books) == 0 {
			return helpers.ErrNotFound
		}

		// store borrow
		borrow := models.Borrow{
			UUID:    borrowUUID,
			StartAt: time.Now(),
			User:    &user,
		}
		if err := tx.Create(&borrow).Error; err != nil {
			return err
		}

		// store borrow detail
		for _, book := range books {
			detail := models.BorrowDetail{
				BorrowID: borrow.ID,
				BookID:   &book.ID,
			}

			if err := tx.Create(&detail).Error; err != nil {
				return err
			}

			// update book's borrowed status
			if err := tx.Model(&book).Update("borrowed", true).Error; err != nil {
				return err
			}
		}

		// get borrow detail
		var detailData []models.BorrowDetail
		var detailRes []models.BorrowDetail
		if err := tx.Where("borrow_id = ?", borrow.ID).Find(&detailData).Error; err != nil {
			return err
		}

		for _, detail := range detailData {
			// collect book data
			var book models.Book
			if err := tx.First(&book, detail.BookID).Error; err != nil {
				return err
			}
			detail.Book = &book
			detailRes = append(detailRes, detail)
		}

		borrowData = borrow
		borrowData.User = &user
		borrowData.BorrowDetail = detailRes

		return nil
	})

	return borrowData, err
}

// ReturnBorrow to finish borrow data season
func (m *conBorrowRepository) ReturnBorrow(ctx context.Context, UUID string) (models.Borrow, error) {
	// delete redis data
	key := redisKey + "_" + UUID
	helpers.DelRedisData(ctx, m.Redis, redisKey)
	helpers.DelRedisData(ctx, m.Redis, key)

	// transaction
	var borrowData models.Borrow
	err := m.Conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// collect borrow data
		var borrow models.Borrow
		if err := tx.Where("uuid = ?", UUID).First(&borrow).Error; err != nil {
			return err
		}

		// update borrow data
		currentTime := time.Now()
		updated := models.Borrow{
			EndAt: &currentTime,
		}
		if err := tx.Model(&borrow).Updates(updated).Error; err != nil {
			return err
		}

		// get borrow detail
		var detailData []models.BorrowDetail
		if err := tx.Where("borrow_id = ?", borrow.ID).Find(&detailData).Error; err != nil {
			return err
		}

		// collect book ids
		var bookIDs []uint
		for _, detail := range detailData {
			if detail.BookID != nil {
				bookIDs = append(bookIDs, *detail.BookID)
			}
		}

		// reset book datas
		var book models.Book
		if err := tx.Model(book).Where("id IN ?", bookIDs).Update("borrowed", false).Error; err != nil {
			return err
		}

		// collect user data
		var user models.User
		if err := tx.First(&user, borrow.UserID).Error; err != nil {
			return err
		}

		// collect borrow detail datas
		var detailRes []models.BorrowDetail

		for _, detail := range detailData {
			// collect book data
			if err := tx.First(&book, detail.BookID).Error; err != nil {
				return err
			}
			detail.Book = &book
			detailRes = append(detailRes, detail)
		}

		borrowData = borrow
		borrowData.User = &user
		borrowData.BorrowDetail = detailRes

		return nil
	})

	return borrowData, err
}
