package dao

import (
	"errors"
	"github.com/Rinai-R/Gocument/DataBase/DB"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Utils/Error"
	"github.com/Rinai-R/Gocument/Utils/encrypt"
	"github.com/Rinai-R/Gocument/models"
	"gorm.io/gorm"
	"strings"
)

func Register(user models.User) error {
	err := DB.Db.Create(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			Logger.Logger.Debug("DAO: User Already Exist" + err.Error())
			return Error.UserExists
		}
		Logger.Logger.Warn("DAO: " + "Internal Error" + err.Error())
		return err
	}
	Logger.Logger.Debug("DAO: User Created Success")
	return nil
}

func Login(user models.User) error {
	var Origin models.User
	err := DB.Db.Where("username = ?", user.Username).First(&Origin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Logger.Logger.Debug("DAO: User Not Found" + err.Error())
			return Error.UsernameOrPassword
		}
		Logger.Logger.Warn("DAO: " + "Internal Error" + err.Error())
		return err
	}
	if !encrypt.ComparePasswords(Origin.Password, user.Password) {
		Logger.Logger.Debug("DAO: User Password Not Match")
		return Error.UsernameOrPassword
	}
	Logger.Logger.Debug("DAO: Login OK ")
	return nil
}
