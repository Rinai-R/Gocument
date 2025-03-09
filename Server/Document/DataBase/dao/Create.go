package dao

import (
	"context"
	"github.com/Rinai-R/Gocument/Server/Document/DataBase/DB"
	"github.com/Rinai-R/Gocument/Server/Document/DataBase/conf/DB"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	models2 "github.com/Rinai-R/Gocument/pkg/models"
	"strconv"
)

func Create(ctx context.Context, document models2.Document) error {
	tx := DB.Db.Begin()
	res := tx.Create(&document)
	if res.Error != nil {
		tx.Rollback()
		Logger.Logger.Debug("Dao: " + res.Error.Error() + " create failed")
		return res.Error
	}

	res = tx.Model(models2.Permission{}).Create(models2.Permission{
		DocumentId: document.Id,
		UserId:     document.UserId,
		Type:       true,
	})
	if res.Error != nil {
		tx.Rollback()
		Logger.Logger.Debug("Dao: Grant Error" + res.Error.Error())
		return res.Error
	}

	//在es里面初始化文档数据
	ESDocument := models2.ElasticDocument{
		Id:        strconv.Itoa(document.Id),
		UserId:    int64(document.UserId),
		Title:     document.Title,
		IsPrivate: document.IsPrivate,
		Content:   "在这里填写你的文档吧！",
		CreateAt:  document.CreateAt,
		UpdateAt:  document.UpdateAt,
	}
	do, err := DB.ES.Index().
		Index(conf.DocDB.ElasticSearch.IndexName).
		Id(ESDocument.Id).
		BodyJson(ESDocument).
		Do(ctx)
	if err != nil {
		tx.Rollback()
		Logger.Logger.Debug("Dao: " + res.Error.Error() + " create failed")
		return res.Error
	}
	Logger.Logger.Debug("Dao: ES Index Created, Id: " + do.Id)
	Logger.Logger.Debug("Document created")
	tx.Commit()
	return nil
}
