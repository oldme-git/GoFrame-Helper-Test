package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"GoFrame-Helper-Test/api/user/v1"
)

func (c *ControllerV1) AA(ctx context.Context, req *v1.AAReq) (res *v1.AARes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
