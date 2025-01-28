package dao

import (
	"context"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Server/Document/DataBase/DB"
	"github.com/Rinai-R/Gocument/Server/Document/DataBase/conf/DB"
	"github.com/Rinai-R/Gocument/models"
)

func Edit(ctx context.Context, document models.ElasticDocument) error {
	tx := DB.Db.Begin()
	var doc models.Document
	err := tx.Model(&doc).Where("id = ?", document.Id).Update("title", document.Title).Error
	if err != nil {
		tx.Rollback()
		Logger.Logger.Error("Dao: MySQL update Error, " + err.Error())
		return err
	}
	document.CreateAt = doc.CreateAt
	document.UpdateAt = doc.UpdateAt
	document.IsPrivate = doc.IsPrivate
	_, err = DB.ES.Update().
		Index(conf.DocDB.ElasticSearch.IndexName).
		Id(document.Id).
		Doc(document).
		Do(ctx)
	if err != nil {
		tx.Rollback()
		Logger.Logger.Debug("ES Update Failed " + err.Error())
		return err
	}
	tx.Commit()
	return nil
}
