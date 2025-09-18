package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	JwtSecret       = []byte("your-secret-key") // 生产环境应该从配置读取
	ErrInvalidToken = errors.New("invalid token")
)

type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成 Access Token
func GenerateAccessToken(id uint, username string) (string, error) {
	claims := Claims{
		Username: username,
		ID:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // 15分钟过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gim-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

// 生成 Refresh Token
func GenerateRefreshToken(id uint, userName string) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 7天过期
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   "refresh",
		Issuer:    "gim-app",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

// 解析 Token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return JwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

// 验证 Refresh Token
func ValidateRefreshToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 检查是否是 refresh token
		if sub, ok := claims["sub"].(string); ok && sub == "refresh" {
			return true, nil
		}
	}

	return false, ErrInvalidToken
}
