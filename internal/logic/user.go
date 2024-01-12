package logic

import (
	"GoFrame-Helper-Test/internal/dao"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gctx"
)

var ctx = gctx.New()

func A() {
	var db *gdb.Model
	db = dao.Article.Ctx(ctx).Where("")
	fmt.Println(db)
}

func B() {
	db := dao.Article.Ctx(ctx).Where("")
	fmt.Println(db)
}

func C() {
	var db = dao.Article.Ctx(ctx).Where("")
	fmt.Println(db)
}

func D() {
	db := dao.Article.Ctx(ctx).Where("")
	db = db.Where("")
	db.Where("")
	db = db.Where("")
}

func E() {
	dao.Article.Ctx(ctx).Where("")
}
