package middleware

import (
	"backend-go/public"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// MyClaims  ...
type MyClaims struct {
	UserID uint64
	RoleID uint64
	jwt.RegisteredClaims
}

// GenerateToken ...
func GenerateToken(userID uint64, roleID uint64) (string, error) {
	claims := MyClaims{
		UserID: userID,
		RoleID: roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour * public.TOKEN_SURVIVAL)), // Expiration time 12 hours
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),                                        // Issue time
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
			Issuer:    public.TOKEN_ISSUER, // 颁发者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(public.TOKEN_SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken ...
func ParseToken(bearerToken string) (*MyClaims, error) {
	token := strings.ReplaceAll(bearerToken, "Bearer ", "")
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(public.TOKEN_SECRET), nil
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
