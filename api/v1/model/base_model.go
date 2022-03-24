package model

import (
	"fmt"
	"gin-base/tool"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Model struct {
	Id         string `json:"id" xorm:"pk index"`
	CreateTime int64  `json:"create_time" xorm:"bigint"`
}

var DB *xorm.Engine

func NewDbEngine() {

	config := tool.GetConfig().Database

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Charset)

	engine, err := xorm.NewEngine(tool.GetConfig().Database.Driver, conn)
	if err != nil {
		panic(err)
	}
	err = engine.Sync2(new(SmsCodeModel), new(UserModel))

	if err != nil {
		panic(err)
	}
	engine.ShowSQL(config.ShowSql)

	DB = engine

}
