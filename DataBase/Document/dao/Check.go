package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/Rinai-R/Gocument/DataBase/DB"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/models"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func Check(ctx context.Context, permission models.Permission) (error, string) {
	var IsPrivate bool
	var Permission bool
	key1 := fmt.Sprintf("DocuId:%v", permission.DocumentId)
	key2 := fmt.Sprintf("Perm:%v:%v", permission.DocumentId, permission.UserId)
	//设置过期时间
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	mul := r.Intn(60) + 60
	//看看是否公开，先缓存读取
	ans, err := DB.Rdb.Get(ctx, key1).Result()
	if err == nil {
		Logger.Logger.Debug("Dao: Redis Cache Read IsPrivate")
		if ans == "false" {
			IsPrivate = false
		} else {
			IsPrivate = true
		}
		DB.Rdb.Expire(ctx, key1, time.Duration(mul)*time.Hour)
	} else {
		//没读取到，换mysql读取
		res := DB.Db.Model(&models.Document{}).Where("id = ?", permission.DocumentId).Select("IsPrivate").First(&IsPrivate)
		if res.Error != nil {
			Logger.Logger.Debug("Dao: document not found " + res.Error.Error())
			return res.Error, "notfound"
		}
		if IsPrivate {
			DB.Rdb.SetEx(ctx, key1, "true", time.Duration(mul)*time.Hour)
			Logger.Logger.Debug("Dao: Redis Cache Write IsPrivate ")
		} else {
			DB.Rdb.SetEx(ctx, key1, "false", time.Duration(mul)*time.Hour)
			Logger.Logger.Debug("Dao: Redis Cache Write IsPrivate ")
		}
	}

	//先读取缓存里面是否存在对应权限
	ans, err = DB.Rdb.Get(ctx, key2).Result()
	if err == nil {
		Logger.Logger.Debug("Dao: Redis Cache Read Permission")
		//此时说明缓存存在权限
		if ans == "true" {
			Permission = true
		} else {
			Permission = false
		}
	} else {
		//缓存没找到，继续在数据库中找，如果redis之前已经确认存在权限，那么就不需要执行这些代码
		//检查对应用户和对应文档的权限
		res := DB.Db.Model(&models.Permission{}).Where("document_id = ? AND user_id = ?", permission.DocumentId, permission.UserId).Select("type").First(&Permission)
		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				if IsPrivate {
					//私有且无权限
					Logger.Logger.Debug("Dao: it's private and you are not permitted")
					return nil, "No Permission"
				} else {
					//公开但没有权限，只读
					Logger.Logger.Debug("Dao: only read")
					return nil, "Only Read"
				}
			}
			//其他问题
			Logger.Logger.Debug("Dao: Internal Error" + res.Error.Error())
			return res.Error, "Internal Error"
		}
	}
	if Permission == true {
		//此时有权限，权限为读写
		Logger.Logger.Debug("Dao: Read And Write")
		DB.Rdb.SetEx(ctx, key2, "true", time.Duration(mul)*time.Hour)
		Logger.Logger.Debug("Dao: Redis Cache Write Permission")
		return nil, "RW"
	}
	//权限为只读.
	Logger.Logger.Debug("Dao: Only Read")
	DB.Rdb.SetEx(ctx, key2, "false", time.Duration(mul)*time.Hour)
	Logger.Logger.Debug("Dao: Redis Cache Write Permission")
	return nil, "Only Read"
}
