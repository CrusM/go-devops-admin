package topic

import (
	"go-devops-admin/pkg/app"
	"go-devops-admin/pkg/database"
	"go-devops-admin/pkg/paginator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func Get(id string) (topics Topic) {
	database.DB.Preload(clause.Associations).Where("id", id).First(&topics)
	return
}

func GetBy(field, value string) (topics Topic) {
	database.DB.Where("? = ?", field, value).First(&topics)
	return
}

func All() (topics []Topic) {
	database.DB.Find(&topics)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Topic{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// 分页查询
func Paginate(c *gin.Context, pageSize int) (Topics []Topic, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Topic{}),
		&Topics,
		app.V1URL(database.TableName(&Topic{})),
		pageSize,
	)
	return
}
