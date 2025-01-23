package dao

import (
	"context"
	"github.com/Rinai-R/Gocument/DataBase/DB"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/models"
)

func PersonalPage(c context.Context, user *models.User) error {
	//key := fmt.Sprintf("u:" + user.Username)
	//if res, err := DB.Rdb.HGetAll(c, key).Result(); err == nil {
	//	if err = mapstructure.Decode(res, &user); err == nil {
	//		Logger.Logger.Info("Dao: Redis Cache Read")
	//		return nil
	//	}
	//	Logger.Logger.Debug("Dao: Redis Cache Read failed")
	//} else {
	//	Logger.Logger.Debug("Dao: Redis Cache User Not Found")
	//}

	err := DB.Db.Where("username = ?", user.Username).First(&user).Error
	if err != nil {
		Logger.Logger.Debug("Dao: SQL User Not Found")
		return err
	}

	var document []models.Document
	DB.Db.Where("user_id = ?", user.Id).Find(&document)

	user.Documents = document

	return nil
}
