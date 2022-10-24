package factories

import (
	topic "go-devops-admin/app/topic/models"

	"github.com/bxcodec/faker/v3"
)

func MakeTopics(count int) []topic.Topic {
	var obj []topic.Topic

	for i := 0; i < count; i++ {
		topicsModel := topic.Topic{
			Title:      faker.Sentence(),
			Body:       faker.Paragraph(),
			CategoryID: "23",
			UserID:     "1",
		}

		obj = append(obj, topicsModel)
	}
	return obj
}
