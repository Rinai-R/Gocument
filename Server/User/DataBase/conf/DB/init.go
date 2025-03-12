package conf

import (
	ElasticSearch2 "github.com/Rinai-R/Gocument/Server/User/DataBase/DB/ElasticSearch"
	"github.com/Rinai-R/Gocument/Server/User/DataBase/DB/MySQL"
	"github.com/Rinai-R/Gocument/Server/User/DataBase/DB/Redis"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	"github.com/spf13/viper"
)

type DBConfig struct {
	MySQL.MySQL
	Redis.Redis
	ElasticSearch2.ElasticSearch
	ElasticSearch2.SKey
}

var UserDB *DBConfig

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("db")
	viper.SetConfigFile("/home/rinai/PROJECTS/Gocument/Server/User/DataBase/conf/DB/db.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		Logger.Logger.Panic("Viper: " + err.Error())
	}
	err = viper.Unmarshal(&UserDB)
	if err != nil {
		Logger.Logger.Panic("Viper: " + err.Error())
	}

	Logger.Logger.Debug("Viper: DataBase OK")

}
