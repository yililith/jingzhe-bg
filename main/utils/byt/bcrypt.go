package byt

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"jingzhe-bg/main/global"
)

// 加密
func HashPassword(str string) (string, error) {
	hashStr, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		global.GVA_LOGGER.Error("密码加密失败", zap.String("err", err.Error()))
		return "", err
	}
	return string(hashStr), nil
}

// 验证
func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
