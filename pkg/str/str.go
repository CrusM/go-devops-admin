package str

// go-pluralize 用来处理英文单复数
// strcase 是用来处理大小写
import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

// 转为复数  user -> users
func Plural(word string) string {
	return pluralize.NewClient().Plural(word)
}

// 转为单数 users -> user
func Singular(word string) string {
	return pluralize.NewClient().Singular(word)
}

// 转为 snake_case 蛇形命名法, TopicComment -> topic_comment
func Snake(word string) string {
	return strcase.ToSnake(word)
}

// 转为 CamelCase 驼峰形式, topic_comment -> TopicComment
func Camel(word string) string {
	return strcase.ToCamel(word)
}

// 转为全部小写
func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
