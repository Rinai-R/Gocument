package DB

import (
	"github.com/Rinai-R/Gocument/DataBase/conf/DB"
	"github.com/Rinai-R/Gocument/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Rdb *redis.Client
var Db *gorm.DB

func init() {
	var err error
	Rdb = redis.NewClient(&redis.Options{
		Addr:     conf.DB.Redis.Addr,
		Password: conf.DB.Redis.Password,
		DB:       conf.DB.Redis.DB,
	})
	dsn := conf.DB.MySQL.UserName + ":" + conf.DB.MySQL.Password + "@tcp(" + conf.DB.MySQL.Addr + ")/" + conf.DB.MySQL.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = Db.AutoMigrate(models.User{})
	if err != nil {
		return
	}
}
