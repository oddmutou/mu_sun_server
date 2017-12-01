package user

import (
    "github.com/go-xorm/xorm"
    _ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

type User struct {
    Id int `xorm:"int pk autoincr 'id'"`
    Name string `xorm:"name"`
    Score int `xorm:"score"`
}

func init() {
    var err error
    engine, err = xorm.NewEngine("mysql", "admin:@/ferry")
    if err != nil {
        panic(err)
    }
}

func Insert(name string, score int) User {
    user := User{Name: name, Score: score}
    engine.Insert(&user)
    return user
}

func Update(id int, name string, score int) User {
    user := User{Id: id, Name: name, Score: score}
    engine.Where("id = ?", id).Update(&user)
    return user
}

func Delete(id int) User {
    user := User{Id: id}
    engine.Where("id = ?", id).Delete(&user)
    return user
}

func Get(id int) User {
    var user = User{Id: id}
    engine.Get(&user)
    return user
}

func GetAll() []User {
    var allUsers []User
    engine.Desc("Id").Limit(50).Find(&allUsers)
    return allUsers
}
