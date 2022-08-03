package captcha

import (
	"errors"
	"go-devops-admin/pkg/app"
	"go-devops-admin/pkg/config"
	"go-devops-admin/pkg/redis"
	"time"
)

// 自定存储驱动，使用 redis 进行作为主要存储器, 替换默认的内存存储, 避免多机部署状态下失效问题
// 实现 base64Captcha.Store interface 即可

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// 实现 base64Captcha.Store interface 的 Set 方法
func (s *RedisStore) Set(key string, value string) error {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))
	// 本地调试添加额外的过期时间
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}

	if ok := s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime); !ok {
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}

// 实现 base64Captcha.Store interface 的 Get 方法
func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

// 实现 base64Captcha.Store interface 的 Verify 方法
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
