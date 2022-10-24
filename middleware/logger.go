package middleware

import (
	"bytes"
	"go-devops-admin/pkg/helpers"
	"go-devops-admin/pkg/logger"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

// 定义http接口响应 数据结构
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// 记录日志请求
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取response响应内容
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: ctx.Writer}
		ctx.Writer = w

		// 获取请求数据
		var requestBody []byte
		if ctx.Request.Body != nil {
			// ctx.Request.Body 是一个Buffer对象，只能读一次
			requestBody, _ = ioutil.ReadAll(ctx.Request.Body)
			// 读取后，重新赋值 ctx.Request.Body,以供后续操作
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 设置开始时间
		start := time.Now()
		// 记录日志不阻塞正常的业务响应
		ctx.Next()

		// 开始记录日志逻辑
		cost := time.Since(start)
		responseStatus := ctx.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", responseStatus),
			zap.String("request", ctx.Request.Method+" "+ctx.Request.URL.String()),
			zap.String("query", ctx.Request.URL.RawQuery),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}

		if ctx.Request.Method == "POST" || ctx.Request.Method == "PUT" || ctx.Request.Method == "DELETE" {
			// 记录请求内容
			logFields = append(logFields, zap.String("Request Body", string(requestBody)))
			// 记录响应内容
			logFields = append(logFields, zap.String("Response Body", w.body.String()))
		}

		if responseStatus > 400 && responseStatus <= 499 {
			logger.Warn("HTTP WARNING "+cast.ToString(responseStatus), logFields...)
		} else if responseStatus > 500 && responseStatus <= 599 {
			logger.Error("HTTP ERROR "+cast.ToString(responseStatus), logFields...)
		} else {
			logger.Debug("HTTP ACCESS Log", logFields...)
		}

	}
}
