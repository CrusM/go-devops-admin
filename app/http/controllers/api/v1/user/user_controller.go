package user

import (
	"go-devops-admin/app/http/controllers/api"
	"go-devops-admin/app/models/user"
	"go-devops-admin/app/requests"

	// "go-devops-admin/app/policies"
	userRequest "go-devops-admin/app/requests/user"
	"go-devops-admin/pkg/auth"
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	api.BaseAPIController
}

// 获取当前用户
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

// 列表查询
func (ctrl *UsersController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, paper := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"paper": paper,
	})
}

// 单挑查询
func (ctrl *UsersController) Show(c *gin.Context) {
	usersModel := user.Get(c.Param("id"))
	if usersModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, usersModel)
}

// 新增接口
func (ctrl *UsersController) Create(c *gin.Context) {
	request := userRequest.UserRequest{}
	if ok := requests.Validate(c, &request, userRequest.UserSave); !ok {
		return
	}

	usersModel := user.User{
		// 填充各个字段的内容
		// FieldName: request.FieldName,
	}
	usersModel.Create()
	if usersModel.ID > 0 {
		response.Data(c, usersModel)
	} else {
		response.Abort500(c, "创建失败, 稍后再试")
	}
}

// 修改接口
func (ctrl *UsersController) Update(c *gin.Context) {
	usersModel := user.Get(c.Param("id"))
	if usersModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// if ok := policies.CanModifyUser(c, usersModel); !ok {
	// 	response.Abort403(c)
	// 	return
	// }

	request := userRequest.UserRequest{}
	if bindOk := requests.Validate(c, &request, userRequest.UserSave); !bindOk {
		return
	}

	// 需要求改的字段内容
	// usersModel.FieldName = request.FieldName

	rowsAffected := usersModel.Save()

	if rowsAffected > 0 {
		response.Data(c, usersModel)
	} else {
		response.Abort500(c, "更新失败, 稍后再试")
	}
}

// 删除接口
func (ctrl *UsersController) Delete(c *gin.Context) {
	usersModel := user.Get(c.Param("id"))
	if usersModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// if ok := policies.CanModifyUser(c, usersModel); !ok {
	// 	response.Abort403(c)
	// 	return
	// }

	rowsAffected := usersModel.Save()

	if rowsAffected > 0 {
		response.SUCCESS(c)
		return
	}

	response.Abort500(c, "删除失败, 稍后再试")

}
