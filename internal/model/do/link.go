// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Link is the golang structure of table link for DAO operations like Where/Data.
type Link struct {
	g.Meta      `orm:"table:link, do:true"`
	Id          interface{} //
	Name        interface{} // 名称
	Description interface{} // 描述
	Link        interface{} // 链接
}
