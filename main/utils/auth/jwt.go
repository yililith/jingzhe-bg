package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"sync"
	"time"
)

type JwtModel struct {
	AvatarID string `json:"avatar_id"`
	UID      uint64 `json:"uid"`
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
func GenerateToken(avatarID, username string, uid uint64) (string, error) {
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
			return nil, errors.New("token expired")
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, errors.New("token malformed")
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return nil, errors.New("token signature invalid")
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, errors.New("token not yet valid")
		default:
			return nil, errors.New("parse token failed")
		}
	}

	claims, ok := token.Claims.(*JwtModel)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return claims, nil
}
