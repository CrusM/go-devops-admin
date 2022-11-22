package v1

import (
	"go-devops-admin/app/base"
	category "go-devops-admin/app/category/models"

	// "go-devops-admin/app/policies"
	categoryRequest "go-devops-admin/app/category/requests"
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type CategoriesController struct {
	base.BaseAPIController
}

// 列表查询
func (ctrl *CategoriesController) List(c *gin.Context) {
	request := base.PaginationRequest{}
	if ok := base.Validate(c, &request, base.Pagination); !ok {
		return
	}

	data, paper := category.Paginate(c, cast.ToInt(request.PageSize))
	response.JSON(c, gin.H{
		"data":  data,
		"paper": paper,
	})

}

// 单挑查询
func (ctrl *CategoriesController) Show(c *gin.Context) {
	categoriesModel := category.Get(c.Param("id"))
	if categoriesModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, categoriesModel)
}

// 新增接口
func (ctrl *CategoriesController) Create(c *gin.Context) {
	request := categoryRequest.CategoryRequest{}
	if ok := base.Validate(c, &request, categoryRequest.CategorySave); !ok {
		return
	}

	categoriesModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
		// 填充各个字段的内容
		// FieldName: request.FieldName,
	}
	categoriesModel.Create()
	if categoriesModel.ID > 0 {
		response.Data(c, categoriesModel)
	} else {
		response.Abort500(c, "创建失败, 稍后再试")
	}
}

// 修改接口
func (ctrl *CategoriesController) Update(c *gin.Context) {
	categoriesModel := category.Get(c.Param("id"))
	if categoriesModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// if ok := policies.CanModifyCategory(c, categoriesModel); !ok {
	//     response.Abort403(c)
	//     return
	// }

	request := categoryRequest.CategoryRequest{}
	if bindOk := base.Validate(c, &request, categoryRequest.CategorySave); !bindOk {
		return
	}

	// 需要求改的字段内容
	categoriesModel.Name = request.Name
	categoriesModel.Description = request.Description

	rowsAffected := categoriesModel.Save()

	if rowsAffected > 0 {
		response.Data(c, categoriesModel)
	} else {
		response.Abort500(c, "更新失败, 稍后再试")
	}
}

// 删除接口
func (ctrl *CategoriesController) Delete(c *gin.Context) {
	categoriesModel := category.Get(c.Param("id"))
	if categoriesModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// if ok := policies.CanModifyCategory(c, categoriesModel); !ok {
	// 	response.Abort403(c)
	// 	return
	// }

	rowsAffected := categoriesModel.Delete()

	if rowsAffected > 0 {
		response.SUCCESS(c)
		return
	}

	response.Abort500(c, "删除失败, 稍后再试")

}
