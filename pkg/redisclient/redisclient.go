package redisclient

import (
	"encoding/json"
	"fmt"
	"time"

	"context"

	"github.com/go-redis/redis/v8"
)

// RedisClient ...
type RedisClient struct {
	Client *redis.Client
}

// Set ...
func (r *RedisClient) Set(key string, value interface{}) error {
	binary, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value for Redis: %v ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	cmd := r.Client.Set(ctx, key, string(binary), 3*time.Hour)
	return cmd.Err()
}
