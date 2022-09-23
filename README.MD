# 从零开始一步一步搭建go api server

## 第三方依

> 使用到的开源库:

* [gin](https://github.com/gin-gonic/gin) —— 路由、路由组、中间件
* [zap](https://github.com/gin-contrib/zap) —— 高性能日志方案
* [gorm](https://github.com/go-gorm/gorm) —— ORM 数据操作
* [cobra](https://github.com/spf13/cobra) —— 命令行结构
* [viper](https://github.com/spf13/viper) —— 配置信息
* [cast](https://github.com/spf13/cast) —— 类型转换
* [redis](https://github.com/go-redis/redis/v8) —— Redis 操作
* [jwt](https://github.com/golang-jwt/jwt) —— JWT 操作
* [base64Captcha](https://github.com/mojocn/base64Captcha) —— 图片验证码
* [govalidator](https://github.com/thedevsaddam/govalidator) —— 请求验证器
* [limiter](https://github.com/ulule/limiter/v3) —— 限流器
* [email](https://github.com/jordan-wright/email) —— SMTP 邮件发送
* [aliyun-communicate](https://github.com/KenmyZhang/aliyun-communicate) —— 发送阿里云短信
* [ansi](https://github.com/mgutz/ansi) —— 终端高亮输出
* [strcase](https://github.com/iancoleman/strcase) —— 字符串大小写操作
* [pluralize](https://github.com/gertd/go-pluralize) —— 英文字符单数复数处理
* [faker](https://learnku.com/courses/go-api/1.19/finish-up/github.com/bxcodec/faker) —— 假数据填充
* [imaging](https://learnku.com/courses/go-api/1.19/finish-up/github.com/disintegration/imaging) —— 图片裁

## 自定义的包

> 自建的库，在 pkg 目录下



* app —— 应用对象
* auth —— 用户授权
* cache —— 缓存
* captcha —— 图片验证码
* config —— 配置信息
* console —— 终端
* database —— 数据库操作
* file —— 文件处理
* hash —— 哈希
* helpers —— 辅助方法
* jwt —— JWT 认证
* limiter —— API 限流
* logger —— 日志记录
* mail —— 邮件发送
* migrate —— 数据库迁移
* paginator —— 分页器
* redis —— Redis 数据库操作
* response —— 响应处理
* seed —— 数据填充
* sms —— 发送短信
* str —— 字符串处理
* verifycode —— 数字验证码
