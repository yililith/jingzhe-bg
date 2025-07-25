package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"jingzhe-bg/main/global"
)

func InitMinio() error {
	conf := global.GVA_CONFIG.Minio

	client, err := minio.New(conf.EndPoInt, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKeyID, conf.SecretAccessKey, ""),
		Secure: conf.UseSLL,
	})

	if err != nil {
		return err
	}
	// 赋值到全局变量
	global.GAV_MINIO_CLIENT = client
	return nil
}
