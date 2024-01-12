// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReplyDao is the data access object for table reply.
type ReplyDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns ReplyColumns // columns contains all the column names of Table for convenient usage.
}

// ReplyColumns defines and stores column names for table reply.
type ReplyColumns struct {
	Id        string //
	Aid       string // 文章id
	Rid       string // 根回复id，一个根可视为一个楼层
	Pid       string // 回复的id
	PName     string // 回复的名称
	Email     string // 回复人邮箱
	Name      string // 回复人名称
	Site      string // 回复人网站
	Content   string // 回复内容
	Status    string // 状态： 1待审核 2审核通过 3审核失败
	Reason    string // 审核失败原因
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// replyColumns holds the columns for table reply.
var replyColumns = ReplyColumns{
	Id:        "id",
	Aid:       "aid",
	Rid:       "rid",
	Pid:       "pid",
	PName:     "p_name",
	Email:     "email",
	Name:      "name",
	Site:      "site",
	Content:   "content",
	Status:    "status",
	Reason:    "reason",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewReplyDao creates and returns a new DAO object for table data access.
func NewReplyDao() *ReplyDao {
	return &ReplyDao{
		group:   "default",
		table:   "reply",
		columns: replyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReplyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReplyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReplyDao) Columns() ReplyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReplyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReplyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReplyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
