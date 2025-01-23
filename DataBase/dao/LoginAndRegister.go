package dao

import (
	"errors"
	"github.com/Rinai-R/Gocument/DataBase/DB"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Utils/Error"
	"github.com/Rinai-R/Gocument/Utils/encrypt"
	"github.com/Rinai-R/Gocument/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

func Register(user models.User) error {
	err := DB.Db.Create(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			Logger.Logger.Debug("DAO: User Already Exist", zap.Error(err))
			return Error.UserExists
		}
		Logger.Logger.Warn("DAO: " + err.Error() + "Internal Error")
		return err
	}
	return nil
}

func Login(user models.User) error {
	var Origin models.User
	err := DB.Db.Where("username = ?", user.Username).First(&Origin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Logger.Logger.Debug("DAO: User Not Found", zap.Error(err))
			return Error.UsernameOrPassword
		}
		Logger.Logger.Debug("DAO: " + err.Error() + "Internal Error")
		return err
	}
	if !encrypt.ComparePasswords(Origin.Password, user.Password) {
		Logger.Logger.Debug("DAO: User Password Not Match")
		return Error.UsernameOrPassword
	}
	return nil
}
