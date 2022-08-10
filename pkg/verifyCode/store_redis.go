package verifyCode

import "go-devops-admin/pkg/redis"

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}


