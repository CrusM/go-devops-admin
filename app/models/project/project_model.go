package project

import (
	"go-devops-admin/app/models"
	"go-devops-admin/pkg/database"
)

type Project struct {
	models.BaseModel

	//  put fields in here
	// FIXME()

	models.CommonTimestampField
}


func (projects *Project) Create() {
	database.DB.Create(&projects)
}

func (projects *Project) Save() (rowsAffected int64){
	result := database.DB.Save(&projects)
	return result.RowsAffected
}

func (projects *Project) Delete() (rowAffected int64) {
	result := database.DB.Delete(&projects)
	return result.RowsAffected
}
