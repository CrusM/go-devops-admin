package validators

import (
	"errors"
	"fmt"
	"go-devops-admin/pkg/database"
	"strings"

	"github.com/thedevsaddam/govalidator"
)

// 自定义校验规则

func init() {
	// 自定义校验规则, 验证请求数据必须不存在于数据库中
	// 常用于保证数据库某个字段的值唯一, 例如 用户名、手机号、邮箱、或者分类的名称
	// not_exists 参数可以有两种, 一种是 2 个参数, 一种是 3 个参数
	// not_exists:users,email 检查数据库表里是否存在同一信息
	// not_exists:users,email,32 排除掉用户 id 为 32 的用户
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		// 第一个参数, 表名称, 如 users
		tableName := rng[0]
		// 第二个参数字段名称, 如 phone、email
		dbField := rng[1]
		// 第三个参数, 排除 ID
		var ExceptID string
		if len(rng) > 2 {
			ExceptID = rng[2]
		}

		// 用户请求过来的数据
		requestValue := value.(string)

		// 拼接 SQL
		query := database.DB.Table(tableName).Where(dbField+" = ?", requestValue)

		if len(ExceptID) > 0 {
			query.Where("id != ?", ExceptID)
		}

		// 查询结果
		var count int64
		query.Count(&count)

		// 数据库有数据, 验证不通过
		if count != 0 {
			// 如果有自定义错误消息的话
			if message != "" {
				return errors.New(message)
			}
			// 默认错误消息
			return fmt.Errorf("%v 已被占用.", requestValue)
		}
		// 验证通过
		return nil
	})
}