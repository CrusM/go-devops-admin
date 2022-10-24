package factories

import (
	user "go-devops-admin/app/user/models"
	"go-devops-admin/pkg/helpers"

	"github.com/bxcodec/faker/v3"
)

func MakeUsers(items int) []user.User {
	var models []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < items; i++ {
		model := user.User{
			Name:     faker.Username(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "abc123456",
		}
		models = append(models, model)
	}
	return models
}
