package redis

import (
	"context"
	"go-devops-admin/pkg/logger"
	"sync"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

// once 确保全局redis对象只创建一次
var once sync.Once

// 全局Redis, 使用db 1
var Redis *RedisClient

// 连接 redis 数据库，设置全局的Redis对象
func ConnectRedis(address string, username string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, username, password, db)
	})
}

// 创建一个新的 redis 连接
func NewClient(address string, username string, password string, db int) *RedisClient {
	// 初始化自定义的 RedisClient 实例
	rds := &RedisClient{}

	// 使用默认 Context
	rds.Context = context.Background()

	// 使用 redis 中的 NewClient 初始化链接
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	// 测试连接
	err := rds.Ping()
	logger.LogIf(err)

	return rds
}

// Ping 测试 redis 连接是否正常
func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

// Set 存储 Key 对应的 value, 切设置 Expiration 过期时间
func (rds RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

// Get 获取 Key 对应的 value
func (rds RedisClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return ""
	}
	return result
}

// Has 判断一个 Key 是否存在
func (rds RedisClient) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Has", err.Error())
		}
		return false
	}
	return true
}

// Del 删除 key, 支持多个 key 批量删除
func (rds RedisClient) Del(key ...string) bool {
	if err := rds.Client.Del(rds.Context, key...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}
	return true
}

// FlushDB 清空 redis db 里所有数据, 慎用
func (rds RedisClient) FlushDB() bool {
	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
		logger.ErrorString("Redis", "FlushDB", err.Error())
		return false
	}
	return true
}

// Increment 计数器操作 参数只有 1 个时, 为 key 的值增加 1
// 当参数有 2 个时, 第一个参数为 key, 第二个参数为要增加的值 int64 类型
func (rds RedisClient) Increment(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
		return true
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
		return true
	default:
		logger.ErrorString("Redis", "Increment", "参数过多, 不能超过2个")
		return false
	}
}

// Decrement , 和 Increment 相反, 计数器减去响应的值
func (rds RedisClient) Decrement(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
		return true
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.DecrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
		return true
	default:
		logger.ErrorString("Redis", "Decrement", "参数过多, 不能超过2个")
		return false
	}
}
