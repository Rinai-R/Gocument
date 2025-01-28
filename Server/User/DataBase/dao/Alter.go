package dao

import (
	"errors"
	"fmt"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Server/User/DataBase/DB"
	"github.com/Rinai-R/Gocument/Utils/Error"
	"github.com/Rinai-R/Gocument/models"
	"gorm.io/gorm"
)

func AlterUserInfo(user models.User) error {
	fmt.Println(user)
	err := DB.Db.Model(&user).Where("username = ?", user.Username).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Logger.Logger.Debug("Dao: User Not Found")
			return Error.UserNotExists
		}
		Logger.Logger.Debug("Dao: " + err.Error())
		return err
	}
	updates := map[string]interface{}{}
	if user.Password != "" {
		updates["password"] = user.Password
	}
	if user.Bio != "" {
		updates["bio"] = user.Bio
	}
	if user.Gender != "" {
		updates["gender"] = user.Gender
	}
	if user.Avatar != "" {
		updates["avatar"] = user.Avatar
	}

	if err = DB.Db.Model(&user).Where("username = ?", user.Username).Updates(updates).Error; err != nil {
		Logger.Logger.Debug("Dao: " + err.Error())
		return err
	}
	return nil
}
