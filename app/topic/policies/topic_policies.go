package policies

import (
	topic "go-devops-admin/app/topic/models"
	"go-devops-admin/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}
