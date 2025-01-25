package dao

import (
	"errors"
	"github.com/Rinai-R/Gocument/DataBase/DB"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/models"
	"gorm.io/gorm"
)

func Grant(permission models.Permission) error {
	var exist models.Permission
	res := DB.Db.Model(&models.Permission{}).Where("document_id = ? AND user_id = ? AND type = ?", permission.DocumentId, permission.UserId, permission.Type).First(&exist)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			Logger.Logger.Debug("Dao: Created")
			DB.Db.Create(&permission)
			return nil
		}
		Logger.Logger.Error("Dao: Internal Error " + res.Error.Error())
		return res.Error
	}
	res = DB.Db.Where("document_id = ? AND user_id = ? AND type = ?", permission.DocumentId, permission.UserId, permission.Type).First(&permission).Delete(&models.Permission{})
	if res.Error != nil {
		Logger.Logger.Debug("Dao: Delete Error " + res.Error.Error())
		return res.Error
	}
	return nil
}
