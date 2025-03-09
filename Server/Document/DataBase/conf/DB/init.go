package conf

import (
	ElasticSearch2 "github.com/Rinai-R/Gocument/Server/Document/DataBase/DB/ElasticSearch"
	"github.com/Rinai-R/Gocument/Server/Document/DataBase/DB/MySQL"
	"github.com/Rinai-R/Gocument/Server/Document/DataBase/DB/Redis"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	"github.com/spf13/viper"
)

type DBConfig struct {
	MySQL.MySQL
	Redis.Redis
	ElasticSearch2.ElasticSearch
	ElasticSearch2.SKey
}

var DocDB *DBConfig

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("db")
	viper.SetConfigFile("./Server/Document/DataBase/conf/DB/db.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		Logger.Logger.Panic("Viper: " + err.Error())
	}
	err = viper.Unmarshal(&DocDB)
	if err != nil {
		Logger.Logger.Panic("Viper: " + err.Error())
	}

	Logger.Logger.Debug("Viper: DataBase OK")

}
