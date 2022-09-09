package project

import (
    "go-devops-admin/pkg/database"
)

func Get(id string) (projects Project) {
    database.DB.Where("id", id).First(&projects)
    return 
}

func GetBy(field, value string) (projects Project) {
    database.DB.Where("? = ?", field, value).First(&projects)
    return
}

func All() (projects []Project) {
    database.DB.Find(&projects)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Project{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}