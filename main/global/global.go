package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"jingzhe-bg/main/internal/model"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG *model.Conf
	GVA_LOGGER *zap.Logger
)
