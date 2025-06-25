package router

import (
	"github.com/gin-gonic/gin"
	"jingzhe-bg/main/app/api"
)

func init() {
	routerAuthRole = append(routerAuthRole, authUserRouter)
}

// 需要认证的路由
func authUserRouter(r *gin.RouterGroup) {
	col := api.NewUserApi()
	group := r.Group("/user")
	{
		// 登录
		group.GET("/login", col.UserApi_login)
	}
}
