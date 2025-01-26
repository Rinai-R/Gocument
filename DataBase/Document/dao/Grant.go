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
	res := DB.Db.Model(&models.Permission{}).Where("document_id = ? AND user_id = ?", permission.DocumentId, permission.UserId).First(&exist)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			if permission.Type != exist.Type {
				//如果权限类型不相同就更改权限类型
				DB.Db.Delete(&exist)
				Logger.Logger.Debug("Dao: Updated")
				DB.Db.Create(&permission)
				return nil
			} else {
				//如果权限类型相同就取消权限
				DB.Db.Delete(&exist)
				Logger.Logger.Debug("Dao: Deleted")
				return nil
			}
		}
		Logger.Logger.Error("Dao: Internal Error " + res.Error.Error())
		return res.Error
	}
	DB.Db.Create(&permission)
	Logger.Logger.Debug("Dao: Created")
	return nil
}
