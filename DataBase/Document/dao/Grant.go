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

func Grant(ctx context.Context, permission models.Permission) error {
	var exist models.Permission
	//设置缓存关键字
	key := fmt.Sprintf("Perm:%v:%v", permission.DocumentId, permission.UserId)
	//设置缓存过期时间
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	mul := r.Intn(60) + 60

	//先看看缓存有没有
	ans, err := DB.Rdb.Get(ctx, key).Result()
	if err != nil {
		//没有，则在mysql里面找
		res := DB.Db.Model(&models.Permission{}).Where("document_id = ? AND user_id = ?", permission.DocumentId, permission.UserId).First(&exist)
		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				DB.Db.Create(&permission)
				if permission.Type == true {
					DB.Rdb.SetEx(ctx, key, "true", time.Duration(mul)*time.Minute)
				} else {
					DB.Rdb.SetEx(ctx, key, "false", time.Duration(mul)*time.Minute)
				}
				Logger.Logger.Debug("Dao: Created")
				return nil
			}
			Logger.Logger.Error("Dao: Internal Error " + res.Error.Error())
			return res.Error
		}
	} else {
		if ans == "true" {
			exist.Type = true
		} else {
			exist.Type = false
		}
	}
	//如果存在，必须要去除原来的权限
	DB.Db.Where("document_id = ? AND user_id = ?", permission.DocumentId, permission.UserId).Delete(&permission)
	if permission.Type != exist.Type {
		//如果权限类型不相同就更改权限类型
		Logger.Logger.Debug("Dao: Updated")
		DB.Db.Create(&permission)
		if permission.Type {
			DB.Rdb.SetEx(ctx, key, "true", time.Duration(mul)*time.Minute)
			Logger.Logger.Debug("Dao: Redis Permission Updated")
		}
	} else {
		//如果权限类型相同就取消权限
		Logger.Logger.Debug("Dao: Deleted")
		DB.Rdb.Del(ctx, key)
	}
	return nil
}
