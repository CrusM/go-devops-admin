package category

import (
    "go-devops-admin/pkg/app"
	"go-devops-admin/pkg/database"
	"go-devops-admin/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(id string) (categories Category) {
    database.DB.Where("id", id).First(&categories)
    return 
}

func GetBy(field, value string) (categories Category) {
    database.DB.Where("? = ?", field, value).First(&categories)
    return
}

func All() (categories []Category) {
    database.DB.Find(&categories)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Category{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

// 分页查询
func Paginate(c *gin.Context, pageSize int) (Categories []Category, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Category{}),
		&Categories,
		app.V1URL(database.TableName(&Category{})),
		pageSize,
	)
	return
}