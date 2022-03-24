package controller

import "github.com/gin-gonic/gin"

type HelloController struct {
}

func (hello *HelloController) Router(r *gin.Engine) {

	apiv1 := r.Group("/api/v1")
	{

		apiv1.GET("/hello", hello.hello)

	}

}

func (hello *HelloController) hello(ctx *gin.Context) {

}
