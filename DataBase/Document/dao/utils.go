package dao

import (
	"context"
	"fmt"
	"github.com/Rinai-R/Gocument/DataBase/DB"
	conf "github.com/Rinai-R/Gocument/DataBase/conf/DB"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/models"
	"github.com/olivere/elastic/v7"
	"math/rand"
	"strconv"
	"time"
)

func GetId(ctx context.Context, username string) (int, error) {
	key := fmt.Sprintf("id:%v", username)
	//设置过期时间
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	mul := r.Intn(60) + 60
	//缓存读取
	ans, err := DB.Rdb.Get(ctx, key).Result()
	if err == nil {
		Logger.Logger.Debug("Dao: Redis Cache Read UserId")
		DB.Rdb.Expire(ctx, key, time.Duration(mul)*time.Minute)
		return strconv.Atoi(ans)
	}
	//redis没读取到，去mysql里面读取
	var user models.User
	res := DB.Db.Where("username = ?", username).Select("id").First(&user)
	if res.Error != nil {
		Logger.Logger.Debug("Dao: InternalError " + res.Error.Error())
		return 0, res.Error
	}
	//设置缓存
	_, err = DB.Rdb.SetEx(ctx, key, user.Id, time.Duration(mul)*time.Minute).Result()
	if err != nil {
		Logger.Logger.Debug("Dao: Redis Cache Write Error " + err.Error())
	}
	Logger.Logger.Debug("Dao: Redis Cache Write UserId")
	return int(user.Id), nil
}

func IsHost(ctx context.Context, username string, DocumentId int) error {
	id, err := GetId(ctx, username)
	if err != nil {
		return err
	}
	res := DB.Db.Model(&models.Document{}).Where("user_id = ? AND id = ?", id, DocumentId).First(&models.Document{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func SensitiveCheck(ctx context.Context, str string) bool {
	matchQuery := elastic.NewMatchQuery("key", str)

	// 根据信息查询，是否命中敏感词汇
	searchResult, err := DB.ES.Search().
		Index(conf.DB.ElasticSearch.Sensitive).
		Query(matchQuery).
		TrackTotalHits(true). // 跟踪总命中数
		Size(100).
		Do(ctx) // 执行查询
	if err != nil {
		Logger.Logger.Debug("Sensitive Check Error " + err.Error())
		return false
	}
	fmt.Println(searchResult.Hits.TotalHits.Value)
	// 处理响应
	totalHits := searchResult.Hits.TotalHits.Value
	if totalHits == 0 {
		return true
	}
	return false
}
