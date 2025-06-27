package api

import (
	"github.com/gin-gonic/gin"
	"jingzhe-bg/main/app/dto"
	"jingzhe-bg/main/app/service"
	"jingzhe-bg/main/utils/r"
	"net/http"
)

type UserApi struct {
	result  *r.Result
	service *service.UserService
}

func NewUserApi() *UserApi {
	return &UserApi{
		service: service.NewUserService(),
	}
}

func (ctx *UserApi) UserApi_login(c *gin.Context) {
	//ctx.r.Success(c, "login success", nil)
	var accept_data *dto.ResUserLoginDto
	if err := c.ShouldBindJSON(&accept_data); err != nil {
		ctx.result.Fail(c, http.StatusBadRequest, err.Error())
	}
	login_data, service_err := ctx.service.UserLoginService(accept_data.Username, accept_data.Password)
	if service_err != nil {
		ctx.result.Fail(c, http.StatusInternalServerError, service_err.Error())
		return
	}
	ctx.result.Success(c, "login success", login_data)
}
