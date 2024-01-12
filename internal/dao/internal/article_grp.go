// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleGrpDao is the data access object for table article_grp.
type ArticleGrpDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns ArticleGrpColumns // columns contains all the column names of Table for convenient usage.
}

// ArticleGrpColumns defines and stores column names for table article_grp.
type ArticleGrpColumns struct {
	Id          string //
	Name        string // 名称
	Tags        string // 标签，依英文逗号隔开
	Description string // 简介
	Onshow      string // 是否显示
	Order       string // 排序，越大越靠前
}

// articleGrpColumns holds the columns for table article_grp.
var articleGrpColumns = ArticleGrpColumns{
	Id:          "id",
	Name:        "name",
	Tags:        "tags",
	Description: "description",
	Onshow:      "onshow",
	Order:       "order",
}

// NewArticleGrpDao creates and returns a new DAO object for table data access.
func NewArticleGrpDao() *ArticleGrpDao {
	return &ArticleGrpDao{
		group:   "default",
		table:   "article_grp",
		columns: articleGrpColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ArticleGrpDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ArticleGrpDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ArticleGrpDao) Columns() ArticleGrpColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ArticleGrpDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ArticleGrpDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ArticleGrpDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
