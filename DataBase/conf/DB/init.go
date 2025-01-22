package conf

import (
	"github.com/Rinai-R/Gocument/DataBase/DB/MySQL"
	"github.com/Rinai-R/Gocument/DataBase/DB/Redis"
	"github.com/spf13/viper"
)

type DBConfig struct {
	MySQL.MySQL
	Redis.Redis
}

var DB *DBConfig

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("db")
	viper.SetConfigFile("./DataBase/conf/DB/db.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&DB)
	if err != nil {
		panic(err)
	}

}
