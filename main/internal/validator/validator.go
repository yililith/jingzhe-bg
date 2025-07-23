package validator

import (
	"github.com/go-playground/validator/v10"
	"jingzhe-bg/main/global"
)

func InitValidator() {
	global.GVA_VALIDATOR = validator.New()
}
