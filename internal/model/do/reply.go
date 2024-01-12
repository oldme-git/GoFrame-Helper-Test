// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Reply is the golang structure of table reply for DAO operations like Where/Data.
type Reply struct {
	g.Meta    `orm:"table:reply, do:true"`
	Id        interface{} //
	Aid       interface{} // 文章id
	Rid       interface{} // 根回复id，一个根可视为一个楼层
	Pid       interface{} // 回复的id
	PName     interface{} // 回复的名称
	Email     interface{} // 回复人邮箱
	Name      interface{} // 回复人名称
	Site      interface{} // 回复人网站
	Content   interface{} // 回复内容
	Status    interface{} // 状态： 1待审核 2审核通过 3审核失败
	Reason    interface{} // 审核失败原因
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
