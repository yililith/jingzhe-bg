package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

var PublicKey *rsa.PublicKey
var PrivateKey *rsa.PrivateKey

func InitKey() error {
	var err error
	PublicKey, err = loadPublicKey()
	if err != nil {
		return err
	}
	PrivateKey, err = loadPrivateKey()
	if err != nil {
		return err
	}
	return nil
}

// 加载公钥
func loadPublicKey() (*rsa.PublicKey, error) {
	data, err_one := os.ReadFile("./public_key.pem")
	if err_one != nil {
		return nil, err_one
	}
	block, _ := pem.Decode(data)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}
	pubKey, err_two := x509.ParsePKIXPublicKey(block.Bytes)
	if err_two != nil {
		fmt.Println(err_two.Error())
		return nil, err_two
	}
	key, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}
	return key, nil
}

// 加载私钥
func loadPrivateKey() (*rsa.PrivateKey, error) {
	data, err_one := os.ReadFile("./private_key.pem")
	if err_one != nil {
		return nil, err_one
	}
	block, _ := pem.Decode(data)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}
	privKey, err_two := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err_two != nil {
		fmt.Println(err_two.Error())
		return nil, err_two
	}
	key, ok := privKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("Error loading private key")
	}
	return key, nil
}
