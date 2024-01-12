package logic

import (
	"GoFrame-Helper-Test/internal/dao"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
)

func ormA() {
	var db *gdb.Model
	db = dao.Article.Ctx(ctx).Where("")
	fmt.Println(db)
}

func ormB() {
	db := dao.Article.Ctx(ctx).Where("")
	fmt.Println(db)
}

func ormC() {
	var db = dao.Article.Ctx(ctx).Where("")
	fmt.Println(db)
}

func ormD() {
	db := dao.Article.Ctx(ctx).Where("")
	db = db.Where("")
	db.Where("")
	db = db.Where("")
}

func ormE() {
	dao.Article.Ctx(ctx).Where("")
}
