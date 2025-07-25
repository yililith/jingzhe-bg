package router

import (
	"github.com/gin-gonic/gin"
	"jingzhe-bg/main/app/api"
)

func init() {
	routerAuthRole = append(routerAuthRole, authUserRouter)
	routerNoAuthRole = append(routerNoAuthRole, noAuthUserRouter)
}

// 需要认证的路由
func authUserRouter(r *gin.RouterGroup) {
	col := api.NewUserApi()
	group := r.Group("/user")
	{
		// 登录
		group.POST("/create/user", col.CreateUserApi)
		group.POST("/put/avatar", col.PutUserImageApi)
	}
}

// noAuthUserRouter
//
//	@Description: 不需认证路由表
//	@param r
func noAuthUserRouter(r *gin.RouterGroup) {
	col := api.NewUserApi()
	group := r.Group("/user")
	{
		group.POST("/login", col.LoginApi)
	}
}
