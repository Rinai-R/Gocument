package dao

import (
	"context"
	"fmt"
	"github.com/Rinai-R/Gocument/Server/Document/DataBase/DB"
	"github.com/Rinai-R/Gocument/Server/Document/DataBase/conf/DB"
	"github.com/Rinai-R/Gocument/pkg/Error"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	"github.com/Rinai-R/Gocument/pkg/models"
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
		Index(conf.DocDB.ElasticSearch.IndexName).
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
