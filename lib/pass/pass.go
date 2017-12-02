package pass

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

type Pass struct {
	Id int `xorm:"int pk autoincr 'id'"`
	UserId string `xorm:"user_id"`
	PointId string `xorm:"point_id"`
}

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "admin:@/ferry")
	if err != nil {
		panic(err)
	}
}

func Insert(userId string, pointId string) Pass {
	pass := Pass{UserId: userId, PointId: pointId}
	engine.Insert(&pass)
	return pass
}

func Get(id int) Pass {
	var pass = Pass{Id: id}
	engine.Get(&pass)
	return pass
}
