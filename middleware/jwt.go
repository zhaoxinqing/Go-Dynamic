package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	TokenKey  = "Authorization" // Token key
	issuer    = "nb-plus"       // 颁发者
	expiresAt = 48
	secret    = "1234567890"
)

// MyClaims  ...
type MyClaims struct {
	UserID uint64 // userid
	jwt.RegisteredClaims
}

// GenerateToken ...
func GenerateToken(userID uint64) (string, error) {
	claims := MyClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour * expiresAt)), // Expiration time 48 hours
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),                            // Issue time
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken ...
func ParseToken(bearerToken string) (*MyClaims, error) {
	token := strings.ReplaceAll(bearerToken, "Bearer ", "")
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// AuthToken token 校验解析
func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := ParseToken(c.GetHeader(TokenKey))
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, "login has expired, please login again",
			)
			return
		}
		c.Next()
	}
}
