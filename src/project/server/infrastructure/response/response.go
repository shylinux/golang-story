// Package response 服务响应处理
package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess             = 100000
	CodeRequestParamInvalid = 100001
	CodeServerErrUnknown    = 200000
)

// BaseRsp 通用响应
type BaseRsp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

// Base 通用响应
type Base struct {
	BaseRsp *BaseRsp    `json:"baseRsp"`
	Data    interface{} `json:"data"`
}

var success = &BaseRsp{Code: CodeSuccess, Msg: "ok"}

// Write 写回数据
func Write(ctx *gin.Context, data interface{}) {
	ctx.SecureJSON(200, data)
}

// WriteBase 写回数据
func WriteBase(ctx *gin.Context, err interface{}) {
	switch err := err.(type) {
	case nil:
		Write(ctx, &Base{BaseRsp: success})

	case *BaseRsp:
		Write(ctx, &Base{BaseRsp: err})

	default:
		Write(ctx, &Base{BaseRsp: &BaseRsp{
			Code: CodeServerErrUnknown,
			Msg:  fmt.Sprintf("%v", err),
		}})
	}
}

// WriteData 写回数据
func WriteData(ctx *gin.Context, data, err interface{}) {
	if err != nil {
		WriteBase(ctx, err)
		return
	}
	Write(ctx, &Base{BaseRsp: success, Data: data})
}

// WriteBaseID 写回带ID的数据
func WriteBaseID(ctx *gin.Context, id int64) error {
	Write(ctx, map[string]interface{}{"baseRsp": success, "ID": id})
	return nil
}

// WriteBasePage 写回带分页的数据
func WriteBasePage(ctx *gin.Context, list interface{}, total int64) error {
	Write(ctx, map[string]interface{}{"baseRsp": success, "total": total, "list": list})
	return nil
}

// WriteParamInvalid 写回参数错误信息
func WriteParamInvalid(ctx *gin.Context, err interface{}) {
	WriteBase(ctx, &BaseRsp{Code: CodeRequestParamInvalid, Msg: fmt.Sprintf("%v", err)})
}
