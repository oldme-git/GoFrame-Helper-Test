// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure for table admin.
type Admin struct {
	Id        uint        `json:"id"        description:""`
	Username  string      `json:"username"  description:"用户名"`
	Password  string      `json:"password"  description:"密码"`
	Nickname  string      `json:"nickname"  description:"昵称"`
	Avatar    string      `json:"avatar"    description:"头像，base64"`
	Register  *gtime.Time `json:"register"  description:"注册时间"`
	Salt      string      `json:"salt"      description:"随机盐值"`
	LastLogin *gtime.Time `json:"lastLogin" description:"最后登录时间"`
}
