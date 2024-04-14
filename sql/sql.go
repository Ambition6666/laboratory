package sql

import (
	"fmt"

	"laboratory/config"
	"laboratory/model"

	re "github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	rediscc *re.Client
)

func InitSQL() {
	initMySQL()
	initRedis()
}

// 注册mysql
func initMySQL() {
	conf := config.GetMySQLConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Pwd, conf.Host, conf.Port, conf.DbName)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	initMODEL()
}

// 注册redis
func initRedis() {
	conf := config.GetRedisConfig()
	addr := conf.Host + ":" + conf.Port
	rediscc = re.NewClient(&re.Options{
		Addr:     addr,
		Password: conf.Pwd,
		DB:       0, // use default DB
	})
}

func GetMySQLDB() *gorm.DB {
	return db
}

func GetRedisDB() *re.Client {
	return rediscc
}

func initMODEL() {
	db.AutoMigrate(&model.Student{}, &model.User{})
	db.AutoMigrate(&model.Teacher{}, &model.User{})
	db.AutoMigrate(&model.Laboratory{}, &model.Teacher{})
}
