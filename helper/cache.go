package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/storage/redis/v3"
)

func CachePrefix(key string) string {
	return fmt.Sprintf("%s:%s", os.Getenv("CACHE_PREFIX"), key)
}

func CacheRemember[T any](cache *redis.Storage, key string, fn func() (T, error), ttls ...time.Duration) (T, error) {
	var zero T
	ttl := time.Hour

	cacheKey := CachePrefix(key)
	cacheData, err := cache.Get(cacheKey)
	if err == nil && cacheData != nil {
		var data T
		if err := json.Unmarshal([]byte(cacheData), &data); err == nil {
			return data, nil
		}
	}

	data, err := fn()
	if err != nil {
		return zero, err
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return zero, err
	}

	if len(ttls) > 0 {
		ttl = ttls[0]
	}
	if err := cache.Set(cacheKey, jsonData, ttl); err != nil {
		return zero, err
	}

	return data, nil
}
