// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Reply is the golang structure for table reply.
type Reply struct {
	Id        uint        `json:"id"        description:""`
	Aid       int         `json:"aid"       description:"文章id"`
	Rid       int         `json:"rid"       description:"根回复id，一个根可视为一个楼层"`
	Pid       int         `json:"pid"       description:"回复的id"`
	PName     string      `json:"pName"     description:"回复的名称"`
	Email     string      `json:"email"     description:"回复人邮箱"`
	Name      string      `json:"name"      description:"回复人名称"`
	Site      string      `json:"site"      description:"回复人网站"`
	Content   string      `json:"content"   description:"回复内容"`
	Status    int         `json:"status"    description:"状态： 1待审核 2审核通过 3审核失败"`
	Reason    string      `json:"reason"    description:"审核失败原因"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
