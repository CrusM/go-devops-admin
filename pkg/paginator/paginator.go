package paginator

import (
	"fmt"
	"go-devops-admin/pkg/config"
	"go-devops-admin/pkg/logger"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 分页数据结构体
type Paging struct {
	CurrentPage int    // 当前页
	PageSize    int    // 每页条数
	TotalPage   int    // 总页数
	TotalCount  int64  //  总条数
	NextPageURL string // 下一页链接
	PrevPageURL string // 上一页链接
}

// 分页操作类
type Paginator struct {
	BaseURL    string // 基础RUL
	PageSize   int    // 每页条数
	Page       int    // 当前页
	Offset     int    // 数据库读取数据时 Offset 值
	TotalCount int64  // 总条数
	TotalPage  int    // 总页数, TotalCount / PageSize
	Sort       string // 排序字段
	Order      string // 排序规则

	query *gorm.DB     // db query 句柄, 查询数据集和总记录数
	ctx   *gin.Context // gin context 句柄, 从 url 中获取分页参数
}

// 分页
// c -- gin.Context 用来获取分页的 URL 参数
// db -- GORM 查询句柄, 查询数据集和总记录数
// data -- 模型数组
// baseURL -- 用以分页链接
// pageSize -- 每页条数, 优先从 url 中取, 否则使用 config 中的默认值
func Paginate(c *gin.Context, db *gorm.DB, data interface{}, baseURL string, pageSize int) Paging {
	// 初始化 Paginator 实例
	p := Paginator{
		query: db,
		ctx:   c,
	}
	// 初始化对象数据
	p.initProperties(pageSize, baseURL)

	// 查询数据库
	err := p.query.Preload(clause.Associations). // 读取关联
							Order(p.Sort + " " + p.Order). //排序查询
							Limit(p.PageSize).
							Offset(p.Offset).
							Find(data).Error
	if err != nil {
		logger.LogIf(err)
		return Paging{}
	}

	return Paging{
		CurrentPage: p.Page,
		PageSize:    p.PageSize,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
		NextPageURL: p.getNextPageURL(),
		PrevPageURL: p.getPrevPageURL(),
	}
}

func (p *Paginator) initProperties(pageSize int, baseURL string) {
	p.BaseURL = p.formatBaseUrl(baseURL)
	p.PageSize = p.getPageSize(pageSize)

	// 排序参数, 在 controller 中验证过来的参数, 无需再判断
	p.Order = p.ctx.DefaultQuery(config.Get("paging.url_query_order"), "asc")
	p.Sort = p.ctx.DefaultQuery(config.Get("paging.url_query_sort"), "id")

	p.TotalCount = p.getTotalCount()
	p.TotalPage = p.getTotalPage()
	p.Page = p.getCurrentPage()
	p.Offset = (p.Page - 1) * p.PageSize

}

func (p *Paginator) getTotalCount() int64 {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}
	return count
}

func (p *Paginator) getTotalPage() int {
	if p.TotalCount == 0 {
		return 0
	}
	// 计算分页页数
	n := int64(math.Ceil(float64(p.TotalCount) / float64(p.PageSize)))
	// 当总条数小于分页数, 则返回1
	if n == 0 {
		return 1
	}
	return int(n)
}

func (p *Paginator) getCurrentPage() int {
	// 优先取 url 中的 page
	page := cast.ToInt(p.ctx.Query(config.Get("paging.url_query_page")))
	if page <= 0 {
		// 默认为1
		page = 1
	}
	// TotalPage 为 0, 代表不够分页
	if p.TotalPage == 0 {
		return 0
	}
	// 请求也数大于总页数，返回总页数
	if page > p.TotalPage {
		return p.TotalPage
	}
	return page

}

func (p *Paginator) getPageLink(page int) string {
	return fmt.Sprintf("%v%v%s=%s&%s=%s&%s=%v",
		p.BaseURL,
		page,
		config.Get("paging.url_query_sort"),
		p.Sort,
		config.Get("paging.url_query_order"),
		p.Order,
		config.Get("paging.url_query_page_size"),
		p.PageSize,
	)
}

func (p *Paginator) getNextPageURL() string {
	if p.TotalPage > p.Page {
		return p.getPageLink(p.Page + 1)
	}
	return ""
}
func (p *Paginator) getPrevPageURL() string {
	if p.Page <= 1 || p.Page > p.TotalPage {
		return ""
	}
	return p.getPageLink(p.Page - 1)
}

func (p *Paginator) formatBaseUrl(baseURL string) string {
	// 格数化 baseURL
	// URL 中有问号 ?, 表示已经带有其他参数, 直接拼接在参数后面
	if strings.Contains(baseURL, "?") {
		baseURL = baseURL + "&" + config.Get("paging.url_query_page") + "="
	} else {
		baseURL = baseURL + "?" + config.Get("paging.url_query_page") + "="
	}
	return baseURL
}

func (p *Paginator) getPageSize(pageSize int) int {
	// 优先使用请求中的 page_size 参数
	queryPageSize := p.ctx.Query(config.Get("paging.url_query_page_size"))
	if len(queryPageSize) > 0 {
		pageSize = cast.ToInt(queryPageSize)
	}

	// 没有传参, 使用默认
	if pageSize <= 0 {
		pageSize = config.GetInt("paging.url_query_page_size")
	}

	return pageSize
}
