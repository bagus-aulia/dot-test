package helpers

import (
	"context"

	"github.com/go-redis/redis"
)

// GetRedisData to get redis cache
func GetRedisData(ctx context.Context, rds *redis.Client, key string) (string, error) {
	redisResult, err := rds.WithContext(ctx).Get(key).Result()
	if err != nil {
		return "", err
	}

	return redisResult, nil
}

// SetRedisData to store redis cache
func SetRedisData(ctx context.Context, rds *redis.Client, key string, data string) error {
	err := rds.WithContext(ctx).Set(key, data, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// DelRedisData to delete redis cache
func DelRedisData(ctx context.Context, rds *redis.Client, key string) error {
	err := rds.WithContext(ctx).Del(key).Err()
	if err != nil {
		return err
	}

	return nil
}
