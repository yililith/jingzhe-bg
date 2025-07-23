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

// LoginApi
//
//	@Description: 登录接口
//	@receiver ctx
//	@param c
func (ctx *UserApi) LoginApi(c *gin.Context) {
	var accept_data *dto.UserLoginDto
	if err := c.ShouldBindJSON(&accept_data); err != nil {
		ctx.result.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	login_data, service_err := ctx.service.UserLoginService(accept_data.Username, accept_data.Password)
	if service_err != nil {
		ctx.result.Fail(c, http.StatusInternalServerError, service_err.Error())
		return
	}
	ctx.result.Success(c, "login success", login_data)
}

// CreateUserApi
//
//	@Description: 创建用户接口
//	@receiver ctx
//	@param c
func (ctx *UserApi) CreateUserApi(c *gin.Context) {
	var accept_data *dto.CreateUserDto

	if err := c.ShouldBindJSON(&accept_data); err != nil {
		ctx.result.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	if create_err := ctx.service.CreateUserService(accept_data); create_err != nil {
		ctx.result.Fail(c, http.StatusInternalServerError, create_err.Error())
		return
	}
	ctx.result.Success(c, "create success", "")
}
