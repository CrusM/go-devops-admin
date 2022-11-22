package backup

import (
	backup "go-devops-admin/app/backup/models"
	backupRequest "go-devops-admin/app/backup/requests"
	"go-devops-admin/app/base"
	"go-devops-admin/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type BackupsController struct {
	base.BaseAPIController
}

// 列表查询
func (ctrl *BackupsController) List(c *gin.Context) {
	request := base.PaginationRequest{}
	if ok := base.Validate(c, &request, base.Pagination); !ok {
		return
	}

	data, paper := backup.Paginate(c, cast.ToInt(request.PageSize))
	response.JSON(c, gin.H{
		"data":  data,
		"paper": paper,
	})

}

// 单挑查询
func (ctrl *BackupsController) Show(c *gin.Context) {
	backupsModel := backup.Get(c.Param("id"))
	if backupsModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, backupsModel)
}

// 新增接口
func (ctrl *BackupsController) Create(c *gin.Context) {
	request := backupRequest.BackupRequest{}
	if ok := base.Validate(c, &request, backupRequest.BackupSave); !ok {
		return
	}

	backupsModel := backup.BackupHost{
		// 填充各个字段的内容
		// FieldName: request.FieldName,
	}
	backupsModel.Create()
	if backupsModel.ID > 0 {
		response.Data(c, backupsModel)
	} else {
		response.Abort500(c, "创建失败, 稍后再试")
	}
}

// 修改接口
func (ctrl *BackupsController) Update(c *gin.Context) {
	backupsModel := backup.Get(c.Param("id"))
	if backupsModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// 设置权限认证 policy

	request := backupRequest.BackupRequest{}
	if bindOk := base.Validate(c, &request, backupRequest.BackupSave); !bindOk {
		return
	}

	// 需要求改的字段内容
	// backupsModel.FieldName = request.FieldName

	rowsAffected := backupsModel.Save()

	if rowsAffected > 0 {
		response.Data(c, backupsModel)
	} else {
		response.Abort500(c, "更新失败, 稍后再试")
	}
}

// 删除接口
func (ctrl *BackupsController) Delete(c *gin.Context) {
	backupsModel := backup.Get(c.Param("id"))
	if backupsModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// 设置权限认证 policy

	rowsAffected := backupsModel.Delete()

	if rowsAffected > 0 {
		response.SUCCESS(c)
		return
	}

	response.Abort500(c, "删除失败, 稍后再试")

}
