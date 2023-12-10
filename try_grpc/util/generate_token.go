package util

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateToken(userID string) (string, error) {
	// 你的密钥，用于签名和验证 JWT
	secretKey := []byte("ledger")

	// 设置过期时间
	expirationTime := time.Now().Add(24 * time.Hour) // 令牌有效期为 24 小时

	// 创建 JWT
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   userID,
	}
	// 基于HS256对称加密算法签名JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
