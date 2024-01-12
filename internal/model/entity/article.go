// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Article is the golang structure for table article.
type Article struct {
	Id          uint        `json:"id"          description:""`
	GrpId       uint        `json:"grpId"       description:"分组id"`
	Title       string      `json:"title"       description:"标题"`
	Author      string      `json:"author"      description:"作者"`
	Thumb       string      `json:"thumb"       description:"图片地址"`
	Tags        string      `json:"tags"        description:"标签，依英文逗号隔开"`
	Description string      `json:"description" description:"简介"`
	Content     string      `json:"content"     description:"内容"`
	Order       int         `json:"order"       description:"排序，越大越靠前"`
	Ontop       uint        `json:"ontop"       description:"是否置顶"`
	Onshow      uint        `json:"onshow"      description:"是否显示"`
	Hist        uint        `json:"hist"        description:"点击数"`
	Post        uint        `json:"post"        description:"评论数"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"更新时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   description:"删除时间"`
	LastedAt    *gtime.Time `json:"lastedAt"    description:"最后浏览时间"`
}
