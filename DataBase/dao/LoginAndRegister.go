package dao

import (
	"errors"
	"fmt"
	"github.com/Rinai-R/Gocument/DataBase/DB"
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
			fmt.Println("进去了")
			return Error.UserExists
		}
		return err
	}
	return nil
}

func Login(user models.User) error {
	var Origin models.User
	err := DB.Db.Where("username = ?", user.Username).First(&Origin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Error.UsernameOrPassword
		}
		return err
	}
	if !encrypt.ComparePasswords(Origin.Password, user.Password) {
		return Error.UsernameOrPassword
	}
	return nil
}
