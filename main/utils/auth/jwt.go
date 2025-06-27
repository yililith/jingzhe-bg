package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"sync"
	"time"
)

type JwtModel struct {
	AvatarID string `json:"avatar_id"`
	UID      string `json:"uid"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 全局配置
var jwtConfig = struct {
	SecretKey     string
	Expiration    time.Duration
	Issuer        string
	Audience      string
	SigningMethod jwt.SigningMethod
}{
	SecretKey:     "jingzhe-bg",
	Expiration:    24 * time.Hour,
	Issuer:        "jingzhe-bg",
	Audience:      "jingzhe-app",
	SigningMethod: jwt.SigningMethodHS256,
}

// 缓存密钥字节（避免重复转换）
var (
	secretKeyBytes     []byte
	secretKeyBytesOnce sync.Once
	parsingKey         jwt.Keyfunc
	parsingKeyOnce     sync.Once
)

func initSecretKey() {
	secretKeyBytesOnce.Do(func() {
		secretKeyBytes = []byte(jwtConfig.SecretKey)
	})
}

func initParsingKey() {
	parsingKeyOnce.Do(func() {
		initSecretKey()
		parsingKey = func(token *jwt.Token) (interface{}, error) {
			return secretKeyBytes, nil
		}
	})
}

// GenerateToken 生成 JWT Token
func GenerateToken(avatarID, uid, username string) (string, error) {
	initSecretKey()

	claims := JwtModel{
		AvatarID: avatarID,
		UID:      uid,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtConfig.Expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    jwtConfig.Issuer,
			Audience:  []string{jwtConfig.Audience},
		},
	}

	token := jwt.NewWithClaims(jwtConfig.SigningMethod, claims)
	return token.SignedString(secretKeyBytes)
}

// ParseToken 安全解析 JWT Token
func ParseToken(tokenString string) (*JwtModel, error) {
	initParsingKey()

	token, err := jwt.ParseWithClaims(tokenString, &JwtModel{}, parsingKey)
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, fmt.Errorf("token expired: %w", err)
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, fmt.Errorf("token malformed: %w", err)
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return nil, fmt.Errorf("token signature invalid: %w", err)
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, fmt.Errorf("token not yet valid: %w", err)
		default:
			return nil, fmt.Errorf("parse token failed: %w", err)
		}
	}

	claims, ok := token.Claims.(*JwtModel)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return claims, nil
}
