package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

// 后端返回格式
type wrappedResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

/*
统一后端返回格式
@param status      统一状态码，若success可赋nil
@param data        返回数据
*/
func (g *Gin) Response(status *Status, data any) {
	resp := new(wrappedResp)
	if status != nil {
		resp.Code = status.Code
		resp.Msg = status.Msg
	} else {
		resp.Code = SUCCESS.Code
		resp.Msg = SUCCESS.Msg
	}
	resp.Data = data

	g.C.JSON(http.StatusOK, resp)
}
