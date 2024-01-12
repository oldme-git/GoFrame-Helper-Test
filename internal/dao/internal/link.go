// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LinkDao is the data access object for table link.
type LinkDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns LinkColumns // columns contains all the column names of Table for convenient usage.
}

// LinkColumns defines and stores column names for table link.
type LinkColumns struct {
	Id          string //
	Name        string // 名称
	Description string // 描述
	Link        string // 链接
}

// linkColumns holds the columns for table link.
var linkColumns = LinkColumns{
	Id:          "id",
	Name:        "name",
	Description: "description",
	Link:        "link",
}

// NewLinkDao creates and returns a new DAO object for table data access.
func NewLinkDao() *LinkDao {
	return &LinkDao{
		group:   "default",
		table:   "link",
		columns: linkColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LinkDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LinkDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LinkDao) Columns() LinkColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LinkDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LinkDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LinkDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
