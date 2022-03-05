package redis

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type RedisStorage interface {
	GetStorage() *Storage
	SetCache(key string, val interface{})
	SetCacheInt(key string, val int)
	GetCacheInt(key string) int
	DeleteCache(key string)
}

type redisStorage struct {
	store *Storage
}

var DefaultCacheExpireTime time.Duration = time.Second * 3

func (r redisStorage) GetStorage() *Storage {
	return r.store
}

func Init() redisStorage {
	port, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))

	store := New(Config{
		Host:      os.Getenv("REDIS_ENDPOINT"),
		Port:      port,
		Username:  "",
		Password:  os.Getenv("REDIS_DB_PASSWORD"),
		URL:       "",
		Database:  0,
		Reset:     false,
		TLSConfig: nil,
	})
	return redisStorage{
		store: store,
	}
}

func (r redisStorage) SetCache(key string, val interface{}) {
	b, marshalErr := json.Marshal(val)
	if marshalErr != nil {
		zap.S().Warn("Cache: Marshal binary failed: " + marshalErr.Error())
	}
	err := r.store.Set(key, b, DefaultCacheExpireTime)
	if err != nil {
		zap.S().Warn("Set Cache error: " + err.Error())
	}
}

// set cache for any int value ex: views
func (r redisStorage) SetCacheInt(key string, val int) {
	err := r.store.Set(key, []byte(fmt.Sprint(val)), 0)
	if err != nil {
		zap.S().Warn("Set View cache error: ", err.Error())
	}
}

// get cache for any int value ex: views
func (r redisStorage) GetCacheInt(key string) int {
	bytes, err := r.store.Get(key)
	if err != nil {
		zap.S().Warn("Get View cache error: ", err.Error())
		return 0
	}

	count, err := strconv.Atoi(string(bytes))
	if err != nil && count != 0 {
		zap.S().Warn("Set View cache string to int error: ", err.Error())
		return 0
	}
	return count
}

func (r redisStorage) DeleteCache(key string) {
	err := r.store.Delete(key)
	if err != nil {
		zap.S().Warn("Delete cache failed, key=", key, ", error: ", err.Error())
	}
}
