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

func IsHost(username string, DocumentId int) error {
	id, err := GetId(username)
	if err != nil {
		return err
	}
	res := DB.Db.Model(&models.Document{}).Where("user_id = ? AND id = ?", id, DocumentId).First(&models.Document{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
