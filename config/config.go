package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	MySQL 
	Redis 
	EmailInfo 
}

type MySQL struct {
	Host   string
	Port   string 
	Pwd    string 
	DbName string 
	User   string 
}

type Redis struct {
	Host   string 
	Port   string 
	Pwd    string 
}

type EmailInfo struct {
	Addr  string 
	Host  string 
	From  string 
	Email string 
	Auth  string 
}

var config Config

func InitConfig() {
	workdir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workdir + "/config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if err = viper.Unmarshal(&config); err!= nil {
		panic(err)
	}
}

func GetMySQLConfig() MySQL{
	return config.MySQL
}

func GetRedisConfig() Redis{
	return config.Redis
}

func GetEmailInfo() EmailInfo {
	return config.EmailInfo
}

func init() {
	InitConfig()
}