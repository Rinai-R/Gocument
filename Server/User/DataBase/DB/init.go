package DB

import (
	"context"
	"github.com/Rinai-R/Gocument/Server/User/DataBase/DB/ElasticSearch"
	"github.com/Rinai-R/Gocument/Server/User/DataBase/conf/DB"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	models2 "github.com/Rinai-R/Gocument/pkg/models"
	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

var Rdb *redis.Client
var Db *gorm.DB
var ES *elastic.Client

func init() {
	var err error
	//redis连接部分
	Rdb = redis.NewClient(&redis.Options{
		Addr:     conf.UserDB.Redis.Addr,
		Password: conf.UserDB.Redis.Password,
		DB:       conf.UserDB.Redis.DB,
	})

	Logger.Logger.Debug("Redis OK")

	//MySQL连接及其初始化
	dsn := conf.UserDB.MySQL.UserName + ":" + conf.UserDB.MySQL.Password + "@tcp(" + conf.UserDB.MySQL.Addr + ")/" + conf.UserDB.MySQL.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.Logger.Panic(err.Error() + "mysql connect fail")
	}
	err = Db.AutoMigrate(models2.User{})
	if err != nil {
		Logger.Logger.Panic(err.Error() + "mysql Document migrate fail")
	}
	err = Db.AutoMigrate(models2.Document{})
	if err != nil {
		Logger.Logger.Panic(err.Error() + "mysql Document migrate fail")
	}
	err = Db.AutoMigrate(models2.Permission{})
	if err != nil {
		Logger.Logger.Panic(err.Error() + "mysql Document migrate fail")
	}

	Logger.Logger.Debug("Mysql OK")

	//ElasticSearch连接及其初始化
	ES, err = elastic.NewClient(
		elastic.SetURL(conf.UserDB.ElasticSearch.Addr),
		elastic.SetSniff(false),       // 禁用嗅探
		elastic.SetHealthcheck(false), // 禁用健康检查
	)

	if err != nil {
		Logger.Logger.Panic(err.Error() + "elastic connect fail")
	}

	// 检查连接是否成功
	_, _, err = ES.Ping(conf.UserDB.ElasticSearch.Addr).Do(context.Background())
	if err != nil {
		Logger.Logger.Panic(err.Error() + "elastic connect fail")
	}
	Logger.Logger.Debug("Elasticsearch Connected")

	IndexName := conf.UserDB.ElasticSearch.IndexName
	exists, err := ES.IndexExists(IndexName).Do(context.Background())
	if err != nil {
		Logger.Logger.Panic(err.Error() + " Document Exists Request Error")
	}

	if !exists {
		_, err2 := ES.CreateIndex(IndexName).Body(models2.EsDocument).Do(context.Background())
		if err2 != nil {
			Logger.Logger.Panic(err2.Error() + " Document Create Index Request Error")
		}
		Logger.Logger.Debug("Create Document Index OK")
	}

	Skeys := conf.UserDB.ElasticSearch.Sensitive
	exists, err = ES.IndexExists(Skeys).Do(context.Background())
	if err != nil {
		Logger.Logger.Panic(err.Error() + " Sensitive Exists Request Error")
	}
	if !exists {
		_, err2 := ES.CreateIndex(Skeys).Body(ElasticSearch.Keys).Do(context.Background())
		if err2 != nil {
			Logger.Logger.Panic(err2.Error() + " Document Create Index Request Error")
		}
		Logger.Logger.Debug("Sensitive Index Created")
	}
	//敏感词插入
	for i := 0; i < len(conf.UserDB.SKey.Keys); i++ {

		KeyDoc := ElasticSearch.KeyDocument{Key: conf.UserDB.SKey.Keys[i]}
		//插入或更新敏感词
		_, err = ES.Index().
			Index(Skeys).
			Id(strconv.Itoa(i + 1)).
			BodyJson(KeyDoc).
			Do(context.Background())
	}

	Logger.Logger.Debug("Elasticsearch OK")
}
