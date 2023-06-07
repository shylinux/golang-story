package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
)

type Base struct {
	Error *pb.Error   `json:"error"`
	Data  interface{} `json:"data"`
}

var success = &pb.Error{Code: enums.Errors.Success, Info: "ok"}

func Write(ctx *gin.Context, data interface{}) {
	ctx.SecureJSON(http.StatusOK, data)
}

func WriteError(ctx *gin.Context, err interface{}) {
	switch err := err.(type) {
	case nil:
		Write(ctx, &Base{Error: success})
	case *errors.ErrorResp:
		Write(ctx, &Base{Error: &pb.Error{
			Code: err.Code,
			Info: err.Info,
		}})
	case *pb.Error:
		Write(ctx, &Base{Error: err})
	default:
		Write(ctx, &Base{Error: &pb.Error{
			Code: enums.Errors.Unknown,
			Info: fmt.Sprintf("%v", err),
		}})
	}
}

func WriteData(ctx *gin.Context, data, err interface{}) {
	if err != nil {
		WriteError(ctx, err)
		return
	}
	Write(ctx, data)
}
