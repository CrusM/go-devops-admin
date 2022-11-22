package validators

import (
	"errors"
	"fmt"
	"go-devops-admin/pkg/database"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/thedevsaddam/govalidator"
)

// 注册自定义校验规则

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
			return fmt.Errorf("%v 已被占用", requestValue)
		}
		// 验证通过
		return nil
	})

	// 自定义规则 exists，确保数据库存在某条数据,和 not_exists 相反
	govalidator.AddCustomRule("exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "exists:"), ",")

		// 第一个参数, 表名称, 如 users
		tableName := rng[0]
		// 第二个参数字段名称, 如 phone、email
		dbField := rng[1]

		// 用户请求过来的数据
		requestValue := value.(string)

		// 查询结果
		var count int64
		database.DB.Table(tableName).Where(dbField+" = ?", requestValue).Count(&count)

		// 数据库有数据, 验证不通过
		if count == 0 {
			// 如果有自定义错误消息的话
			if message != "" {
				return errors.New(message)
			}
			// 默认错误消息
			return fmt.Errorf("%v 不存在", requestValue)
		}
		// 验证通过
		return nil
	})

	// 中文字符最大长度
	govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
		if valLength > l {
			// 如果有自定义错误，使用自定义错误消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度不能超过 %d 个字", l)
		}
		return nil
	})

	// 中文字符最小长度
	govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))
		if valLength < l {
			// 如果有自定义错误，使用自定义错误消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度不能小于 %d 个字", l)
		}
		return nil
	})

}
