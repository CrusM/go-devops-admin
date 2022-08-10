package user

import (
	"go-devops-admin/pkg/hash"

	"gorm.io/gorm"
)

// 实现 GORM 提供的模型钩子

func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}
	return
}
