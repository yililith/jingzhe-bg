package oss

import (
	"context"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"io"
	"jingzhe-bg/main/global"
	"jingzhe-bg/main/utils/er"
	"time"
)

// PutObject
//
//	@Description: 上传文件
//	@param client
//	@param bucketName
//	@param objectName
//	@param file
//	@param size
//	@return *minio.UploadInfo
//	@return error
func PutObject(client *minio.Client, bucketName string, objectName string, file io.Reader, size int64) (*minio.UploadInfo, error) {
	info, err := client.PutObject(context.Background(), bucketName, objectName, file, size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		global.GVA_LOGGER.Error("上传文件失败", zap.String("err", err.Error()))
		return nil, er.JZError("上传文件失败")
	}
	return &info, nil
}

func DeleFile() {

}

// GeneratePresignedURL
//
//	@Description: 生成签名文件URL
//	@param client
//	@param bucketName
//	@param objectName
//	@param expiry
//	@return string
//	@return error
func GeneratePresignedURL(client *minio.Client, bucketName, objectName string, expiry time.Duration) (string, error) {

	// 生成签名URL
	presignedURL, err := client.PresignedGetObject(context.Background(), bucketName, objectName, expiry, nil)
	if err != nil {
		global.GVA_LOGGER.Error("获取签名链接失败", zap.String("err", err.Error()))
		return "", er.JZError("获取头像链接失败")
	}
	return presignedURL.String(), nil
}
