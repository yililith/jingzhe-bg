package byt

import "golang.org/x/crypto/bcrypt"

// 加密
func HashPassword(str string) (string, error) {
	hashStr, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashStr), nil
}

// 验证
func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
