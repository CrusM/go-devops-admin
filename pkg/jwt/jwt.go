package jwt

import (
	"errors"
	"go-devops-admin/pkg/app"
	"go-devops-admin/pkg/config"
	"go-devops-admin/pkg/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwtPkg "github.com/golang-jwt/jwt"
)

var (
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求令牌格式有误")
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
)

type JWT struct {
	// 密钥, 用以加密 jwt, 读取配置信息app.key
	SignKey []byte

	// jwt header 关键字 Authorization

	// 刷新 Token 的最大时间
	MaxRefresh time.Duration
}

// 自定义 JWT 结构体
type JWTCustomClaims struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_at_time"`

	// StandardClaims 结构体实现了 Claims 接口继承了 Valid() 方法
	// JWT 提供了 7 个官方字段, 提供使用:
	// - iss (issuer): 发布者
	// - sub (subject): 主题
	// - iat (Issued At): 生成签名时间
	// - exp (expiration time): 签名过期时间
	// - aud (audience): 观众, 相当于接收者
	// - nbf (Not Before): 生效时间
	// - jti (JWT ID): 编号
	jwtPkg.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

// 解析 Token, 中间件调用
func (jwt *JWT) ParserToken(c *gin.Context) (*JWTCustomClaims, error) {
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	// 调用 jwt 库解析用户传参的 token
	token, err := jwt.parseTokenString(tokenString)
	// 解析出错
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtPkg.ValidationErrorMalformed {
				return nil, ErrHeaderMalformed
			} else if validationErr.Errors == jwtPkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}

	// 将 token 中的 claims 信息解析出来, 和 JWTCustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

// 刷新 jwt token
func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}
	// 调用 jwt 库解析用户传参的 token
	token, err := jwt.parseTokenString(tokenString)
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		if !ok || validationErr.Errors != jwtPkg.ValidationErrorExpired {
			return "", err
		}
	}
	// 解析 JWTCustomClaims 数据
	claims := token.Claims.(*JWTCustomClaims)

	// 检查是否超过 【最大允许刷新时间】
	x := app.TimeNowInTimezone().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt > x {
		// 修改过期时间
		claims.StandardClaims.ExpiresAt = jwt.ExpireAtTime()
		return jwt.createToken(*claims)
	}
	return "", ErrTokenExpiredMaxRefresh
}

// 生成 Token, 在登陆时调用
func (jwt *JWT) IssueToken(userID string, username string) string {
	// 构建用户 claims 信息(载荷)
	expireAtTime := jwt.ExpireAtTime()
	claims := JWTCustomClaims{
		UserID:       userID,
		UserName:     username,
		ExpireAtTime: expireAtTime,
		StandardClaims: jwtPkg.StandardClaims{
			NotBefore: app.TimeNowInTimezone().Unix(), // 签名生效时间
			IssuedAt:  app.TimeNowInTimezone().Unix(), // 首次签名时间 (刷新 Token 不会更新)
			ExpiresAt: expireAtTime,
			Issuer:    config.GetString("app.name"), // 颁发者

		},
	}

	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return token
}

func (jwt *JWT) createToken(claims JWTCustomClaims) (string, error) {
	// 使用 HS256 算法生成 token
	token := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

func (jwt *JWT) ExpireAtTime() int64 {
	timeNow := app.TimeNowInTimezone()

	var expireTime int64
	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}

	expire := time.Duration(expireTime) * time.Minute
	return timeNow.Add(expire).Unix()
}

func (jwt *JWT) parseTokenString(tokenString string) (*jwtPkg.Token, error) {
	return jwtPkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(t *jwtPkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}
	// 按照空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}
