package cache

import (
	"go-devops-admin/pkg/config"
	"go-devops-admin/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func NewRedisStore(address string, username string, password string, db int) *RedisStore {
	rs := &RedisStore{
		RedisClient: redis.NewClient(address, username, password, db),
		KeyPrefix:   config.GetString("app.name") + ":cache:",
	}
	return rs
}

func (s *RedisStore) Set(key string, value string, expireTime time.Duration) {
	s.RedisClient.Set(s.KeyPrefix+key, value, expireTime)
}
func (s *RedisStore) Get(key string) string {
	return s.RedisClient.Get(s.KeyPrefix + key)
}
func (s *RedisStore) Has(key string) bool {
	return s.RedisClient.Has(s.KeyPrefix + key)
}
func (s *RedisStore) Forget(key string) {
	s.RedisClient.Del(s.KeyPrefix + key)
}
func (s *RedisStore) Forever(key string, value string) {
	s.RedisClient.Set(s.KeyPrefix+key, value, 0)
}
func (s *RedisStore) Flush() {
	s.RedisClient.FlushDB()
}

func (s *RedisStore) IsAlive() error {
	return s.RedisClient.Ping()
}

func (s *RedisStore) Increment(params ...interface{}) {
	s.RedisClient.Increment(params...)
}

func (s *RedisStore) Decrement(params ...interface{}) {
	s.RedisClient.Decrement(params...)
}
