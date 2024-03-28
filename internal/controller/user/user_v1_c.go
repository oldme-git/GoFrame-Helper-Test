package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"GoFrame-Helper-Test/api/user/v1"
)

func (c *ControllerV1) C(ctx context.Context, req *v1.CReq) (res *v1.CRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
