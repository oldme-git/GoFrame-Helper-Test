// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleDao is the data access object for table article.
type ArticleDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ArticleColumns // columns contains all the column names of Table for convenient usage.
}

// ArticleColumns defines and stores column names for table article.
type ArticleColumns struct {
	Id          string //
	GrpId       string // 分组id
	Title       string // 标题
	Author      string // 作者
	Thumb       string // 图片地址
	Tags        string // 标签，依英文逗号隔开
	Description string // 简介
	Content     string // 内容
	Order       string // 排序，越大越靠前
	Ontop       string // 是否置顶
	Onshow      string // 是否显示
	Hist        string // 点击数
	Post        string // 评论数
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 删除时间
	LastedAt    string // 最后浏览时间
}

// articleColumns holds the columns for table article.
var articleColumns = ArticleColumns{
	Id:          "id",
	GrpId:       "grp_id",
	Title:       "title",
	Author:      "author",
	Thumb:       "thumb",
	Tags:        "tags",
	Description: "description",
	Content:     "content",
	Order:       "order",
	Ontop:       "ontop",
	Onshow:      "onshow",
	Hist:        "hist",
	Post:        "post",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	LastedAt:    "lasted_at",
}

// NewArticleDao creates and returns a new DAO object for table data access.
func NewArticleDao() *ArticleDao {
	return &ArticleDao{
		group:   "default",
		table:   "article",
		columns: articleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ArticleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ArticleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ArticleDao) Columns() ArticleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ArticleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ArticleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ArticleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
