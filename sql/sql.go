package sql

import (
	"fmt"

	"laboratory/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitSQL() {
	initMySQL()
}
// 注册mysql
func initMySQL() {
	conf := config.GetMySQLDB()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Pwd, conf.Host, conf.Port, conf.DbName)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func GetMySQLDB() *gorm.DB {
	return db
}
