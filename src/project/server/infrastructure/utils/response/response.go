package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
)

const (
	CodeSuccess             = 100000
	CodeRequestParamInvalid = 100001
	CodeServerErrUnknown    = 200000
)

type Base struct {
	BaseResp *pb.BaseReply `json:"baseResp"`
	Data     interface{}   `json:"data"`
}

var success = &pb.BaseReply{Code: CodeSuccess, Info: "ok"}

func Write(ctx *gin.Context, data interface{}) {
	ctx.SecureJSON(200, data)
}

func WriteBase(ctx *gin.Context, err interface{}) {
	switch err := err.(type) {
	case nil:
		Write(ctx, &Base{BaseResp: success})
	case *pb.BaseReply:
		Write(ctx, &Base{BaseResp: err})
	default:
		Write(ctx, &Base{BaseResp: &pb.BaseReply{
			Code: CodeServerErrUnknown,
			Info: fmt.Sprintf("%v", err),
		}})
	}
}

func WriteData(ctx *gin.Context, data, err interface{}) {
	if err != nil {
		WriteBase(ctx, err)
		return
	}
	Write(ctx, data)
}

func WriteBaseID(ctx *gin.Context, id int64) error {
	Write(ctx, map[string]interface{}{"baseResp": success, "ID": id})
	return nil
}

func WriteBasePage(ctx *gin.Context, list interface{}, total int64) error {
	Write(ctx, map[string]interface{}{"baseResp": success, "total": total, "list": list})
	return nil
}

func WriteParamInvalid(ctx *gin.Context, err interface{}) {
	WriteBase(ctx, &pb.BaseReply{Code: CodeRequestParamInvalid, Info: fmt.Sprintf("%v", err)})
}
