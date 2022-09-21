package link

import (
	"go-devops-admin/app/http/controllers/api"
	"go-devops-admin/app/models/link"
	"go-devops-admin/app/requests"

	// "go-devops-admin/app/policies"
	linkRequest "go-devops-admin/app/requests/link"
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

type LinksController struct {
	api.BaseAPIController
}

// 列表查询
func (ctrl *LinksController) List(c *gin.Context) {
	// request := requests.PaginationRequest{}
	// if ok := requests.Validate(c, &request, requests.Pagination); !ok {
	// 	return
	// }

	// data, paper := link.Paginate(c, cast.ToInt(request.PageSize))
	// response.JSON(c, gin.H{
	// 	"data":  data,
	// 	"paper": paper,
	// })
	response.Data(c, link.AllCache())

}

// 单挑查询
func (ctrl *LinksController) Show(c *gin.Context) {
	linksModel := link.Get(c.Param("id"))
	if linksModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, linksModel)
}

// 新增接口
func (ctrl *LinksController) Create(c *gin.Context) {
	request := linkRequest.LinkRequest{}
	if ok := requests.Validate(c, &request, linkRequest.LinkSave); !ok {
		return
	}

	linksModel := link.Link{
		// 填充各个字段的内容
		Name: request.Name,
		URL:  request.URL,
	}
	linksModel.Create()
	if linksModel.ID > 0 {
		response.Data(c, linksModel)
	} else {
		response.Abort500(c, "创建失败, 稍后再试")
	}
}

// 修改接口
func (ctrl *LinksController) Update(c *gin.Context) {
	linksModel := link.Get(c.Param("id"))
	if linksModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// if ok := policies.CanModifyLink(c, linksModel); !ok {
	//     response.Abort403(c)
	//     return
	// }

	request := linkRequest.LinkRequest{}
	if bindOk := requests.Validate(c, &request, linkRequest.LinkSave); !bindOk {
		return
	}

	// 需要求改的字段内容
	linksModel.Name = request.Name
	linksModel.URL = request.URL

	rowsAffected := linksModel.Save()

	if rowsAffected > 0 {
		response.Data(c, linksModel)
	} else {
		response.Abort500(c, "更新失败, 稍后再试")
	}
}

// 删除接口
func (ctrl *LinksController) Delete(c *gin.Context) {
	linksModel := link.Get(c.Param("id"))
	if linksModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// if ok := policies.CanModifyLink(c, linksModel); !ok {
	//     response.Abort403(c)
	//     return
	// }

	rowsAffected := linksModel.Delete()

	if rowsAffected > 0 {
		response.SUCCESS(c)
		return
	}

	response.Abort500(c, "删除失败, 稍后再试")

}
