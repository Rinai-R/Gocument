package dao

import (
	"context"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Server/User/DataBase/DB"
	"github.com/Rinai-R/Gocument/models"
)

func PersonalPage(_ context.Context, user *models.User) error {

	err := DB.Db.Where("id = ?", user.Id).First(&user).Error
	if err != nil {
		Logger.Logger.Debug("Dao: SQL User Not Found")
		return err
	}

	var document []models.Document
	DB.Db.Where("user_id = ?", user.Id).Find(&document)

	user.Documents = document

	return nil
}
