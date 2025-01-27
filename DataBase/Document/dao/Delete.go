package dao

import (
	"context"
	"fmt"
	"github.com/Rinai-R/Gocument/DataBase/DB"
	conf "github.com/Rinai-R/Gocument/DataBase/conf/DB"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Utils/Error"
	"github.com/Rinai-R/Gocument/models"
	"strconv"
)

func Delete(ctx context.Context, document models.Document) error {
	tx := DB.Db.Begin()
	fmt.Println(document)
	res := tx.Where("id = ? AND user_id = ?", document.Id, document.UserId).Delete(&models.Document{})
	if res.RowsAffected == 0 {
		tx.Rollback()
		Logger.Logger.Debug("Dao: no document found with the this token")
		return Error.NoDocumentFoundWithToken
	}
	_, err := DB.ES.Delete().
		Index(conf.DB.ElasticSearch.IndexName).
		Id(strconv.Itoa(document.Id)).
		Do(ctx)
	if err != nil {
		tx.Rollback()
		Logger.Logger.Debug("Dao: ES Delete Error " + err.Error())
		return err
	}
	tx.Commit()
	return nil
}
