package global

import (
	"crypto/rsa"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"jingzhe-bg/main/internal/model"
)

var (
	GVA_DB          *gorm.DB        // 数据库连接
	GVA_CONFIG      *model.Conf     // 全局配置
	GVA_LOGGER      *zap.Logger     // 日志
	GVA_PRIVATE_KEY *rsa.PrivateKey // 全局私钥
	GVA_PUBLIC_KEY  *rsa.PublicKey  // 全局公钥
	GVA_VALIDATOR   *validator.Validate
)
