package middleware

import (
	"backend-go/config"
	"backend-go/pkg/db"
	"backend-go/public"
	"context"
	"fmt"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v9"
	"github.com/golang-jwt/jwt/v4"
)

// Limiter 全局限流 ...根据 IP
func Limiter(velocity int, timeUnit string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// KEY：
		var limterKey = "limiter:" + strings.Replace(c.Request.URL.Path, "/", "-", -1) + ":"
		userID := ObtainUserIDFromResolvedToken(c)
		if userID == 0 {
			limterKey = limterKey + fmt.Sprint(userID)
		} else {
			limterKey = limterKey + fmt.Sprint(c.ClientIP())
		}

		// 限流：时间单位-流速
		var limit redis_rate.Limit
		switch timeUnit {
		case public.TIME_HOUR:
			limit = redis_rate.PerHour(velocity) // hour
		case public.TIME_MINUTE:
			limit = redis_rate.PerMinute(velocity) // minute
		case public.TIME_SECOND:
			limit = redis_rate.PerSecond(velocity) // second
		}

		// CHECK：
		res, err := redis_rate.NewLimiter(db.Redis).Allow(context.Background(), limterKey, limit)
		if err != nil || res.Allowed == 0 {
			c.AbortWithStatusJSON(http.StatusTooManyRequests,
				gin.H{"code": "0", "message": "请求过快，请稍后再试"})
		}
	}
}

// ObtainUserIDFromResolvedToken 从已解析的的token 中获取 userID
func ObtainUserIDFromResolvedToken(c *gin.Context) uint64 {
	userId, _ := c.Get("UserID")
	if userId == nil {
		return 0
	}
	return userId.(uint64)
}

// GetUserID ExtractClaims help to extract the JWT claims
func GetUserID(c *gin.Context) uint64 {
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println(r)
		}
	}()
	t := c.GetHeader(public.TOKEN_KEY)
	if t == "" {
		return 0
	}
	t = strings.ReplaceAll(t, "Bearer ", "")
	token, _ := jwt.ParseWithClaims(t, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJwtEnv().Secret), nil
	})
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims.UserID
	} else {
		return 0
	}
}
