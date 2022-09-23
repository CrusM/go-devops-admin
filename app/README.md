# api接口开发步骤(概述)

## 1 添加数据模型 model

> app/models/user/user.go

```go
// 用户模型
// json:"-",指定在JSON解析器忽略字段
type User struct {
	models.BaseModel

	Name string `json:"name,omitempty"`

	City         string `json:"city,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Avatar       string `json:"avatar,omitempty"`

	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampField
}
```

## 2 添加请求验证器 Request

> app/requests/user/user_request.go

```go
// http 请求传过来的参数
type UserUpdateProfileRequest struct {
	// request 字段
	Name         string `json:"name" valid:"name"`
	City         string `json:"city" valid:"city"`
	Introduction string `json:"introduction " valid:"introduction"`
}

// 参数验证规则
func UserUpdateProfile(data interface{}, c *gin.Context) map[string][]string {
	uid := auth.CurrentUID(c)
	rules := govalidator.MapData{
		"name":         []string{"required", "alpha_num", "between:3,20", "not_exists:user,name," + uid},
		"city":         []string{"min_cn:2", "max_cn:20"},
		"introduction": []string{"min_cn:4", "max_cn:240"},
	}

	messages := govalidator.MapData{
		"name": []string{
			"required:用户名必须填写",
			"alpha_num:用户名格式错误,只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
			"not_exists:用户名已经被占用",
		},
		"city": []string{
			"min_cn:城市名称需大于2个字",
			"max_cn:城市名称需大于20个字",
		},
		"introduction": []string{
			"min_cn:描述长度需大于4个字",
			"max_cn:描述长度需小于240个字",
		},
	}

	return requests.ValidateData(data, rules, messages)
}
```

## 3 添加权限验证器

> app/policies/user/user_policies.go

```go
func CanModifyUser(c *gin.Context, _user user.User) bool {
	return auth.CurrentUID(c) == cast.ToString(_user.ID)
}
```

## 4 添加控制器 controller

> (http 请求入口) app/http/controllers/api/v1/user/user_controller.go

```go
// 修改个人资料
func (ctrl *UsersController) UpdateProfile(c *gin.Context) {
	// 判断用户权限 policies
	userModel := auth.CurrentUser(c)
	if ok := userPolicies.CanModifyUser(c, userModel); !ok {
		response.Abort403(c)
		return
	}

	request := userRequest.UserUpdateProfileRequest{}
	if bindOk := requests.Validate(c, &request, userRequest.UserUpdateProfile); !bindOk {
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
```

## 5 添加路由

> routes/api.go

```go
v1.Use(middleware.LimitIP("200-H"))
{
   .
   .
   .
   // user 接口
   uc := new(user.UsersController)
	// 获取当前用户
	v1.GET("/user", uc.CurrentUser)
	userGroup := v1.Group("/users")
	{
		.
		.
		.
		userGroup.PUT("/update-profile", middleware.AuthJWT(), uc.UpdateProfile)

	}
}
```
