package factories

import (
	link "go-devops-admin/app/link/models"

	"github.com/bxcodec/faker/v3"
)

func MakeLinks(count int) []link.Link {
	var obj []link.Link

	for i := 0; i < count; i++ {
		linksModel := link.Link{
			Name: faker.Username(),
			URL:  faker.URL(),
		}

		obj = append(obj, linksModel)
	}
	return obj
}
