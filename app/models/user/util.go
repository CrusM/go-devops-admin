package user

import "go-devops-admin/pkg/database"

// 判断邮箱是否被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// 判断电话号码是否被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// 通过手机号获取用户
func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

// 通过邮箱获取用户
func GetByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)
	return
}

// 通过 Email/Phone/Username 获取用户
func GetByMulti(loginId string) (userModel User) {
	database.DB.
		Where("phone = ?", loginId).
		Or("email = ?", loginId).
		Or("name = ? ", loginId).First(&userModel)
	return
}

// 通过 ID 获取用户
func Get(id string) (userModel User) {
	database.DB.Where("id = ?", id).First(&userModel)
	return
}

func All() (userModels []User) {
	database.DB.Find(&userModels)
	return
}
