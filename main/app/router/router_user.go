package router

import (
	"github.com/gin-gonic/gin"
	"jingzhe-bg/main/app/api"
)

func init() {
	routerNoAuthRole = append(routerNoAuthRole, noAuthUserRouter)
}

//// 需要认证的路由
//func authUserRouter(r *gin.RouterGroup) {
//	col := api.NewUserApi()
//	group := r.Group("/user")
//	{
//		// 登录
//		group.POST("/login", col.UserApi_login)
//	}
//}

func noAuthUserRouter(r *gin.RouterGroup) {
	col := api.NewUserApi()
	group := r.Group("/user")
	{
		group.POST("/login", col.UserApi_login)
	}
}
