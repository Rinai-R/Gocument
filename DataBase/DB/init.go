package DB

import (
	"context"
	"github.com/Rinai-R/Gocument/DataBase/conf/DB"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/models"
	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Rdb *redis.Client
var Db *gorm.DB
var ES *elastic.Client

func init() {
	var err error
	//redis连接部分
	Rdb = redis.NewClient(&redis.Options{
		Addr:     conf.DB.Redis.Addr,
		Password: conf.DB.Redis.Password,
		DB:       conf.DB.Redis.DB,
	})

	Logger.Logger.Debug("Redis OK")

	//MySQL连接及其初始化
	dsn := conf.DB.MySQL.UserName + ":" + conf.DB.MySQL.Password + "@tcp(" + conf.DB.MySQL.Addr + ")/" + conf.DB.MySQL.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.Logger.Panic(err.Error() + "mysql connect fail")
	}
	err = Db.AutoMigrate(models.User{})
	if err != nil {
		Logger.Logger.Panic(err.Error() + "mysql Document migrate fail")
	}
	err = Db.AutoMigrate(models.Document{})
	if err != nil {
		Logger.Logger.Panic(err.Error() + "mysql Document migrate fail")
	}

	Logger.Logger.Debug("Mysql OK")

	//ElasticSearch连接及其初始化
	ES, err = elastic.NewClient(
		elastic.SetURL(conf.DB.ElasticSearch.Addr),
		elastic.SetSniff(false),       // 禁用嗅探
		elastic.SetHealthcheck(false), // 禁用健康检查
	)

	if err != nil {
		Logger.Logger.Panic(err.Error() + "elastic connect fail")
	}

	// 检查连接是否成功
	_, _, err = ES.Ping(conf.DB.ElasticSearch.Addr).Do(context.Background())
	if err != nil {
		Logger.Logger.Panic(err.Error() + "elastic connect fail")
	}
	Logger.Logger.Debug("Elasticsearch Connected")

	IndexName := conf.DB.ElasticSearch.IndexName
	exists, err := ES.IndexExists(IndexName).Do(context.Background())
	if err != nil {
		Logger.Logger.Panic(err.Error() + " Document Exists Request Error")
	}

	if !exists {
		_, err2 := ES.CreateIndex(IndexName).Body(models.EsDocument).Do(context.Background())
		if err2 != nil {
			Logger.Logger.Panic(err2.Error() + " Document Create Index Request Error")
		}
	}
	Logger.Logger.Debug("Elasticsearch OK")

}
