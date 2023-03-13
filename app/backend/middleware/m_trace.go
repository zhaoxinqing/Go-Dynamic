package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	XRequestIDKey = "X-Request-ID"
)

// RequestID 是一个 Gin 中间件，用来在每一个 HTTP 请求的 context, response 中注入 `X-Request-ID` 键值对.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查请求头中是否有 `X-Request-ID`
		requestID := c.Request.Header.Get(XRequestIDKey)

		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Set(XRequestIDKey, requestID)

		// 将 RequestID 保存在 HTTP 返回头中，Header 的键为 `X-Request-ID`
		c.Writer.Header().Set(XRequestIDKey, requestID)
		// logx.NewContext(c, zap.Any("x-request-id", requestID))
		c.Next()
	}
}

// ObtainRequestID ...
func ObtainRequestID(c *gin.Context) string {
	requestID, _ := c.Get(XRequestIDKey)
	if requestID == nil {
		return ""
	}
	return requestID.(string)
}
