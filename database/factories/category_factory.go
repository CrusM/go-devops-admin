package factories

import (
	category "go-devops-admin/app/category/models"

	"github.com/bxcodec/faker/v3"
)

func MakeCategories(count int) []category.Category {
	var obj []category.Category

	for i := 0; i < count; i++ {
		categoriesModel := category.Category{
			Name:        faker.Username(),
			Description: faker.Sentence(),
		}

		obj = append(obj, categoriesModel)
	}
	return obj
}
