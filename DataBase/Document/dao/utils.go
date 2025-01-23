package dao

import (
	"github.com/Rinai-R/Gocument/DataBase/DB"
	"github.com/Rinai-R/Gocument/models"
)

func GetId(username string) (int, error) {
	var user models.User
	res := DB.Db.Where("username = ?", username).Select("id").First(&user)
	if res.Error != nil {
		return 0, res.Error
	}
	return int(user.Id), nil
}
