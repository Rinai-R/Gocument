package dao

import (
	"github.com/Rinai-R/Gocument/DataBase/DB"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/models"
)

func Create(document models.Document) error {
	err := DB.Db.Create(&document).Error
	if err != nil {
		Logger.Logger.Debug(err.Error() + "elastic connect fail")
		return err
	}
	Logger.Logger.Debug("Document created")
	return nil
}
