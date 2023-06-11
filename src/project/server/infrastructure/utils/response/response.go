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
	Error *pb.UserError `json:"error,omitempty"`
	Data  interface{}   `json:"data,omitempty"`
}

var success = &pb.UserError{Code: enums.Errors.Success, Info: "ok"}

func Write(ctx *gin.Context, data interface{}) {
	ctx.SecureJSON(http.StatusOK, data)
}

func WriteError(ctx *gin.Context, err interface{}) {
	switch err := err.(type) {
	case nil:
		Write(ctx, &Base{Error: success})
	case *errors.ErrorResp:
		if err == nil {
			Write(ctx, &Base{})
		} else {
			Write(ctx, &Base{Error: &pb.UserError{Code: err.Code, Info: err.Info}})
		}
	case *pb.UserError:
		Write(ctx, &Base{Error: err})
	default:
		Write(ctx, &Base{Error: &pb.UserError{Code: enums.Errors.Unknown, Info: fmt.Sprintf("%v", err)}})
	}
}

func WriteData(ctx *gin.Context, data interface{}, err error) {
	if err != nil && err.Error() != "" {
		WriteError(ctx, err)
	} else {
		Write(ctx, data)
	}
}
