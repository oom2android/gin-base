package controller

import (
	"gin-base/api/v1/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (c *UserController) Router(r *gin.Engine) {

	api := r.Group("/api/v1/user")
	{
		api.GET("/sendcode", c.sendSmsCode)
	}
}

func (c *UserController) smsLogin(ctx *gin.Context) {

}

func (c *UserController) sendSmsCode(ctx *gin.Context) {
	phone, exits := ctx.GetQuery("phone")
	if !exits {
		ctx.JSON(200, map[string]interface{}{
			"code": 200,
			"msg":  "参数解析失败",
		})
		return
	}
	service := service.UserService{}
	isSend := service.SendCode(phone)
	if isSend {
		ctx.JSON(200, map[string]interface{}{
			"code": 200,
			"msg":  "发送成功",
		})
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"code": 200,
		"msg":  "发送失败",
	})
}
