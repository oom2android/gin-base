package tool

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 200
	FAILED  = 400
)

func Succss(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  "成功",
	})
}
func SuccessObj(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  "成功",
		"data": v,
	})
}
func Failed(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"msg":  "失败",
	})
}
func FailedStr(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"msg":  msg,
	})
}
func ParseFalied(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"msg":  "参数解析失败",
	})
}
