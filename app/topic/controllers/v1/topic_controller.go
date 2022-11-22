package v1

import (
	"go-devops-admin/app/base"
	topic "go-devops-admin/app/topic/models"
	topicPolicies "go-devops-admin/app/topic/policies"

	// "go-devops-admin/app/policies"
	topicRequest "go-devops-admin/app/topic/requests"
	"go-devops-admin/pkg/auth"
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type TopicsController struct {
	base.BaseAPIController
}

// 列表查询
func (ctrl *TopicsController) List(c *gin.Context) {
	request := base.PaginationRequest{}
	if ok := base.Validate(c, &request, base.Pagination); !ok {
		return
	}

	data, paper := topic.Paginate(c, cast.ToInt(request.PageSize))
	response.JSON(c, gin.H{
		"data":  data,
		"paper": paper,
	})

}

// 单挑查询
func (ctrl *TopicsController) Show(c *gin.Context) {
	topicsModel := topic.Get(c.Param("id"))
	if topicsModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, topicsModel)
}

// 新增接口
func (ctrl *TopicsController) Create(c *gin.Context) {
	request := topicRequest.TopicRequest{}
	if ok := base.Validate(c, &request, topicRequest.TopicSave); !ok {
		return
	}

	topicsModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserID:     auth.CurrentUID(c),
	}
	topicsModel.Create()
	if topicsModel.ID > 0 {
		response.Data(c, topicsModel)
	} else {
		response.Abort500(c, "创建失败, 稍后再试")
	}
}

// 修改接口
func (ctrl *TopicsController) Update(c *gin.Context) {
	topicsModel := topic.Get(c.Param("id"))
	if topicsModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := topicPolicies.CanModifyTopic(c, topicsModel); !ok {
		response.Abort403(c)
		return
	}

	request := topicRequest.TopicRequest{}
	if bindOk := base.Validate(c, &request, topicRequest.TopicSave); !bindOk {
		return
	}

	// 需要求改的字段内容
	// topicsModel.FieldName = request.FieldName
	topicsModel.Title = request.Title
	topicsModel.Body = request.Body
	topicsModel.CategoryID = request.CategoryID

	rowsAffected := topicsModel.Save()

	if rowsAffected > 0 {
		response.Data(c, topicsModel)
	} else {
		response.Abort500(c, "更新失败, 稍后再试")
	}
}

// 删除接口
func (ctrl *TopicsController) Delete(c *gin.Context) {
	topicsModel := topic.Get(c.Param("id"))
	if topicsModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := topicPolicies.CanModifyTopic(c, topicsModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := topicsModel.Delete()

	if rowsAffected > 0 {
		response.SUCCESS(c)
		return
	}

	response.Abort500(c, "删除失败, 稍后再试")

}
