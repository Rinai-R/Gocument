package dao

import (
	"context"
	"errors"
	"github.com/Rinai-R/Gocument/DataBase/DB"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/models"
	"gorm.io/gorm"
)

func Check(_ context.Context, permission models.Permission) (error, string) {
	var IsPrivate bool
	var Permission bool
	res := DB.Db.Model(&models.Document{}).Where("id = ?", permission.DocumentId).Select("IsPrivate").First(&IsPrivate)
	if res.Error != nil {
		Logger.Logger.Debug("Dao: document not found " + res.Error.Error())
		return res.Error, "notfound"
	}

	res = DB.Db.Model(&models.Permission{}).Where("document_id = ? AND user_id = ?", permission.DocumentId, permission.UserId).Select("type").First(&Permission)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			if IsPrivate {
				Logger.Logger.Debug("Dao: it's private and you are not permitted")
				return nil, "No Permission"
			} else {
				Logger.Logger.Debug("Dao: only read")
				return nil, "Only Read"
			}
		}
		Logger.Logger.Debug("Dao: Internal Error" + res.Error.Error())
		return res.Error, "Internal Error"
	}
	if Permission == true {

		Logger.Logger.Debug("Dao: Read And Write")
		return nil, "RW"
	}
	Logger.Logger.Debug("Dao: Only Read")
	return nil, "Only Read"
}
