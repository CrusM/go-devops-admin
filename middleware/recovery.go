package middleware

import (
	"go-devops-admin/pkg/logger"
	"go-devops-admin/pkg/response"
	"net"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Recovery 使用 zap.Error() 来记录 Panic 和 call stack
func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 获取用户请求信息
				httpRequest, _ := httputil.DumpRequest(ctx.Request, true)

				// 链接中断，客户端中断连接得正常行为，不需要记录堆栈信息
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						errStr := strings.ToLower(se.Error())
						if strings.Contains(errStr, "broken pipe") || strings.Contains(errStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// 中断连接的情况
				if brokenPipe {
					logger.Error("recover from panic",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					ctx.Error(err.(error))
					ctx.Abort()
					// 链接已经断开，无法写状态码
					return
				}

				logger.Error("recover from panic",
					zap.Time("time", time.Now()),
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.Stack("stacktrace"),
				)

				// 返回500状态码
				response.Abort500(ctx)
			}
		}()
		ctx.Next()
	}
}
