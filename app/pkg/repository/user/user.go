package user

import (
	"context"
	"encoding/json"

	"github.com/bagus-aulia/dot-test/app/helpers"
	"github.com/bagus-aulia/dot-test/app/pkg/models"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// User interface for User Repository
type User interface {
	GetUserList(ctx context.Context) ([]models.User, error)
	GetUserByUUID(ctx context.Context, UUID string) (models.User, error)
	CreateUser(ctx context.Context, name string, address string) (models.User, error)
	UpdateUser(ctx context.Context, UUID string, name string, address string) (models.User, error)
}

type conUserRepository struct {
	Conn  *gorm.DB
	Redis *redis.Client
}

var (
	redisKey = "dot-user"
)

// NewConUserRepository is a publisher of User repository
func NewConUserRepository(Conn *gorm.DB, Redis *redis.Client) User {
	return &conUserRepository{Conn, Redis}
}

// GetUserList to show all User
func (m *conUserRepository) GetUserList(ctx context.Context) ([]models.User, error) {
	var users []models.User

	// check redis data
	redisData, _ := helpers.GetRedisData(ctx, m.Redis, redisKey)
	if redisData != "" {
		if err := json.Unmarshal([]byte(redisData), &users); err != nil {
			return users, err
		}

		return users, nil
	}

	// get data from db
	result := m.Conn.WithContext(ctx).Find(&users)

	// set redis data
	dataJSON, _ := json.Marshal(users)
	err := helpers.SetRedisData(ctx, m.Redis, redisKey, string(dataJSON))
	if err != nil {
		return users, err
	}

	return users, result.Error
}

// GetUserByUUID to show single User data by uuid
func (m *conUserRepository) GetUserByUUID(ctx context.Context, UUID string) (models.User, error) {
	var user models.User

	// check redis data
	key := redisKey + "_" + UUID
	redisData, _ := helpers.GetRedisData(ctx, m.Redis, key)
	if redisData != "" {
		if err := json.Unmarshal([]byte(redisData), &user); err != nil {
			return user, err
		}

		return user, nil
	}

	// get data from db
	result := m.Conn.WithContext(ctx).Where("uuid = ?", UUID).First(&user)

	// set redis data
	dataJSON, _ := json.Marshal(user)
	err := helpers.SetRedisData(ctx, m.Redis, key, string(dataJSON))
	if err != nil {
		return user, err
	}

	return user, result.Error
}

// CreateUser to insert User data
func (m *conUserRepository) CreateUser(ctx context.Context, name string, address string) (models.User, error) {
	// delete redis data
	helpers.DelRedisData(ctx, m.Redis, redisKey)

	user := models.User{
		Name:    name,
		Address: &address,
		UUID:    helpers.GenerateUUID("User"),
	}

	result := m.Conn.WithContext(ctx).Create(&user)

	return user, result.Error
}

// UpdateUser to update User data
func (m *conUserRepository) UpdateUser(ctx context.Context, UUID string, name string, address string) (models.User, error) {
	// delete redis data
	key := redisKey + "_" + UUID
	helpers.DelRedisData(ctx, m.Redis, redisKey)
	helpers.DelRedisData(ctx, m.Redis, key)

	user := models.User{}
	result := m.Conn.WithContext(ctx).Where("uuid = ?", UUID).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	updated := models.User{
		Name:    name,
		Address: &address,
	}
	result = m.Conn.WithContext(ctx).Model(&user).Updates(updated)

	return user, result.Error
}
