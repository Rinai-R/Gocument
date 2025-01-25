package dao

import (
	"context"
	"github.com/Rinai-R/Gocument/DataBase/DB"
	conf "github.com/Rinai-R/Gocument/DataBase/conf/DB"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/models"
)

func Edit(ctx context.Context, document models.ElasticDocument) error {
	_, err := DB.ES.Update().
		Index(conf.DB.ElasticSearch.IndexName).
		Id(document.Id).
		Doc(document).
		Do(ctx)
	if err != nil {
		Logger.Logger.Debug("ES Update Failed " + err.Error())
		return err
	}
	return nil
}
