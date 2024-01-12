// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"GoFrame-Helper-Test/api/user/v1"
)

type IUserV1 interface {
	Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
}
