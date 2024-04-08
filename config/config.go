package config

import (
	"os"
	"github.com/spf13/viper"
)

type Config struct {
	Mysql MySQL `yaml:"mysql"`
}

type MySQL struct {
	Port   string `yaml:"port"`
	Host   string `yaml:"host"`
	Pwd    string `yaml:"password"`
	DbName string `yaml:"dbname"`
	User   string `yaml:"user"`
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

func GetMySQLDB() MySQL{
	return config.Mysql
}

func init() {
	InitConfig()
}