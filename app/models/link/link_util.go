package link

import (
	"go-devops-admin/pkg/app"
	"go-devops-admin/pkg/cache"
	"go-devops-admin/pkg/database"
	"go-devops-admin/pkg/helpers"
	"go-devops-admin/pkg/paginator"
	"time"

	"github.com/gin-gonic/gin"
)

func Get(id string) (links Link) {
	database.DB.Where("id", id).First(&links)
	return
}

func GetBy(field, value string) (links Link) {
	database.DB.Where("? = ?", field, value).First(&links)
	return
}

func All() (links []Link) {
	database.DB.Find(&links)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Link{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// 分页查询
func Paginate(c *gin.Context, pageSize int) (Links []Link, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Link{}),
		&Links,
		app.V1URL(database.TableName(&Link{})),
		pageSize,
	)
	return
}

func AllCache() (links []Link) {
	cacheKey := "links:all"

	expireTime := 120 * time.Minute

	// 取数据
	cache.GetObject(cacheKey, &links)

	// 如果数据为空
	if helpers.Empty(links) {
		links = All()
		if helpers.Empty(links) {
			// fmt.Printf("")
			return links
		}
		cache.Set(cacheKey, links, expireTime)
	}
	return
}
