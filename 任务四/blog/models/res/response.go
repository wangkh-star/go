package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	success = 0
	Error   = 400
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
func ResultOk(data any, msg string, c *gin.Context) {
	Result(success, data, msg, c)
}
func ResultFail(msg string, c *gin.Context) {
	Result(Error, nil, msg, c)
}
