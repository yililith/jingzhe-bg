package ead

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

func EncryptWithPublicKeyOAEP(publicKey *rsa.PublicKey, message string) (string, error) {
	if message == "" {
		return "", nil
	}

	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		[]byte(message),
		nil, // 可选的标签
	)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

// DecryptWithPrivateKey 使用RSA私钥解密消息
func DecryptWithPrivateKey(privateKey *rsa.PrivateKey, message string) (string, error) {
	if message == "" {
		return "", nil // 或返回错误
	}

	decryptedBytes, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", err
	}

	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decryptedBytes)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}
