package v1

import (
	"go-devops-admin/app/base"
	user "go-devops-admin/app/user/models"

	userPolicies "go-devops-admin/app/user/policies"
	userRequest "go-devops-admin/app/user/requests"
	"go-devops-admin/pkg/auth"
	"go-devops-admin/pkg/config"
	"go-devops-admin/pkg/file"
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	base.BaseAPIController
}

// 获取当前用户
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

// 列表查询
func (ctrl *UsersController) List(c *gin.Context) {
	request := base.PaginationRequest{}
	if ok := base.Validate(c, &request, base.Pagination); !ok {
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
	if ok := base.Validate(c, &request, userRequest.UserSave); !ok {
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
	if bindOk := base.Validate(c, &request, userRequest.UserSave); !bindOk {
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

// 修改个人资料
func (ctrl *UsersController) UpdateProfile(c *gin.Context) {
	// 判断用户权限 policies
	userModel := auth.CurrentUser(c)
	if ok := userPolicies.CanModifyUser(c, userModel); !ok {
		response.Abort403(c)
		return
	}

	request := userRequest.UserUpdateProfileRequest{}
	if bindOk := base.Validate(c, &request, userRequest.UserUpdateProfile); !bindOk {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Name = request.Name
	currentUser.City = request.City
	currentUser.Introduction = request.Introduction

	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Data(c, currentUser)
	} else {
		response.Abort500(c, "更新失败, 稍后再试")
	}
}

// 修改邮箱
func (ctrl *UsersController) UpdateUserEmail(c *gin.Context) {
	// if ok := policies.CanModifyUser(c, usersModel); !ok {
	// 	response.Abort403(c)
	// 	return
	// }

	request := userRequest.UserUpdateEmailRequest{}
	if bindOk := base.Validate(c, &request, userRequest.UserUpdateEmail); !bindOk {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Email = request.Email

	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Data(c, currentUser)
	} else {
		response.Abort500(c, "更新失败, 稍后再试")
	}
}

// 修改邮箱
func (ctrl *UsersController) UpdateUserPhone(c *gin.Context) {
	// if ok := policies.CanModifyUser(c, usersModel); !ok {
	// 	response.Abort403(c)
	// 	return
	// }

	request := userRequest.UserUpdatePhoneRequest{}
	if bindOk := base.Validate(c, &request, userRequest.UserUpdatePhone); !bindOk {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Phone = request.Phone

	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Data(c, currentUser)
	} else {
		response.Abort500(c, "更新失败, 稍后再试")
	}
}

// 修改邮箱
func (ctrl *UsersController) UpdateUserPassword(c *gin.Context) {
	// if ok := policies.CanModifyUser(c, usersModel); !ok {
	// 	response.Abort403(c)
	// 	return
	// }

	request := userRequest.UserUpdatePasswordRequest{}
	if bindOk := base.Validate(c, &request, userRequest.UserUpdatePassword); !bindOk {
		return
	}

	currentUser := auth.CurrentUser(c)

	_, err := auth.Attempt(currentUser.Name, request.Password)
	if err != nil {
		// 密码验证失败
		response.Unauthorized(c, "原始密码不正确")
	} else {
		currentUser.Password = request.Password

		rowsAffected := currentUser.Save()

		if rowsAffected > 0 {
			response.Data(c, currentUser)
		} else {
			response.Abort500(c, "更新失败, 稍后再试")
		}
	}
}

// 上传用户头像
func (ctrl *UsersController) UpdateUserAvatar(c *gin.Context) {

	request := userRequest.UserUpdateAvatarRequest{}
	if bindOk := base.Validate(c, &request, userRequest.UserUpdateAvatar); !bindOk {
		return
	}

	avatar, err := file.SaveUploadAvatar(c, request.Avatar)
	if err != nil {
		// 密码验证失败
		response.Abort500(c, "上传头像失败, 请稍后再试")
	} else {
		currentUser := auth.CurrentUser(c)
		currentUser.Avatar = config.GetString("app.url") + avatar

		rowsAffected := currentUser.Save()

		if rowsAffected > 0 {
			response.Data(c, currentUser)
		} else {
			response.Abort500(c, "更新失败, 稍后再试")
		}
	}
}
