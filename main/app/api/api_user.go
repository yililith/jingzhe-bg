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
	var req *dto.UserLoginDto
	if err := c.ShouldBindJSON(&req); err != nil {
		ctx.result.Fail(c, http.StatusBadRequest, "参数解析失败")
		return
	}
	loginData, err := ctx.service.UserLoginService(req.Username, req.Password)
	if err != nil {
		ctx.result.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.result.Success(c, "login success", loginData)
}

// CreateUserApi
//
//	@Description: 创建用户接口
//	@receiver ctx
//	@param c
func (ctx *UserApi) CreateUserApi(c *gin.Context) {
	var req *dto.CreateUserDto

	if err := c.ShouldBindJSON(&req); err != nil {
		ctx.result.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := ctx.service.CreateUserService(req); err != nil {
		ctx.result.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.result.Success(c, "create success", "")
}

// GetUserPagingApi
//
//	@Description: 分页获取用户数据
//	@receiver ctx
//	@param c
func (ctx *UserApi) GetUserPagingApi(c *gin.Context) {
	var req dto.GetUserPagingDto

	if err := c.ShouldBindQuery(&req); err != nil {
		ctx.result.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

}

// UploadUserImageApi
//
//	@Description: 上传用户图片
//	@receiver ctx
//	@param c
func (ctx UserApi) PutUserImageApi(c *gin.Context) {
	// 接收关联uid
	var req *dto.PutUserAvatarDto
	if err := c.ShouldBindJSON(&req); err != nil {
		ctx.result.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	// 接收文件
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		ctx.result.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer func() {
		_ = file.Close()

	}()
	if err = ctx.service.PutUserAvatarService(req.UID, file, header); err != nil {
		ctx.result.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.result.Success(c, "success", "")
}
