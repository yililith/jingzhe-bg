package api

import (
	"github.com/gin-gonic/gin"
	"jingzhe-bg/main/utils"
	"net/http"
)

type UserApi struct {
	result *utils.Result
}

func NewUserApi() *UserApi {
	return &UserApi{}
}

func (ctx *UserApi) UserApi_login(c *gin.Context) {
	//ctx.result.Success(c, "login success", nil)
	ctx.result.Fail(c, http.StatusInternalServerError, "登录失败")
}
