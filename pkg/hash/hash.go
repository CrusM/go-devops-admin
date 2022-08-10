package hash

import (
	"go-devops-admin/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

// 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	// GenerateFromPassword 的第二个参数是 cost 值. 建议大于12, 数值越大爆破消耗时间越长
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogIf(err)

	return string(bytes)
}

// 对比明文密码和 hash 值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// 判断字符串是否是哈希过的数据
func BcryptIsHashed(str string) bool {
	// 加密后的长度等于 60, 用于简单判断
	return len(str) == 60
}
