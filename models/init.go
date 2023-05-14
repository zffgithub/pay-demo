package models

import (
	"fmt"
	"shop/models/system"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbHost, _ := web.AppConfig.String("mysql::dbhost")
	dbUser, _ := web.AppConfig.String("mysql::dbuser")
	dbPassword, _ := web.AppConfig.String("mysql::dbpasswd")
	dbBase, _ := web.AppConfig.String("mysql::dbbase")
	dbPort, _ := web.AppConfig.String("mysql::dbport")

	link := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbBase)

	logs.Info("mysql init.....", link)

	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", link)
	orm.RegisterModel(new(system.CustomerInfo))
}
