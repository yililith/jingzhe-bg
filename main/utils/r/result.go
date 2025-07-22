package r

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct{}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 成功返回范式
func (ctx *Result) Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: http.StatusOK,
		Msg:  msg,
		Data: data,
	})
}

// 失败返回范式
func (ctx *Result) Fail(c *gin.Context, code int, msg string) {
	c.Set("code", code)
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
